package preorderRoutes

import (
	"github.com/gin-gonic/gin"

	preorderControllers "github.com/tyler-mcmullin/go-backend/controllers"
)

func GetPreorderControllers(r *gin.Engine) {
	group := r.Group("/preorders")
	group.GET("/latest", preorderControllers.GetLatest)
	group.GET("/test", preorderControllers.PreorderTest)
}
