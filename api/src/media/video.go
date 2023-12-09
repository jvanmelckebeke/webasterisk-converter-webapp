package media

import (
	"fmt"
	"github.com/tidbyt/go-libwebp/webp"
	"jvanmelckebeke/anyconverter-api/tools"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func ffmpegProcess(inputFilePath string, outputFilePath string, args ...string) error {

	// show the command that is being executed
	fmt.Println("ffmpeg", args)

	cmd := exec.Command("ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	go func() {
		inputSize := tools.GetHumanFileSize(inputFilePath)

		for {
			if err := cmd.Wait(); err != nil {
				fmt.Println(err)
				return
			}

			outputSize := tools.GetHumanFileSize(outputFilePath)
			fmt.Printf("Converted %s / %s\n", outputSize, inputSize)
		}
	}()

	return cmd.Wait()
}

func webmToMp4(inputPath string) (string, error) {
	outputFilePath := tools.PrepareOutputFile(inputPath)

	args := []string{
		"-y", // overwrite output file if it exists
		"-v", "error",
		"-i", inputPath,
		"-map", "V:0?",
		"-map", "0:a?",
		"-c:v", "libx264",
		"-movflags", "+faststart",
		"-preset", "veryslow",
		"-pix_fmt", "yuv420p",
		"-vf", "pad=ceil(iw/2)*2:ceil(ih/2)*2",
		outputFilePath,
	}

	if err := ffmpegProcess(inputPath, outputFilePath, args...); err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("ffmpeg: error converting to mp4")
	}

	return outputFilePath, nil

}

func gifToMp4(inputPath string) (string, error) {
	// conversion steps for gif to mp4 are the same as webm to mp4
	return webmToMp4(inputPath)
}

func webpToGif(inputPath string) (string, error) {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Printf("Error reading file: %s\n", err)
		return "", fmt.Errorf("error reading file")
	}

	dec, err := webp.NewAnimationDecoder(data)
	if err != nil {
		log.Printf("Error creating animation decoder: %s\n", err)
		return "", fmt.Errorf("error creating animation decoder")
	}
	defer dec.Close()

	anim, err := dec.Decode()
	if err != nil {
		log.Printf("Error decoding animation: %s\n", err)
		return "", fmt.Errorf("error decoding animation")
	}

	timestamps := anim.Timestamp
	images := anim.Image

	frameRate := (timestamps[len(timestamps)-1] - timestamps[0]) / (len(timestamps))

	outputFilePath := tools.PrepareOutputFile(inputPath)
	frameDir := tools.PrepareFrameDirectory(inputPath)

	// create gif from frames
	for i, img := range images {
		framePath := fmt.Sprintf("%s/frame_%d.png", outputFilePath, i)

		if err := imageToPng(framePath, img); err != nil {
			return "", err
		}
	}

	// create gif from frames
	args := []string{
		"-y", // overwrite output file if it exists
		"-v", "error",
		"-framerate", fmt.Sprintf("%d", frameRate),
		"-pattern_type", "glob",
		"-i", fmt.Sprintf("%s/frame_*.png", frameDir),
		"-loop", "0",
		"-c:v", "libx264",
		"-preset", "veryslow",
		"-pix_fmt", "yuv420p",
		"-vf", "pad=ceil(iw/2)*2:ceil(ih/2)*2",
		outputFilePath,
	}

	return outputFilePath, ffmpegProcess(inputPath, outputFilePath, args...)
}

func animatedWebpToMp4(inputPath string) (string, error) {
	// strategy: convert animated webp to gif, then gif to mp4

	gifPath, err := webpToGif(inputPath)

	if err != nil {
		return "", err
	}

	return gifToMp4(gifPath)
}
func ToMp4(inputPath string) (string, error) {
	if inputPath == "" {
		return "", fmt.Errorf("empty inputPath provided")
	}

	extenstion := filepath.Ext(inputPath)

	switch extenstion {
	case ".webm":
		return webmToMp4(inputPath)
	case ".gif":
		return gifToMp4(inputPath)
	case ".webp":
		return animatedWebpToMp4(inputPath)
	}

	return "", fmt.Errorf("unsupported file type")
}
