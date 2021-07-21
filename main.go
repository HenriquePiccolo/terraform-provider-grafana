package main

import (
	"context"
	"flag"
	"log"

	"github.com/HenriquePiccolo/terraform-provider-grafana-v2/grafana_v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{ProviderFunc: grafana-v2.Provider(version)}

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/HenriquePiccolo/grafana_v2", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}
