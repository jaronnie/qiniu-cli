package initalize

import (
	"fmt"
	"os"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/spf13/viper"
)


func InitMAC() *qbox.Mac {
	ak, sk := viper.GetString("qiniu.ak"), viper.GetString("qiniu.sk")
	mac := qbox.NewMac(ak, sk)
	return mac
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read error")
		os.Exit(0)
	}
}
