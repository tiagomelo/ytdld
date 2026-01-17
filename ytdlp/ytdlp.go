// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package ytdlp

import (
	"context"
	_ "embed"
	"fmt"
	"runtime"
	"strings"

	"github.com/tiagomelo/ytdld/syscall"
	"github.com/tiagomelo/ytdld/ytdlp/fp"
	"github.com/tiagomelo/ytdld/ytdlp/os"
)

//go:embed yt-dlp_macos
var ytDlpMacOS []byte // ytDlpMacOS is the embedded yt-dlp executable for macOS.

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

// osOperationsProvider is a variable that holds the OS operations provider.
var osOperationsProvider os.OSOperations = &os.OSOperationsProvider{}

// osName holds the name of the operating system.
var osName = runtime.GOOS

// osCommandExecutorProvider is a variable that holds the function
// that executes a command with arguments.
var osCommandExecutorProvider osCommandExecutor = &defaultOSCommandExecutor{}

// filePathOperationsProvider is a variable that holds the file path operations provider.
var filePathOperationsProvider fp.FilePathOperations = fp.FilePathOperationsProvider{}

// DownloadVideo downloads a video from the given URL using yt-dlp.
// Only macOS is supported at the moment.
func DownloadVideo(ctx context.Context, url, outputPath string) (string, error) {
	tool, err := ytDlpPath()
	if err != nil {
		return "", err
	}

	template := fmt.Sprintf("%s.%%(ext)s", outputPath)
	out, err := osCommandExecutorProvider.ExecCommand(
		ctx,
		tool,
		"-o", template,
		"--print", "after_move:filepath",
		"--no-progress",
		url,
	)
	if err != nil {
		return "", err
	}

	// yt-dlp prints the path with a trailing newline
	finalPath := strings.TrimSpace(out)
	if finalPath == "" {
		return "", fmt.Errorf("yt-dlp did not return final file path")
	}

	return finalPath, nil
}

// ytDlpPath returns the path to the yt-dlp executable for macOS.
func ytDlpPath() (string, error) {
	const darwinOS = "darwin"

	if osName != darwinOS {
		return "", fmt.Errorf("unsupported OS: %s", osName)
	}
	dir, err := osOperationsProvider.MkdirTemp("", "ytdld-*")
	if err != nil {
		return "", err
	}
	p := filePathOperationsProvider.Join(dir, "yt-dlp_macos")
	if err := osOperationsProvider.WriteFile(p, ytDlpMacOS, 0o755); err != nil {
		_ = osOperationsProvider.RemoveAll(dir)
		return "", err
	}
	return p, nil
}
