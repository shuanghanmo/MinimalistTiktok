package utils

import (
	"MinimalistTiktok/config"
	"fmt"
	"os/exec"
	"strings"
)

func GetCover(videoFileName string) (string, string) {
	// videoFileName: 1.mp4
	videoNameList := strings.Split(videoFileName, ".")
	filePath := config.VideosImagePath + videoFileName
	coverPath := config.VideosImagePath + videoNameList[0] + ".jpeg"
	cmd := exec.Command("ffmpeg", "-i", filePath, "-ss", "1", "-f",
		"image2", "-frames:v", "1", coverPath)

	err := cmd.Run()
	if err != nil {
		println(err)
		return "", ""
	}
	filePath = fmt.Sprintf("http://%s:%s/%s",
		config.ServerIP, config.ServerPort, filePath[2:])
	coverPath = fmt.Sprintf("http://%s:%s/%s",
		config.ServerIP, config.ServerPort, coverPath[2:])
	return filePath, coverPath
}
