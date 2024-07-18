package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	ln "gateway/genprotos/learn"
)


// StartLearningLanguage godoc
// @Summary Start learning a language
// @Description Start learning a language
// @Tags languages
// @Accept json
// @Produce json
// @Param start body ln.CreateLanguageRequest true "Start Learning Language"
// @Success 200 {object} ln.CreateLanguageResponse
// @Router /languages/start [post]
func (h *LearnHandler) StartLearningLanguage(c *gin.Context) {
	var req ln.CreateLanguageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Language.CreateLanguage(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CompleteLesson godoc
// @Summary Complete a lesson
// @Description Mark a lesson as completed
// @Tags languages
// @Accept json
// @Produce json
// @Param lesson_id path string true "Lesson ID"
// @Param complete body ln.GetLessonRequest true "Complete Lesson"
// @Success 200 {object} ln.GetLessonResponse
// @Router /languages/lesson/{lesson_id}/complete [post]
func (h *LearnHandler) CompleteLesson(c *gin.Context) {
	lessonID := c.Param("lesson_id")
	req := &ln.GetLessonRequest{Id: lessonID}	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.LessonS.GetLesson(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// SubmitExercise godoc
// @Summary Submit an exercise
// @Description Submit an exercise for evaluation
// @Tags languages
// @Accept json
// @Produce json
// @Param exercise_id path string true "Exercise ID"
// @Param submit body ln.GetQuestionRequest true "Submit Exercise"
// @Success 200 {object} ln.GetQuestionResponse
// @Router /languages/exercise/{exercise_id}/submit [post]
func (h *LearnHandler) SubmitExercise(c *gin.Context) {
	exerciseID := c.Param("exercise_id")
	req := &ln.GetQuestionRequest{Id: exerciseID}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Exercises.GetQuestion(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// AddVocabulary godoc
// @Summary Add vocabulary
// @Description Add new vocabulary to a language
// @Tags languages
// @Accept json
// @Produce json
// @Param language_code path string true "Language Code"
// @Param vocabulary body ln.AddVocabularyRequest true "Add Vocabulary"
// @Success 200 {object} ln.AddVocabularyResponse
// @Router /languages/vocabulary/{language_code}/add [post]
func (h *LearnHandler) AddVocabulary(c *gin.Context) {
	languageCode := c.Param("language_code")
	var req ln.AddVocabularyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.LanguageCode = languageCode

	res, err := h.Vocabulary.AddVocabulary(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetLanguages godoc
// @Summary Get languages
// @Description Get a list of languages
// @Tags languages
// @Produce json
// @Success 200 {object} ln.GetLanguageResponse
// @Router /languages [get]
func (h *LearnHandler) GetLanguages(c *gin.Context) {
	res, err := h.Language.GetLanguage(context.Background(), &ln.GetLanguageRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetLessons godoc
// @Summary Get lessons
// @Description Get lessons for a language
// @Tags languages
// @Produce json
// @Param language_code path string true "Language Code"
// @Success 200 {object} ln.GetLessonResponse
// @Router /languages/lessons/{language_code} [get]
func (h *LearnHandler) GetLessons(c *gin.Context) {
	languageCode := c.Param("language_code")
	req := &ln.GetLessonRequest{Id: languageCode}

	res, err := h.LessonS.GetLesson(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetLessonDetails godoc
// @Summary Get lesson details
// @Description Get details for a specific lesson
// @Tags languages
// @Produce json
// @Param lesson_id path string true "Lesson ID"
// @Success 200 {object} ln.GetLessonResponse
// @Router /languages/lesson/{lesson_id} [get]
func (h *LearnHandler) GetLessonDetails(c *gin.Context) {
	lessonID := c.Param("lesson_id")
	req := &ln.GetLessonRequest{Id: lessonID}

	res, err := h.LessonS.GetLesson(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetExercises godoc
// @Summary Get exercises
// @Description Get exercises for a language
// @Tags languages
// @Produce json
// @Param language_code path string true "Language Code"
// @Success 200 {object} ln.GetQuestionResponse
// @Router /languages/exercises/{language_code} [get]
func (h *LearnHandler) GetExercises(c *gin.Context) {
	languageCode := c.Param("language_code")
	req := &ln.GetQuestionRequest{Id: languageCode}

	res, err := h.Exercises.GetQuestion(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetVocabulary godoc
// @Summary Get vocabulary
// @Description Get vocabulary for a language
// @Tags languages
// @Produce json
// @Param language_code path string true "Language Code"
// @Success 200 {object} ln.GetVocabularyResponse
// @Router /languages/vocabulary/{language_code} [get]
func (h *LearnHandler) GetVocabulary(c *gin.Context) {
	languageCode := c.Param("language_code")
	req := &ln.GetVocabularyRequest{Id: languageCode}

	res, err := h.Vocabulary.GetVocabulary(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
