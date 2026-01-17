// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package syscall

import (
	"context"

	"github.com/pkg/errors"
	"github.com/tiagomelo/ytdld/syscall/fs"
)

// fsProvider is a variable that holds the file system implementation.
var fsProvider fs.FileSystem = fs.OSFileSystem{}

// ExecCommand executes a command with arguments.
func ExecCommand(ctx context.Context, cmd string, args ...string) (string, error) {
	output, err := fsProvider.CommandContext(ctx, cmd, args...)
	if err != nil {
		return "", errors.WithMessagef(err, "error when executing command [%s] with args %v: stdout: [%v] stderr: [%v]", cmd, args, output.Stdout, output.Stderr)
	}
	return output.Stdout, nil
}
