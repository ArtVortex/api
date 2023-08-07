package https

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer(port string) error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/query", GraphQLQuery)
	e.GET("/playground", GraphQLPlayground)
	e.POST("/api/v1/create-prediction", CreatePrediction)
	e.GET("/api/v1/predictions/:predictionId", GetPrediction)
	e.GET("/api/v1/predictions", ListPredictions)

	//ipfs apis just for demo
	//eg. http://localhost:8000/add-ipfs/https://placekitten.com/200/300
	e.GET("/add-ipfs/:url", IPFSAddFromURLHandler)
	//copy the cid returned from "/add-ipfs/:url"
	//eg. http://localhost:8000/get-ipfs/:cid
	e.GET("/get-ipfs/:cid", IPFSGetHandler)

	return e.Start(":" + port)

}
