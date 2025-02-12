package main

import (
	"log"
	"myapp/server"
)

func main() {
	// Создаём сервер, передавая путь к директории с шаблонами и путь для логирования
	srv, err := server.NewServer("templates", "server.log")
	if err != nil {
		log.Fatalf("Ошибка при создании сервера: %v", err)
	}

	// Запускаем сервер на порту 8080
	if err := srv.Start(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
