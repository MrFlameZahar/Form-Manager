package main

import (
	"FormManager/internal/config"
	"fmt"
)

func main() {
	cfg, err := config.NewConfigFromFile("./config/local.yaml")
	if err != nil {
		panic(fmt.Sprintf("cant load config %v", err))
	}
	fmt.Println(cfg)

	// Инициализировать логгер

	// Инициализировать репозиторий

	// Инициализировать роутер

	// Запуск сервера
}
