package ytdlp

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadVideo(t *testing.T) {
	originalExecutor := new(defaultOSCommandExecutor)

	t.Run("successful download", func(t *testing.T) {
		defer func() {
			osCommandExecutorProvider = originalExecutor
		}()
		osCommandExecutorProvider = &mockOSCommandExecutor{
			execCommandOutput: "Download successful",
		}

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

func Test_ytDlpPath(t *testing.T) {
	originalOsName := osName
	originalOsOperationsProvider := osOperationsProvider
	originalFilePathOperationsProvider := filePathOperationsProvider

	t.Run("returns path on success", func(t *testing.T) {
		defer func() {
			osName = originalOsName
			osOperationsProvider = originalOsOperationsProvider
			filePathOperationsProvider = originalFilePathOperationsProvider
		}()
		osName = "darwin"
		osOperationsProvider = &mockOsOperationsProvider{}
		filePathOperationsProvider = &mockFilePathOperationsProvider{
			joinOutput: "/tmp/yt-dlp",
		}

		path, err := ytDlpPath()
		require.NoError(t, err)
		require.Equal(t, "/tmp/yt-dlp", path)
	})

	t.Run("returns error on unsupported OS", func(t *testing.T) {
		defer func() {
			osName = originalOsName
		}()
		osName = "linux"

		_, err := ytDlpPath()
		require.Equal(t, "unsupported OS: linux", err.Error())
	})

	t.Run("returns error on MkdirTemp failure", func(t *testing.T) {
		defer func() {
			osName = originalOsName
			osOperationsProvider = originalOsOperationsProvider
		}()
		osName = "darwin"
		osOperationsProvider = &mockOsOperationsProvider{
			errMkdirTemp: errors.New("MkdirTemp failed"),
		}

		_, err := ytDlpPath()
		require.Equal(t, "MkdirTemp failed", err.Error())
	})

	t.Run("returns error on WriteFile failure", func(t *testing.T) {
		defer func() {
			osName = originalOsName
			osOperationsProvider = originalOsOperationsProvider
			filePathOperationsProvider = originalFilePathOperationsProvider
		}()
		osName = "darwin"
		osOperationsProvider = &mockOsOperationsProvider{
			errWriteFile: errors.New("WriteFile failed"),
		}
		filePathOperationsProvider = &mockFilePathOperationsProvider{
			joinOutput: "/tmp/yt-dlp",
		}

		_, err := ytDlpPath()
		require.Equal(t, "WriteFile failed", err.Error())
	})

}

type mockOSCommandExecutor struct {
	execCommandOutput string
	execCommandErr    error
}

func (m *mockOSCommandExecutor) ExecCommand(ctx context.Context, name string, arg ...string) (string, error) {
	return m.execCommandOutput, m.execCommandErr
}

type mockOsOperationsProvider struct {
	errMkdirTemp error
	errWriteFile error
	errRemoveAll error
}

func (m *mockOsOperationsProvider) MkdirTemp(dir string, pattern string) (string, error) {
	return "", m.errMkdirTemp
}

func (m *mockOsOperationsProvider) WriteFile(name string, data []byte, perm os.FileMode) error {
	return m.errWriteFile
}

func (m *mockOsOperationsProvider) RemoveAll(path string) error {
	return m.errRemoveAll
}

type mockFilePathOperationsProvider struct {
	joinOutput string
}

func (m *mockFilePathOperationsProvider) Join(elem ...string) string {
	return m.joinOutput
}
