// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package ytdlp

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadVideo(t *testing.T) {
	originalExecutor := new(defaultOSCommandExecutor)

	t.Run("successful download", func(t *testing.T) {
		defer func() {
			osCommandExecutorProvider = originalExecutor
		}()
		osCommandExecutorProvider = &mockOSCommandExecutor{}

		_, err := DownloadVideo(context.Background(), "https://example.com/video", "/path/to/output")
		require.NoError(t, err)
	})

	t.Run("failed download", func(t *testing.T) {
		defer func() {
			osCommandExecutorProvider = originalExecutor
		}()
		osCommandExecutorProvider = &mockOSCommandExecutor{
			execCommandErr: errors.New("command failed"),
		}

		_, err := DownloadVideo(context.Background(), "https://example.com/video", "/path/to/output")
		require.Equal(t, "command failed", err.Error())
	})
}

type mockOSCommandExecutor struct {
	execCommandErr error
}

func (m *mockOSCommandExecutor) ExecCommand(ctx context.Context, name string, arg ...string) (string, error) {
	return "", m.execCommandErr
}
