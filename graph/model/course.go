package model

type Course struct {
	ID                     string   `json:"_id" bson:"_id"`
	CourseNumber           string   `json:"course_number" bson:"course_number"`
	SubjectPrefix          string   `json:"subject_prefix" bson:"subject_prefix"`
	Title                  string   `json:"title" bson:"title"`
	Descriptions           string   `json:"descriptions" bson:"descriptions"`
	School                 string   `json:"school" bson:"school"`
	CreditHours            string   `json:"credit_hours" bson:"credit_hours"`
	ClassLevel             string   `json:"class_level" bson:"class_level"`
	ActivityType           string   `json:"activity_type" bson:"activity_type"`
	Grading                string   `json:"grading" bson:"grading"`
	InternalCourseNumber   string   `json:"internal_course_number" bson:"internal_course_number"`
	Sections               []string `json:"sections" bson:"sections"`
	LectureContactHours    string   `json:"lecture_contact_hours" bson:"letcure_contact_hours"`
	LaboratoryContactHours string   `json:"laboratory_contact_hours" bson:"laboratory_contact_hours"`
	OfferingFrequency      string   `json:"offering_frequency" bson:"offering_frequency"`
	V                      int32    `json:"__v" bson:"__v"`
}
