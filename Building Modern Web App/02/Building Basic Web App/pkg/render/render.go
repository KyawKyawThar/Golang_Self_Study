package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tpl string) {
	parsedTemplated, _ := template.ParseFiles("./templates/" + tpl)
	err := parsedTemplated.Execute(w, nil)

	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}
