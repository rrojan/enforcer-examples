package enforcements

import (
	"strconv"

	"github.com/rrojan/enforcer"
)

func NewProductEnforcementsWithBase() enforcer.CustomEnforcements {
	e := enforcer.CustomEnforcements{{
		"FitsTitleFormat": func (title string) string {
			isValidTitleFormat := true // validation logic goes here
			if isValidTitleFormat {
				return ""
			}
			return "Invalid product title format"
		},
		"IsNotOverpriced": func (priceStr string) string {
			price, _ := strconv.Atoi(priceStr)
			if price > 1000 {
				return "Product is overpriced"
			}
			return ""
		},
	}}
	return append(BASE_ENFORCEMENTS, e...)
}