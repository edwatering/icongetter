/*
 * @Author: Edwater
 * @Description:
 * @File: file
 * @Version: 1.0.0
 * @Data: 2025/2/21 16:40
 */

package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const SavePathKey = "SetOutPutPath"

func DirectoryExists(dir string) bool {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func FileExists(dir string) bool {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}
	return info.Mode().IsRegular()
}

func SetOutPutPath(output string) {
	if output != "" && filepath.IsAbs(output) {
		outputInfo, err := os.Stat(output)
		CheckErr(err)
		if outputInfo.IsDir() {
			viper.Set(SavePathKey, output)
			fmt.Println("SetOutPutPath:", output)
		} else {
			log.Fatalln("Output is not a directory!")
		}
	} else {
		dirPath, err := os.Getwd()
		CheckErr(err)
		viper.Set(SavePathKey, dirPath)
	}
}

func SavePath(path string) (string, error) {
	savePath := viper.Get(SavePathKey).(string)
	savePath = filepath.Join(savePath, path+"-icon")
	err := os.MkdirAll(savePath, fs.ModePerm)
	fmt.Println("SavePath:", savePath)
	return savePath, err
}
