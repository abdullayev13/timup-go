package main

import (
	"abdullayev13/timeup/internal/config"
	"abdullayev13/timeup/internal/handler"
	"abdullayev13/timeup/internal/models"
	"abdullayev13/timeup/internal/pkg/postgresdb"
	"abdullayev13/timeup/internal/repo"
	"abdullayev13/timeup/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	db := postgresdb.New()
	models.AutoMigrate(db)

	repos := repo.New(db)
	services := service.New(repos)
	handlers := handler.New(services)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	r.Static("api/v1/media", "./media")

	initApi(r, handlers)

	log.Fatalln(r.Run(":8080"))
}

func initApi(r *gin.Engine, handlers *handler.Handlers) {
	v1 := r.Group("/api/v1")

	sms := v1.Group("/sms")
	{
		h := handlers.SmsCode
		sms.POST("/send", h.SendSms)
		sms.POST("/verify", h.VerifySmsCode)

		sms.POST("/last-sent-sms", h.LastSentSms)
	}

	auth := v1.Group("/auth")
	{
		h := handlers.Auth
		auth.POST("/register", h.Register)
	}

	dev := v1.Group("/dev")
	{
		dev.GET("/domain", func(c *gin.Context) {
			config.Domain = c.Param("domain")
		})
	}

}

/*






















 */
