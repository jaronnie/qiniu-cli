package fs

import (
	"fmt"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/spf13/viper"
)

type Account struct {
	Ak, Sk string
}

func (acc *Account) GetMac() *qbox.Mac {
	acc.Ak = viper.GetString("qn.ak")
	acc.Sk = viper.GetString("qn.sk")
	return qbox.NewMac(acc.Ak, acc.Sk)
}

func (acc *Account) GetAccount() Account {
	return Account{
		Ak: viper.GetString("qn.ak"),
		Sk: viper.GetString("qn.sk"),
	}
}

func (acc *Account) InitAccount(params []string) {
	viper.Set("qn.ak", params[0])
	viper.Set("qn.sk", params[1])
	viper.Set("qn.bucket", params[2])
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("write config failed: ", err)
	}
}

func (acc *Account) ValidAccount() bool {
	return true
}
