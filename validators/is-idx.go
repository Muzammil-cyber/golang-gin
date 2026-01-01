package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func IsIdx(fl validator.FieldLevel) bool {
	return strings.HasPrefix(fl.Field().String(), "idx_")
}
