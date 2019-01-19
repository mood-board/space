package common

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/ofonimefrancis/spaceship/common/must"
)

func GenerateRandomToken() string {
	// 128 bits / 16 bytes as per:
	//  https://www.owasp.org/index.php/Session_Management_Cheat_Sheet#Session_ID_Length
	token := make([]byte, 16)
	must.DoF(func() error {
		_, err := rand.Read(token)
		return err
	})
	return hex.EncodeToString(token)
}
