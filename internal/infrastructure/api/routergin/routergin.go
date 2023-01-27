package routergin

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ninja-dark/QSOFT-task/internal/infrastructure/api/handler"
)

type RouterGin struct {
	*gin.Engine
	hs *handler.Handlers
}

func NewRouterGin(hs *handler.Handlers) *RouterGin {
	r := gin.Default()
	ret := &RouterGin{
		hs: hs,
	}
	r.GET("/when/:year", ret.GetCountDays)
	ret.Engine = r
	return ret
}

func (rt *RouterGin) GetCountDays(c *gin.Context) {
	inputYear, err := strconv.Atoi(c.Param("year"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	d, err := rt.hs.GetCountDays(c.Request.Context(), inputYear)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "%s%d", d.Message, d.Count)
}