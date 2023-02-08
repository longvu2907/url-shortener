package controllers

import (
	"net/http"
	"os"
	"url-shortener/models"
	"url-shortener/utils/shortener"
	"url-shortener/utils/token"

	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	Url string `json:"url" binding:"required"`
}

type UrlCreationResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// @Summary Create short url
// @Description Create short url
// @Tags url-shortener
// @Accept json
// @Produce json
// @Param data body UrlCreationRequest true "Create short url data"
// @Success 200 {object} UrlCreationResponse
// @Router /url-shortener/create-short-url [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func CreateShortUrl(c *gin.Context) {
	var input UrlCreationRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := token.ExtractUID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	su := models.ShortUrl{
		Url:      input.Url,
		Uid:      uid,
		ShortUrl: os.Getenv("SV_HOST") + shortener.GenerateShortUrl(input.Url, uid),
	}
	_, err = su.CreateShortUrl()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := gin.H{
		"short_url": su.ShortUrl,
	}
	c.JSON(200, UrlCreationResponse{
		Message: "short url created successfully",
		Data:    res,
	})
}

// @Summary Redirect short url
// @Description Redirect short url
// @Param shortUrl path string true "Short Url"
// @Router /{shortUrl}  [get]
func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	url, err := models.GetUrl(shortUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusPermanentRedirect, url)
}

type GetUrlResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// @Summary Get short url list
// @Description Get short url list
// @Tags url-shortener
// @Accept json
// @Produce json
// @Success 200 {object} GetUrlResponse
// @Router /url-shortener/url-list  [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func ListShortUrl(c *gin.Context) {
	uid, err := token.ExtractUID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	urlList, err := models.GetUrlListByUid(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := gin.H{
		"url_list": urlList,
	}
	c.JSON(http.StatusOK, GetUrlResponse{
		Message: "get url list successfully",
		Data:    res,
	})
}
