package zaico

import (
	gozaico "github.com/fukata/zaico-go"
	"github.com/mark3labs/mcp-go/server"
)

func NewServer(client *gozaico.Client) *server.MCPServer {
	s := server.NewMCPServer(
		"zaico-mcp-server",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging())

	// Add zaico Resources
	s.AddResourceTemplate(getInventoryResourceContent(client))

	// Add zaico tools - inventories
	s.AddTool(getInventories(client))

	return s
}
