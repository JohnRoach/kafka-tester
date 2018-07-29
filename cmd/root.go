package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cfgFile is the config file path
var cfgFile string

// AppVersion is the application version
var AppVersion = "1.0.0"

// Debug allows printing of bunch of DEBUG values in log
var Debug bool

// DisableColor disables the color for a given run
var DisableColor bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kafkaTester",
	Short: "...",
	Long: `...

For example:
    kafkaTester ...
    kafkaTester ... --output json
`,
}

// Execute is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(versionCmd)
	//rootCmd.AddCommand(testConfigCmd)

	// Persistent flags. Flags that will live for all subcommands.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kafkaTester.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "", false, "adding debug to the logging")
	rootCmd.PersistentFlags().BoolVarP(&DisableColor, "disable-color", "", false, "disable color for logging output")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kafkaTester" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".kafkaTester")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kafkaTester",
	Long:  `All software has versions. This is kafkaTesters's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(printVersion())
	},
}

// printVersion prints out version which is defined globally
func printVersion() string {
	return AppVersion
}
