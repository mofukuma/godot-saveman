package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mofukuma/golang-gorm-postgres/controllers"
	"github.com/mofukuma/golang-gorm-postgres/middleware"
)

// ユーザ認証用

// AuthRouteController は認証に関連するルート
type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

// ginの謎設定ここまで

// AuthRoute は認証に関連するルートを設定します。
// @param rg *gin.RouterGroup ルーターグループ
// @return なし
func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/refresh", rc.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(), rc.authController.LogoutUser)
}
