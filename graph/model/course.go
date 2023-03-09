package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	ID                     string                 `json:"_id" bson:"_id"`
	CourseNumber           string                 `json:"course_number" bson:"course_number"`
	SubjectPrefix          string                 `json:"subject_prefix" bson:"subject_prefix"`
	Title                  string                 `json:"title"`
	Description            string                 `json:"description"`
	School                 string                 `json:"school"`
	CreditHours            string                 `json:"credit_hours"`
	ClassLevel             string                 `json:"class_level"`
	ActivityType           string                 `json:"activity_type"`
	Grading                string                 `json:"grading"`
	InternalCourseNumber   string                 `json:"internal_course_number"`
	Prerequisites          *CollectionRequirement `json:"prerequisites"`
	Corequisites           *CollectionRequirement `json:"corequisites"`
	CoOrPreRequisites      *CollectionRequirement `json:"co_or_pre_requisites"`
	Sections               []*primitive.ObjectID  `json:"sections"`
	LectureContactHours    string                 `json:"lecture_contact_hours"`
	LaboratoryContactHours string                 `json:"laboratory_contact_hours"`
	OfferingFrequency      string                 `json:"offering_frequency"`
}

func (Course) IsOutcome() {}
