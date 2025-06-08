package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mycli",                                   // ชื่อของโปรแกรม CLI ของคุณเมื่อรัน (เช่น `mycli`)
	Short: "A brief description of your application", // คำอธิบายสั้นๆ ที่แสดงใน `mycli --help`
	Long: `A longer description that spans multiple lines and
likely contains examples and usage of using your application.
For example:

mycli init         # Initialize something
mycli serve --port 8080 # Start a server`,
	// Run: จะถูกเรียกใช้เมื่อไม่มี sub-command ใดๆ ถูกระบุ หรือเมื่อ command นี้ถูกเรียกโดยตรง
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to my CLI Application!")
		fmt.Println("Try 'mycli --help' for more information.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err) // พิมพ์ error ไปยัง stderr
		os.Exit(1)                                 // ออกจากโปรแกรมด้วย exit code 1 (บ่งบอกว่ามี error)
	}
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Display help information",
	Long:  "Display detailed help information for the CLI application.",
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags which will be available to all subcommands of this command.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mycli.yaml)")

	// Cobra also supports local flags which will only run when this command is called directly
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(helpCmd)     // เพิ่มคำสั่ง help เข้าไปใน root command
	rootCmd.SetHelpCommand(helpCmd) // กำหนดคำสั่ง help ให้กับ root command
}
