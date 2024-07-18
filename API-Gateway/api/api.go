package api

import (
	"gateway/api/handlers"

	_ "gateway/docs"

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
	// auth_H := handlers.NewAuthHangler(l_conn)
	// user_H := handlers.NewUserHandler(l_conn)
	// game_H := handlers.NewGameHandler(l_conn)
	progress_H := handlers.NewProgressHandler(l_conn)
	learn_H := handlers.NewLearnHandler(l_conn)


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
	

	learn := router.Group("/languages")
	{
		learn.POST("/start", learn_H.StartLearningLanguage)
		learn.POST("/lesson/:lesson_id/complete", learn_H.CompleteLesson)
		learn.POST("/exercise/:exercise_id/submit", learn_H.SubmitExercise)
		learn.POST("/vocabulary/{language_code}/add", learn_H.AddVocabulary)

		learn.GET("", learn_H.GetLanguages)
		learn.GET("/lessons/:language_code", learn_H.GetLessons)
		learn.GET("/lesson/:lesson_id", learn_H.GetLessonDetails)
        learn.GET("/exercises/:language_code", learn_H.GetExercises)
		learn.GET("/vocabulary/:language_code", learn_H.GetVocabulary)
	}
	

	progress := router.Group("/progress")
	{
		progress.POST("/set-goal", progress_H.SetGoal)
		progress.PUT("/update-goal/:goal_id", progress_H.UpdateGoal)
		progress.GET("/:language_code", progress_H.GetLanguageProgress)
		progress.GET("/daily", progress_H.GetDailyProgress)
		progress.GET("/weekly", progress_H.GetWeeklyProgress)
		progress.GET("/monthly", progress_H.GetMonthlyProgress)
		progress.GET("/achievements", progress_H.GetAchievements)
		progress.GET("/leaderboard/:language_code", progress_H.GetLeaderboard)
		progress.GET("/skills/:language_code", progress_H.GetSkills)
		progress.GET("/goals", progress_H.GetGoals)
		progress.GET("/statistics/:language_code", progress_H.GetStatistics)
		progress.DELETE("/delete-goal/:goal_id", progress_H.DeleteGoal)
	}

	return router
}

// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }
