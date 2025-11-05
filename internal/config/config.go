package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// парсинг конфига
type Config struct {
	//полностью соответствует yaml файлу
	Env string `yaml:"env" env-default:"local"  ` //окружение
	//env-requared потому что хотим гарантированно иметь заданный определенный путь до storage
	StoragePath string `yaml:"storage_path" env-requared:"true"` //путь к хранилищу
	//встраиваем объект HTTPServer в общую структуру конфига
	HTTPServer `yaml:"http_server"`
}

// http сервер , который будет в виде отдельного объекта,
// потому что у него есть вложенные параметры
type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

// ф-ия кот-аяфайл с конфига и создаст, и заполнит объект конфиг
// используется Must, когда ф-ия вместо возврата ошибки сразу паникует
func MustLoad() *Config {
	//разьираемся откуда считывать файл с конфигом
	//берем его из переменной окружения
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		//то уронем приложение в фаталом
		//используем стандартный логгер, потому что не инициализировали еще наш
		log.Fatal("CONFIG_PATH is not set")
	}

	//проверяем существует ли файл по этому пути
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}
	var cfg Config
	//считываем файл по пути, кот-ый указали
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}
	return &cfg
}
