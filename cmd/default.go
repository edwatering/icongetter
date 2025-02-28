/*
 * @Author: Edwater
 * @Description:
 * @File: default
 * @Version: 1.0.0
 * @Data: 2025/2/27 16:39
 */

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"icongetter/utils"
	"os"
	"path/filepath"
)

var defaultCmd = &cobra.Command{
	Use:    "default",
	Short:  "extract icon from files of current directory",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		dir := ""
		output := ""
		utils.SetOutPutPath(output)

		exePath, err := os.Executable()
		utils.CheckErr(err)
		exeName := filepath.Base(exePath)

		dir = "./"
		if utils.DirectoryExists(dir) {
			err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
				if info.Name() == exeName {
					return nil
				}
				if info.IsDir() && path != dir {
					return filepath.SkipDir
				}
				if utils.FileExists(path) {
					err := extractFile(path)
					fmt.Println(err, ":", path)
				}
				return nil
			})
			utils.CheckErr(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(defaultCmd)
}
