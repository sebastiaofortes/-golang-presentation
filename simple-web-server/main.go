package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Configuração do roteamento
	http.HandleFunc("/", handler)

	// Inicia o servidor na porta 8080
	port := 8080
	fmt.Printf("Servidor web iniciado em http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}

// Função de manipulação de requisição
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Olá, mundo!")
}