package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:           "tools",
	Long:          `定制组升级包工具`,
	SilenceErrors: true,
	SilenceUsage:  true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func SetVersion(version string) {
	rootCmd.SetVersionTemplate(`{{printf "Version: %s" .Version}}`)
	rootCmd.Version = version
}

func init() {
	var debug bool
	rootCmd.PersistentFlags().SortFlags = false
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "i", false, "show debug info")
}
