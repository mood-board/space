package account

import (
	"context"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/spaceship/common"
	"github.com/ofonimefrancis/spaceship/common/log"
)

const (
	BasePath            = "account"
	LoginPath           = "login"
	RegisterPath        = "signup"
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

type facade struct {
	accountHandler *Handler
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	TokenString string `json:"tokenString"`
	User        *User  `json:"user"`
}

func NewHTMLFacade(handler *Handler) *facade {
	return &facade{handler}
}

func (facade *facade) RegisterRoutes(r *gin.RouterGroup) {
	r.GET(TestRoute, func(c *gin.Context) {
		fmt.Println("Test Routes....")
		c.JSON(200, gin.H{"hello": "Everything works well"})
	})

	r.POST(LoginPath, func(c *gin.Context) {
		var requestBody LoginRequest

		if err := c.Bind(&requestBody); err != nil {
			log.Info("[LOGIN] Error decoding json into payload struct")
			c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "There was an issue during login"})
			return
		}
		ds := facade.accountHandler.datastore.OpenSession(context.Background())

		user, err := ds.FindUserByEmail(requestBody.Email)
		if err != nil {
			log.Info("Account doesn't exists")
			c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Account doesn't exists"})
			return
		}

		if !ds.IsValidLoginCredentials(requestBody.Email, requestBody.Password) {
			log.Info("Incorrect Email or Password...")
			c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Incorrect Email or password"})
			return
		}

		if !ds.IsVerifiedAccount(requestBody.Email) {
			log.Info("Verify your account and try again")
			c.JSON(http.StatusNotAcceptable, common.ErrorResponse{Message: "Verify your account and try again."})
			return
		}

		token, tokenString, err := common.GetTokenAuth().Encode(jwt.MapClaims{})
		if err != nil {
			log.Debug("There was an error encoding jwt claims")
			c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "There was an error, try again later..."})
			return
		}

		log.Info(token, tokenString)
		c.JSON(http.StatusOK, LoginResponse{TokenString: tokenString, User: user})
	})

	r.POST(RegisterPath, func(c *gin.Context) {
		var user User
		if err := c.Bind(&user); err != nil {
			log.Info("[REGISTER] Error decoding json into payload struct")
			c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "There was an issue during registration"})
			return
		}

		ds := facade.accountHandler.datastore.OpenSession(context.Background())
		if ds.AccountExists(user.Email) {
			log.Info("Account already exists")
			c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Account already exists"})
			return
		}

		passwordHash, err := NewPasswordHash(user.Password)
		if err != nil {
			log.Info("Error hashing password")
			c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Something went wrong. Try again later"})
			return
		}
		user.Password = ""
		user.PasswordHash = *passwordHash
		user.ID = bson.NewObjectId()
		user.CreatedAt = time.Now()
		user.IsActive = false
		user.IsVerified = false
		user.AvatarURL = "/static/img/default_avatar.png"

		userAccount, err := ds.CreateUser(user)
		if err != nil {
			log.Info("[REGISTER] Error inserting new user")
			c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Error Creating new account"})
			return
		}

		_, tokenString, err := common.GetTokenAuth().Encode(jwt.MapClaims{
			"id":          userAccount.ID,
			"email":       userAccount.Email,
			"name":        userAccount.Name,
			"is_verified": userAccount.IsVerified,
			"created_at":  userAccount.CreatedAt,
		})
		if err != nil {
			log.Info("[REGISTER] Error encoding user claims")
			c.JSON(http.StatusBadRequest, common.ErrorResponse{Message: "Error creating new account"})
			return
		}

		c.JSON(http.StatusOK, LoginResponse{TokenString: tokenString, User: userAccount})

	})

}
