package security

import "golang.org/x/crypto/bcrypt"

func Hash(value string) ([]byte, error) {
	//Não passar um valor de custo muito elevado, caso contrário vai gastar muito processamento
	return bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
}

func ComparePassword(hash, uncrypted string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(uncrypted))
}
