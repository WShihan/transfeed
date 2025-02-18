package api

import (
	"log"
	"net/http"
	"time"
	"transfeed/internal/app/form"
	"transfeed/internal/app/model"
	"transfeed/internal/app/store"
	"transfeed/internal/app/web/jwt"
	"transfeed/internal/app/web/worker"

	"github.com/labstack/echo/v4"
)

// @Summary get feed
// @Tags feed
// @Accept json
// @Param id path int true "ID"
// @Success 200 {object} Response "Success"
// @Router /api/v1/feed/get/{id} [get]
func feedGet(c echo.Context) error {
	id := c.Param("id")
	user, err := jwt.ParseJWT(c.Get("user"))
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	feed := model.Feed{}
	err = store.DB.Preload("Entries").Model(&feed).Where("user_id = ?, id = ?", user.ID, id).Error
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	} else {
		return ApiSuccess(c, feed)
	}
}

// @Summary get feed
// @Tags feed
// @Accept json
// @Param id path int true "ID"
// @Success 200  "feed"
// @Router /feed/rss/{id} [get]
func feedGenRSS(c echo.Context) error {
	id := c.Param("id")
	rss, err := worker.GenFeedRss(id)

	if err != nil {
		return c.String(http.StatusBadGateway, err.Error())
	} else {
		atom, err := rss.ToAtom()
		if err != nil {
			log.Fatal(err)
		}
		// 设置响应头
		c.Response().Header().Set(echo.HeaderContentType, "application/xml")
		return c.String(http.StatusOK, atom)

	}
}

// @Summary get feed
// @Tags feed
// @Accept json
// @Param userid path int true "userid"
// @Success 200 {object} Response "Success"
// @Router /api/v1/feed/all/{userid} [get]
func feedAll(c echo.Context) error {
	user, err := jwt.ParseJWT(c.Get("user"))
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	err = store.DB.Preload("Feeds").Find(user, user.ID).Error
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	} else {
		return ApiSuccess(c, user.Feeds)
	}
}

type FeedPaginationData struct {
	Count   int          `json:"count"`
	Feeds   []model.Feed `json:"feeds"`
	Keyword string       `json:"keyword"`
	Page    int          `json:"page"`
}

// @Summary get feed
// @Tags feed
// @Accept json
// @Param pagination query form.Pagination true "pagination 分页"
// @Success 200 {object} Response "Success"
// @Router /api/v1/feed/public/pagination [get]
func feedPaginationPublic(c echo.Context) error {
	var reqData form.Pagination
	err := c.Bind(&reqData)
	if err != nil {
		return ApiFailed(c, http.StatusBadRequest, "bad request")
	}
	if reqData.Limit == 0 {
		reqData.Limit = 20
	}
	if reqData.Page > 0 {
		reqData.Page--
	}
	offset := reqData.Page * reqData.Limit
	condition := "%" + reqData.Keyword + "%"
	feeds := []model.Feed{}
	availableFeeds := []model.Feed{}
	err = store.DB.Where("Title LIKE ? OR Description LIKE ?", condition, condition).Where("Public = ?", true).Order("ID DESC").Limit(reqData.Limit).Offset(int(offset)).Find(&feeds).Error
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	err = store.DB.Where("Title LIKE ? OR Description LIKE ?", condition, condition).Where("Public = ?", true).Order("ID DESC").Find(&availableFeeds).Error
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}

	return ApiSuccess(c, FeedPaginationData{
		Count:   len(availableFeeds),
		Feeds:   feeds,
		Keyword: reqData.Keyword,
		Page:    reqData.Page,
	})

}

