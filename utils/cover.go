package utils

import (
	"MinimalistTiktok/config"
	"os/exec"
	"strings"
)

func GetCover(videoFileName string) (string, string) {
	// videoFileName: 1.mp4
	videoNameList := strings.Split(videoFileName, ".")
	filePath := config.VideosImagePath + videoFileName
	coverPath := config.VideosImagePath + videoNameList[0] + ".jpeg"
	println(filePath)
	println(coverPath)
	cmd := exec.Command("ffmpeg", "-i", filePath, "-ss", "1", "-f",
		"image2", "-frames:v", "1", coverPath)

	err := cmd.Run()
	if err != nil {
		println(err)
		return "", ""
	}

	println("after cmd")
	return filePath, coverPath
}


