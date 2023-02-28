package model

type Course struct {
	ID            string `json:"_id" bson:"_id"`
	CourseNumber  string `json:"course_number" bson:"course_number"`
	SubjectPrefix string `json:"subject_prefix" bson:"subject_prefix"`
	Title         string `json:"title" bson:"title"`
}
