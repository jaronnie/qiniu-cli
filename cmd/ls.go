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
	"strings"
)

type FileInfo struct {
	filename string

	size string

	putTime string
}

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all files in qiniu server bucket",
	Long:  `If you want to get all files in qiniu server bucket, you can use qn ls, then you will get a list which contains all files`,
	Run:   ls,
}

func ls(cmd *cobra.Command, params []string) {
	accessKey := viper.GetString("ak")
	secretKey := viper.GetString("sk")
	bucket := viper.GetString("bucket")
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		UseHTTPS: true,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	limit := 10000
	prefix := ""
	delimiter := ""
	marker := ""
	for {
		entries, _, nextMarker, hasNext, err := bucketManager.ListFiles(bucket, prefix, delimiter, marker, limit)
		if err != nil {
			fmt.Println("list error,", err)
			break
		}
		for _, entry := range entries {
			if len(params) == 0 {
				fmt.Println(entry.Key)
			} else if len(params) == 1 && string(params[0][0]) == "*" {
				if strings.HasSuffix(entry.Key, params[0][1:]) {
					filename := entry.Key[0:strings.LastIndexAny(entry.Key, ".")]
					fmt.Println(filename + params[0][1:])
				}
			}
		}

		if hasNext {
			marker = nextMarker
		} else {
			break
		}
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
