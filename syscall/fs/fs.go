// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package fs

import (
	"context"
	"os/exec"
)

// ExecCmd interface abstracts the command execution.
type ExecCmd interface {
	// CombinedOutput runs the command and returns its
	// combined standard output and standard error.
	CombinedOutput() ([]byte, error)
}

// FileSystem interface abstracts the file system operations.
type FileSystem interface {
	// CommandContext creates a new execCmd for the given command.
	CommandContext(ctx context.Context, name string, arg ...string) ExecCmd
}

// OSFileSystem struct implements the fileSystem interface using
// the standard library's os package. This is the real implementation
// that interacts with the actual file system.
type OSFileSystem struct{}

func (OSFileSystem) CommandContext(ctx context.Context, name string, arg ...string) ExecCmd {
	return exec.CommandContext(ctx, name, arg...)
}
