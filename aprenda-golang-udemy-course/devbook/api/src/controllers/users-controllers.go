package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, R *http.Request) {
	w.Write([]byte("Criando usuários"))
}
func ListUsers(w http.ResponseWriter, R *http.Request) {
	w.Write([]byte("Listar usuários"))
}
func FindUser(w http.ResponseWriter, R *http.Request) {
	w.Write([]byte("Buscar usuário"))
}
func UpdateUser(w http.ResponseWriter, R *http.Request) {
	w.Write([]byte("Atualizar Usuário"))
}
func DeleteUser(w http.ResponseWriter, R *http.Request) {
	w.Write([]byte("Deletar Usuário"))
}
