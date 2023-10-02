package main

import (
	"abdullayev13/timeup/internal/config"
	"abdullayev13/timeup/internal/handler"
	"abdullayev13/timeup/internal/handler/middleware"
	"abdullayev13/timeup/internal/models"
	"abdullayev13/timeup/internal/pkg/postgresdb"
	"abdullayev13/timeup/internal/repo"
	"abdullayev13/timeup/internal/service"
	"abdullayev13/timeup/internal/utill"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	db := postgresdb.New()
	models.AutoMigrate(db)

	jwtToken := utill.NewToken(config.JwtSignKey, config.JwtExpiringDuration)

	repos := repo.New(db)
	services := service.New(repos, jwtToken)
	handlers := handler.New(services, jwtToken)
	mw := middleware.New(jwtToken, services)

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

	initApi(r, handlers, mw)

	log.Fatalln(r.Run(":" + config.Port))
}

func initApi(r *gin.Engine, handlers *handler.Handlers, mw *middleware.MW) {
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

	user := v1.Group("/user")
	{
		h := handlers.User
		user.GET("/me", mw.UserIDFromToken, h.UserMe)
		user.PUT("/edit-me", mw.UserIDFromToken, h.EditMe)
		user.DELETE("/delete-me", mw.UserIDFromToken, h.DeleteMe)
	}

	business := v1.Group("/business")
	{
		h := handlers.Business
		business.POST("/create", mw.UserIDFromToken, h.Create)
		business.GET("/get-me", mw.UserIDFromToken, h.GetMe)
		business.DELETE("/delete-me", mw.UserIDFromToken, h.DeleteMe)
	}

	category := v1.Group("/category")
	{
		h := handlers.Category
		category.POST("/create", mw.UserIDFromToken, h.Create)
		category.GET("/get", h.Get)
		category.DELETE("/delete/:id", mw.UserIDFromToken, h.Delete)
	}

	dev := v1.Group("/dev")
	{
		dev.GET("/domain", func(c *gin.Context) {
			config.Domain = c.Query("domain")
		})
	}

}

func init() {
	godotenv.Load("default.env")
	config.LoadVarsFromEnv()
}
