/*
 * @Author: Edwater
 * @Description:
 * @File: extract
 * @Version: 1.0.0
 * @Data: 2025/2/21 10:22
 */

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"icongetter/internal"
	"icongetter/utils"
	"os"
	"path/filepath"
	"strings"
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "extract icon from file",
	Run: func(cmd *cobra.Command, args []string) {
		filename := cmd.Flag("file").Value.String()
		dir := cmd.Flag("dir").Value.String()
		output := cmd.Flag("output").Value.String()
		utils.SetOutPutPath(output)

		exePath, err := os.Executable()
		utils.CheckErr(err)
		exeName := filepath.Base(exePath)

		if filename != "" {
			// file
			if utils.FileExists(filename) {
				fmt.Printf("File: %s\n", filename)
				err := extractFile(filename)
				utils.CheckErr(err)
			}
		} else if dir != "" {
			// dir
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
						utils.CheckErr(err)
					}
					return nil
				})
				utils.CheckErr(err)
			}
		} else {
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
						utils.CheckErr(err)
					}
					return nil
				})
				utils.CheckErr(err)
			}
		}

		if output != "" {
			fmt.Println("Output:", output)
		}
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	extractCmd.Flags().StringP("file", "f", "", "file to extract")
	extractCmd.Flags().StringP("dir", "d", "", "directory to extract")
	extractCmd.Flags().StringP("output", "o", "", "output directory")

}

func extractFile(filename string) error {
	var extractor internal.Extractor
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".apk":
		extractor = &internal.AndroidExtractor{}
		break
	case ".ipa":
		extractor = &internal.IOSExtractor{}
		break
	case ".exe":
		extractor = &internal.WindowsExtractor{}
		break
	default:
		extractor = nil
		return utils.ErrUnsupportedExt
	}
	if extractor != nil {
		return extractor.ExtractIcon(filename)
	}

	return nil
}
