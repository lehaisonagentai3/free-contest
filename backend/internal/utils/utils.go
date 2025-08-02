package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lehaisonagentai3/free-contest/backend/internal/model"
	"github.com/xuri/excelize/v2"
)

func ImportQuestionFromJson(filePath string) ([]model.Question, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var questions []model.Question
	if err := json.Unmarshal(data, &questions); err != nil {
		return nil, err
	}
	return questions, nil
}

func LoadQuestionFromExcel(filePath string) ([]model.Question, error) {
	// 1. Mở file Excel
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if f.SheetCount == 0 {
		return nil, errors.New("no sheets found in the Excel file")
	}

	sheetName := f.GetSheetList()[0]
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	if len(rows)%7 != 0 && (len(rows)+1)%7 != 0 {
		return nil, errors.New("invalid row count, expected multiple of 7")
	}

	var questions []model.Question

	// 3. Duyệt qua từng row, gom mỗi nhóm 7 dòng thành 1 Question
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 2 {
			return nil, errors.New("row must have 2 columns, got " + fmt.Sprint(len(row)))
		}
		// Dòng bắt đầu bằng "Câu X"
		if strings.HasPrefix(row[0], "Câu") {

			q := model.Question{
				ID:      i + 1, // Sử dụng chỉ số dòng làm ID tạm thời
				Content: row[1],
			}

			// đảm bảo còn đủ các dòng A, B, C, D, Đáp án
			if i+5 < len(rows) {
				optA := rows[i+1]
				optB := rows[i+2]
				optC := rows[i+3]
				optD := rows[i+4]
				ans := rows[i+5]

				if len(optA) >= 2 {
					q.AnswerA = optA[1]
				}
				if len(optB) >= 2 {
					q.AnswerB = optB[1]
				}
				if len(optC) >= 2 {
					q.AnswerC = optC[1]
				}
				if len(optD) >= 2 {
					q.AnswerD = optD[1]
				}
				if len(ans) >= 2 {
					q.Correct = ans[1]
				}
			} else {
				return nil, errors.New("not enough rows for question options and answer, row " + fmt.Sprint(i+1) + " not enough")
			}

			questions = append(questions, q)
			// bỏ qua nhóm 7 dòng (5 dòng data + 1 dòng blank)
			i += 6
		}
	}

	// 4. In ra kiểm tra
	// for _, q := range questions {
	// 	fmt.Printf("Q%d: %s\n", q.ID, q.Content)
	// 	fmt.Printf("  A. %s\n", q.AnswerA)
	// 	fmt.Printf("  B. %s\n", q.AnswerB)
	// 	fmt.Printf("  C. %s\n", q.AnswerC)
	// 	fmt.Printf("  D. %s\n", q.AnswerD)
	// 	fmt.Printf("  -> Đáp án: %s\n\n", q.Correct)
	// }
	// This function should implement the logic to read questions from an Excel file
	// For now, we return an empty slice and nil error
	return questions, nil
}

// /Users/maianhnguyen/go/src/github.com/lehaisonagentai3/free-contest/backend/Kỳ thi sĩ quan phân đội
func LoadContestInfo(path string) (*model.Contest, error) {
	folderName := filepath.Base(path)
	contest := &model.Contest{
		ID:         1,
		Name:       folderName,
		Subjects:   []model.Subject{},
		FolderPath: path,
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for id, entry := range entries {
		if entry.IsDir() {
			subjectPath := filepath.Join(path, entry.Name())
			subjectParts := strings.Split(entry.Name(), "-")
			if len(subjectParts) < 3 {
				return nil, fmt.Errorf("invalid subject folder name: %s", entry.Name())
			}
			// Giả sử tên thư mục là "Tên đề thi - Thời gian - phút"
			subjectName := strings.TrimSpace(subjectParts[0])
			testTime, _ := strconv.ParseInt(strings.TrimSpace(subjectParts[1]), 10, 64)
			subject := model.Subject{
				Name:        subjectName,
				Description: subjectName,
				FolderPath:  subjectPath,
				TestTime:    int(testTime),
				ID:          id + 1, // ID bắt đầu từ 1
				ContestID:   contest.ID,
			}
			// Đọc các chương trong thư mục
			chapterEntries, err := os.ReadDir(subject.FolderPath)
			if err != nil {
				return nil, err
			}
			for chapterID, chapterEntry := range chapterEntries {
				if chapterEntry.IsDir() {
					chapterPath := chapterEntry.Name()
					// Giả sử tên chương là "chương X - số câu hỏi - câu"
					chapterParts := strings.Split(chapterPath, "-")
					if len(chapterParts) < 3 {
						return nil, fmt.Errorf("invalid chapter folder name: %s", chapterPath)
					}
					chapterName := strings.TrimSpace(chapterParts[0])
					numberTestQuestion, _ := strconv.Atoi(strings.TrimSpace(chapterParts[1]))
					if numberTestQuestion <= 0 {
						return nil, fmt.Errorf("invalid number of questions in chapter: %s", chapterName)
					}
					chapter := model.Chapter{
						ID:              chapterID + 1, // ID bắt đầu từ 1
						SubjectID:       subject.ID,
						Name:            chapterName,
						NumQuestionTest: numberTestQuestion, // Giả sử mỗi chương có 20 câu
						FolderPath:      filepath.Join(subject.FolderPath, chapterPath),
					}
					subject.NumQuestionTest += chapter.NumQuestionTest
					// Đọc các câu hỏi trong chương (câu hỏi là cá file *.xlsx trong thư mục chương)
					questionFiles, err := os.ReadDir(chapter.FolderPath)
					if err != nil {
						return nil, err
					}
					for _, questionFile := range questionFiles {
						if filepath.Ext(questionFile.Name()) == ".xlsx" {
							listQuestion, err := LoadQuestionFromExcel(filepath.Join(chapter.FolderPath, questionFile.Name()))
							if err != nil {
								return nil, err
							}
							if len(listQuestion) < chapter.NumQuestionTest {
								return nil, fmt.Errorf("not enough questions in chapter %s, expected %d, got %d", chapter.Name, chapter.NumQuestionTest, len(listQuestion))
							}
							chapter.Questions = append(chapter.Questions, listQuestion...)
							chapter.TotalQuestions += len(listQuestion)
						}
					}
					subject.Chapters = append(subject.Chapters, chapter)

				}
			}

			contest.Subjects = append(contest.Subjects, subject)
		}
	}
	return contest, nil
}
