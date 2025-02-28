/*
 * @Author: Edwater
 * @Description:
 * @File: root
 * @Version: 1.0.0
 * @Data: 2025/2/20 16:12
 */

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"icongetter/utils"
	"time"
)

var rootCmd = cobra.Command{
	Use:     "IconGetter",
	Version: "v1.0.0",
	Short:   "A tool to extract application icons.",
	Long: `IconGetter is a command-line tool for extracting icons from Android(.apk), iOS(.ipa), and Windows(.exe) applications.
It supports various application formats and can automatically identify and extract icon files.
.___                      ________        __    __                
|   | ____  ____   ____  /  _____/  _____/  |__/  |_  ___________ 
|   |/ ___\/  _ \ /    \/   \  ____/ __ \   __\   __\/ __ \_  __ \
|   \  \__(  <_> )   |  \    \_\  \  ___/|  |  |  | \  ___/|  | \/
|___|\___  >____/|___|  /\______  /\___  >__|  |__|  \___  >__|   
         \/           \/        \/     \/                \/       
Welcome to IconGetter!
`,
	Run: func(cmd *cobra.Command, args []string) {
		// no command
		err := cmd.Help()
		utils.CheckErr(err)
		fmt.Println("\nNo command provided, running default command...\nExtracting icon from files of current directory")
		defaultArg := []string{"default"}
		find, _, err := cmd.Find(defaultArg)
		utils.CheckErr(err)
		find.Run(cmd, args)

		time.Sleep(cobra.MousetrapDisplayDuration)
	},
}

func Execute() {
	cobra.MousetrapHelpText = ""                       // enable to start from explorer.exe
	rootCmd.CompletionOptions.DisableDefaultCmd = true // delete completion
	err := rootCmd.Execute()
	utils.CheckErr(err)
}
