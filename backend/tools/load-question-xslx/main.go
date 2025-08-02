package main

import (
	"fmt"

	"github.com/lehaisonagentai3/free-contest/backend/internal/utils"
)

func main() {
	contestPath := "/Users/maianhnguyen/go/src/github.com/lehaisonagentai3/free-contest/backend/Kỳ thi sĩ quan phân đội"
	contest, err := utils.LoadContestInfo(contestPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Contest loaded successfully!")
	fmt.Printf("Contest ID: %d\n", contest.ID)
	fmt.Printf("Contest Name: %s\n", contest.Name)
	fmt.Printf("Contest Folder Path: %s\n", contest.FolderPath)
	fmt.Println("Total Subjects:", len(contest.Subjects))
	for _, subject := range contest.Subjects {
		fmt.Printf("Subject ID: %d, Name: %s, Description: %s, Total Chapters: %d, Test time: %d, Total Questions In Test: %d\n", subject.ID, subject.Name, subject.Description, len(subject.Chapters), subject.TestTime, subject.NumQuestionTest)
		for _, chapter := range subject.Chapters {
			fmt.Printf("  Chapter ID: %d, Name: %s, NumQuestionTest: %d, Total question: %d\n", chapter.ID, chapter.Name, chapter.NumQuestionTest, chapter.TotalQuestions)
		}
	}

}
