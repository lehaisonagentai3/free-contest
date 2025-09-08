package model

type Officer struct {
	ID             int           `json:"id,omitempty"`
	Name           string        `json:"name,omitempty"`
	Unit           string        `json:"unit,omitempty"`
	Score          float32       `json:"score,omitempty"`
	Rank           string        `json:"rank,omitempty"`
	Position       string        `json:"position,omitempty"`
	ListSubmission []*Submission `json:"list_submission,omitempty"` // list of submissions made by the officer
}

type Contest struct {
	ID          int        `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Subjects    []*Subject `json:"subjects,omitempty"`    // list of subjects in the contest
	FolderPath  string     `json:"folder_path,omitempty"` // path to the folder containing questions
}

type Test struct {
	ID            int         `json:"id,omitempty"`
	Name          string      `json:"name,omitempty"`
	ContestID     string      `json:"contest_id,omitempty"`
	Duration      int         `json:"duration,omitempty"` // in seconds
	Subject       *Subject    `json:"subject,omitempty"`
	Officer       *Officer    `json:"officer,omitempty"`
	Questions     []*Question `json:"questions,omitempty"`      // list of questions in the test
	RemainingTime int         `json:"remaining_time,omitempty"` // time left for the test in seconds
	IsFinished    bool        `json:"is_finished,omitempty"`    // whether the test is finished
	StartTime     int64       `json:"start_time,omitempty"`     // timestamp when the test started
}

type Question struct {
	ID      int    `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
	AnswerA string `json:"answer_a,omitempty"`
	AnswerB string `json:"answer_b,omitempty"`
	AnswerC string `json:"answer_c,omitempty"`
	AnswerD string `json:"answer_d,omitempty"`
	Correct string `json:"correct,omitempty"` // correct answer (A, B, C, or D)
}

type Subject struct {
	ID              int        `json:"id,omitempty"`
	Name            string     `json:"name,omitempty"`
	Description     string     `json:"description,omitempty"`
	ContestID       int        `json:"contest_id,omitempty"`
	Chapters        []*Chapter `json:"chapters,omitempty"`          // list of chapters in the subject
	NumQuestionTest int        `json:"num_question_test,omitempty"` // number of questions of subject in the test
	TestTime        int        `json:"test_time,omitempty"`         // time limit for the test in minutes
	FolderPath      string     `json:"folder_path,omitempty"`       // path to the folder containing questions
}

type Chapter struct {
	ID              int         `json:"id,omitempty"`
	SubjectID       int         `json:"subject_id,omitempty"`
	Name            string      `json:"name,omitempty"`
	Questions       []*Question `json:"questions,omitempty"`
	NumQuestionTest int         `json:"num_question_test,omitempty"` // number of questions of chapter in the test
	FolderPath      string      `json:"folder_path,omitempty"`       // path to the folder containing questions
	TotalQuestions  int         `json:"total_questions"`             // total number of questions in the chapter
}

type Submission struct {
	ID          int               `json:"id,omitempty"`
	OfficerID   int               `json:"officer_id,omitempty"`
	TestID      int               `json:"test_id,omitempty"`
	Answers     map[string]string `json:"answers,omitempty"` // question ID to answer mapping
	Score       float32           `json:"score,omitempty"`
	SubmittedAt int64             `json:"submitted_at,omitempty"` // timestamp of submission
	SubjectID   int               `json:"subject_id,omitempty"`   // ID of the subject for which the test was taken
	SubjectName string            `json:"subject_name,omitempty"` // name of the subject for which the test was taken
}

type ContestMetaInfo struct {
	RootPath string   `json:"root_path,omitempty"` // root path of the contest
	Contest  *Contest `json:"contest,omitempty"`   // contest information
}
