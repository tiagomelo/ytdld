// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package fs

import (
	"bytes"
	"context"
	"os/exec"
)

// ExecCmdOutput struct holds the output of an executed command.
type ExecCmdOutput struct {
	Stdout string // Stdout is the standard output of the command.
	Stderr string // Stderr is the standard error output of the command.
}

// FileSystem interface abstracts the file system operations.
type FileSystem interface {
	// CommandContext creates a new execCmd for the given command.
	CommandContext(ctx context.Context, name string, arg ...string) (ExecCmdOutput, error)
}

// OSFileSystem struct implements the fileSystem interface using
// the standard library's os package. This is the real implementation
// that interacts with the actual file system.
type OSFileSystem struct{}

// CommandContext creates a new execCmd for the given command.
func (OSFileSystem) CommandContext(ctx context.Context, name string, arg ...string) (ExecCmdOutput, error) {
	c := exec.CommandContext(ctx, name, arg...)

	var outBuf, errBuf bytes.Buffer
	c.Stdout = &outBuf
	c.Stderr = &errBuf

	err := c.Run()

	return ExecCmdOutput{
		Stdout: outBuf.String(),
		Stderr: errBuf.String(),
	}, err
}
