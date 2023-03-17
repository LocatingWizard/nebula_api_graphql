package model

type Professor struct {
	ID          string     `json:"_id" bson:"_id"`
	FirstName   string     `json:"first_name" bson:"first_name"`
	LastName    string     `json:"last_name" bson:"last_name"`
	Titles      []string   `json:"titles"`
	Email       string     `json:"email"`
	PhoneNumber *string    `json:"phone_number" bson:"phone_number"`
	Office      *Location  `json:"office"`
	ProfileURI  *string    `json:"profile_uri" bson:"profile_uri"`
	ImageURI    *string    `json:"image_uri" bson:"image_uri"`
	OfficeHours []*Meeting `json:"office_hours" bson:"office_hours"`
	Sections    []*Section `json:"sections"`
}
