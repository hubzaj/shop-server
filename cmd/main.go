package main

import (
	"github.com/hubzaj/golang-component-test/cmd/shop"
	"github.com/hubzaj/golang-component-test/pkg/utils"
	"github.com/spf13/cobra"
)

var mainCommand = &cobra.Command{
	Use:   "server",
	Short: "",
}

func main() {
	mainCommand.AddCommand(
		shop.NewShopCommand(),
	)

	if err := mainCommand.Execute(); err != nil {
		utils.ExitOnError(err, "error executing new shop cmd")
	}
}
