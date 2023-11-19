package main

import (
 "embed"
 "net/http"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
 // Cria um sistema de arquivos a partir dos arquivos incorporados.
 fileServer := http.FileServer(http.FS(staticFiles))

 // Registra o sistema de arquivos como manipulador HTTP.
 http.Handle("/", fileServer)

 // Inicia o servidor HTTP na porta 8080
 http.ListenAndServe(":8080", nil)
}