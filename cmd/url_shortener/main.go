package main

import (
	"RestAPIURLShortener/internal/config"
	"fmt"
)

const (
	envLoval = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	//в этом файле нам нужен конфиг: библиотека cleanenv
	cfg := config.MustLoad()
	fmt.Println(cfg)
	//потом logger: библиотека slog(import log/slog)
	//инициируем storage:библиотека sqlite
	//иницилизируем router:библиотека chi,"chi render"
	//запускаем сервер
}

// вынесим конфигурацию логгера в отдельную ф-ию,т.к. его установка будет зависить от параметра env
// потому что локально хотим видеть текстовые логги , а на сервере т.е. в окруженние dev или prod
// хотим видеть json на dev уровня дебаггинга, а на prod - json без отладочной информации
// func setupLogger(env string) {
// 	//объявляем логгер
// 	var log *slog.Logger
// 	switch env {

// 	}

// }
