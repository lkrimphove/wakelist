package wolconfig

import (
	_ "embed"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/good
var goodConfig []byte

//go:embed testdata/invalid_device
var invalidNodeConfig []byte

func TestParseFile(t *testing.T) {
	t.Run("GoodConfig", func(t *testing.T) {
		devices, err := ParseFile("testdata/good.conf")
		require.NoError(t, err)
		require.Len(t, devices, 1)
	})

	t.Run("InvalidNodeConfig", func(t *testing.T) {
		_, err := ParseFile("testdata/invalid_node.conf")
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid line on device")
	})

	t.Run("NonExistentConfig", func(t *testing.T) {
		_, err := ParseFile("testdata/nope.conf")
		require.ErrorIs(t, err, os.ErrNotExist)
	})
}
