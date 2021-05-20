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
	Long:  `Get/Set AccessKey and SecretKey and Bucket`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 && len(args) != 3 {
			return fmt.Errorf("command account receives zero or four args, received %d\n", len(args))
		}
		return nil
	},
	Run: Account,
}

func welcome() {
	fmt.Println("Welcome you to use qiniu-cli")
	fmt.Println("help you hava a good enjoy")
	fmt.Println("               --by nj-jay")
	a := "        _       _             "
	fmt.Println(a)
	b := " _ __  (_)     (_) __ _ _   _ "
	fmt.Println(b)
	c := "| '_ \\ | |_____| |/ _` | | | |"
	fmt.Println(c)
	d := "| | | || |_____| | (_| | |_| |"
	fmt.Println(d)
	e := "|_| |_|/ |    _/ |\\__,_|\\__, |"
	fmt.Println(e)
	f := "     |__/    |__/       |___/ "
	fmt.Println(f)
}

func Account(cmd *cobra.Command, params []string) {
	ak := viper.GetString("ak")
	sk := viper.GetString("sk")
	bucket := viper.GetString("bucket")
	fmt.Println(ak)
	if len(params) == 0 {
		if ak == "" || sk == "" || bucket == "" {
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
		welcome()
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
