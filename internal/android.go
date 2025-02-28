/*
 * @Author: Edwater
 * @Description: Extractor of Android
 * @File: android
 * @Version: 1.0.0
 * @Data: 2025/2/21 15:36
 */

package internal

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/avast/apkparser"
	"icongetter/utils"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type Manifest struct {
	Package     string      `xml:"package,attr"`
	Application Application `xml:"application"`
}

type Application struct {
	Icon string `xml:"icon,attr"`
}

type AndroidExtractor struct{}

func (a *AndroidExtractor) ExtractIcon(path string) error {
	fmt.Println("AndroidExtractor:", path)

	var buffer bytes.Buffer
	enc := xml.NewEncoder(&buffer)
	enc.Indent("", "\t")
	zipErr, resErr, manErr := apkparser.ParseApk(path, enc)
	utils.CheckErr(zipErr)
	utils.CheckErr(resErr)
	utils.CheckErr(manErr)

	var manifest Manifest
	// parse AndroidManifest.xml
	if err := xml.Unmarshal(buffer.Bytes(), &manifest); err != nil {
		log.Fatalf("Error parsing AndroidManifest.xml: %v", err)
	}

	r, err := zip.OpenReader(path)
	utils.CheckErr(err)

	iconFilePath := manifest.Application.Icon
	if iconFilePath != "" {
		for _, file := range r.File {
			if file.Name == iconFilePath {
				savePath, err := utils.SavePath(path)
				utils.CheckErr(err)

				rc, err := file.Open()
				utils.CheckErr(err)
				iconData, err := io.ReadAll(rc)
				utils.CheckErr(err)
				err = os.WriteFile(filepath.Join(savePath, "icon.png"), iconData, fs.ModePerm)
				utils.CheckErr(err)

				defer utils.CheckErr(rc.Close())
			}
		}
	}
	defer utils.CheckErr(r.Close())

	return nil
}
