/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
)

func sso() {
	// Create a session from shared config
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	if err != nil {
		fmt.Println("Error creating session")
		fmt.Println(err)
	}

	// create environment variables from SSO
	accessKey, err := sess.Config.Credentials.Get()

	if err != nil {
		fmt.Println("Error getting credentials")
		fmt.Println(err)
	}

	// create environment variables
	os.Setenv("AWS_ACCESS_KEY_ID", accessKey.AccessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", accessKey.SecretAccessKey)
}

// awsCmd represents the aws command
var awsCmd = &cobra.Command{
	Use: "aws",
	// 	Short: "A brief description of your command",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	//
	//		// Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("aws called")
	//	},
}

func init() {
	rootCmd.AddCommand(awsCmd)

	sso()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// awsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// awsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
