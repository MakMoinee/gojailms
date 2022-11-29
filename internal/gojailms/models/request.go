package models

type CreateRequest struct {
	LocalUser    Users   `json:"user"`
	LocalVisitor Visitor `json:"visitor"`
}

type VisitorHistoryRequest struct {
	VisitorID string `json:"visitorID,omitempty"`
	Remarks   string `json:"remarks,omitempty"`
}
