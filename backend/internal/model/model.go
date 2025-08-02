package model

type Officer struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	UnitID   int     `json:"unit_id"`
	Unit     *Unit   `json:"unit,omitempty"`
	Score    float32 `json:"score"`
	Rank     string  `json:"rank"`
	Position string  `json:"position"`
}

type Unit struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ParentUnit *Unit  `json:"parent_unit,omitempty"`
}

type Contest struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Subjects    []Subject `json:"subjects"`    // list of subjects in the contest
	FolderPath  string    `json:"folder_path"` // path to the folder containing questions
}

type Test struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	ContestID     string     `json:"contest_id"`
	Duration      int        `json:"duration"` // in seconds
	Subject       Subject    `json:"subject"`
	Officer       Officer    `json:"officer"`
	Questions     []Question `json:"questions"`      // list of questions in the test
	RemainingTime int        `json:"remaining_time"` // time left for the test in seconds
}

type Question struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	AnswerA string `json:"answer_a"`
	AnswerB string `json:"answer_b"`
	AnswerC string `json:"answer_c"`
	AnswerD string `json:"answer_d"`
	Correct string `json:"correct"` // correct answer (A, B, C, or D)
}

type Subject struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	ContestID       int       `json:"contest_id"`
	Chapters        []Chapter `json:"chapters"`          // list of chapters in the subject
	NumQuestionTest int       `json:"num_question_test"` // number of questions of subject in the test
	TestTime        int       `json:"test_time"`         // time limit for the test in minutes
	FolderPath      string    `json:"folder_path"`       // path to the folder containing questions
}

type Chapter struct {
	ID              int        `json:"id"`
	SubjectID       int        `json:"subject_id"`
	Name            string     `json:"name"`
	Questions       []Question `json:"questions"`
	NumQuestionTest int        `json:"num_question_test"` // number of questions of chapter in the test
	FolderPath      string     `json:"folder_path"`       // path to the folder containing questions
	TotalQuestions  int        `json:"total_questions"`   // total number of questions in the chapter
}

type Submission struct {
	ID          int               `json:"id"`
	OfficerID   int               `json:"officer_id"`
	TestID      int               `json:"test_id"`
	Answers     map[string]string `json:"answers"` // question ID to answer mapping
	Score       float32           `json:"score"`
	SubmittedAt int64             `json:"submitted_at"` // timestamp of submission
}

type ContestMetaInfo struct {
	RootPath string  `json:"root_path"` // root path of the contest
	Contest  Contest `json:"contest"`   // contest information
}
