package api

import (
	"gateway/api/handlers"
	// "getway/api/middlerware"

	_ "getway/docs"
	_ "getway/genproto"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// @LinguaLeap System API
// @version 1.0
// @description API for managing LinguaLeap System resources
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(l_conn *grpc.ClientConn) *gin.Engine {
	handler := handlers.NewHandler(l_conn)

	router := gin.Default()
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// router.Use(middlerware.Authorizations)
	router.Use(cors.Default())
	// router.Use(corsMiddleware())
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"}, // Adjust for your specific origins
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))
	

	laung := router.Group("/languages")
	{
		laung.POST("/start", handler.StartLearningLanguage)
		laung.POST("/pronunciation/submit", handler.SubmitPronunciation)
		laung.POST("/lesson/:lesson_id/complete", handler.CompleteLesson)
		laung.POST("/exercise/:exercise_id/submit", handler.SubmitExercise)
		laung.POST("/vocabulary/{language_code}/add", handler.AddVocabulary)
		laung.POST("/listening/{exercise_id}/submit", handler.SubmitListeningExercise)

		laung.GET("", handler.GetLanguages)
		laung.GET("/lessons/:language_code", handler.GetLessons)
		laung.GET("/lesson/:lesson_id", handler.GetLessonDetails)
        laung.GET("/exercises/:language_code", handler.GetExercises)
		laung.GET("/vocabulary/:language_code", handler.GetVocabulary)
		laung.GET("/grammar/:language_code", handler.GetGrammarRules)
		laung.GET("/pronunciation/{language_code}", handler.GetPronunciationExercises)
		laung.GET("/listening/{language_code}", handler.GetListeningExercises)
	}


	auth := router.Group("/auth")
	{
		auth.POST("/register", handler.RegisterUser)
		auth.POST("/login", handler.LoginUser)
		auth.POST("/logout", handler.LogoutUser)
		auth.POST("/forgot-password", handler.ForgotPassword)
		auth.POST("/reset-password", handler.ResetPassword)
	}


	user := router.Group("/user")
	{
		user.GET("/profile", handler.GetUserProfile)
		user.PUT("/profile", handler.UpdateUserProfile)
		user.POST("/change-password", handler.ChangeUserPassword)
		user.GET("/settings", handler.GetUserSettings)
		user.PUT("/settings", handler.UpdateUserSettings)
	}
	

	progress := router.Group("/progress")
	{
		progress.POST("/set-goal", handler.SetGoal)
		progress.PUT("/update-goal/:goal_id", handler.UpdateGoal)
		progress.GET("/:language_code", handler.GetLanguageProgress)
		progress.GET("/daily", handler.GetDailyProgress)
		progress.GET("/weekly", handler.GetWeeklyProgress)
		progress.GET("/monthly", handler.GetMonthlyProgress)
		progress.GET("/achievements", handler.GetAchievements)
		progress.GET("/leaderboard/:language_code", handler.GetLeaderboard)
		progress.GET("/skills/:language_code", handler.GetSkills)
		progress.GET("/goals", handler.GetGoals)
		progress.GET("/statistics/:language_code", handler.GetStatistics)
		progress.DELETE("/delete-goal/:goal_id", handler.DeleteGoal)
	}

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
