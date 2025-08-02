package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/lehaisonagentai3/free-contest/backend/internal/config"
	"github.com/lehaisonagentai3/free-contest/backend/internal/model"
	"github.com/lehaisonagentai3/free-contest/backend/internal/utils"
)

type ContestService struct {
	conf                    *config.AppConfig
	mapUnits                map[int]model.Unit
	mapSubjects             map[int]model.Subject
	mapOfficers             map[int]model.Officer
	contest                 *model.Contest              // Contest info loaded from config
	mapOfficerToSubjectTest map[int]map[int]*model.Test // map[officerID][subjectID]Test
}

func NewContestService(conf *config.AppConfig) (*ContestService, error) {
	mapSubjects := make(map[int]model.Subject)
	mapOfficers := make(map[int]model.Officer)
	for _, officer := range conf.ListOfficer {
		mapOfficers[officer.ID] = officer
	}
	mapUnits := make(map[int]model.Unit)
	for _, unit := range conf.ListUnit {
		mapUnits[unit.ID] = unit
	}

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

// Retrive list question from a subject, total question is field NumQuestionTest and each chapter has NumQuestionTest, NumQuestionTest is total question of Chapter.NumQuestionTest
func (s *ContestService) GetSubjectTestForOfficer(officerID int, subjectID int) (*model.Test, error) {
	if s.mapOfficerToSubjectTest == nil {
		s.mapOfficerToSubjectTest = make(map[int]map[int]*model.Test)
	}
	if _, ok := s.mapOfficerToSubjectTest[officerID]; !ok {
		s.mapOfficerToSubjectTest[officerID] = make(map[int]*model.Test)
	}
	if _, ok := s.mapOfficerToSubjectTest[officerID][subjectID]; ok {
		return nil, errors.New("officer already has a test for this subject") // Return existing test if it exists
	}
	if _, ok := s.mapSubjects[subjectID]; !ok {
		return nil, fmt.Errorf("no subjects available")
	}

	rand.NewSource(time.Now().UnixNano())

	subject, ok := s.mapSubjects[subjectID]
	if !ok {
		return nil, fmt.Errorf("subject not found")
	}
	officer, ok := s.mapOfficers[officerID]
	if !ok {
		return nil, fmt.Errorf("officer not found")
	}
	if subject.NumQuestionTest <= 0 {
		return nil, fmt.Errorf("subject does not have enough questions for test")
	}
	listQuestions := []model.Question{}
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

	return &model.Test{
		Subject:   subject,
		Questions: listQuestions,
		Officer:   officer,
		Duration:  subject.TestTime * 60, // Convert minutes to seconds
		ID:        rand.Intn(1000),       // Random ID for the test
	}, nil
}
