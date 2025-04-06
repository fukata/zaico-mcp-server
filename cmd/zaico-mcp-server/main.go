package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	gozaico "github.com/fukata/zaico-go"
	"github.com/fukata/zaico-mcp-server/pkg/zaico"
	"github.com/mark3labs/mcp-go/server"
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

		if err := runStdioServer(); err != nil {
			logrus.Fatal(err)
		}
	},
	Version: "1.0.0",
}

func runStdioServer() error {
	// Create app context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	c := gozaico.NewClientWithBaseURL(viper.GetString("zaico-api-key"), viper.GetString("zaico-api-endpoint"))
	stdioServer := server.NewStdioServer(zaico.NewServer(c))

	// Start listening for messages
	errC := make(chan error, 1)
	go func() {
		in, out := io.Reader(os.Stdin), io.Writer(os.Stdout)

		// if cfg.logCommands {
		// 	loggedIO := iolog.NewIOLogger(in, out, cfg.logger)
		// 	in, out = loggedIO, loggedIO
		// }

		// if cfg.prettyPrintJSON {
		// 	out = JSONPrettyPrintWriter{writer: out}
		// }
		errC <- stdioServer.Listen(ctx, in, out)
	}()

	// Output github-mcp-server string
	_, _ = fmt.Fprintf(os.Stderr, "GitHub MCP Server running on stdio\n")

	// Wait for shutdown signal
	select {
	case <-ctx.Done():
		// cfg.logger.Infof("shutting down server...")
		_, _ = fmt.Fprintf(os.Stderr, "shutting down server...\n")
	case err := <-errC:
		if err != nil {
			return fmt.Errorf("error running server: %w", err)
		}
	}

	return nil
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose logging")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	rootCmd.PersistentFlags().StringP("zaico-api-key", "", "", "Zaico API Key")
	viper.BindPFlag("zaico-api-key", rootCmd.PersistentFlags().Lookup("zaico-api-key"))
	rootCmd.PersistentFlags().StringP("zaico-api-endpoint", "", "https://web.zaico.co.jp/api/v1/", "Zaico API Endpoint")
	viper.BindPFlag("zaico-api-endpoint", rootCmd.PersistentFlags().Lookup("zaico-api-endpoint"))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
