package ffmpeg

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"os/exec"

	_ "github.com/disintegration/imaging"
	"github.com/nfnt/resize"
)

func transcodeVideo(inputFile, outputFile string) error {
	// Check if FFmpeg is installed
	_, err := exec.LookPath("ffmpeg")
	if err != nil {
		return fmt.Errorf("FFmpeg not found. Please install FFmpeg and ensure it's in your PATH")
	}

	cmd := exec.Command("ffmpeg",
		"-i", inputFile, // Input file
		"-c:v", "libx264", // Video codec (H.264)
		"-crf", "23", // Constant Rate Factor for quality (adjust as needed)
		"-c:a", "aac", // Audio codec (AAC)
		"-strict", "experimental",
		"-y",
		outputFile,
	)

	// Run FFmpeg command
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("FFmpeg command failed: %v", err)
	}

	return nil
}

func transcodeImg(inputPath, outputPath string) error {
	maxWidth := 800

	// Open the input image file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Decode the input image
	inputImage, _, err := image.Decode(inputFile)
	if err != nil {
		return err
	}

	// Resize the image while maintaining the aspect ratio
	resizedImage := resize.Resize(uint(maxWidth), 0, inputImage, resize.Lanczos3)

	// Compress and save the image as JPEG
	err = saveAsJPEG(resizedImage, outputPath, 90) // 90 is the JPEG quality, adjust as needed
	if err != nil {
		return err
	}

	return nil
}

func saveAsJPEG(img image.Image, outputPath string, quality int) error {
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Encode and save the image as JPEG with the specified quality
	err = jpeg.Encode(outputFile, img, &jpeg.Options{Quality: quality})
	if err != nil {
		return err
	}

	return nil
}
