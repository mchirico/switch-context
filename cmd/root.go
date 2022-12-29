/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mchirico/switch-context/constants"
	"github.com/mchirico/switch-context/profile"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "switch-context",
	Short: "switch-context is a CLI tool to switch between AWS profiles",
	Long: fmt.Sprintf(`switch-context is a CLI tool to switch between AWS profiles
and kubernetes contexts. (version: %s)

   switch-context usprod -f ~/.switchcontext/switchcontext

`, constants.VERSION),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Profiles:")
			profile.SetPath(scFile)

			profiles := profile.ProfilesAvailable()
			last, err := profile.LastKey()
			if err != nil {
				for _, p := range profiles {
					fmt.Printf("  %s\n", p)
				}
			} else {
				for _, p := range profiles {
					if p == last {
						color.Red("  %s\n", p)
					} else {
						fmt.Printf("  %s\n", p)
					}
				}
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
