package assets

import "text/template"

func TextTemplate(descriptor string) (*template.Template, error) {
	temp, err := Asset("static/" + descriptor + ".txt")
	if err != nil {
		return nil, err
	}
	t, err := template.New("text").Parse(string(temp))
	if err != nil {
		return nil, err
	}
	return t, nil
}
func PHPTemplate(descriptor string) (*template.Template, error) {
	temp, err := Asset("static/" + descriptor + ".php")
	if err != nil {
		return nil, err
	}
	t, err := template.New("text").Parse(string(temp))
	if err != nil {
		return nil, err
	}
	return t, nil
}
