package main

import "gin/initRouter"

func main()  {
	engine := initRouter.SetupRouter()
	engine.Run(":8099")
}