package cmd

import (
	"fmt"
	"strings"
)

func formatCommand(programName string, others ...string) string {
	return fmt.Sprintf(
		"%s %s",
		programName,
		strings.Join(others, " "),
	)
}

func printAvailableTools(toolsMap *VideoToolsMapping) {
	fmt.Println("Here is a list of available tools:")
	for toolName, tool := range *toolsMap {
		fmt.Printf("- %s: %s\n", toolName, tool.Description)
	}
}
