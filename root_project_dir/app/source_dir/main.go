package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var client *redis.Client

func main() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Адрес Redis-сервера
		Password: "",               // Пароль Redis-сервера (если требуется)
		DB:       0,                // Индекс базы данных Redis

	})

	router := mux.NewRouter()
	router.HandleFunc("/get_key/{key}", GetValue).Methods("GET")
	router.HandleFunc("/set_key/{key}/{value}", SetValue).Methods("SET")
	router.HandleFunc("/del_key/{key}", DeleteValue).Methods("DELETE")
	router.HandleFunc("/{any}", Forbidden).Methods("GET")

	// certFile := "redis_user.crt"
	// keyFile := "redis_user_private.key"

	//Загрузка сертификата и закрытого ключа
	// cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	// if err != nil {
	// 	log.Fatalf("Ошибка загрузки сертификата: %s", err)
	// }

	// //Создание конфигурации TLS
	// tlsConfig := &tls.Config{
	// 	Certificates: []tls.Certificate{cert},
	// }

	//Загрузка корневого сертификата
	// caCertFile := "redis_ca.pem"
	// caCert, err := os.ReadFile(caCertFile)
	// if err != nil {
	// 	log.Fatalf("Ошибка загрузки корневого сертификата: %s", err)
	// }

	//Добавление корневого сертификата к конфигу TLS

	// tlsConfig.RootCAs = x509.NewCertPool()
	// tlsConfig.RootCAs.AppendCertsFromPEM(caCert)

	//Создание сервера HTTP с поддержкой TLS
	server := &http.Server{
		Addr:    "localhost:8000",
		Handler: router,
		// TLSConfig: tlsConfig,
	}

	log.Println("Сервер запущен на порту :8000 с поддержкой")
	log.Fatal(server.ListenAndServe())
}

// GetValue возвращает значение по указанному ключу
func GetValue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["get_key"]

	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		http.Error(w, "Ключ не найден", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Ошибка сервера Redis", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Значение ключа '%s': %s", key, val)
}

// SetValue создает или перезаписывает пару ключ-значение
func SetValue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["set_key"]
	value := params["value"]

	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		http.Error(w, "Ошибка сервера Redis", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Значение '%s' успешно сохранено для ключа '%s'", value, key)
}

// DeleteValue удаляет пару ключ-значение
func DeleteValue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["del_key"]

	affected, err := client.Del(ctx, key).Result()
	if err != nil {
		http.Error(w, "Ошибка сервера Redis", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Ключ '%s' успешно удален. Количество удаленных значений: %d", key, affected)
}

// Forbidden возвращает ошибку 403 для всех других запросов
func Forbidden(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Доступ запрещен 403", http.StatusForbidden)
}
