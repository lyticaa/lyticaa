package helpers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var (
	cwd, _    = os.Getwd()
	baseTmpl  = "app"
	baseFiles = []string{
		filepath.Join(cwd, "./web/dist/"+baseTmpl+".html"),
		filepath.Join(cwd, "./web/templates/partials/_nav.gohtml"),
		filepath.Join(cwd, "./web/templates/partials/_footer.gohtml"),
	}
)

func RenderTemplate(w http.ResponseWriter, templates []string, data interface{}) {
	files := templateList(templates)
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, baseTmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func templateList(fileList []string) []string {
	var container []string
	container = append(container, baseFiles...)

	for _, file := range fileList {
		container = append(container, filepath.Join(cwd, "./web/templates/"+file+".gohtml"))
	}

	return container
}
