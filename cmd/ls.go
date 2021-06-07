package cmd

import (
	"fmt"
	"strings"

	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/cobra"
)

type FileInfo struct {
	filename string
	size     string
	putTime  string
}

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all files in qiniu server bucket",
	Long:  `If you want to get all files in qiniu server bucket, you can use qn ls, then you will get a list which contains all files`,
	Run:   ls,
}

func ls(cmd *cobra.Command, params []string) {
	bucket := Bm.GetBucket()
	mac := Acc.GetMac()
	cfg := storage.Config{
		UseHTTPS: true,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	limit := 1000
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
}
