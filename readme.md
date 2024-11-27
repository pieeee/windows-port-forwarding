### **Port Forwarding App**

This Go application is designed to simplify and automate the management of **port forwarding** and **firewall rules** on a Windows system. It provides a lightweight, user-friendly interface for adding, removing, and listing port proxy rules along with the necessary firewall configurations, making it ideal for developers or system administrators managing local or network-bound applications.

---

### **Why This App Was Built**

Managing port forwarding and firewall rules on Windows can be cumbersome, especially when:
- Applications require specific port forwarding rules (e.g., Docker services, web servers).
- Manual `netsh` commands for port proxy and firewall management are error-prone.
- Existing solutions like Task Scheduler or startup scripts don't provide the flexibility of a centralized tool for dynamic port management.

This app was created to:
1. **Automate Port Proxy and Firewall Rules**: Streamline the process of forwarding ports and ensuring proper access permissions.
2. **Reduce Repetition**: Avoid repeatedly entering commands for frequently used ports.
3. **Enable Dynamic Rules**: Allow the creation of forwarding rules for any port dynamically, without hardcoding.
4. **Provide a Persistent UI**: Offer a lightweight interface that tracks and manages rules, even after reboots.
5. **Run Seamlessly**: Operate silently in the background, making it unobtrusive during usage.

---

### **Features**

- **Add Port Forwarding Rules**: Dynamically forward traffic from a specific `listen address` and `listen port` to a `connect address` and `connect port`.
- **Firewall Management**: Automatically create corresponding firewall rules to allow traffic for forwarded ports.
- **List Existing Rules**: Display all current port proxy and firewall rules in a simple UI.
- **Remove Rules**: Clean up port proxy and firewall rules with a single action.
- **Run Without Console**: Operates as a hidden application to avoid cluttering the user's workflow.

---

### **How It Works**

1. **Port Proxy Rules**: Uses the `netsh interface portproxy` command to add, delete, and list port forwarding rules.
2. **Firewall Rules**: Leverages `netsh advfirewall` to create or remove corresponding firewall rules for open ports.
3. **Persistent Storage**: Tracks rule names and configurations using a lightweight local SQLite database.
4. **Startup Option**: Can be configured to run automatically at system startup.

---

### **Prerequisites**

- **Go 1.20 or higher** (for building the application).
- **Windows 10/11** (or a compatible version with `netsh` available).
- **Admin Privileges**: Required to create port proxy and firewall rules.

---

### **Installation**

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-username/port-forwarding-app.git
   cd port-forwarding-app
   ```

2. **Build the Executable**:
   ```bash
   go build -o port-forwarding.exe
   ```

3. **Run the Application**:
   ```bash
   port-forwarding.exe
   ```

4. **(Optional) Add to Startup**:
   - Use the **Startup Folder** or create a scheduled task for automatic startup.

---

### **Usage**

#### **Add a Rule**
1. Open the application in your browser (e.g., `http://localhost:2233`).
2. Enter:
   - **Rule Name**: A descriptive name for the rule.
   - **Listen Address**: The IP address to listen for incoming traffic (default: local machine's router IP).
   - **Listen Port**: The port to forward traffic from.
   - **Connect Address**: The destination IP address (default: `127.0.0.1`).
   - **Connect Port**: The port to forward traffic to.
3. Click **Add Rule**.

#### **List Rules**
- View all current port proxy and firewall rules in the table on the main page.

#### **Remove a Rule**
- Click the **Remove** button next to a rule to delete both the port proxy and firewall rule.

---

### **Technical Details**

- **Backend**: Go with `netsh` commands for port proxy and firewall management.
- **Database**: SQLite for persistent rule storage.
- **Frontend**: Minimal HTML UI with [htmx](https://htmx.org/) for dynamic updates.
- **Logging**: Logs application events to a file for debugging.

---

### **Why This App Is Useful**

- **For Developers**: Simplifies managing Docker and other local development tools requiring specific port forwarding.
- **For Network Admins**: Provides a centralized way to manage rules across various ports.
- **For Power Users**: Avoids repetitive tasks like manually entering `netsh` commands.
- **For Automation**: Can be run as a background service or at startup to ensure rules are always active.

---

### **Future Improvements**

- **TLS Support**: Secure the web UI with HTTPS.
- **System Tray Integration**: Add a system tray icon for easier access and status updates.
- **Cross-Platform Support**: Extend functionality to Linux and macOS.
- **Advanced Firewall Rules**: Allow specifying IP ranges or protocols for advanced firewall configurations.

---

### **Credits**

- Built with ❤️ using Go and `netsh`.
- SQLite database for persistent storage.
- `htmx` for dynamic frontend updates.

---
