package hunit

import (
	"hunit/expr"
	"hunit/test"
)

// Interpolate if required
func interpolateIfRequired(c Context, s string) (string, error) {
	if (c.Options & test.OptionInterpolateVariables) == test.OptionInterpolateVariables {
		return expr.Interpolate(s, c.Variables)
	} else {
		return s, nil
	}
}
