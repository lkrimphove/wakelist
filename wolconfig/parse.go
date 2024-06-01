// Package sshconfig can parse a wol config file into a list of devices.
package wolconfig

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/lkrimphove/wakelist/device"
)

// NamedReader is an io.Reader that also has a name, usually an os.File.
type NamedReader interface {
	io.Reader
	Name() string
}

// PraseFile reads and parses the file in the given path.
func ParseFile(path string) ([]device.Device, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config: %w", err)
	}
	defer f.Close()
	return ParseReader(f)
}

// ParseReader reads and parses the given reader.
func ParseReader(r NamedReader) ([]device.Device, error) {
	scanner := bufio.NewScanner(r)

	var devices []device.Device
	var newDevice *device.Device

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue // skip empty lines
		}

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue // skip comments
		}

		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			if newDevice != nil {
				return nil, fmt.Errorf("invalid line on device %q", newDevice.Name)
			}
			return nil, fmt.Errorf("invalid line: %q", line)
		}

		var key = parts[0]
		var value = parts[1]

		switch strings.ToLower(key) {
		case "device":
			if newDevice != nil {
				if newDevice.MacAddr == "" {
					return nil, fmt.Errorf("missing mac adress on device %q: %q", newDevice.Name, line)
				}

				devices = append(devices, *newDevice)
			}

			if newDevice == nil {
				newDevice = &device.Device{}
			}
			(*newDevice).Name = value
		case "macaddr":
			(*newDevice).MacAddr = value
		case "bcastinterface":
			(*newDevice).BcastInterface = value
		case "bcastip":
			(*newDevice).BcastIp = value
		case "udpport":
			(*newDevice).UDPPort = value
		case "ping":
			(*newDevice).Ping = value
		default:
			if newDevice != nil {
				return nil, fmt.Errorf("invalid line on device %q: %q", newDevice.Name, line)
			}
			return nil, fmt.Errorf("invalid line: %q", line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	if newDevice != nil {
		if newDevice.MacAddr == "" {
			return nil, fmt.Errorf("missing mac adress on device %q", newDevice.Name)
		}

		devices = append(devices, *newDevice)
	}

	return devices, nil
}
