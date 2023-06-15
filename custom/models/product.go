package models

type Product struct {
	// Here, FilterDirtyWords is from the globally shared base enforcement
	// FitsTitleFormat is from the product enforcement
	Name        string `json:"name" enforce:"required custom:ContainsDirtyWords,FitsTitleFormat"`
	Description string `json:"description" enforce:"wordCount:5,50"`
	Price       string `json:"price" enforce:"custom:IsNotOverpriced"`
}