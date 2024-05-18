/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package versionCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCMDVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Returns the version of the cli app",
		Short: "Something",
		Long:  "Somethig really really long",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Something cool.")
		},
	}
	return cmd
}
