// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package syscall

import (
	"context"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tiagomelo/ytdld/syscall/fs"
)

func TestExecCommand(t *testing.T) {
	originalFSProvider := fsProvider
	t.Run("executes command successfully", func(t *testing.T) {
		defer func() {
			fsProvider = originalFSProvider
		}()
		mockCmdOutput := fs.ExecCmdOutput{
			Stdout: "command output",
			Stderr: "",
		}
		fsProvider = &mockOSFileSystem{
			execCmdOutput: mockCmdOutput,
		}
		output, err := ExecCommand(context.TODO(), "echo", "hello")
		require.NotNil(t, output)
		require.NoError(t, err)
	})

	t.Run("error during command execution", func(t *testing.T) {
		defer func() {
			fsProvider = originalFSProvider
		}()
		mockCmdOutput := fs.ExecCmdOutput{
			Stdout: "",
			Stderr: "",
		}
		fsProvider = &mockOSFileSystem{
			execCmdOutput:     mockCmdOutput,
			errCommandContext: exec.ErrNotFound,
		}
		output, err := ExecCommand(context.TODO(), "some-nonexistent-command", "arg1")
		require.Empty(t, output)
		require.Equal(t, `error when executing command [some-nonexistent-command] with args [arg1]: stdout: [] stderr: []: executable file not found in $PATH`, err.Error())
	})
}

type mockOSFileSystem struct {
	execCmdOutput     fs.ExecCmdOutput
	errCommandContext error
}

func (m *mockOSFileSystem) CommandContext(ctx context.Context, name string, arg ...string) (fs.ExecCmdOutput, error) {
	return m.execCmdOutput, m.errCommandContext
}
