/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sclevine/agouti"
)

type LoginInfo struct {
    email string
    password string
}

var (
    loginInfo = &LoginInfo{}
)

// kintaiCmd represents the kintai command
var kintaiCmd = &cobra.Command{
	Use:   "kintai",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("勤怠ページを開く..")
		agoutiDriver := agouti.ChromeDriver()
		agoutiDriver.Start()		

		// リンク先のページを開く
		page, err := agoutiDriver.NewPage()
		if err != nil {
			fmt.Printf("Failed to open a new page. %s\n", err)
			return
		}
		// 自動操作
		page.Navigate("https://id.jobcan.jp/account/profile")
		page.First("#user_email").Fill(loginInfo.email)
		page.First("#user_password").Fill(loginInfo.password)
	
		page.FindByID("login_button").Click()
		page.FindByLink("勤怠").Click()
	},
}

func init() {
	rootCmd.AddCommand(kintaiCmd)

	kintaiCmd.Flags().StringVarP(&loginInfo.email, "email", "e", "default", "email in login")
   	kintaiCmd.Flags().StringVarP(&loginInfo.password, "pass", "p", "default", "password in login")
}
