package validator

import (
	"strings"

	"github.com/gobuffalo/validate"
)

func FormatErrors(error *validate.Errors) string {
	var res string
	for _, v := range error.Errors {
		res = res + strings.Join(v, ", ") + ", "
	}
	return res
}
