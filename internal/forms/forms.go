package forms

import (
	"net/http"
	"net/url"
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

//Has checks if a form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool{
	if  r.Form.Get("field") == ""{
		return false
	} else {
		return true
	}
}