package api

import (
	"net/http"

	"github.com/NuttapolCha/simple-covid-grouping/custom_error"
	"github.com/gin-gonic/gin"
)

func init() {
	apis = append(apis, API{
		Method:      http.MethodGet,
		Path:        "/covid/summary",
		HandlerFunc: covidSummaryHandler,
	})
}

func covidSummaryHandler(c *gin.Context) {
	service.Logger.Debugf("client IP: %s", c.ClientIP())
	ctx := c.Request.Context()

	summary, err := service.GetCovidCasesSummary(ctx)
	if err != nil {
		if uErr, ok := err.(*custom_error.UserError); ok {
			c.AbortWithStatus(uErr.StatusCode)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}
	c.JSON(http.StatusOK, summary)
}
