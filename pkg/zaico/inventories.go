package zaico

import (
	"context"
	"encoding/json"
	"fmt"

	gozaico "github.com/fukata/zaico-go"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// getInventories creates a tool to get inventories from zaico.
func getInventories(client *gozaico.Client) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_inventories",
			mcp.WithDescription("search zaico inventories"),
			mcp.WithString("title",
				mcp.Description("search title"),
			),
			mcp.WithString("code",
				mcp.Description("search code"),
			),
			mcp.WithString("place",
				mcp.Description("search place"),
			),
			mcp.WithString("category",
				mcp.Description("search category"),
			),
			mcp.WithString("optional_attributes_name",
				mcp.Description("search optional_attributes_name"),
			),
			mcp.WithString("optional_attributes_value",
				mcp.Description("search optional_attributes_value"),
			),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			inventories, err := client.Inventory.List(ctx, &gozaico.InventoryListOptions{
				Title:                   request.Params.Arguments["title"].(string),
				Code:                    request.Params.Arguments["code"].(string),
				Place:                   request.Params.Arguments["place"].(string),
				Category:                request.Params.Arguments["category"].(string),
				OptionalAttributesName:  request.Params.Arguments["optional_attributes_name"].(string),
				OptionalAttributesValue: request.Params.Arguments["optional_attributes_value"].(string),
				Page:                    1,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to get inventories: %w", err)
			}

			r, err := json.Marshal(inventories)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal inventories: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}
