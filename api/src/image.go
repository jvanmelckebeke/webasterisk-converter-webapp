package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/tidbyt/go-libwebp/webp"
)

func imageToJpg(inputPath string, img image.Image) (string, error) {
	outputPath := convertPath(inputPath, "jpg")

	output, err := os.Create(outputPath)
	if err != nil {
		log.Printf("Error creating output file: %s\n", err)
		return "", fmt.Errorf("error creating output file")
	}
	defer output.Close()

	if err := jpeg.Encode(output, img, nil); err != nil {
		log.Printf("Error encoding jpeg file: %s\n", err)
		return "", fmt.Errorf("error encoding jpeg file")
	}

	return outputPath, nil
}

func animatedWebpToJpg(inputPath string) (string, error) {
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

	return imageToJpg(inputPath, anim.Image[0])
}

func webpToJpg(inputPath string) (string, error) {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Printf("Error reading file: %s\n", err)
		return "", fmt.Errorf("error reading file")
	}

	options := &webp.DecoderOptions{}

	img, err := webp.DecodeRGBA(data, options)
	if err != nil {
		log.Printf("Error decoding webp file: %s\n", err)
		log.Printf("trying as animation")

		return animatedWebpToJpg(inputPath)

	}
	log.Printf("Decoded webp file")

	return imageToJpg(inputPath, img)
}

func gifToJpg(inputPath string) (string, error) {
	return "", fmt.Errorf("gifToJpg not implemented")
}

func mediaToJpg(inputPath string) (string, error) {
	if inputPath == "" {
		return "", fmt.Errorf("inputPath is empty")
	}

	extension := filepath.Ext(inputPath)

	switch extension {
	case ".webp":
		return webpToJpg(inputPath)
	case ".jpg":
	case ".jpeg":
		return inputPath, nil
	case ".gif":
		return gifToJpg(inputPath)
	}

	return "", fmt.Errorf("unsupported file type")
}
