package helper

import (
	"fmt"
	"html/template"
)

type DropdownInterface interface {
	GenerateDropdownTemplate(string, string) *template.HTML
}

type Option struct {
	Value string
	Name  string
}

type Dropdown []Option

func (d *Dropdown) GenerateDropdownTemplate(key string, selected string) template.HTML {
	t := `<select class='form-control' name='%s'>%s</select>`

	opt := ``
	for _, v := range *d {
		sel := ``
		if v.Value == selected {
			sel = `selected`
		}
		tmp := `<option value='%s' %s>%s</option>`
		opt += fmt.Sprintf(tmp, v.Value, v.Name, sel)
	}

	s := fmt.Sprintf(t, key, opt)
	return template.HTML(s)
}
