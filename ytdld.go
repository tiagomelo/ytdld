// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package ytdld

import (
	"context"

	"github.com/tiagomelo/ytdld/ytdlp"
)

// videoDownloader defines an interface for downloading videos.
type videoDownloader interface {
	DownloadVideo(ctx context.Context, url, outputPath string) (string, error)
}

// ytdlpDownloader is the default implementation of videoDownloader using yt-dlp.
type ytdlpDownloader struct{}

// SetVideoDownloader allows setting a custom video downloader implementation.
var downloader videoDownloader = &ytdlpDownloader{}

// DownloadVideo downloads a video from the given URL using yt-dlp.
func (d *ytdlpDownloader) DownloadVideo(ctx context.Context, url, outputPath string) (string, error) {
	return ytdlp.DownloadVideo(ctx, url, outputPath)
}

// DownloadVideo downloads a video from the given URL and saves it to the specified output path.
// Only macOS is supported at the moment.
func DownloadVideo(ctx context.Context, url, outputPath string) (string, error) {
	return downloader.DownloadVideo(ctx, url, outputPath)
}
