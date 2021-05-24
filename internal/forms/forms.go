package forms

import (
	"fmt"
	//"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

//embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

//New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//Has checks if a form field is in post and not empty
func (f *Form) Has(field string) bool {
	if f.Get(field) == "" {
		return false
	} else {
		return true
	}
}

// Required checks for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MinLength check for minimum length
func (f *Form) MinLength(field string, length int) bool {
	input := f.Get(field)
	if len(input) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

//Checks if given field is an email adress
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email adress")
	}
}
