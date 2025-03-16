package tools

import (
	"fmt"
	"os"
	"video-utils/cmd"
	"video-utils/utils"
	"video-utils/utils/flags"
)

var MergeVideo *cmd.VideoTool = &cmd.VideoTool{
	Description: "Splits video into smaller chunks",
	Usage: func(programName string) {
		fmt.Print(
			fmt.Sprintf("Usage: %s <...many input video path...>\n\n", programName),
			"Note: you must need at least 2 video paths, well, why you want to merge only 1 video right?\n\n",
			"Example: ",
			fmt.Sprintf("\t%s /path/to/first-video.mp4 /path/to/second-video.mp4\n", programName),
		)
	},
	Parse: func(arguments cmd.VideoToolArguments) bool {
		return !(len(arguments) < 2)
	},
	Run: func(arguments cmd.VideoToolArguments) {
		fileListPath := writeFileList(arguments)
		savedVideoPath := "./out/merged-video-output.mp4"

		utils.RunFFmpeg(
			flags.FORCE_FORMAT, "concat",
			flags.ENABLE_SAFE_MODE, "0",
			flags.INPUT_FILE, fileListPath,
			flags.CODEC_NAME, "copy",
			savedVideoPath,
		)

		os.Remove(fileListPath)
		println("Merged video has been saved at:", savedVideoPath)
	},
}

func writeFileList(inputFiles []string) string {
	content := ""
	for i := 0; i < len(inputFiles); i++ {
		content += fmt.Sprintf("file '%s'\n", inputFiles[i])
	}

	fileListPath := "./out/video-list.txt"
	os.WriteFile(fileListPath, []byte(content), 0644 /*magic number, I don't understand what that is*/)
	return fileListPath
}
