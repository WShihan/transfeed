package api

import (
	"net/http"
	"transfeed/internal/app/form"
	"transfeed/internal/app/model"
	"transfeed/internal/app/store"
	"transfeed/internal/app/web/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @Summary get translator
// @Tags translator
// @Accept json
// @Param id path int true "ID"
// @Success 200 {object} Response "Success"
// @Router /api/v1/translator/get/{id} [get]
func translatorGet(c echo.Context) error {
	user, err := jwt.ParseJWT(c.Get("user"))
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	id := c.Param("id")
	translator, err := store.GetTranslator(user, id)
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	} else {
		return ApiSuccess(c, translator)
	}
}

// @Summary get translator
// @Tags translator
// @Accept json
// @Success 200 {object} Response "Success"
// @Router /api/v1/translator/all [get]
func translatorAll(c echo.Context) error {
	user, err := jwt.ParseJWT(c.Get("user"))
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	trans, err := store.GetTranslators(user)
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	} else {
		return ApiSuccess(c, trans)
	}
}

// @Summary get translator
// @Tags translator
// @Accept json
// @Param id path int true "id"
// @Success 200 {object} Response "Success"
// @Router /api/v1/translator/feeds/{id} [get]
func translatorFeeds(c echo.Context) error {
	id := c.Param("id")
	translator := model.Translator{}
	err := store.DB.Preload("Feeds").First(&translator, id).Error
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	} else {
		return ApiSuccess(c, translator.Feeds)
	}
}

// @Summary add translator
// @Tags translator
// @Accept json
// @Param translatorAddForm body form.TranslatorAddForm true "TranslatorAddForm"
// @Success 200 {object} Response
// @Router /api/v1/translator/add [post]
func translatorAdd(c echo.Context) error {
	user, err := jwt.ParseJWT(c.Get("user"))
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	var formData form.TranslatorAddForm
	err = c.Bind(&formData)
	if err != nil {
		return ApiFailed(c, http.StatusBadRequest, "bad request")
	}
	if err := store.DB.First(&model.Feed{}, "url = ?", formData.Url).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			translator, err := store.CreateTranslator(user, formData)
			if err != nil {
				return ApiFailed(c, 201, err.Error())
			}
			return ApiSuccess(c, translator)

		} else {
			return ApiFailed(c, 201, "translator exists")
		}
	} else {
		return ApiFailed(c, 201, err.Error())
	}

}

// @Summary add translator
// @Tags translator
// @Param data body form.TranslatorUpdateForm true "数据"
// @Success 200 {object} Response
// @Router /api/v1/translator/update [post]
func translatorUpdate(c echo.Context) error {
	user, err := jwt.ParseJWT(c.Get("user"))
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	var formData form.TranslatorUpdateForm
	err = c.Bind(&formData)
	if err != nil {
		return ApiFailed(c, http.StatusBadRequest, user.Username)
	}
	trans, err := store.UpdateTranslator(user, formData)
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	return ApiSuccess(c, trans)

}

// @Summary delete translator
// @Tags translator
// @Param id path int true "ID"
// @Success 200 {object} Response
// @Router /api/v1/translator/delete/{id} [get]
func translatorDel(c echo.Context) error {
	id := c.Param("id")
	user, err := jwt.ParseJWT(c.Get("user"))
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	translator := model.Translator{}
	store.DB.First(&translator, id)
	store.DB.Where("user_id = ?", user.ID).First(&translator, id)
	err = store.DB.Unscoped().Delete(&translator).Error
	if err != nil {
		return ApiFailed(c, 201, "translator not exist")
	}
	if translator.Url == "" {
		return ApiFailed(c, 201, "translator not exist")
	}
	err = store.DB.Unscoped().Delete(&translator).Error
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	return ApiSuccess(c, translator)

}

func TranslatorAttach(protectGroup *echo.Group, publickGroup *echo.Group) {
	protectGroup.POST("/api/v1/translator/add", translatorAdd)
	protectGroup.GET("/api/v1/translator/all", translatorAll)
	protectGroup.GET("/api/v1/translator/get/:id", translatorGet)
	protectGroup.GET("/api/v1/translator/delete/:id", translatorDel)
	protectGroup.GET("/api/v1/translator/feeds/:id", translatorFeeds)
	protectGroup.POST("/api/v1/translator/update", translatorUpdate)
}
