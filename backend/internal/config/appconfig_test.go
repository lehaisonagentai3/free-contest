package config

import (
	"testing"

	"github.com/lehaisonagentai3/free-contest/backend/internal/model"
)

func TestSaveConfig(t *testing.T) {
	conf := &AppConfig{

		ListOfficer: []*model.Officer{
			{ID: 1, Name: "Nguyễn Văn A", Unit: "Phòng Tham mưu"},
			{ID: 2, Name: "Trần Thị B", Unit: "Phòng Chính trị"},
			{ID: 3, Name: "Lê Văn C", Unit: "Phòng Hậu cần - Kỹ thuật"},
		},
		ContestPath: "/Users/maianhnguyen/go/src/github.com/lehaisonagentai3/free-contest/backend/Kỳ thi sĩ quan phân đội",
	}
	configFile := "config.json"
	if err := SaveAppConfig(configFile, conf); err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

}
