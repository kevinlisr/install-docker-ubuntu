/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "版本信息",
	Long: `关于版本的长信息`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.0.1Alpha")
		author, err := cmd.Flags().GetString("author")
		if err != nil {
			fmt.Println("请输入正确的作者信息")
			return
		}
		fmt.Println("作者是：", author)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	versionCmd.Flags().StringP("author","a","绿巨人","作者信息")
}
