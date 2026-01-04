package utils

import "strings"

/*
 Key: 'Video.Author.Name' Error:Field validation for 'Name' failed on the 'required' tag
Key: 'Video.Author.Email' Error:Field validation for 'Email' failed on the 'required' tag
*/

/*
[
{"KEY":"Video.Author.Name", "ERROR":"Field validation for 'Name' failed on the 'required' tag"},
{"KEY":"Video.Author.Email", "ERROR":"Field validation for 'Email' failed on the 'required' tag"}
]
*/
//

type ValidationError struct {
	Key   string `json:"key" example:"Video.Author.Name"`
	Error string `json:"error" example:"Field validation for 'Name' failed on the 'required' tag"`
}

func FormatValidationError(errs error) []ValidationError {
	ve := strings.Split(errs.Error(), "\n")
	out := make([]ValidationError, len(ve))
	for i, fe := range ve {
		parts := strings.SplitN(fe, " Error:", 2)
		parts[0] = strings.TrimPrefix(parts[0], "Key: '")
		parts[0] = strings.TrimSuffix(parts[0], "'")
		if len(parts) == 2 {
			out[i] = ValidationError{
				Key:   strings.TrimSpace(parts[0]),
				Error: strings.TrimSpace(parts[1]),
			}
		}
	}
	return out

}
