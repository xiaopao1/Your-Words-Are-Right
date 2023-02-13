package Your_Words_Are_Right

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	initRouter(r)
	r.GET("ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong"})
	})

	err := r.Run()
	if err != nil {
		return
	}

}
