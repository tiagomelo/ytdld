// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/tiagomelo/ytdld/ytdlp"
)

func run(args []string, log *slog.Logger) error {
	if len(args) < 2 {
		return errors.New("usage: ytdld <video-url> [output-file]")
	}

	ctx := context.Background()
	defer log.InfoContext(ctx, "completed")

	videoURL := strings.TrimSpace(args[0])
	if videoURL == "" {
		return errors.New("video url is required")
	}

	outputFile := strings.TrimSpace(args[1])
	if outputFile == "" {
		return errors.New("output file is required")
	}

	log.InfoContext(ctx, "downloading...", "url", videoURL, "output", outputFile)

	outputFilePath, err := ytdlp.DownloadVideo(ctx, videoURL, outputFile)
	if err != nil {
		return fmt.Errorf("failed to download video: %w", err)
	}

	log.InfoContext(ctx, "video downloaded successfully", "output", outputFilePath)

	return nil
}

func main() {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)

	if err := run(os.Args[1:], log); err != nil {
		log.Error("error", "err", err)
		os.Exit(1)
	}
}
