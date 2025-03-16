package tools

import (
	"fmt"
	"video-utils/cmd"
	"video-utils/utils"
	"video-utils/utils/flags"
)

var SplitVideo *cmd.VideoTool = &cmd.VideoTool{
	Description: "Splits video into smaller chunks",
	Usage: func(programName string) {
		fmt.Print(
			fmt.Sprintf("Usage: %s <input file> <time to split video>\n\n", programName),
			"Example: split some video into a 5-minute chunk\n",
			fmt.Sprintf("\t%s /path/to/video-file.mp4 00:05:00\n", programName),
		)
	},
	Parse: func(arguments cmd.VideoToolArguments) bool {
		if !utils.IsElementExist(arguments, 0) {
			return false
		}

		if !utils.IsElementExist(arguments, 1) {
			return false
		}

		return true
	},
	Run: func(arguments cmd.VideoToolArguments) {
		inputFile, splitAfterTime := arguments[0], arguments[1]

		originalName := utils.GetFileName(inputFile)
		outputVideoPath := utils.MergeDir("out", originalName+"-%03d.mp4")

		utils.RunFFmpeg(
			flags.INPUT_FILE, inputFile,
			flags.CODEC_NAME, "copy",
			flags.INPUT_STREAM_MAPPING, "0",
			flags.SEGMENT_DURATION, splitAfterTime,
			flags.FORCE_FORMAT, "segment",
			outputVideoPath,
		)

		println("Splited videos has been saved at:", outputVideoPath)
	},
}
