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
	"context"
	"encoding/json"
	"fmt"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type upToken struct {
	Token string `json:"token"`
}

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "put file to my qiniu server",
	Run: put,
}


func put(cmd *cobra.Command, args []string) {
	if len(args) <= 0 {
		fmt.Println("please set filepath...")
		os.Exit(0)
	}
	upToken := getToken()
	cfg := storage.Config{
		ApiHost:"http://api.qiniu.com",
	}
	cfg.Zone = &storage.ZoneHuanan
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "picture or some other data",
		},
	}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutPolicy{}
	for _, path := range args {
		index := strings.LastIndexAny(path, "/")
		var upload string
		if index != -1 {
			upload = path[index+1:]
		} else {
			upload = path
		}
		err := formUploader.PutFile(context.Background(), &ret, upToken, upload, path, &putExtra)
		if err != nil {
			fmt.Println(err)
			fmt.Println("上传失败")
			return
		}
		fmt.Println("https://pic.gocloudcoder.com/" + upload)
	}
}

func getToken() string {
	var (
		response *http.Response
		data []byte
		err error
	)

	if response, err = http.Get("http://gocloudcoder.com:8081/upload"); err != nil {
		log.Fatal("请求token失败")
	}
	if data, err = ioutil.ReadAll(response.Body); err != nil {
		log.Fatal("请求token失败")
	}
	uptoken := upToken{}
	if err = json.Unmarshal(data, &uptoken); err != nil {
		log.Fatal("请求token失败")
	}
	return uptoken.Token
}

func init() {
	rootCmd.AddCommand(putCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// putCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// putCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
