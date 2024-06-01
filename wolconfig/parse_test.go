package wolconfig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFile(t *testing.T) {
	t.Run("GoodConfig1", func(t *testing.T) {
		devices, err := ParseFile("testdata/good1")
		require.NoError(t, err)
		require.Len(t, devices, 1)
	})

	t.Run("GoodConfig2", func(t *testing.T) {
		devices, err := ParseFile("testdata/good2")
		require.NoError(t, err)
		require.Len(t, devices, 2)
	})

	t.Run("EmptyConfig", func(t *testing.T) {
		devices, err := ParseFile("testdata/empty")
		require.NoError(t, err)
		require.Len(t, devices, 0)
	})

	t.Run("BadConfig", func(t *testing.T) {
		_, err := ParseFile("testdata/bad")
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid line:")
	})

	t.Run("InvalidDeviceConfig", func(t *testing.T) {
		_, err := ParseFile("testdata/invalid_device")
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid line on device")
	})

	t.Run("NonExistentConfig", func(t *testing.T) {
		_, err := ParseFile("testdata/nope")
		require.ErrorIs(t, err, os.ErrNotExist)
	})
}
