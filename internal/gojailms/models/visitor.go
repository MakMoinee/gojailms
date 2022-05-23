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

type UserVisitor struct {
	UserID           int    `json:"userID,omitempty"`
	UserName         string `json:"userName,omitempty"`
	UserPassword     string `json:"userPassword,omitempty"`
	UserType         int    `json:"userType,omitempty"`
	VisitorID        int    `json:"visitorID,omitempty"`
	FirstName        string `json:"firstName,omitempty"`
	LastName         string `json:"lastName,omitempty"`
	MiddleName       string `json:"middleName,omitempty"`
	Address          string `json:"address,omitempty"`
	BirthPlace       string `json:"birthPlace,omitempty"`
	BirthDate        string `json:"birthDate,omitempty"`
	LastModifiedDate string `json:"lastModifiedDate,omitempty"`
	CreatedDate      string `json:"createdDate,omitempty"`
	Token            string `json:"authToken,omitempty"`
}
