package models

type Registration struct {
	SpiffeID string `json:"spiffeid" gorm:"primary_key"`
	Selectors string `json:"selectors"`
}
