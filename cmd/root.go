package cmd

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qn",
	Short: "simple qiniu-cli",
	Long:  `qiniu-cli is based on qiniu product about go sdk, I just simplify it in order to use it easily`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	homeDir, hErr := homedir.Dir()
	cfgFile = homeDir + "/.qn.toml"
	if hErr != nil {
		fmt.Printf("get current home directory: %v\n", hErr)
		os.Exit(1)
	}
	if !PathExists(cfgFile) {
		viper.SetConfigName(".qn")
		viper.SetConfigType("toml")
		viper.AddConfigPath(homeDir)
		err := viper.SafeWriteConfig()
		if err != nil {
			log.Fatal("write config failed: ", err)
		}
	}
	viper.SetConfigName(".qn")
	viper.SetConfigType("toml")
	viper.AddConfigPath(homeDir)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error")
	}
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {

		return true
	}
	if os.IsNotExist(err) {

		return false
	}
	return false
}
