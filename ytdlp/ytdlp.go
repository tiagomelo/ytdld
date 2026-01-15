// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package ytdlp

import (
	"context"
	"fmt"

	"github.com/tiagomelo/ytdld/syscall"
)

// ytDlpToolName is the name of the yt-dlp executable for macOS.
const ytDlpToolName = "bin/yt-dlp_macos"

// osCommandExecutor defines an interface for executing OS commands.
type osCommandExecutor interface {
	ExecCommand(ctx context.Context, name string, arg ...string) (string, error)
}

// defaultOSCommandExecutor is the default implementation of osCommandExecutor.
type defaultOSCommandExecutor struct{}

// ExecCommand executes a command with arguments.
func (d *defaultOSCommandExecutor) ExecCommand(ctx context.Context, name string, arg ...string) (string, error) {
	return syscall.ExecCommand(ctx, name, arg...)
}

// osCommandExecutorProvider is a variable that holds the function
// that executes a command with arguments.
var osCommandExecutorProvider osCommandExecutor = &defaultOSCommandExecutor{}

// DownloadVideo downloads a video from the given URL using yt-dlp.
// Only macOS is supported at the moment.
func DownloadVideo(ctx context.Context, url, outputPath string) (string, error) {
	outputPath = fmt.Sprintf("%s.%%(ext)s", outputPath)
	if _, err := osCommandExecutorProvider.ExecCommand(
		ctx,
		ytDlpToolName,
		"-o",
		outputPath,
		url,
	); err != nil {
		return "", err
	}
	return outputPath, nil
}