// @Summary get feed
// @Tags feed
// @Accept json
// @Param pagination query form.Pagination true "pagination 分页"
// @Success 200 {object} Response "Success"
// @Router /api/v1/feed/pagination [get]
func feedPagination(c echo.Context) error {
	var reqData form.Pagination
	err := c.Bind(&reqData)
	if err != nil {
		return ApiFailed(c, http.StatusBadRequest, "bad request")
	}
	if reqData.Limit == 0 {
		reqData.Limit = 20
	}
	if reqData.Page > 0 {
		reqData.Page--
	}
	offset := reqData.Page * reqData.Limit
	condition := "%" + reqData.Keyword + "%"
	feeds := []model.Feed{}
	availableFeeds := []model.Feed{}
	user, err := jwt.ParseJWT(c.Get("user"))
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	err = store.DB.Where("user_id = ?", user.ID).Where("Title LIKE ? OR Description LIKE ?", condition, condition).Order("ID DESC").Limit(reqData.Limit).Offset(int(offset)).Find(&feeds).Error
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	err = store.DB.Where("user_id = ?", user.ID).Where("Title LIKE ? OR Description LIKE ?", condition, condition).Order("ID DESC").Find(&availableFeeds).Error
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}

	return ApiSuccess(c, FeedPaginationData{
		Count:   len(availableFeeds),
		Feeds:   feeds,
		Keyword: reqData.Keyword,
		Page:    reqData.Page,
	})

}

// @Summary add feed
// @Tags feed
// @Accept json
// @Param feedAddForm body form.FeedAddForm true "FeedAddForm"
// @Success 200 {object} Response
// @Router /api/v1/feed/add [POST]
func feedAdd(c echo.Context) error {
	user, err := jwt.ParseJWT(c.Get("user"))
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	var reqData form.FeedAddForm
	err = c.Bind(&reqData)
	if err != nil {
		return ApiFailed(c, http.StatusBadRequest, "bad request")
	}
	feed, err := store.CreateFeed(reqData)
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	worker.PostProcessEntries(feed, feed.Entries, 10)
	err = store.DB.Model(&user).Association("Feeds").Append(feed)
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	refreshTime := time.Now()
	feed.RefreshTime = &refreshTime

	store.DB.Save(feed)
	return ApiSuccess(c, feed)

}

// @Summary add feed
// @Tags feed
// @Param data body form.FeedUpdateForm true "数据"
// @Success 200 {object} Response
// @Router /api/v1/feed/update [get]
func feedUpdate(c echo.Context) error {
	user, err := jwt.ParseJWT(c.Get("user"))
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	var reqData form.FeedUpdateForm
	err = c.Bind(&reqData)
	if err != nil {
		return ApiFailed(c, http.StatusBadRequest, "bad request")
	}

	feed, err := store.UpdateFeed(*user, reqData)
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	return ApiSuccess(c, feed)

}

// @Summary add feed
// @Tags feed
// @Param id path int true "ID"
// @Success 200 {object} Response
// @Router /api/v1/feed/refresh/{id} [get]
func feedRefresh(c echo.Context) error {
	id := c.Param("id")
	user, err := jwt.ParseJWT(c.Get("user"))
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	feed := model.Feed{}
	store.DB.Preload("Translator").Preload("Entries").Where("user_id = ?", user.ID).First(&feed, id)
	_, err = worker.RefreshFeed(&feed)
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	return ApiSuccess(c, feed)

}

// @Summary delete feed
// @Tags feed
// @Param id path int true "ID"
// @Success 200 {object} Response
// @Router /api/v1/feed/delete/{id} [get]
func feedDel(c echo.Context) error {
	id := c.Param("id")
	user, err := jwt.ParseJWT(c.Get("user"))
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	feed := model.Feed{}
	store.DB.First(&feed, id)
	store.DB.Where("user_id = ?", user.ID).First(&feed, id)
	err = store.DB.Unscoped().Delete(&feed).Error
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	if feed.Url == "" {
		return ApiFailed(c, 201, "feed not exist")
	}
	err = store.DB.Select("Entries").Unscoped().Delete(&feed).Error
	if err != nil {
		return ApiFailed(c, 201, err.Error())
	}
	return ApiSuccess(c, feed)

}

func FeedAttach(protectGroup *echo.Group, publickGroup *echo.Group) {
	publickGroup.GET("/rss/:id", feedGenRSS)
	publickGroup.GET("/api/v1/feed/public/pagination", feedPaginationPublic)

	protectGroup.GET("/api/v1/feed/pagination", feedPagination)
	protectGroup.GET("/api/v1/feed/refresh/:id", feedRefresh)
	protectGroup.POST("/api/v1/feed/add", feedAdd)
	protectGroup.GET("/api/v1/feed/all", feedAll)
	protectGroup.GET("/api/v1/feed/get/:id", feedGet)
	protectGroup.GET("/api/v1/feed/delete/:id", feedDel)
	protectGroup.POST("/api/v1/feed/update", feedUpdate)
}
