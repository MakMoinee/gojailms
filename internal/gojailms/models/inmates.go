package models

type Inmates struct {
	InmateID         int    `json:"inmateID"`
	CrimeID          int    `json:"crimeID"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	MiddleName       string `json:"middleName"`
	Address          string `json:"address"`
	BirthPlace       string `json:"birthPlace"`
	BirthDate        string `json:"birthDate"`
	LastModifiedDate string `json:"lastModifiedDate"`
	CreatedDate      string `json:"createdDate"`
}
