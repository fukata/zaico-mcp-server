package main
import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "zaico-mcp-server",
	Short: "Zaico MCP Server",
	Long:  "Zaico MCP Server is a server application for MCP protocol",
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("verbose") {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}
		logrus.Info("Starting Zaico MCP Server...")
	},
	Version: "1.0.0",
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose logging")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
