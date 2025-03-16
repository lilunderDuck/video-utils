package cmd

import (
	"fmt"
	"os"
	"video-utils/utils"
)

type ThisCommandUtils struct {
	ProgramName    string
	SelectedTool   string
	OtherArguments VideoToolArguments
	ToolsMapping   *VideoToolsMapping
}

func New(toolsMapping *VideoToolsMapping) *ThisCommandUtils {
	args := os.Args

	programName := utils.GetFileName(args[0])
	toolToUse, notFoundErr := utils.GetArrayElementAt(args, 1)
	if notFoundErr != nil {
		fmt.Printf(
			"You need to select a tool to use, the syntax is:\n\t%s",
			formatCommand(programName, "<tools>", "<other arguments>\n\n"),
		)

		printAvailableTools(toolsMapping)
		return nil
	}

	// remove the program name first
	args = utils.RemoveElement(args, 0)
	// then the tool name
	args = utils.RemoveElement(args, 0)

	return &ThisCommandUtils{
		SelectedTool:   toolToUse,
		OtherArguments: args,
		ToolsMapping:   toolsMapping,
		ProgramName:    programName + " " + toolToUse,
	}
}

func (c *ThisCommandUtils) Run() {
	mapping := *c.ToolsMapping
	selectedTool, isExist := mapping[c.SelectedTool]
	if !isExist {
		fmt.Printf("Tool %s did not exist", c.SelectedTool)
		return
	}

	parseOk := selectedTool.Parse(c.OtherArguments)
	if !parseOk {
		selectedTool.Usage(c.ProgramName)
		return
	}

	selectedTool.Run(c.OtherArguments)
	fmt.Println("Finished!")
}
