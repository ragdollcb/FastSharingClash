package router

import (
	"FastSharingClash/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
)
import "FastSharingClash/router/api/v1"

func InitRouter() *gin.Engine {
	uuid := pkg.GetUUid()
	r := gin.Default()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET(fmt.Sprintf("/%s/profiles.yaml", uuid), v1.GetProfiles)
	}
	return r
}
