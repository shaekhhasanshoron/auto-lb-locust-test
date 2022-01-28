package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/locust-test/config"
	"strconv"
)

var count =1
// Router api/v1 base router
func Router(g *echo.Group) {
	BlockRouter(g.Group("/blocks"))
}

func BlockRouter(g *echo.Group) {
	g.GET("/", generateRandomBlocks)
}
func generateRandomBlocks(context echo.Context) error {
	config.InitChain()
	config.Chain.AddBlock(strconv.Itoa(count)+" Block after Genesis")
	return context.JSON(200,config.Chain.Blocks)
}
