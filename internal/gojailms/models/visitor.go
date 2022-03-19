package models

type Visitor struct {
	VisitorID        int    `json:"visitorID"`
	UserID           int    `json:"userID"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	MiddleName       string `json:"middleName,omitempty"`
	Address          string `json:"address"`
	BirthPlace       string `json:"birthPlace"`
	BirthDate        string `json:"birthDate"`
	LastModifiedDate string `json:"lastModifiedDate"`
	CreatedDate      string `json:"createdDate"`
}
