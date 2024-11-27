package main

import (
	"database/sql"
	"html/template"
	"net"
	"net/http"
	"os/exec"
	"strings"

	_ "modernc.org/sqlite"
)

type PortProxy struct {
	RuleName       string
	ListenAddress  string
	ListenPort     string
	ConnectAddress string
	ConnectPort    string
}

var templates = template.Must(template.ParseGlob("templates/*.html"))
var db *sql.DB

func main() {
	// Initialize the database
	initDB()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/remove", removeHandler)
	http.HandleFunc("/list", listHandler)

	http.ListenAndServe(":2233", nil)
}

// Initialize SQLite Database
func initDB() {
	var err error
	// Use "sqlite" instead of "sqlite3" for the driver name
	db, err = sql.Open("sqlite", "./rules.db")
	if err != nil {
		panic(err)
	}

	// Create table if not exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS rules (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			rule_name TEXT NOT NULL,
			listen_address TEXT NOT NULL,
			listen_port TEXT NOT NULL,
			connect_address TEXT NOT NULL,
			connect_port TEXT NOT NULL
		)
	`)
	if err != nil {
		panic(err)
	}
}

// Render home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		RouterIP string
	}{
		RouterIP: getRouterIP(),
	}
	if err := templates.ExecuteTemplate(w, "home.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Add rule to database + firewall + portproxy
func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ruleName := r.FormValue("rule_name")
		listenAddress := r.FormValue("listen_address")
		listenPort := r.FormValue("listen_port")
		connectAddress := r.FormValue("connect_address")
		connectPort := r.FormValue("connect_port")

		// Add to database
		_, err := db.Exec(`INSERT INTO rules (rule_name, listen_address, listen_port, connect_address, connect_port) VALUES (?, ?, ?, ?, ?)`,
			ruleName, listenAddress, listenPort, connectAddress, connectPort)
		if err != nil {
			http.Error(w, "Failed to save rule to database", http.StatusInternalServerError)
			return
		}

		// Add portproxy rule
		addPortProxy := exec.Command("netsh", "interface", "portproxy", "add", "v4tov4",
			"listenaddress="+listenAddress, "listenport="+listenPort,
			"connectaddress="+connectAddress, "connectport="+connectPort)
		if err := addPortProxy.Run(); err != nil {
			http.Error(w, "Failed to add port proxy", http.StatusInternalServerError)
			return
		}

		// Add firewall rule
		addFirewall := exec.Command("netsh", "advfirewall", "firewall", "add", "rule",
			"name="+ruleName, "protocol=TCP", "dir=in",
			"localport="+listenPort, "action=allow")
		if err := addFirewall.Run(); err != nil {
			http.Error(w, "Failed to add firewall rule", http.StatusInternalServerError)
			return
		}

		listHandler(w, r)
	}
}

// Remove rule from database + firewall + portproxy
func removeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ruleName := r.FormValue("rule_name")

		// Fetch rule from database
		row := db.QueryRow(`SELECT listen_address, listen_port FROM rules WHERE rule_name = ?`, ruleName)
		var listenAddress, listenPort string
		err := row.Scan(&listenAddress, &listenPort)
		if err != nil {
			http.Error(w, "Failed to find rule in database", http.StatusInternalServerError)
			return
		}

		// Remove from database
		_, err = db.Exec(`DELETE FROM rules WHERE rule_name = ?`, ruleName)
		if err != nil {
			http.Error(w, "Failed to remove rule from database", http.StatusInternalServerError)
			return
		}

		// Remove portproxy rule
		removePortProxy := exec.Command("netsh", "interface", "portproxy", "delete", "v4tov4",
			"listenaddress="+listenAddress, "listenport="+listenPort)
		if err := removePortProxy.Run(); err != nil {
			http.Error(w, "Failed to remove port proxy", http.StatusInternalServerError)
			return
		}

		// Remove firewall rule
		removeFirewall := exec.Command("netsh", "advfirewall", "firewall", "delete", "rule",
			"name="+ruleName)
		if err := removeFirewall.Run(); err != nil {
			http.Error(w, "Failed to remove firewall rule", http.StatusInternalServerError)
			return
		}

		listHandler(w, r)
	}
}

// List all rules
func listHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT rule_name, listen_address, listen_port, connect_address, connect_port FROM rules`)
	if err != nil {
		http.Error(w, "Failed to fetch rules from database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var rules []PortProxy
	for rows.Next() {
		var rule PortProxy
		if err := rows.Scan(&rule.RuleName, &rule.ListenAddress, &rule.ListenPort, &rule.ConnectAddress, &rule.ConnectPort); err != nil {
			http.Error(w, "Failed to parse rules", http.StatusInternalServerError)
			return
		}
		rules = append(rules, rule)
	}

	if err := templates.ExecuteTemplate(w, "list.html", rules); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getRouterIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ip := ipNet.IP.String()
			// Verify if it's in a private IP range (e.g., 192.168.x.x)
			if strings.HasPrefix(ip, "192.168.") || strings.HasPrefix(ip, "10.") || strings.HasPrefix(ip, "172.") {
				return ip
			}
		}
	}
	return "127.0.0.1"
}
