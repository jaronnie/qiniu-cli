/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put <local file> [flags]",
	Short: "put file to qiniu server",
	Long: `use put command, you can put your local file to qiniu server, and get a url.`,
	Run: put,
}

func put(cmd *cobra.Command, params []string){

	if len(params) <= 0 {

		fmt.Println("please set filepath...")

		os.Exit(0)

	}

	path := params[0]

	index := strings.LastIndexAny(path, "/")

	var upload string

	if index != -1 {

		upload = path[index+1:]

	} else {

		upload = path

	}

	accessKey := viper.GetString("ak")

	secretKey := viper.GetString("sk")

	bucket := viper.GetString("bucket")

	//fmt.Println(accessKey)
	//
	//fmt.Println(secretKey)
	//
	//fmt.Println(bucket)

	putPolicy := storage.PutPolicy{

		Scope: bucket,
	}

	mac := qbox.NewMac(accessKey, secretKey)

	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}

	//bm := storage.NewBucketManager(mac, &cfg)
	//
	//domains, err := bm.ListBucketDomains(bucket)
	//
	//fmt.Println(domains)

	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}

	putExtra := storage.PutExtra{

		Params: map[string]string{

			"x:name": "picture or some other data",
		},
	}


	err := formUploader.PutFile(context.Background(), &ret, upToken, upload, path, &putExtra)

	if err != nil {

		fmt.Println("上传失败")

		return
	}

	fmt.Println("upload successfully")

}


func init() {

	rootCmd.AddCommand(putCmd)

}
