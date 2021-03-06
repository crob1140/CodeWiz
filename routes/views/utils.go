package views

import (
	"github.com/crob1140/codewiz-server/log"
	"html/template"
	"net/http"
	"path/filepath"
)


const (
	resourceDirectory = "routes/views/resources"
	templateDirectory = resourceDirectory + "/templates"
)

func render(w http.ResponseWriter, templateName string, data interface{}) {
	path, _ := filepath.Abs(templateDirectory + "/" + templateName)
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Error("Failed to parse template", log.Fields{"template" : templateName, "error" : err})
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Error("Failed to execute template", log.Fields{"template" : templateName, "data" : data, "error" : err})
	}
}