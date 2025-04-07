package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	stdlog "log"
	"os"
	"os/signal"
	"syscall"

	gozaico "github.com/fukata/zaico-go"
	"github.com/fukata/zaico-mcp-server/pkg/zaico"
	"github.com/mark3labs/mcp-go/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "zaico-mcp-server",
	Short: "Zaico MCP Server",
	Long:  "Zaico MCP Server is a server application for MCP protocol",
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("verbose") {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.InfoLevel)
		}
		log.Info("Starting Zaico MCP Server...")

		logFile := viper.GetString("log-file")
		prettyPrintJSON := viper.GetBool("pretty-print-json")
		logger, err := initLogger(logFile)
		if err != nil {
			stdlog.Fatal("Failed to initialize logger:", err)
		}
		logCommands := viper.GetBool("enable-command-logging")

		cfg := runConfig{
			logger:          logger,
			logCommands:     logCommands,
			prettyPrintJSON: prettyPrintJSON,
		}
		if err := runStdioServer(cfg); err != nil {
			log.Fatal(err)
		}
	},
	Version: "1.0.0",
}

func initLogger(outPath string) (*log.Logger, error) {
	if outPath == "" {
		return log.New(), nil
	}

	file, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	logger.SetOutput(file)

	return logger, nil
}

type runConfig struct {
	logger          *log.Logger
	logCommands     bool
	prettyPrintJSON bool
}

// JSONPrettyPrintWriter is a Writer that pretty prints input to indented JSON
type JSONPrettyPrintWriter struct {
	writer io.Writer
}

func (j JSONPrettyPrintWriter) Write(p []byte) (n int, err error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, p, "", "\t"); err != nil {
		return 0, err
	}
	return j.writer.Write(prettyJSON.Bytes())
}

func runStdioServer(cfg runConfig) error {
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
		cfg.logger.Info("shutting down server...")
	case err := <-errC:
		if err != nil {
			return fmt.Errorf("error running server: %w", err)
		}
	}

	return nil
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose logging")
	rootCmd.PersistentFlags().String("log-file", "", "Path to log file")
	rootCmd.PersistentFlags().Bool("enable-command-logging", false, "When enabled, the server will log all command requests and responses to the log file")
	rootCmd.PersistentFlags().Bool("pretty-print-json", false, "Pretty print JSON output")
	rootCmd.PersistentFlags().StringP("zaico-api-key", "", "", "Zaico API Key")
	rootCmd.PersistentFlags().StringP("zaico-api-endpoint", "", "https://web.zaico.co.jp/api/v1/", "Zaico API Endpoint")

	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("log-file", rootCmd.PersistentFlags().Lookup("log-file"))
	viper.BindPFlag("enable-command-logging", rootCmd.PersistentFlags().Lookup("enable-command-logging"))
	viper.BindPFlag("pretty-print-json", rootCmd.PersistentFlags().Lookup("pretty-print-json"))
	viper.BindPFlag("zaico-api-key", rootCmd.PersistentFlags().Lookup("zaico-api-key"))
	viper.BindPFlag("zaico-api-endpoint", rootCmd.PersistentFlags().Lookup("zaico-api-endpoint"))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
