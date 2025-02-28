/*
 * @Author: Edwater
 * @Description: Extractor of Windows
 * @File: windows
 * @Version: 1.0.0
 * @Data: 2025/2/21 15:36
 */

package internal

import (
	"fmt"
	"github.com/tc-hib/winres"
	"icongetter/utils"
	"os"
	"path/filepath"
)

type WindowsExtractor struct{}

func (w *WindowsExtractor) ExtractIcon(path string) error {
	fmt.Println("WindowsExtractor:", path)

	file, err := os.Open(path)
	utils.CheckErr(err)

	resources, err := winres.LoadFromEXE(file)
	utils.CheckErr(err)

	savePath, err := utils.SavePath(path)
	utils.CheckErr(err)

	outputFile, err := os.Create(filepath.Join(savePath, "icon.ico"))
	utils.CheckErr(err)

	resources.Walk(func(typeID, resID winres.Identifier, langID uint16, data []byte) bool {
		switch typeID {
		case winres.RT_GROUP_ICON:
			icon, err := resources.GetIconTranslation(resID, langID)
			utils.CheckErr(err)
			err = icon.SaveICO(outputFile)
			utils.CheckErr(err)
		}
		return true
	})

	defer utils.CheckErr(file.Close())
	return nil
}
