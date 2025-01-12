package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	r.GET("/ipxe", handleIPXERequest)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func handleIPXERequest(c *gin.Context) {
	mac := c.Query("mac")
	dhcpClass := c.Query("dhcp_class")
	subnet := c.Query("subnet")

	config, err := getIPXEConfig(mac, dhcpClass, subnet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, config)
}

func getIPXEConfig(mac, dhcpClass, subnet string) (string, error) {
	db, err := sql.Open("sqlite3", "./ipxe.db")
	if err != nil {
		return "", fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	config, err := queryConfigByMAC(db, mac)
	if err == nil {
		return config, nil
	}

	config, err = queryConfigByDHCPClass(db, dhcpClass)
	if err == nil {
		return config, nil
	}

	config, err = queryConfigBySubnet(db, subnet)
	if err == nil {
		return config, nil
	}

	return "", fmt.Errorf("no matching configuration found")
}

func queryConfigByMAC(db *sql.DB, mac string) (string, error) {
	var config string
	err := db.QueryRow("SELECT content FROM configurations WHERE id = (SELECT configuration_id FROM system_configurations WHERE system_id = (SELECT id FROM systems WHERE mac_address = ?))", mac).Scan(&config)
	if err != nil {
		return "", fmt.Errorf("failed to query configuration by MAC: %v", err)
	}
	return config, nil
}

func queryConfigByDHCPClass(db *sql.DB, dhcpClass string) (string, error) {
	var config string
	err := db.QueryRow("SELECT content FROM configurations WHERE id = (SELECT configuration_id FROM system_configurations WHERE system_id = (SELECT id FROM systems WHERE dhcp_class = ?))", dhcpClass).Scan(&config)
	if err != nil {
		return "", fmt.Errorf("failed to query configuration by DHCP class: %v", err)
	}
	return config, nil
}

func queryConfigBySubnet(db *sql.DB, subnet string) (string, error) {
	var config string
	err := db.QueryRow("SELECT content FROM configurations WHERE id = (SELECT configuration_id FROM system_configurations WHERE system_id = (SELECT id FROM systems WHERE subnet = ?))", subnet).Scan(&config)
	if err != nil {
		return "", fmt.Errorf("failed to query configuration by subnet: %v", err)
	}
	return config, nil
}
