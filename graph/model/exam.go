package model

import "go.mongodb.org/mongo-driver/bson"

type Exam interface {
	IsExam()
}

type Outcome interface {
	IsOutcome()
}

type PossibleOutcomes struct {
	Requirement      *bson.Raw     `json:"requirement"`
	PossibleOutcomes [][]*bson.Raw `json:"outcome" bson:"outcome"`
}

type ALEKSExam struct {
	ID        string             `json:"_id" bson:"_id"`
	Placement []PossibleOutcomes `json:"placement"`
}

func (ALEKSExam) IsExam() {}

type APExam struct {
	ID     string             `json:"_id" bson:"_id"`
	Name   string             `json:"name"`
	Yields []PossibleOutcomes `json:"yields"`
}

func (APExam) IsExam() {}

type CLEPExam struct {
	ID     string             `json:"_id" bson:"_id"`
	Name   string             `json:"name"`
	Yields []PossibleOutcomes `json:"yields"`
}

func (CLEPExam) IsExam() {}

type CSPlacementExam struct {
	ID     string             `json:"_id" bson:"_id"`
	Yields []PossibleOutcomes `json:"yields"`
}

func (CSPlacementExam) IsExam() {}

type IBExam struct {
	ID     string             `json:"_id" bson:"_id"`
	Name   string             `json:"name"`
	Level  string             `json:"level"`
	Yields []PossibleOutcomes `json:"yields"`
}

func (IBExam) IsExam() {}
