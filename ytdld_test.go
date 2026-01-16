// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package ytdld

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadVideo(t *testing.T) {
	originalDownloader := downloader

	t.Run("successful download", func(t *testing.T) {
		defer func() { downloader = originalDownloader }()

		mockPath := "/path/to/downloaded/video.mp4"
		downloader = &mockVideoDownloader{
			fileOutputPath: mockPath,
			err:            nil,
		}

		output, err := DownloadVideo(context.TODO(), "http://example.com/video", "/output/path")
		require.NoError(t, err)
		require.Equal(t, mockPath, output)
	})

	t.Run("download error", func(t *testing.T) {
		defer func() { downloader = originalDownloader }()

		mockErr := "download failed"
		downloader = &mockVideoDownloader{
			fileOutputPath: "",
			err:            errors.New(mockErr),
		}

		output, err := DownloadVideo(context.TODO(), "http://example.com/video", "/output/path")
		require.Error(t, err)
		require.Contains(t, err.Error(), mockErr)
		require.Equal(t, "", output)
	})

}

type mockVideoDownloader struct {
	fileOutputPath string
	err            error
}

func (m *mockVideoDownloader) DownloadVideo(ctx context.Context, url, outputPath string) (string, error) {
	return m.fileOutputPath, m.err
}
