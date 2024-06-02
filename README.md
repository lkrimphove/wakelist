# Wakelist

Wakelist is a tui application designed to easily start devices using Wake-on-LAN (WoL).
This project is my first foray into Go programming, and is mainly for educational purposes.
I discovered [Wishlist](https://github.com/charmbracelet/wishlist) while searching for inspiration, and Wakelist is hugely inspired by it.

## Goal

The primary goal of Wakelist is to provide a user-friendly terminal application that allows users to quickly and easily send Wake-on-LAN packets to start their devices.
The interface aims to be intuitive, making it accessible even to users who may not be familiar with the intricacies of WoL.

## Motivation

As a newcomer to the Go programming language, I wanted to try and build something that would not only help me learn the language but also result in a useful tool.
While exploring potential projects, I came across [Wishlist](https://github.com/charmbracelet/wishlist), a TUI application for managing SSH endpoints.
Its clean design and practical functionality inspired me to create something similar for Wake-on-LAN.
As I'm still learning and wanted to get something working in a short time span I relied on [go-wol](https://github.com/sabhiram/go-wol) to do the heavy lifting and sending the magic packets.

## Features

- **Device Management**: Easily add and manage devices that support Wake-on-LAN.
- **Simple Interface**: Navigate through your devices and send WoL packets with minimal effort.

## Work in Progress

Wakelist is still a work in progress, and there are several features and improvements planned:

- **Feedback after Selection**: Provide clear feedback to the user after a device is selected and a WoL packet is sent.
- **Visual Enhancements**: Make the interface more visually appealing, drawing more inspiration from Wishlist's clean and modern look.
- **Improved Configuration Parsing**: Enhance the configuration file parsing to support more complex setups and error handling.

## How to Use

1. **Clone the Repository**:
    ```sh
    git clone https://github.com/yourusername/wakelist.git
    cd wakelist
    ```

2. **Build the Application**:
    ```sh
    go build -o wakelist
    ```

3. **Run the Application**:
    ```sh
    ./wakelist
    ```

4. **Configuration**:
    - The configuration file works similarly to an SSH config file.
    - By default, the app looks for a configuration file in `.wol/config` in the user's home directory.
    - You can change the default location of the configuration file by setting the `WOL_CONFIG_PATH` environment variable.
    - **Example Configuration File**:
      ```plaintext
      # Sample wakelist configuration

      Host my-desktop
        Mac AA:BB:CC:DD:EE:FF

      Host media-server
        Mac 11:22:33:44:55:66

      Host office-pc
        Mac 77:88:99:AA:BB:CC
      ```

## Contributing

Contributions are welcome! If you have suggestions for new features or improvements, please open an issue or submit a pull request. Feedback and collaboration are greatly appreciated as I continue to develop and refine Wakelist.

## License

Wakelist is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

Thank you for checking out Wakelist!
If you like this project please consider checking out [Wishlist](https://github.com/charmbracelet/wishlist) and [go-wol](https://github.com/sabhiram/go-wol).

If you like my work you can also buy me a coffee. Thank you!

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/lkrimphove)
