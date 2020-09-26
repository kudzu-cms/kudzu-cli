package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// RunCmd runs the ponzu server.
var RunCmd = &cobra.Command{
	Use:   "run [flags] <service(,service)>",
	Short: "starts the 'ponzu' HTTP server for the JSON API and or Admin System.",
	Long: `Starts the 'ponzu' HTTP server for the JSON API, Admin System, or both.
The segments, separated by a comma, describe which services to start, either
'admin' (Admin System / CMS backend) or 'api' (JSON API), and, optionally,
if the server should utilize TLS encryption - served over HTTPS, which is
automatically managed using Let's Encrypt (https://letsencrypt.org)

Defaults to 'run --port=8080 admin,api' (running Admin & API on port 8080, without TLS)

Note:
Admin and API cannot run on separate processes unless you use a copy of the
database, since the first process to open it receives a lock. If you intend
to run the Admin and API on separate processes, you must call them with the
'ponzu' command independently.`,
	Example: `$ ponzu run
(or)
$ ponzu run --port=8080 --https admin,api
(or)
$ ponzu run admin
(or)
$ ponzu run --port=8888 api`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var addTLS string
		if https {
			addTLS = "--https"
		} else {
			addTLS = "--https=false"
		}

		if devhttps {
			addTLS = "--dev-https"
		}

		var addDocs string
		if docs {
			addDocs = "--docs"
		} else {
			addDocs = "--docs=false"
		}

		var services string
		if len(args) > 0 {
			services = args[0]
		} else {
			services = "admin,api"
		}

		name := buildOutputName()
		buildPathName := strings.Join([]string{".", name}, string(filepath.Separator))
		serve := exec.Command(buildPathName,
			"serve",
			services,
			fmt.Sprintf("--bind=%s", bind),
			fmt.Sprintf("--port=%d", port),
			fmt.Sprintf("--https-port=%d", httpsport),
			fmt.Sprintf("--docs-port=%d", docsport),
			addDocs,
			addTLS,
		)
		serve.Stderr = os.Stderr
		serve.Stdout = os.Stdout

		return serve.Run()
	},
}

func init() {

	RunCmd.Flags().StringVar(&bind, "bind", "localhost", "address for ponzu to bind the HTTP(S) server")
	RunCmd.Flags().IntVar(&httpsport, "https-port", 443, "port for ponzu to bind its HTTPS listener")
	RunCmd.Flags().IntVar(&port, "port", 8080, "port for ponzu to bind its HTTP listener")
	RunCmd.Flags().IntVar(&docsport, "docs-port", 1234, "[dev environment] override the documentation server port")
	RunCmd.Flags().BoolVar(&docs, "docs", false, "[dev environment] run HTTP server to view local HTML documentation")
	RunCmd.Flags().BoolVar(&https, "https", false, "enable automatic TLS/SSL certificate management")
	RunCmd.Flags().BoolVar(&devhttps, "dev-https", false, "[dev environment] enable automatic TLS/SSL certificate management")

	RootCmd.AddCommand(RunCmd)

}
