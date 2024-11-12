package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Адрес сервера
	serverAddress := "localhost:8080"

	// Подключение к серверу
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Не удалось подключиться к серверу: %v", err)
	}
	defer conn.Close()

	// Ввод сообщения
	fmt.Print("Введите сообщение для сервера: ")
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')

	// Отправка сообщения серверу
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
		return
	}

	// Ожидание ответа от сервера
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Ошибка при чтении ответа от сервера: %v", err)
		return
	}

	// Вывод ответа
	fmt.Printf("Ответ от сервера: %s\n", string(buffer[:n]))
}
