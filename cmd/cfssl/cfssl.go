/*
cfssl is the command line tool to issue/sign/bundle client certificate. It's
also a tool to start a HTTP server to handle web requests for signing, bundling
and verification.

Usage:
	cfssl command [-flags] arguments

	The commands are

	bundle	 create a certificate bundle
	sign	 signs a certificate signing request (CSR)
	serve	 starts a HTTP server handling sign and bundle requests
	version	 prints the current cfssl version
	genkey   generates a key and an associated CSR
	gencert  generates a key and a signed certificate
	gencsr   generates a certificate request
	selfsign generates a self-signed certificate

Use "cfssl [command] -help" to find out more about a command.
*/
package main

import (
	"flag"
	"os"

	"github.com/redNixon/cfssl/cli"
	"github.com/redNixon/cfssl/cli/bundle"
	"github.com/redNixon/cfssl/cli/certinfo"
	"github.com/redNixon/cfssl/cli/crl"
	"github.com/redNixon/cfssl/cli/gencert"
	"github.com/redNixon/cfssl/cli/gencrl"
	"github.com/redNixon/cfssl/cli/gencsr"
	"github.com/redNixon/cfssl/cli/genkey"
	"github.com/redNixon/cfssl/cli/info"
	"github.com/redNixon/cfssl/cli/ocspdump"
	"github.com/redNixon/cfssl/cli/ocsprefresh"
	"github.com/redNixon/cfssl/cli/ocspserve"
	"github.com/redNixon/cfssl/cli/ocspsign"
	"github.com/redNixon/cfssl/cli/printdefault"
	"github.com/redNixon/cfssl/cli/revoke"
	"github.com/redNixon/cfssl/cli/scan"
	"github.com/redNixon/cfssl/cli/selfsign"
	"github.com/redNixon/cfssl/cli/serve"
	"github.com/redNixon/cfssl/cli/sign"
	"github.com/redNixon/cfssl/cli/version"

	_ "github.com/go-sql-driver/mysql" // import to support MySQL
	_ "github.com/lib/pq"              // import to support Postgres
	_ "github.com/mattn/go-sqlite3"    // import to support SQLite3
)

// main defines the cfssl usage and registers all defined commands and flags.
func main() {
	// Add command names to cfssl usage
	flag.Usage = nil // this is set to nil for testabilty
	// Register commands.
	cmds := map[string]*cli.Command{
		"bundle":         bundle.Command,
		"certinfo":       certinfo.Command,
		"crl":            crl.Command,
		"sign":           sign.Command,
		"serve":          serve.Command,
		"version":        version.Command,
		"genkey":         genkey.Command,
		"gencert":        gencert.Command,
		"gencsr":         gencsr.Command,
		"gencrl":         gencrl.Command,
		"ocspdump":       ocspdump.Command,
		"ocsprefresh":    ocsprefresh.Command,
		"ocspsign":       ocspsign.Command,
		"ocspserve":      ocspserve.Command,
		"selfsign":       selfsign.Command,
		"scan":           scan.Command,
		"info":           info.Command,
		"print-defaults": printdefaults.Command,
		"revoke":         revoke.Command,
	}

	// If the CLI returns an error, exit with an appropriate status
	// code.
	err := cli.Start(cmds)
	if err != nil {
		os.Exit(1)
	}
}
