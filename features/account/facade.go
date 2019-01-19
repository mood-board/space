package account

import "fmt"

const (
	BasePath            = "account"
	LoginPath           = "login"
	MePath              = "me"
	DeleteAccountPath   = "delete-account"
	UpdateAccountPath   = "update-account"
	ActivateAccountPath = "activate-account"
)

func BuildPath(path string) string {
	return fmt.Sprintf("/%s/%s", BasePath, path)
}

func Paths() interface{} {
	return struct {
		Base            string
		Login           string
		Me              string
		DeleteAccount   string
		UpdateAccount   string
		ActivateAccount string
	}{
		BasePath,
		BuildPath(LoginPath),
		BuildPath(MePath),
		BuildPath(DeleteAccountPath),
		BuildPath(UpdateAccountPath),
		BuildPath(ActivateAccountPath),
	}
}

func GetAccountMePath() string {
	return BuildPath(MePath)
}
