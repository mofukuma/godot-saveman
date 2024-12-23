package controllers

import (
	"net/http"
	"strings"

	//"time"

	"github.com/gin-gonic/gin"
	"github.com/mofukuma/golang-gorm-postgres/models"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB}
}

func (pc *PostController) CreatePost(ctx *gin.Context) {
	//currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreatePostRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newPost := models.Post{
		Userid: payload.Userid,
		K:      payload.K,
		V:      payload.V,
	}

	result := pc.DB.Create(&newPost)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "ok"})
}

func (pc *PostController) UpdatePost(ctx *gin.Context) {
	postId := ctx.Param("userid")
	key := ctx.Param("k")

	var payload *models.CreatePostRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedPost models.Post
	result := pc.DB.First(&updatedPost, "userid = ? and k = ?", postId, key)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not exists"})
		return
	}
	postToUpdate := models.Post{
		Userid: payload.Userid,
		K:      payload.K,
		V:      payload.V,
	}

	pc.DB.Model(&updatedPost).Updates(postToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (pc *PostController) UpsertPost(ctx *gin.Context) {
	var payload *models.CreatePostRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var existingPost models.Post
	result := pc.DB.First(&existingPost, "userid = ? and k = ?", payload.Userid, payload.K)

	if result.Error != nil {
		// 投稿が存在しない場合は新規作成
		newPost := models.Post{
			Userid: payload.Userid,
			K:      payload.K,
			V:      payload.V,
		}

		createResult := pc.DB.Create(&newPost)
		if createResult.Error != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": createResult.Error.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"status": "ok"})
	} else {
		// 投稿が存在する場合は更新
		existingPost.V = payload.V

		updateResult := pc.DB.Save(&existingPost)
		if updateResult.Error != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": updateResult.Error.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func (pc *PostController) FindByUserId(ctx *gin.Context) {
	postId := ctx.Param("userid")

	var post models.Post
	result := pc.DB.First(&post, "userid = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": post})
}

func (pc *PostController) FindPosts(ctx *gin.Context) {
	var payload *models.GetPostRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var post models.Post
	result := pc.DB.First(&post, "userid = ? and k = ?", payload.Userid, payload.K)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"userid": post.Userid, "k": post.K, "v": post.V})
}

func (pc *PostController) DeletePost(ctx *gin.Context) {
	postId := ctx.Param("userid")
	key := ctx.Param("k")

	result := pc.DB.Delete(&models.Post{}, "userid = ? and k = ?", postId, key)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
