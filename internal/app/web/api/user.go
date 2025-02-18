package api

import (
	"net/http"
	"time"
	"transfeed/internal/app/form"
	"transfeed/internal/app/model"
	"transfeed/internal/app/store"
	"transfeed/internal/app/web/config"
	"transfeed/internal/util"

	jwter "transfeed/internal/app/web/jwt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type userinfoData struct {
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
}

// @Summary get user
// @Description get user
// @Tags user
// @Accept json
// @Param id path int true "ID"
// @Success 200 {object} Response "Success"
// @Router /api/v1/user/info [get]
func UserInfo(c echo.Context) error {
	user, err := jwter.ParseJWT(c.Get("user"))
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	return ApiSuccess(c, userinfoData{Username: user.Username, Admin: user.Admin})
}

// @Summary register user
// @Tags user
// @Accept json
// @Param user body form.UserRegisterForm true "UserRegisterForm"
// @Success 200 {object} Response
// @Router /api/v1/user/register [POST]
func UserRegister(c echo.Context) error {
	administrator, err := jwter.ParseJWT(c.Get("user"))
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}

	if !administrator.Admin {
		return ApiFailed(c, 201, "permission denied")
	}

	u := new(form.UserRegisterForm)
	if err := c.Bind(u); err != nil {
		return err
	}
	user := model.User{}
	store.DB.Where("username = ?", u.Username).Find(&user)
	if user.Username == "" {
		hasedPass, err := util.HashMessage(u.Password)
		if err != nil {
			return ApiFailed(c, 201, err.Error())
		}
		user.Username = u.Username
		user.Password = hasedPass
		user.Admin = u.Admin
		err = store.DB.Create(&user).Error
		if err != nil {
			return ApiFailed(c, 201, err.Error())
		}
		return ApiSuccess(c, user)
	} else {
		return ApiFailed(c, 201, "user exist")
	}

}

type LoginResult struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
}

// @Summary  user login
// @Tags user
// @Param user body form.UserLoginForm true "UserLoginForm"
// @Success 200 {object} Response
// @Router /api/v1/user/login [POST]
func UserLogin(c echo.Context) error {
	u := new(form.UserLoginForm)
	if err := c.Bind(u); err != nil {
		return err
	}
	config := c.Get("config").(config.Config)
	user := model.User{}
	err := store.DB.Where("username = ?", u.Username).Find(&user).Error
	if err == nil && user.Username != "" {
		if util.ValidateHash(user.Password, u.Password) != nil {
			return ApiFailed(c, 201, "password incorrect")
		}
		claims := jwt.MapClaims{
			"id":       user.ID,
			"username": user.Username,
			"admin":    user.Admin,
			"exp":      time.Now().Add(time.Hour * 5).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte(config.JWTScrect))
		if err != nil {
			return err
		}
		user.Token = &t
		store.DB.Save(&user)
		return ApiSuccess(c, LoginResult{Token: t, Username: user.Username, Admin: user.Admin})
	} else {
		return ApiFailed(c, 201, "user not exist")
	}

}

// @Summary  user login
// @Tags user
// @Success 200 {object} Response
// @Router /api/v1/user/logout [GET]
func UserLogout(c echo.Context) error {
	user, err := jwter.ParseJWT(c.Get("user"))
	if err == nil && user.Username != "" {
		user.Token = nil
		store.DB.Save(&user)
		return ApiSuccess(c, nil)
	} else {
		return ApiFailed(c, 201, "user not exist")
	}

}

// @Summary add user
// @Description add user
// @Tags user
// @Param data body form.UserUpdateForm true "data"
// @Success 200 {object} Response
// @Router /api/v1/user/update [post]
func UserUpdate(c echo.Context) error {
	user, err := jwter.ParseJWT(c.Get("user"))
	if err != nil {
		return ApiFailed(c, http.StatusForbidden, err.Error())
	}
	var formData form.UserUpdateForm
	err = c.Bind(&formData)
	if err != nil {
		return ApiFailed(c, http.StatusBadRequest, user.Username)
	}
	hasedPass, err := util.HashMessage(formData.PasswordOld)
	if err != nil {
		return ApiFailed(c, 201, "password error")
	}

	if util.ValidateHash(hasedPass, formData.PasswordOld) != nil {
		return ApiFailed(c, 201, "password incorrect")
	}

	newHashedPass, err := util.HashMessage(formData.Password)
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}

	user.Password = newHashedPass
	user.Token = nil
	err = store.DB.Save(&user).Error

	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	return ApiSuccess(c, nil)

}

// @Summary delete user
// @Description delete user
// @Tags user
// @Param id path int true "ID"
// @Success 200 {object} Response
// @Router /api/v1/user/delete/{id} [get]
func UserDel(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")

}

func UserAttach(protectGroup *echo.Group, publickGroup *echo.Group) {
	publickGroup.POST("/api/v1/user/login", UserLogin)

	protectGroup.POST("/api/v1/user/register", UserRegister)
	protectGroup.GET("/api/v1/user/logout", UserLogout)
	protectGroup.GET("/api/v1/user/info", UserInfo)
	protectGroup.GET("/api/v1/user/delete/:id", UserDel)
	protectGroup.POST("/api/v1/user/update", UserUpdate)
}
