package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/qnfnypen/changuan/internal/domain/param"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "changuan",
	Short: "蝉观",
	Long:  "自用论坛",
}

var versionCmd = cobra.Command{
	Use:   "version",
	Short: "打印版本号",
	Run: func(cmd *cobra.Command, args []string) {
		osv := runtime.GOOS + "/" + runtime.GOARCH
		gov := runtime.Version()
		fmt.Printf("Version: %s %s %s\n", "1.0.0", osv, gov)
	},
}

func init() {
	rootCmd.AddCommand(&versionCmd)
	rootCmd.Flags().StringVar(&param.ConfType, "mode", "test", "程序运行模式，只支持: test debug release三种模式")
}

func cmdExec() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "changuan execute error: %v", err)
		os.Exit(1)
	}
}
