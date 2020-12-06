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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "get account information",
	Long: `Get AccessKey and SecretKey`,
	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) != 0 {

			return fmt.Errorf("command account receives zero or four args, received %d\n", len(args))
		}

		return nil

	},

	Run: Account,
}

func Account(cmd *cobra.Command, params []string) {

	ak := viper.GetString("ak")

	sk := viper.GetString("sk")

	bucket := viper.GetString("bucket")

	if len(params) == 0 {

		if  ak== "" ||  sk== "" || bucket == "" {

			fmt.Println("请在配置文件中设置你的ak, sk, bucket")

		} else {

			fmt.Println("ak:	", ak)
			fmt.Println("sk:	", sk)
			fmt.Println("bucket	", bucket)

		}

	}

}

func init() {

	rootCmd.AddCommand(accountCmd)

}
