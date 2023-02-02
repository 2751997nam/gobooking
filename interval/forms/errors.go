package forms

type Errors map[string][]string

func (err Errors) Add(field, message string) {
	err[field] = append(err[field], message)
}

func (err Errors) First(field string) string {
	values := err[field]
	if len(values) == 0 {
		return ""
	}

	return values[0]
}

func (err Errors) Get(field string) []string {
	values := err[field]
	if len(values) == 0 {
		return make([]string, 0)
	}

	return values
}
