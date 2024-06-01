package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lkrimphove/wakelist/device"
	"github.com/lkrimphove/wakelist/wol"
	"github.com/lkrimphove/wakelist/wolconfig"
)

const defaultConfigPath = "/.wol/config"

func getConfigPath() string {
	// Check if the environment variable is set
	path := os.Getenv("WOL_CONFIG_PATH")
	if path != "" {
		return path
	}

	// Get the current user's home directory
	usr, _ := user.Current()
	if usr != nil {
		// Fallback to the default path if unable to get the user's home directory
		return filepath.Join(usr.HomeDir, defaultConfigPath)
	}

	return ""
}

type model struct {
	devices []device.Device
	cursor  int
}

func initialModel(devices []device.Device) model {
	return model{
		devices: devices,
		cursor:  0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.devices)-1 {
				m.cursor++
			}
		case "enter":
			wol.Wake(m.devices[m.cursor])
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Device List:\n\n"

	for i, device := range m.devices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, device.Name)
	}

	s += "\nPress q to quit.\n"
	return s
}

func main() {
	devicePath := getConfigPath()
	devices, err := wolconfig.ParseFile(devicePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading devices configuration: %v\n", err)
		os.Exit(1)
	}

	p := tea.NewProgram(initialModel(devices))
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
