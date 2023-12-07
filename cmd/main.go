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
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	println("Started...")
	db := postgresdb.New()
	go models.AutoMigrate(db)

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

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Server is running!!!",
		})
	})

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
		auth.POST("/log-out", mw.UserIDFromToken, h.LogOut)
	}

	user := v1.Group("/user")
	{
		h := handlers.User
		user.GET("/me", mw.UserIDFromToken, h.UserMe)
		user.PUT("/edit-me", mw.UserIDFromToken, h.EditMe)
		user.PUT("/edit-photo", mw.UserIDFromToken, h.EditPhoto)
		user.DELETE("/delete-me", mw.UserIDFromToken, h.DeleteMe)
	}

	business := v1.Group("/business")
	{
		h := handlers.Business
		business.POST("/create", mw.UserIDFromToken, h.Create)
		business.GET("/get-me", mw.UserIDFromToken, h.GetMe)
		business.GET("/profile/:id", mw.UserIDFromToken, h.GetProfile)
		business.GET("/get-by-category/:id", mw.UserIDFromToken, h.GetByCategory)
		business.PUT("/update-me", mw.UserIDFromToken, h.UpdateMe)
		business.DELETE("/delete-me", mw.UserIDFromToken, h.DeleteMe)
		// about following
		fh := handlers.Following
		business.POST("/:id/follow", mw.UserIDFromToken, fh.Create)
		business.DELETE("/:id/unfollow", mw.UserIDFromToken, fh.DeleteByFollower)
		business.GET("/followed/list", mw.UserIDFromToken, fh.GetBusinessList)
	}

	category := v1.Group("/category")
	{
		h := handlers.Category
		category.POST("/create", mw.UserIDFromToken, h.Create)
		category.GET("/get", h.Get)
		category.PUT("/update/:id", h.Update)
		category.DELETE("/delete/:id", mw.UserIDFromToken, h.Delete)
	}

	booking := v1.Group("/booking")
	{
		h := handlers.Booking
		booking.POST("/client/create", mw.UserIDFromToken, h.Create)

		booking.GET("/client/get-list", mw.UserIDFromToken, h.GetListByClient)
		booking.GET("/my-bookings-with-business-profiles", mw.UserIDFromToken, h.GetListByClient)
		booking.GET("/business/get-list/:business_id", h.GetListByBusinessId)
		booking.GET("/get-clients-of-business-profile/:business_id", h.GetListByBusinessId)

		booking.PUT("/update/:id", mw.UserIDFromToken, h.Update)

		booking.DELETE("/client/delete/:id", mw.UserIDFromToken, h.DeleteByClient)
		booking.DELETE("/business/delete/:id", mw.UserIDFromToken, h.DeleteByBusiness)
		booking.DELETE("/delete/:id", mw.UserIDFromToken, h.DeleteById)
	}

	post := v1.Group("/post")
	{
		h := handlers.Post
		post.POST("/create", mw.UserIDFromToken, h.Create)
		post.GET("/get/:id", mw.UserIDFromToken, h.GetDetail)
		post.GET("/list/:business_id", h.GetList)
		post.GET("/list/followed-profiles", mw.UserIDFromToken, h.GetListFollowed)
		post.PUT("/update/:id", mw.UserIDFromToken, h.Update)
		post.DELETE("/delete/:id", mw.UserIDFromToken, h.DeleteById)
	}

	region := v1.Group("/region")
	{
		h := handlers.Region
		region.GET("/get", h.Get)
	}

	dev := v1.Group("/dev")
	{
		dev.GET("/domain", func(c *gin.Context) {
			config.Domain = c.Query("domain")
		})
		dev.GET("/booking/get-list", handlers.Booking.GetList)
		dev.GET("/eskiz/set-data", handlers.Dev.EskizSetData)
		dev.POST("/eskiz/set-data", handlers.Dev.EskizSetData)
		dev.GET("/eskiz/refresh-token", handlers.Dev.EskizRefreshToken)
	}
}

func init() {
	godotenv.Load(".env")
	godotenv.Load("default.env")
	config.LoadVarsFromEnv()
}
