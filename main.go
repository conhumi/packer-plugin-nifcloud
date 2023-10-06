// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"
	"packer-plugin-nifcloud/builder/scaffolding"
	scaffoldingData "packer-plugin-nifcloud/datasource/scaffolding"
	scaffoldingPP "packer-plugin-nifcloud/post-processor/scaffolding"
	scaffoldingProv "packer-plugin-nifcloud/provisioner/scaffolding"
	scaffoldingVersion "packer-plugin-nifcloud/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder("my-builder", new(scaffolding.Builder))
	pps.RegisterProvisioner("my-provisioner", new(scaffoldingProv.Provisioner))
	pps.RegisterPostProcessor("my-post-processor", new(scaffoldingPP.PostProcessor))
	pps.RegisterDatasource("my-datasource", new(scaffoldingData.Datasource))
	pps.SetVersion(scaffoldingVersion.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
