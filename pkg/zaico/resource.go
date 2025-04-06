package zaico

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/fukata/zaico-go/zaico"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func getInventoryResourceContent(client *zaico.Client) (mcp.ResourceTemplate, server.ResourceTemplateHandlerFunc) {
	return mcp.NewResourceTemplate(
			"https://web.zaico.co.jp/inventories/{inventoryId}",
			"Inventory Content", // Resource template
		),
		inventoryResourceContentsHandler(client)
}

func inventoryResourceContentsHandler(client *zaico.Client) func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	return func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		inventoryId, ok := request.Params.Arguments["inventoryId"].(string)
		if !ok || len(inventoryId) == 0 {
			return nil, errors.New("inventory id is required")
		}
		inventoryIdInt, err := strconv.Atoi(inventoryId)
		if err != nil {
			return nil, errors.New("inventory id is not a number")
		}

		inventory, err := client.Inventory.Get(ctx, inventoryIdInt)
		if err != nil {
			return nil, err
		}

		var contents []mcp.ResourceContents
		content := inventoryResourceContent(inventory)
		contents = append(contents, mcp.TextResourceContents{
			URI:      "https://web.zaico.co.jp/inventories/" + strconv.Itoa(inventory.ID),
			MIMEType: "application/json",
			Text:     content,
		})
		return contents, nil
	}
}

// inventoryResourceContent returns the json styled content of the inventory.
func inventoryResourceContent(inventory *zaico.Inventory) string {
	jsonStr, err := json.Marshal(inventory)
	if err != nil {
		return ""
	}
	return string(jsonStr)
}
