package config

import (
	"testing"

	"github.com/lehaisonagentai3/free-contest/backend/internal/model"
)

func TestSaveConfig(t *testing.T) {
	conf := &AppConfig{
		ListUnit: []model.Unit{
			{ID: 1, Name: "Phòng Tham mưu"},
			{ID: 2, Name: "Phòng Chính trị"},
			{ID: 3, Name: "Phòng Hậu cần - Kỹ thuật"},
			{ID: 4, Name: "Tiểu đoàn 1"},
			{ID: 5, Name: "Tiểu đoàn 2"},
			{ID: 6, Name: "Tiểu đoàn 3"},
			{ID: 7, Name: "Đại đội 60"},
			{ID: 8, Name: "Đại đội 70"},
			{ID: 9, Name: "Kho 36"},
		},
		ListOfficer: []model.Officer{
			{ID: 1, Name: "Nguyễn Văn A", UnitID: 1},
			{ID: 2, Name: "Trần Thị B", UnitID: 2},
			{ID: 3, Name: "Lê Văn C", UnitID: 3},
		},
		ContestPath: "/Users/maianhnguyen/go/src/github.com/lehaisonagentai3/free-contest/backend/Kỳ thi sĩ quan phân đội",
	}
	configFile := "config.json"
	if err := SaveAppConfig(configFile, conf); err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

}
