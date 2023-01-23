package models

type VisitorHistory struct {
	VisitorHistID int     `json:"visitorHistoryID,omitempty"`
	UserID        int     `json:"userID,omitempty"`
	FirstName     string  `json:"firstName,omitempty"`
	LastName      string  `json:"lastName,omitempty"`
	MiddleName    string  `json:"middleName,omitempty"`
	Remarks       string  `json:"remarks,omitempty"`
	VisitDate     string  `json:"visitDateTime,omitempty"`
	VisitOut      *string `json:"visitOut"`
}
