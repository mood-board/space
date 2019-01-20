package account

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	BasePath            = "account"
	LoginPath           = "login"
	MePath              = "me"
	DeleteAccountPath   = "delete-account"
	UpdateAccountPath   = "update-account"
	ActivateAccountPath = "activate-account"
	TestRoute           = "test"
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

type HTMLFacade struct {
	handler *Handler
}

func NewHTMLFacade(handler *Handler) *HTMLFacade {
	return &HTMLFacade{handler}
}

func (facade *HTMLFacade) RegisterRoutes(r *gin.RouterGroup) {

	r.GET(TestRoute, func(c *gin.Context) {
		fmt.Println("Test Routes....")
		c.JSON(200, gin.H{"hello": "Everything works well"})
	})
}
