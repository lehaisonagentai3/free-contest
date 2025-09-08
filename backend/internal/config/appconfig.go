package config

import (
	"encoding/json"
	"os"

	"github.com/lehaisonagentai3/free-contest/backend/internal/model"
)

type AppConfig struct {
	ListOfficer []*model.Officer `json:"list_officer,omitempty"`
	ContestPath string           `json:"contest_path,omitempty"` // Path to the contest data directory contain multi subjects
	OfficerPath string           `json:"officer_path,omitempty"` // Path to the officer data directory
}

func LoadAppConfig(configFileJson string) (*AppConfig, error) {
	data, err := os.ReadFile(configFileJson)
	if err != nil {
		return nil, err
	}

	var config AppConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func SaveAppConfig(configFileJson string, config *AppConfig) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	if err := os.WriteFile(configFileJson, data, 0644); err != nil {
		return err
	}

	return nil
}
