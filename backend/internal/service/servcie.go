package service

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/lehaisonagentai3/free-contest/backend/internal/config"
	"github.com/lehaisonagentai3/free-contest/backend/internal/model"
	"github.com/lehaisonagentai3/free-contest/backend/internal/utils"
)

type ContestService struct {
	conf                    *config.AppConfig
	mapUnits                map[string]string
	mapSubjects             map[int]*model.Subject
	mapOfficers             map[int]*model.Officer
	contest                 *model.Contest              // Contest info loaded from config, include all subjecs and chapters and questions
	mapOfficerToSubjectTest map[int]map[int]*model.Test // map[officerID][subjectID]Test
}

func NewContestService(conf *config.AppConfig) (*ContestService, error) {
	var err error
	conf.ListOfficer, err = utils.LoadOfficers(conf.OfficerPath)
	if err != nil {
		return nil, err
	}
	mapSubjects := make(map[int]*model.Subject)
	mapOfficers := make(map[int]*model.Officer)
	for _, officer := range conf.ListOfficer {
		mapOfficers[officer.ID] = officer
	}
	fmt.Printf("Loaded %d officers from %s\n", len(conf.ListOfficer), conf.OfficerPath)
	mapUnits := make(map[string]string)

	contestInfo, err := utils.LoadContestInfo(conf.ContestPath)
	if err != nil {
		return nil, err
	}

	for _, subject := range contestInfo.Subjects {
		mapSubjects[subject.ID] = subject
	}

	return &ContestService{
		conf:                    conf,
		mapSubjects:             mapSubjects,
		mapOfficers:             mapOfficers,
		mapUnits:                mapUnits,
		contest:                 contestInfo,
		mapOfficerToSubjectTest: make(map[int]map[int]*model.Test),
	}, nil
}

func (s *ContestService) GetContestInfo() *model.Contest {
	return s.contest
}

// GetAllUnits returns all units
func (s *ContestService) GetAllUnits() []string {
	units := make([]string, 0, len(s.mapUnits))
	for unitName := range s.mapUnits {
		units = append(units, unitName)
	}
	return units
}

// GetAllOfficers returns all officers with their unit information
func (s *ContestService) GetAllOfficers() []*model.Officer {
	officers := make([]*model.Officer, 0, len(s.mapOfficers))
	for _, officer := range s.mapOfficers {
		officer.Score = s.caculateTotalScoreOfOffices(officer)
		officers = append(officers, officer)
	}
	return officers
}

func (s *ContestService) caculateTotalScoreOfOffices(office *model.Officer) float32 {
	if office == nil || len(office.ListSubmission) == 0 {
		return 0
	}
	var totalScore float32
	for _, submission := range office.ListSubmission {
		totalScore += submission.Score
	}
	return totalScore
}

// GetOfficerByID returns an officer by ID with unit information
func (s *ContestService) GetOfficerByID(officerID int) (*model.Officer, error) {
	officer, exists := s.mapOfficers[officerID]
	if !exists {
		return nil, fmt.Errorf("officer not found")
	}

	// Add unit information to officer
	if unit, exists := s.mapUnits[officer.Unit]; exists {
		officer.Unit = unit
	}

	return officer, nil
}

// GetAllSubjects returns all subjects without questions and chapters
func (s *ContestService) GetAllSubjects() []*model.Subject {
	subjects := make([]*model.Subject, 0, len(s.mapSubjects))
	for _, subject := range s.mapSubjects {
		// Create a copy without chapters (which contain questions)
		subjectInfo := &model.Subject{
			ID:              subject.ID,
			Name:            subject.Name,
			Description:     subject.Description,
			ContestID:       subject.ContestID,
			NumQuestionTest: subject.NumQuestionTest,
			TestTime:        subject.TestTime,
			Chapters:        nil, // Explicitly set to nil to exclude chapters and questions
		}
		subjects = append(subjects, subjectInfo)
	}
	return subjects
}

