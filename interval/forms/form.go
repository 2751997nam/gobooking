package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors Errors
}

func (form *Form) Valid() bool {
	return len(form.Errors) == 0
}

// init a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		Errors(map[string][]string{}),
	}
}

func (form *Form) Required(req *http.Request, fields ...string) {
	for _, field := range fields {
		form.Has(field, req)
	}
}

func (form *Form) Has(field string, req *http.Request) bool {
	value := strings.TrimSpace(req.Form.Get(field))
	if value == "" {
		form.Errors.Add(field, fmt.Sprintf("the %s field is required", field))
		return false
	}

	return true
}
