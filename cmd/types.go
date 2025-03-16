package cmd

type VideoToolArguments []string
type VideoToolFn func(arguments VideoToolArguments)
type VideoTool struct {
	Description string
	Usage       func(programName string)
	Run         func(arguments VideoToolArguments)
	Parse       func(arguments VideoToolArguments) (ok bool)
}

type VideoToolsMapping map[string]*VideoTool
