module github.com/bobbygryzynger/ponzu

go 1.15

require (
	content v0.0.0-00010101000000-000000000000
	github.com/blevesearch/bleve v1.0.10
	github.com/boltdb/bolt v1.3.1
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/gorilla/schema v1.2.0
	github.com/nilslice/email v0.1.0
	github.com/nilslice/jwt v1.0.0
	github.com/spf13/cobra v1.0.0
	github.com/tidwall/gjson v1.6.1
	github.com/tidwall/sjson v1.1.1
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	golang.org/x/text v0.3.3
)

replace content => ./content
