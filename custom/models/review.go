package models

type Review struct {
	Stars int `json:"stars" enforce:"min:1 max:5"`
	Comment string `enforce:"custom:ContainsDirtyWords,NotFromCompetitor"`
}