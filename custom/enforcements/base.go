package enforcements

import (
	"strings"

	"github.com/rrojan/enforcer"
)

var BASE_ENFORCEMENTS enforcer.CustomEnforcements = enforcer.CustomEnforcements{
	{
		"ContainsDirtyWords": func (s string) string {
			if strings.Contains(s, "dirty") {
				return "Field contains inappropriate words!"
			}
			return ""
		},
	},
} 