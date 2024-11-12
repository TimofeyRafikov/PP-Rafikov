package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		// Чтение данных от клиента
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Ошибка при чтении данных: %v", err)
			return
		}

		// Вывод полученного сообщения
		message := string(buffer[:n])
		fmt.Printf("Получено сообщение: %s\n", message)

		// Отправка подтверждения клиенту
		_, err = conn.Write([]byte("Сообщение получено"))
		if err != nil {
			log.Printf("Ошибка при отправке ответа: %v", err)
			return
		}
	}
}

func startServer(address string) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
	defer ln.Close()

	var wg sync.WaitGroup

	// Для graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Сервер запущен. Ожидание подключений...")

	for {
		// Ожидание нового подключения
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-signalChan:
				// Если сигнал завершения, завершить сервер
				fmt.Println("Сервер завершает работу...")
				return
			default:
				log.Printf("Ошибка при подключении: %v", err)
				continue
			}
		}

		// Обработка соединения в горутине
		wg.Add(1)
		go handleConnection(conn, &wg)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
}

func main() {
	// Указание адреса для прослушивания
	address := ":8080"
	startServer(address)
}
