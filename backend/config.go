package main

import (
	"database/sql"
	"fmt"
	"log"
)

type SystemConfiguration struct {
	ID              int
	MACAddress      string
	DHCPClass       string
	Subnet          string
	OneTimeOverride bool
}

func loadConfiguration(db *sql.DB, id int) (*SystemConfiguration, error) {
	var config SystemConfiguration
	err := db.QueryRow("SELECT id, mac_address, dhcp_class, subnet, one_time_override FROM systems WHERE id = ?", id).Scan(&config.ID, &config.MACAddress, &config.DHCPClass, &config.Subnet, &config.OneTimeOverride)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %v", err)
	}
	return &config, nil
}

func saveConfiguration(db *sql.DB, config *SystemConfiguration) error {
	_, err := db.Exec("INSERT OR REPLACE INTO systems (id, mac_address, dhcp_class, subnet, one_time_override) VALUES (?, ?, ?, ?, ?)", config.ID, config.MACAddress, config.DHCPClass, config.Subnet, config.OneTimeOverride)
	if err != nil {
		return fmt.Errorf("failed to save configuration: %v", err)
	}
	return nil
}

func mergeConfigurations(baseConfig string, configBlocks []string) string {
	mergedConfig := baseConfig
	for _, block := range configBlocks {
		mergedConfig += "\n" + block
	}
	return mergedConfig
}
