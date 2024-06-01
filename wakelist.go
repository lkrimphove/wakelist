package wakelist

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Device struct {
	name string
	mac  string
}

type model struct {
	devices []Device
	cursor  int
}

func initialModel(devices []Device) model {
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
			sendWOLPacket(m.devices[m.cursor].mac)
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
		s += fmt.Sprintf("%s %s\n", cursor, device.name)
	}

	s += "\nPress q to quit.\n"
	return s
}

func sendWOLPacket(macAddress string) {
	fmt.Fprintf(os.Stderr, "Error calling %v", macAddress)
	// macAddr, err := wol.ParseMAC(macAddress)
	// if err != nil {
	//     fmt.Println("Invalid MAC address:", err)
	//     return
	// }

	// if err := wol.SendMagicPacket(macAddr); err != nil {
	//     fmt.Println("Failed to send WOL packet:", err)
	// } else {
	//     fmt.Println("WOL packet sent to", macAddress)
	// }
}

func parseConfig(filename string) ([]Device, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var devices []Device
	scanner := bufio.NewScanner(file)
	var currentDevice Device

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Host") {
			if currentDevice.name != "" {
				devices = append(devices, currentDevice)
			}
			currentDevice = Device{}
			currentDevice.name = strings.TrimSpace(strings.TrimPrefix(line, "Host"))
		} else if strings.HasPrefix(line, "Hostname") {
			currentDevice.mac = strings.TrimSpace(strings.TrimPrefix(line, "Hostname"))
		}
	}

	if currentDevice.name != "" {
		devices = append(devices, currentDevice)
	}

	return devices, scanner.Err()
}

func main() {
	devices, err := parseConfig("devices.conf")
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
