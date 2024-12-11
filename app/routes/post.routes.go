package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mofukuma/golang-gorm-postgres/controllers"
	//"github.com/mofukuma/golang-gorm-postgres/middleware" //ログインチェック
)

// PostRouteController は、投稿に関連するルート
type PostRouteController struct {
	postController controllers.PostController
}

func NewRoutePostController(postController controllers.PostController) PostRouteController {
	return PostRouteController{postController}
}

// ルートグループを受け取り、KeyValueの作成、取得、更新、削除の各エンドポイントを設定
func (pc *PostRouteController) PostRoute(rg *gin.RouterGroup) {

	router := rg.Group("save")
	//router.Use(middleware.DeserializeUser()) //ログインチェック。必要があれば。
	router.POST("/", pc.postController.UpsertPost)
	router.POST("/insert", pc.postController.CreatePost)
	router.GET("/", pc.postController.FindPosts)
	router.PUT("/", pc.postController.UpdatePost)
	router.GET("/:postId", pc.postController.FindByUserId)
	router.DELETE("/:postId", pc.postController.DeletePost)
}
