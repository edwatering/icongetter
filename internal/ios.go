/*
 * @Author: Edwater
 * @Description: Extractor of iOS
 * @File: ios
 * @Version: 1.0.0
 * @Data: 2025/2/21 15:36
 */

package internal

import (
	"archive/zip"
	"fmt"
	"icongetter/utils"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type IOSExtractor struct{}

func (i *IOSExtractor) ExtractIcon(path string) error {
	fmt.Println("IOSExtractor:", path)

	savePath, err := utils.SavePath(path)
	utils.CheckErr(err)

	res, err := zip.OpenReader(path)
	utils.CheckErr(err)
	for _, f := range res.File {
		if strings.Contains(f.Name, ".app/") && strings.Contains(f.Name, "Icon") && strings.HasSuffix(f.Name, ".png") {
			rc, err := f.Open()
			utils.CheckErr(err)
			iconData, err := io.ReadAll(rc)
			utils.CheckErr(err)
			err = os.WriteFile(filepath.Join(savePath, filepath.Base(f.Name)), iconData, fs.ModePerm)
			utils.CheckErr(err)
			defer utils.CheckErr(rc.Close())
		}
	}

	defer utils.CheckErr(res.Close())
	return nil
}
