package server

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// TemplateRenderer описывает интерфейс для рендеринга HTML-шаблонов.
type TemplateRenderer interface {
	Render(w http.ResponseWriter, name string, data interface{}) error
}

// renderer — конкретная реализация TemplateRenderer.
type renderer struct {
	tmpl *template.Template
}

// NewRenderer создаёт новый рендерер, парся все файлы с расширением .html из указанной директории.
func NewRenderer(templateDir string) (TemplateRenderer, error) {
	pattern := filepath.Join(templateDir, "*.html")
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		return nil, err
	}
	return &renderer{tmpl: tmpl}, nil
}

// Render выполняет шаблон с заданным именем.
func (r *renderer) Render(w http.ResponseWriter, name string, data interface{}) error {
	return r.tmpl.ExecuteTemplate(w, name, data)
}
