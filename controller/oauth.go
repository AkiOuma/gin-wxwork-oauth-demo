package controller

import (
	"errors"
	"fmt"
	"net/http"
	"oauth/dto"
	"oauth/utils"

	"github.com/gin-gonic/gin"
)

func OAuthController(app *gin.Engine) {
	app.GET("/oauth/redirect", OAuth)
}

func OAuth(c *gin.Context) {
	user := dto.UserInfoDTO{}
	code := c.Query("code")
	if code == "" {
		c.Status(http.StatusInternalServerError)
		return
	}
	token, err := GetAccessToken()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s", token.AccessToken, code)
	result, err := http.Get(url)
	if err != nil {
		return
	}
	err = utils.JsonHandler(result, &user)
	if user.ErrCode != 0 {
		err = errors.New(user.Errmsg)
	}
	if err != nil {
		return
	}
	c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/home?name=%s", user.UserId))
}

func GetAccessToken() (result dto.AccessTokenDTO, err error) {
	corpid := "corpid of you company"
	corpsecret := "secret of your app"
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corpid, corpsecret)
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	err = utils.JsonHandler(response, &result)
	if result.ErrCode != 0 {
		err = errors.New(result.Errmsg)
	}
	return
}
