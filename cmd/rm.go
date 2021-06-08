package cmd

import (
	"fmt"

	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/qiniu/go-sdk/v7/cdn"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm filename [filename]",
	Short: "rm file in qiniu server",
	Long:  `you can use rm command to delete file in qiniu server`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("please choose file you delete")
		}
		return nil
	},
	Run: rm,
}

func rm(cmd *cobra.Command, args []string) {
	mac := Bm.GetMac()
	cfg := storage.Config{}
	bm := storage.NewBucketManager(mac, &cfg)
	bucket := Bm.GetBucket()
	err := bm.Delete(bucket, args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("delete %s successfully\n", args[0])
	//删除只有还需要刷新cdn 几秒后就能看到效果！ 还是挺不错的!
	urlsToRefresh := []string{
		"https://picture.nj-jay.com/banner.txt",
	}

	cdnManager := cdn.NewCdnManager(mac)
	_, err = cdnManager.RefreshUrls(urlsToRefresh)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
