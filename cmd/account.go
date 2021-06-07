package cmd

import (
	_ "embed"
	"fmt"

	"github.com/jaronnie/qiniu-cli/cmd/fs"
	"github.com/spf13/cobra"
)

var (
	accountOver bool
)

//go:embed ..\banner.txt
var welcome string

var (
	Acc fs.Account
	Bm  = fs.BucketManger{}
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account [<AccessKey> <SecretKey> <Bucket>]",
	Short: "get/set account information",
	Long:  `Get/Set AccessKey and SecretKey and Bucket`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 && len(args) != 3 {
			return fmt.Errorf("command account receives zero or four args, received %d\n", len(args))
		}
		return nil
	},
	Run: Account,
}

func Account(cmd *cobra.Command, params []string) {
	ak := Acc.GetAccount().Ak
	sk := Acc.GetAccount().Sk
	bucket := Bm.GetBucket()
	if len(params) == 0 {
		if ak == "" || sk == "" || bucket == "" {
			fmt.Println("please set your ak, sk, bucket in config file(default is ~/.qn.toml)")
		} else {
			fmt.Println(welcome)
			fmt.Println("Reading from ~/.qn.toml")
			fmt.Println("ak:	", ak)
			fmt.Println("sk:	", sk)
			fmt.Println("bucket:	", bucket)
		}
	} else if len(params) == 3 && accountOver {
		Acc.InitAccount(params)
		fmt.Println(welcome)
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
