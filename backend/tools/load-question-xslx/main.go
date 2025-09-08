package main

import (
	"fmt"

	"github.com/lehaisonagentai3/free-contest/backend/internal/utils"
)

func main() {
	contestPath := "officers.xlsx"
	officers, err := utils.LoadOfficers(contestPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Officers loaded successfully!")
	for _, officer := range officers {
		fmt.Printf("ID: %d, Name: %s, Rank: %s, Position: %s, Unit: %s\n", officer.ID, officer.Name, officer.Rank, officer.Position, officer.Unit)
	}
}
