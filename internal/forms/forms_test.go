package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/anything", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/anything", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/anything", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/anything", nil)

	postedData := url.Values{}
	postedData.Add("a", "")

	r.PostForm = postedData

	form := New(r.PostForm)
	has := form.Has("a")

	if has {
		t.Error("Form shows valid when required field is empty")
	}

	r, _ = http.NewRequest("POST", "/anything", nil)

	postedData = url.Values{}
	postedData.Add("a", "a")

	r.PostForm = postedData

	form = New(r.PostForm)
	has = form.Has("a")

	if !has {
		t.Error("Form shows invalid when required field is populated")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/anything", nil)
	form := New(r.PostForm)
	form.MinLength("a", 2)

	if form.Valid() {
		t.Error("Form shows valid min length for non-existent field")
	}

	// Get error function testing
	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("Should have an error, but did not get one")
	}

	postedData := url.Values{}
	postedData.Add("a", "")

	r.PostForm = postedData

	form = New(r.PostForm)
	form.MinLength("a", 2)

	if form.Valid() {
		t.Error("Form shows valid when less than min length")
	}

	r, _ = http.NewRequest("POST", "/anything", nil)

	postedData = url.Values{}
	postedData.Add("a", "aaaaaaaa")

	r.PostForm = postedData

	form = New(r.PostForm)
	form.MinLength("a", 2)

	if !form.Valid() {
		t.Error("Form shows invalid when field min length is met")
	}

	isError = form.Errors.Get("a")
	if isError != "" {
		t.Error("Should not get error but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("a", "")

	form := New(postedData)
	form.IsEmail("a")

	if form.Valid() {
		t.Error("Form shows valid when no email entered")
	}

	postedData = url.Values{}
	postedData.Add("a", "a@a.com")

	form = New(postedData)
	form.IsEmail("a")

	if !form.Valid() {
		t.Error("Form shows invalid when email entered is valid")
	}
}
