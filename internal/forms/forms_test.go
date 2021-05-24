package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
	form.Errors.Add("testfield", "this is a test error")
	isValid = form.Valid()
	if isValid {
		t.Error("got valid when should have been invalid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_MinLength(t *testing.T) {
	//r := httptest.NewRequest("POST", "/whatever", nil)
	postedData := url.Values{}
	postedData.Add("a", "abc")
	//r.PostForm = postedData
	form := New(postedData)
	form.MinLength("a", 3)

	if !form.Valid() {
		t.Error("detects length of field input insufficient when it is long enough")
	}

	//r = httptest.NewRequest("POST", "/whatever", nil)
	postedData = url.Values{}
	postedData.Add("b", "bc")
	form = New(postedData)
	form.MinLength("b", 3)
	if form.Valid() {
		t.Error("does not detect when field input is shorter than required")
	}
}

func TestForm_Has(t *testing.T) {
	//r := httptest.NewRequest("POST", "/whatever", nil)
	postedData := url.Values{}
	postedData.Add("a", "abc")
	//r.PostForm = postedData
	form := New(postedData)
	formHasField := form.Has("a")
	if !formHasField {
		t.Error("says that form does not have field, when it does ")
	}
	formHasField = form.Has("b")
	if formHasField {
		t.Error("says that form does have field, when it doesn't")
	}
}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	postedData := url.Values{}
	postedData.Add("a", "me@here.com")
	postedData.Add("b", "@abc.")
	r.PostForm = postedData
	form := New(r.PostForm)
	form.IsEmail("a")
	isValid := form.Valid()
	if !isValid {
		t.Error("shows email is not valid when it is")
	}
	form = New(r.PostForm)
	form.IsEmail("b")
	isValid = form.Valid()
	if isValid {
		t.Error("shows valid email when it is not valid")
	}
}

func TestErrors_Get(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	if form.Errors.Get("hdga") != "" {
		t.Error("function returned an error when there was no error in the form")
	}
	form.Errors.Add("error1", "this is a testerror")
	if form.Errors.Get("error1") == "" {
		t.Error("did not return error when it should have")
	}
}