// Retrive list question from a subject, total question is field NumQuestionTest and each chapter has NumQuestionTest, NumQuestionTest is total question of Chapter.NumQuestionTest
func (s *ContestService) GetSubjectTestForOfficer(officerID int, subjectID int) (*model.Test, error) {
	if s.mapOfficerToSubjectTest == nil {
		s.mapOfficerToSubjectTest = make(map[int]map[int]*model.Test)
	}
	if _, ok := s.mapOfficerToSubjectTest[officerID]; !ok {
		s.mapOfficerToSubjectTest[officerID] = make(map[int]*model.Test)
	}
	if _, ok := s.mapOfficerToSubjectTest[officerID][subjectID]; ok {
		// Check if test is started and not expired
		if s.mapOfficerToSubjectTest[officerID][subjectID].StartTime > 0 {
			elapsed := time.Now().Unix() - s.mapOfficerToSubjectTest[officerID][subjectID].StartTime
			if elapsed > int64(s.mapOfficerToSubjectTest[officerID][subjectID].Duration) {
				return nil, fmt.Errorf("test is expired")
			}
		}
		return s.mapOfficerToSubjectTest[officerID][subjectID], nil // Return existing test if it exists
	}

	subject, ok := s.mapSubjects[subjectID]
	if !ok {
		return nil, fmt.Errorf("subject not found")
	}

	rand.NewSource(time.Now().UnixNano())

	officer, ok := s.mapOfficers[officerID]
	if !ok {
		return nil, fmt.Errorf("officer not found")
	}
	if subject.NumQuestionTest <= 0 {
		return nil, fmt.Errorf("subject does not have enough questions for test")
	}
	listQuestions := []*model.Question{}
	for _, chapter := range subject.Chapters {
		if chapter.NumQuestionTest <= 0 {
			continue
		}
		rand.Shuffle(len(chapter.Questions), func(i, j int) {
			chapter.Questions[i], chapter.Questions[j] = chapter.Questions[j], chapter.Questions[i]
		})
		if len(chapter.Questions) < chapter.NumQuestionTest {
			return nil, fmt.Errorf("not enough questions in chapter %s", chapter.Name)
		}
		listQuestions = append(listQuestions, chapter.Questions[:chapter.NumQuestionTest]...)
	}

	test := &model.Test{
		Subject: &model.Subject{
			ID:              subject.ID,
			Name:            subject.Name,
			Description:     subject.Description,
			ContestID:       subject.ContestID,
			NumQuestionTest: subject.NumQuestionTest,
			TestTime:        subject.TestTime,
		},
		Questions:     listQuestions,
		Officer:       officer,
		Duration:      subject.TestTime * 60, // Convert minutes to seconds
		RemainingTime: subject.TestTime * 60, // Initially same as duration
		ID:            rand.Intn(1000),       // Random ID for the test
	}

	// Store the test for this officer and subject
	s.mapOfficerToSubjectTest[officerID][subjectID] = test

	return test, nil
}

// StartTest starts a test for an officer by setting the start time
func (s *ContestService) StartTest(officerID int, testID int) (*model.Test, error) {
	// Find the test for the officer
	if s.mapOfficerToSubjectTest == nil {
		return nil, fmt.Errorf("no tests found")
	}

	if _, ok := s.mapOfficerToSubjectTest[officerID]; !ok {
		return nil, fmt.Errorf("no tests found for officer")
	}

	// Search for the test with the given testID
	var foundTest *model.Test
	for _, test := range s.mapOfficerToSubjectTest[officerID] {
		if test.ID == testID {
			foundTest = test
			break
		}
	}

	if foundTest == nil {
		return nil, fmt.Errorf("test not found")
	}

	// Check if test is already started
	if foundTest.StartTime > 0 {
		return nil, fmt.Errorf("test already started")
	}

	// Set the start time
	foundTest.StartTime = time.Now().Unix()

	return foundTest, nil
}

// SubmitTest submits test answers and calculates the score
func (s *ContestService) SubmitTest(officerID int, testID int, answers map[string]string) (*model.Submission, error) {
	// Find the test for the officer
	if s.mapOfficerToSubjectTest == nil {
		return nil, fmt.Errorf("no tests found")
	}

	if _, ok := s.mapOfficerToSubjectTest[officerID]; !ok {
		return nil, fmt.Errorf("no tests found for officer")
	}
	// validate officerID and testID
	if _, ok := s.mapOfficers[officerID]; !ok {
		return nil, fmt.Errorf("officer not found")
	}

	// Search for the test with the given testID
	var foundTest *model.Test
	for _, test := range s.mapOfficerToSubjectTest[officerID] {
		if test.ID == testID {
			foundTest = test
			break
		}
	}

	if foundTest == nil {
		return nil, fmt.Errorf("test not found")
	}

	// Check if test has started
	if foundTest.StartTime == 0 {
		return nil, fmt.Errorf("test has not been started yet")
	}

	// Check if test is already finished
	if foundTest.IsFinished {
		return nil, fmt.Errorf("test has already been submitted")
	}

	// Calculate score
	totalQuestions := len(foundTest.Questions)
	correctAnswers := 0

	// Create a map of question IDs to correct answers for easy lookup
	questionAnswers := make(map[string]string)
	for _, question := range foundTest.Questions {
		questionAnswers[fmt.Sprintf("%d", question.ID)] = question.Correct
	}

	// Check answers and count correct ones
	for questionIDStr, userAnswer := range answers {
		if correctAnswer, exists := questionAnswers[questionIDStr]; exists {
			if strings.EqualFold(userAnswer, correctAnswer) {
				correctAnswers++
			}
		}
	}

	// Calculate score as percentage
	score := float32(correctAnswers) / float32(totalQuestions) * 10

	// Mark test as finished
	foundTest.IsFinished = true

	// Create submission record
	submission := &model.Submission{
		ID:          rand.Intn(10000), // Random ID for submission
		OfficerID:   officerID,
		TestID:      testID,
		Answers:     answers,
		Score:       score,
		SubmittedAt: time.Now().Unix(),
		SubjectID:   foundTest.Subject.ID,
		SubjectName: foundTest.Subject.Name,
	}
	// Add submission to officer's list
	if officer, exists := s.mapOfficers[officerID]; exists {
		officer.ListSubmission = append(officer.ListSubmission, submission)
	}
	return submission, nil
}
