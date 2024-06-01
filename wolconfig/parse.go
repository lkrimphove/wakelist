// Package sshconfig can parse a wol config file into a list of devices.
package wolconfig

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/lkrimphove/wakelist"
)

// NamedReader is an io.Reader that also has a name, usually an os.File.
type NamedReader interface {
	io.Reader
	Name() string
}

// PraseFile reads and parses the file in the given path.
func ParseFile(path string) ([]*wakelist.Device, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config: %w", err)
	}
	defer f.Close() //nolint:errcheck
	return ParseReader(f)
}

// ParseReader reads and parses the given reader.
func ParseReader(r NamedReader) ([]*wakelist.Device, error) {
	scanner := bufio.NewScanner(r)

	var devices []wakelist.Devices
	var newDevice wakelist.Device

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue // skip empty lines
		}

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue // skip comments
		}

		parts := strings.SplitN(line, " ", 2) //nolint:gomnd
		if len(parts) != 2 {                  //nolint:gomnd
			if newDevice != nil {
				return nil, fmt.Errorf("invalid line on device %q: %q", newDevice.Name, line)
			}
			return nil, fmt.Errorf("invalid line: %q", line)
		}

		var key = parts[0]
		var value = parts[1]

		switch strings.ToLower(key) {
		case "host":
			if newDevice != nil {
				devices = append(devices, newDevice)
			}

			newDevice = wakelist.Device
			newDevice.Name = value
		case "mac":
			newDevice.Mac = value
		case "ping":
			newDevice.Ping = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	return devices, nil
}
