/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: fetch,
}

func fetch(cmd *cobra.Command, params []string) {
	remoteUrl := params[0]
	accessKey := viper.GetString("ak")
	secretKey := viper.GetString("sk")
	bucket := viper.GetString("bucket")
	if accessKey == "" || secretKey == "" || bucket == "" {
		fmt.Println("请设置ak, sk, bucket")
		os.Exit(0)
	}

	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		ApiHost:"http://api.qiniu.com",
	}
	bm := storage.NewBucketManager(mac, &cfg)
	res, err := bm.Fetch(remoteUrl, bucket, "test123456.png")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Key)

}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
