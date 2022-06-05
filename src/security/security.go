package security

import "golang.org/x/crypto/bcrypt"

//pega uma string "senha" e coloca hash nela
//o hash não tem como reverter para normal enquanto que a criptografia
//tem como ser descriptografada
func Hash(password string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//compara uma senha e um hash e retorna se elas sáo iguais
func VerifyPassword(passwordHashed string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
}
