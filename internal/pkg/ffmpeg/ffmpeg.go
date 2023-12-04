package ffmpeg

import (
	"abdullayev13/timeup/internal/config"
	"sync/atomic"
	"time"
)

var funcs = make(chan func(), 10_000)

func TranscodeVideo(inputPath, outputPath string, callback func(error)) {
	Do(func() {
		err := transcodeVideo(inputPath, outputPath)
		callback(err)
	})
}

func TranscodeImg(inputPath, outputPath string, callback func(error)) {
	Do(func() {
		err := transcodeImg(inputPath, outputPath)
		callback(err)
	})
}

func Do(fun func()) {
	funcs <- fun
}

func Logic() {
	var countRunningFuncs atomic.Int32
	for fun := range funcs {
		for int(countRunningFuncs.Load()) > config.FfmpegRunLimit {
			time.Sleep(time.Second)
		}

		go func() {
			countRunningFuncs.Add(1)
			fun()
			countRunningFuncs.Add(-1)
		}()
	}
}

func init() {
	go Logic()
}
