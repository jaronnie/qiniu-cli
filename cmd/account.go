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

var (

	accountOver bool

)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account [<AccessKey> <SecretKey> <Bucket>]",
	Short: "get/set account information",
	Long: `Get/Set AccessKey and SecretKey and Bucket`,
	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) != 0 && len(args) != 3 {

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

			fmt.Println("please set your ak, sk, bucket in config file(default is ~/.qn.toml)")

		} else {

			fmt.Println("Reading from ~/.qn.toml....")
			fmt.Println("ak:	", ak)
			fmt.Println("sk:	", sk)
			fmt.Println("bucket:	", bucket)
		}

	} else if len(params) == 3 && accountOver {

		viper.Set("ak", params[0])

		viper.Set("sk", params[1])

		viper.Set("bucket", params[2])

		err := viper.WriteConfig()

		if err != nil {

			fmt.Println("write config failed: ", err)
		}

	} else if len(params) == 3 && !accountOver {


		fmt.Println("please use -w flags. For example, qn account -w ak sk bucket")

	}

}

func init() {

	accountCmd.Flags().BoolVarP(
		&accountOver,
		"overwrite",
		"w",
		false,
		"overwrite account or not when account exists in local config file(default ~/.qn.toml), by default not overwrite",

	)

	rootCmd.AddCommand(accountCmd)

}
