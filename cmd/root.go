/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/mchirico/switch-context/profile"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "switch-context",
	Short: "A brief description of your application",
	Long: `switch-context is a CLI tool to switch between AWS profiles
and kubernetes contexts.

   switch-context usprod -f ~/.switchcontext/switchcontext

`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Profiles:")
			profile.SetPath(scFile)

			profiles := profile.ProfilesAvailable()
			for _, p := range profiles {
				fmt.Printf("  %s\n", p)
			}
			return
		}
		if d, err := profile.PR(args[0]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			if d != "" {
				if scFile != "" {
					os.WriteFile(scFile, []byte(d), 0644)
				}
			}

		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var scFile string

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&scFile, "file", "f", "", "output file (default is ~/.switchcontext/switchcontext)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
