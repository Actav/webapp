package server

import (
	"myapp/logger" // Импортируем пакет логирования
	"net/http"
)

// PageData — структура для передачи данных в шаблоны.
type PageData struct {
	Title string
}

// Server описывает интерфейс веб-сервера.
type Server interface {
	// Start запускает сервер на указанном адресе.
	Start(addr string) error
}

// server — конкретная реализация интерфейса Server.
type server struct {
	mux      *http.ServeMux
	renderer TemplateRenderer
	logger   *logger.Logger
}

// NewServer создаёт новый сервер, используя рендерер шаблонов, загруженный из директории templateDir.
func NewServer(templateDir, logFile string) (Server, error) {
	// Создаём логгер
	l, err := logger.NewLogger(logFile)
	if err != nil {
		return nil, err
	}

	// Создаём рендерер шаблонов
	r, err := NewRenderer(templateDir)
	if err != nil {
		return nil, err
	}

	// Создаём сервер
	mux := http.NewServeMux()

	s := &server{
		mux:      mux,
		renderer: r,
		logger:   l,
	}

	// Регистрируем маршруты
	s.Routes()

	return s, nil
}

// Start запускает HTTP-сервер на указанном адресе.
func (s *server) Start(addr string) error {
	// Логируем запуск сервера
	s.logger.Printf("Сервер запущен на %s", addr)

	return http.ListenAndServe(addr, s.mux)
}
