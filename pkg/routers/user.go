package routers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"simple-blog/pkg/models"
	"simple-blog/pkg/utils"
)

type loginReqModel struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	params := loginReqModel{}
	c.ShouldBindJSON(&params)
	user, err := models.GetUserByName(params.Name)
	if err != nil {
		utils.GetCommonError(err, c)
		return
	}
	if user == nil {
		utils.GetMessageError("ACCOUNT_NOT_EXIST", "Account not exist", c)
		return
	}
	fmt.Println(params.Password, user.Password)
	matchPass := models.CheckPass(params.Password, user.Password)
	if !matchPass {
		utils.GetMessageError("PASSWORD_ERROR", "Password error", c)
		return
	}

	// Save user session
	session := sessions.Default(c)
	session.Set("name", user.Name)
	session.Save()
	utils.GetCommonSuccess(c)
}

func GetCurrentUser(c *gin.Context) {
	session := sessions.Default(c)
	name := session.Get("name").(string)
	user, err := models.GetUserByName(name)
	if err != nil {
		utils.GetCommonError(err, c)
		return
	}
	utils.GetCommonResponse(c, user)
}
