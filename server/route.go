package server

import (
	"log"
	"net/http"
)

// Routes определяет все маршруты веб-сервера.
func (s *server) Routes() {
	// Главная страница
	s.mux.HandleFunc("/", s.handleIndex)

	// Статические файлы (CSS, JS)
	s.routeStaticFiles()
}

// handleIndex обрабатывает запросы на главную страницу.
func (s *server) handleIndex(w http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Главная страница"}

	// Логирование запроса
	s.logger.Printf("Запрос на главную страницу от %s", r.RemoteAddr)

	// Выполняем шаблон "layout", внутри которого определён блок "content"
	if err := s.renderer.Render(w, "layout", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Ошибка рендеринга шаблона:", err)
	}
}

// Функция для роутинга статических файлов
func (s *server) routeStaticFiles() {
	s.mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	s.mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	s.mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
}
