package v1

import "github.com/gin-gonic/gin"
import "FastSharingClash/models"

func GetProfiles(c *gin.Context) {
	profiles := models.ReadProfiles()
	c.String(200, profiles)
}
