package adminRouters

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"simple-blog/pkg/models"
	"simple-blog/pkg/utils"
	"time"
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
	matchPass := models.CheckPass(params.Password, user.Password)
	if !matchPass {
		utils.GetMessageError("PASSWORD_ERROR", "Password error", c)
		return
	}

	user.LastLoginTime = time.Now()
	user.LastLoginIp = c.GetHeader("X-Real-IP")
	models.UpdateUser(*user)

	// Save user session
	session := sessions.Default(c)
	session.Set("id", user.Id)
	session.Save()
	utils.GetCommonSuccess(c)
}

func GetCurrentUser(c *gin.Context) {
	id := c.GetInt("id")
	user, err := models.GetUserById(int32(id))
	if err != nil {
		utils.GetCommonError(err, c)
		return
	}
	utils.GetCommonResponse(c, user)
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	utils.GetCommonSuccess(c)
}
