/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/mchirico/switch-context/profile"
	"os"

	"github.com/spf13/cobra"
)

// updatefileCmd represents the updatefile command
var updatefileCmd = &cobra.Command{
	Use:   "updatefile",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			return
		}
		if d, err := profile.PR(args[0]); err != nil {
			os.Exit(1)
		} else {
			if d != "" {
				os.WriteFile(args[1], []byte(d), 0644)
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(updatefileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updatefileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updatefileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
