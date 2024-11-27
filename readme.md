# **Windows Port Forwarding App**

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


This Go application provides an intuitive way to manage **port forwarding** and **firewall rules** on Windows, making it easy to expose local application ports within your local network. It automates the otherwise cumbersome process of using `netsh` commands to enable port forwarding, ensuring a smoother experience for developers and system administrators.

---

## **Why This App Was Needed**

When working with applications that expose ports (e.g., Docker containers, web servers, or custom services), it’s common to face challenges when trying to access these services from other devices on the same network. Windows does not natively support dynamic or persistent port forwarding rules through a GUI, requiring manual configuration of `netsh` commands and firewall rules every time.

This app was built to solve the following issues:
1. **Access Exposed Ports from Other Devices**: Enable seamless access to local services (e.g., `localhost:6300`) from other devices in the network (e.g., `192.168.0.230:6300`).
2. **Avoid Manual Configurations**: Simplify repetitive `netsh` and firewall configuration steps.
3. **Dynamic and Persistent Rules**: Allow dynamic creation and management of forwarding rules that persist across reboots.
4. **Centralized Management**: Provide a user-friendly interface to manage port forwarding and firewall rules without relying on multiple tools or scripts.

---

## **Features**

1. **Add Port Forwarding Rules**:
   - Dynamically forward traffic from a specific IP (`listen address`) and port (`listen port`) to another IP (`connect address`) and port (`connect port`).
   - Automatically create corresponding firewall rules to allow traffic for forwarded ports.

2. **Remove Port Forwarding Rules**:
   - Easily clean up forwarding rules and associated firewall rules from the app.

3. **List Existing Rules**:
   - View all active port forwarding and firewall rules in a centralized interface.

4. **Simple UI**:
   - A lightweight, browser-based interface built with Go and `htmx` for easy management.

5. **Hidden Operation**:
   - Can run as a background process or startup application without requiring a visible console window.

---

## **How It Works**

1. **Port Proxy Rules**:
   - Uses the `netsh interface portproxy` command to add, delete, and list port forwarding rules.
2. **Firewall Rules**:
   - Automatically creates or removes firewall rules using `netsh advfirewall` to ensure traffic is allowed for forwarded ports.
3. **Database**:
   - Tracks all rules in a local SQLite database to provide persistence and easy retrieval.

---

## **Installation**

### **Prerequisites**
- **Go 1.20 or higher** (for building the application).
- **Windows 10/11** (or a compatible version with `netsh` available).
- **Admin Privileges**: Required to create port proxy and firewall rules.

### **Clone the Repository**
```bash
git clone https://github.com/pieeee/windows-port-forwarding.git
cd windows-port-forwarding
```

### **Build the Application**
```bash
go build -o port-forwarding.exe
```

### **Run the Application**
```bash
port-forwarding.exe
```

By default, the app runs on `http://localhost:2233`. Open this URL in your browser to manage port forwarding rules.

---

## **Usage**

### **Add a Rule**
1. Open the app at `http://localhost:2233`.
2. Fill in the required fields:
   - **Rule Name**: A descriptive name for the rule.
   - **Listen Address**: The IP address to listen for incoming traffic (default: local machine's IP).
   - **Listen Port**: The port to forward traffic from.
   - **Connect Address**: The destination IP address (default: `127.0.0.1`).
   - **Connect Port**: The port to forward traffic to.
3. Click **Add Rule** to create the rule.

### **List Rules**
- View all active port proxy and firewall rules in the table on the main page.

### **Remove a Rule**
- Click the **Remove** button next to a rule to delete both the port proxy and firewall rule.

---

## **Startup Option**

To automatically start the app at system boot, add the executable to the **Startup Folder**:
1. Press `Win + R` and type:
   ```plaintext
   shell:startup
   ```
2. Copy the `port-forwarding.exe` file to this folder.

---

## **Technical Details**

- **Backend**: Go with `netsh` commands for port proxy and firewall management.
- **Database**: SQLite for persistent rule storage.
- **Frontend**: Minimal HTML UI with [htmx](https://htmx.org/) for dynamic updates.
- **Port Forwarding**: Uses `netsh interface portproxy` for traffic redirection.
- **Firewall Rules**: Configures Windows Firewall automatically for each port forwarding rule.

---

## **Limitations**

1. **Local Network Only**: This app is designed for managing ports within the local network (LAN).
2. **Admin Rights Required**: Creating port proxy and firewall rules requires administrative privileges.

---

## **Future Improvements**

1. **System Tray Integration**: Add a tray icon for quick access and notifications.
2. **TLS Support**: Secure the web interface with HTTPS.
3. **Cross-Platform Support**: Extend functionality to Linux and macOS.
4. **Advanced Firewall Rules**: Add options for IP ranges and protocols.

---

## **Contributing**

Contributions are welcome! Feel free to open issues or submit pull requests for improvements or new features.

---

## **License**

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

