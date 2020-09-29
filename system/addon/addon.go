package addon

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/kudzu-cms/kudzu/system/db"
	"github.com/kudzu-cms/kudzu/system/item"

	"github.com/tidwall/sjson"
)

var (
	// Types is a record of addons, like content types, of addon_reverse_dns:interface{}
	Types = make(map[string]func() interface{})
)

const (
	// StatusEnabled defines string status for Addon enabled state
	StatusEnabled = "enabled"
	// StatusDisabled defines string status for Addon disabled state
	StatusDisabled = "disabled"
)

// Meta contains the basic information about the addon
type Meta struct {
	kudzuAddonName       string `json:"addon_name"`
	kudzuAddonAuthor     string `json:"addon_author"`
	kudzuAddonAuthorURL  string `json:"addon_author_url"`
	kudzuAddonVersion    string `json:"addon_version"`
	kudzuAddonReverseDNS string `json:"addon_reverse_dns"`
	kudzuAddonStatus     string `json:"addon_status"`
}

// Addon contains information about a provided addon to the system
type Addon struct {
	item.Item
	Meta
}

// Register constructs a new addon and registers it with the system. Meta is a
// addon.Meta and fn is a closure returning a pointer to your own addon type
func Register(m Meta, fn func() interface{}) Addon {
	// get or create the reverse DNS identifier
	if m.kudzuAddonReverseDNS == "" {
		revDNS, err := reverseDNS(m)
		if err != nil {
			panic(err)
		}

		m.kudzuAddonReverseDNS = revDNS
	}

	Types[m.kudzuAddonReverseDNS] = fn

	a := Addon{Meta: m}

	err := register(a)
	if err != nil {
		panic(err)
	}

	return a
}

// register sets up the system to use the Addon by:
// 1. Validating the Addon struct
// 2. Saving it to the __addons bucket in DB with id/key = addon_reverse_dns
func register(a Addon) error {
	if a.kudzuAddonName == "" {
		return fmt.Errorf(`Addon must have valid Meta struct embedded: missing %s field.`, "kudzuAddonName")
	}
	if a.kudzuAddonAuthor == "" {
		return fmt.Errorf(`Addon must have valid Meta struct embedded: missing %s field.`, "kudzuAddonAuthor")
	}
	if a.kudzuAddonAuthorURL == "" {
		return fmt.Errorf(`Addon must have valid Meta struct embedded: missing %s field.`, "kudzuAddonAuthorURL")
	}
	if a.kudzuAddonVersion == "" {
		return fmt.Errorf(`Addon must have valid Meta struct embedded: missing %s field.`, "kudzuAddonVersion")
	}

	if _, ok := Types[a.kudzuAddonReverseDNS]; !ok {
		return fmt.Errorf(`Addon "%s" has no record in the addons.Types map`, a.kudzuAddonName)
	}

	// check if addon is already registered in db as addon_reverse_dns
	if db.AddonExists(a.kudzuAddonReverseDNS) {
		return nil
	}

	// convert a.Item into usable data, Item{} => []byte(json) => map[string]interface{}
	kv := make(map[string]interface{})

	data, err := json.Marshal(a.Item)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &kv)
	if err != nil {
		return err
	}

	// save new addon to db
	vals := make(url.Values)
	for k, v := range kv {
		vals.Set(k, fmt.Sprintf("%v", v))
	}

	vals.Set("addon_name", a.kudzuAddonName)
	vals.Set("addon_author", a.kudzuAddonAuthor)
	vals.Set("addon_author_url", a.kudzuAddonAuthorURL)
	vals.Set("addon_version", a.kudzuAddonVersion)
	vals.Set("addon_reverse_dns", a.kudzuAddonReverseDNS)
	vals.Set("addon_status", StatusDisabled)

	// db.SetAddon is like SetContent, but rather than the key being an int64 ID,
	// we need it to be a string based on the addon_reverse_dns
	kind, ok := Types[a.kudzuAddonReverseDNS]
	if !ok {
		return fmt.Errorf("Error: no addon to set with id: %s", a.kudzuAddonReverseDNS)
	}

	err = db.SetAddon(vals, kind())
	if err != nil {
		return err
	}

	return nil
}

// Deregister removes an addon from the system. `key` is the addon_reverse_dns
func Deregister(key string) error {
	err := db.DeleteAddon(key)
	if err != nil {
		return err
	}

	delete(Types, key)
	return nil
}

// Enable sets the addon status to `enabled`. `key` is the addon_reverse_dns
func Enable(key string) error {
	err := setStatus(key, StatusEnabled)
	if err != nil {
		return err
	}

	return nil
}

// Disable sets the addon status to `disabled`. `key` is the addon_reverse_dns
func Disable(key string) error {
	err := setStatus(key, StatusDisabled)
	if err != nil {
		return err
	}

	return nil
}

// KeyFromMeta creates a unique string identifier for an addon based on its url and name
func KeyFromMeta(meta Meta) (string, error) {
	return reverseDNS(meta)
}

func setStatus(key, status string) error {
	a, err := db.Addon(key)
	if err != nil {
		return err
	}

	a, err = sjson.SetBytes(a, "addon_status", status)
	if err != nil {
		return err
	}

	kind, ok := Types[key]
	if !ok {
		return fmt.Errorf("Error: no addon to set with id: %s", key)
	}

	// convert json => map[string]interface{} => url.Values
	var kv map[string]interface{}
	err = json.Unmarshal(a, &kv)
	if err != nil {
		return err
	}

	vals := make(url.Values)
	for k, v := range kv {
		switch v.(type) {
		case []string:
			s := v.([]string)
			for i := range s {
				if i == 0 {
					vals.Set(k, s[i])
				}

				vals.Add(k, s[i])
			}
		default:
			vals.Set(k, fmt.Sprintf("%v", v))
		}
	}

	err = db.SetAddon(vals, kind())
	if err != nil {
		return err
	}

	return nil
}

func reverseDNS(meta Meta) (string, error) {
	u, err := url.Parse(meta.kudzuAddonAuthorURL)
	if err != nil {
		return "", nil
	}

	if u.Host == "" {
		return "", fmt.Errorf(`Error parsing Addon Author URL: %s. Ensure URL is formatted as "scheme://hostname/path?query" (path & query optional)`, meta.kudzuAddonAuthorURL)
	}

	name := strings.Replace(meta.kudzuAddonName, " ", "", -1)

	// reverse the host name parts, split on '.', ex. bosssauce.it => it.bosssauce
	parts := strings.Split(u.Host, ".")
	strap := make([]string, 0, len(parts))
	for i := len(parts) - 1; i >= 0; i-- {
		strap = append(strap, parts[i])
	}

	return strings.Join(append(strap, name), "."), nil
}

// String returns the addon name and overrides the item String() method in
// item.Identifiable interface
func (a *Addon) String() string {
	return a.kudzuAddonName
}
