package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mofukuma/golang-gorm-postgres/controllers"
	"github.com/mofukuma/golang-gorm-postgres/middleware"
)

// UserRouteController は、ユーザに関連するルート
type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("users")
	router.GET("/me", middleware.DeserializeUser(), uc.userController.GetMe)
}
