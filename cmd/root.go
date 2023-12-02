/*
Copyright © 2023 Velmurugan Ramasamy
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cf",
	Short: "Create a new folder or empty file",
	Long:  `Createf-cli is a CLI application used to create a new folder or empty file using the cf command`,
	Run: func(cmd *cobra.Command, args []string) {
		isFile, _ := cmd.Flags().GetBool("file")
		isFolder, _ := cmd.Flags().GetBool("folder")

		if len(args) == 0 {
			fmt.Println("Please provide a file or foldername. Use -h or --help to know more about the command")
			return
		}

		if !isFile && !isFolder {
			if strings.Contains(args[0], ".") {
				isFile = true
			} else {
				isFolder = true
			}
		}

		if isFile && isFolder {
			fmt.Println("Please specify only one of --file or --folder.")
			return
		}

		if isFile {
			if len(args) == 0 {
				fmt.Println("Please provide a filename.")
				return
			}

			filename := args[0]
			err := createFile(filename)
			if err != nil {
				fmt.Printf("Error creating file: %v\n", err)
			}
		} else if isFolder {
			if len(args) == 0 {
				fmt.Println("Please provide a folder name.")
				return
			}

			foldername := args[0]
			err := createFolder(foldername)
			if err != nil {
				fmt.Printf("Error creating folder: %v\n", err)
			}
		}
	},
}

func createFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func createFolder(foldername string) error {
	err := os.Mkdir(foldername, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.createf-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("file", "f", false, "Use this flag to create a new file")
	rootCmd.Flags().BoolP("folder", "d", false, "Use this flag to create a new folder")
}
