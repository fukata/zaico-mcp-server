package zaico

import (
	"github.com/fukata/zaico-go/zaico"
	"github.com/mark3labs/mcp-go/server"
)

func NewServer(client *zaico.Client) *server.MCPServer {
	s := server.NewMCPServer(
		"zaico-mcp-server",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging())

	// Add zaico Resources
	s.AddResourceTemplate(getInventoryResourceContent(client))

	// Add zaico tools - inventories

	return s
}
