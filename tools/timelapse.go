package tools

import (
	"fmt"
	"video-utils/cmd"
	"video-utils/utils"
	"video-utils/utils/flags"
)

var MakeTimelapse *cmd.VideoTool = &cmd.VideoTool{
	Description: "Make a timelapse video.",
	Usage: func(programName string) {
		fmt.Print(
			fmt.Sprintf("Usage: %s <input file>\n\n", programName),
			"Example:\n",
			fmt.Sprintf("\t%s /path/to/video-file.mp4\n\n", programName),
		)
	},
	Parse: func(arguments cmd.VideoToolArguments) bool {
		return utils.IsElementExist(arguments, 0)
	},
	Run: func(arguments cmd.VideoToolArguments) {
		input := arguments[0]

		originalName := utils.GetFileName(input)
		outputVideoPath := utils.MergeDir("out", originalName)

		utils.RunFFmpeg(
			flags.INPUT_FILE, input,
			flags.SET_FRAME_RATE, "60",
			flags.DISABLE_AUDIO,
			flags.VIDEO_FILTER, "setpts=PTS/30",
			originalName,
		)

		println("Timelapsed video has been saved at:", outputVideoPath)
	},
}
