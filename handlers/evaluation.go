package handlers

import (
	"errors"
	"go-fiber-jwt/database"
	"go-fiber-jwt/dto"
	"go-fiber-jwt/models"
	"go-fiber-jwt/utils"
	"go-fiber-jwt/utils/response"

	// "fmt"
	// "strings"
	// "time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

var validate = validator.New()

// Master Evaluation Category
func CreateMasterEvalCategory(c *fiber.Ctx) error {
	var req dto.MasterEvalCategoryCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := validate.Struct(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request payload", err.Error())
	}

	category := models.MasterEvalCategory{
		Title: req.Title,
		Order: req.Order,
	}

	err := database.DB.WithContext(c.Context()).Transaction(func(tx *gorm.DB) error {
		return tx.Create(&category).Error
	})

	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Could not create master category", err.Error())
	}

	return response.Success(c, "master category created successfully", category)
}

func UpdateMasterEvalCategory(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid Master Category ID", err.Error())
	}

	var req dto.MasterEvalCategoryUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := validate.Struct(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request payload", err.Error())
	}

	var category models.MasterEvalCategory
	if err := database.DB.First(&category, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(c, fiber.StatusNotFound, "Course not found", err.Error())
		}
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch master category", err.Error())
	}

	if req.Title != nil {
		category.Title = *req.Title
	}

	if req.Order != nil {
		category.Order = *req.Order
	}

	if err := database.DB.Save(&category).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not update master category", err.Error())
	}

	return response.Success(c, "Master Category updated successfully", category)
}

func DeleteMasterEvalCategory(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid Category ID", err.Error())
	}

	var category models.MasterEvalCategory
	if err := database.DB.First(&category, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(c, fiber.StatusNotFound, "Category not found", err.Error())
		}
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch category", err.Error())
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not delete master category", err.Error())
	}

	return response.Success(c, "Master Category delete successfully", category)

}

func toMasterEvalCategory(m models.MasterEvalCategory) dto.MasterEvalCategory {
	return dto.MasterEvalCategory{
		ID:    m.ID,
		Title: m.Title,
		Order: m.Order,
	}
}

func GetAllMasterEvalCategory(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 10)
	search := c.Query("search", "")

	if page < 1 {
		page = 1
	}

	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	var master_eval_category []models.MasterEvalCategory
	var totalItems int64

	query := database.DB.Model(&models.MasterEvalCategory{})

	if search != "" {
		searchTerm := "%" + search + "%"
		query = query.Where("title ILIKE ?", searchTerm)
	}

	query.Count(&totalItems)

	if err := query.
		Order(`"order"ASC`).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&master_eval_category).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch source", err.Error())
	}

	result := make([]dto.MasterEvalCategory, 0, len(master_eval_category))
	for _, master_eval_category := range master_eval_category {
		result = append(result, toMasterEvalCategory(master_eval_category))
	}

	return response.PaginatedSuccess(c, "Category retrieved successfully", result, page, pageSize, totalItems)
}

// Master Evaluation Question
func CreateMasterEvalQuestion(c *fiber.Ctx) error {
	var req dto.MasterEvalQuestionCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := validate.Struct(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request payload", err.Error())
	}

	var category models.MasterEvalCategory
	if err := database.DB.First(&category, "id = ?", req.CategoryID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(c, fiber.StatusNotFound, "Category not found", err.Error())
		}
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch category", err.Error())
	}

	if req.ParentID != nil {
		var parent models.MasterEvalQuestion
		if err := database.DB.First(&parent, "id = ?", *req.ParentID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return response.Error(c, fiber.StatusNotFound, "Parent question not found", err.Error())
			}
			return response.Error(c, fiber.StatusInternalServerError, "Could not fetch parent question", err.Error())
		}
	}

	question := models.MasterEvalQuestion{
		Title:      req.Title,
		ParentID:   req.ParentID,
		CategoryID: req.CategoryID,
	}

	err := database.DB.WithContext(c.Context()).Transaction(func(tx *gorm.DB) error {
		return tx.Create(&question).Error
	})

	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Could not create master evaluation question", err.Error())
	}

	return response.Success(c, "master evaluation question created successfully", question)
}

func UpdateMasterEvalQuestion(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid Master Evaluation Question ID", err.Error())
	}

	var req dto.MasterEvalQuestionUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := validate.Struct(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request payload", err.Error())
	}

	var question models.MasterEvalQuestion
	if err := database.DB.First(&question, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(c, fiber.StatusNotFound, "Master Eval Question not found", err.Error())
		}
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch master eval question", err.Error())
	}

	if req.Title != nil {
		question.Title = *req.Title
	}

	if req.ParentID != nil {
		if *req.ParentID == question.ID {
			return response.Error(c, fiber.StatusBadRequest, "Invalid parent question", "a question cannot be its own parent")
		}

		var parent models.MasterEvalQuestion
		if err := database.DB.First(&parent, "id = ?", *req.ParentID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return response.Error(c, fiber.StatusNotFound, "Parent question not found", err.Error())
			}
			return response.Error(c, fiber.StatusInternalServerError, "Could not fetch parent question", err.Error())
		}

		question.ParentID = req.ParentID
	}

	if req.CategoryID != nil {
		var category models.MasterEvalCategory
		if err := database.DB.First(&category, "id = ?", *req.CategoryID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return response.Error(c, fiber.StatusNotFound, "Category not found", err.Error())
			}
			return response.Error(c, fiber.StatusInternalServerError, "Could not fetch category", err.Error())
		}

		question.CategoryID = *req.CategoryID
	}

	if err := database.DB.Save(&question).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not update master eval question", err.Error())
	}

	return response.Success(c, "Master Eval Question updated successfully", question)
}

func DeleteMasterEvalQuestion(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid Master Evaluation Question ID", err.Error())
	}

	var question models.MasterEvalQuestion
	if err := database.DB.First(&question, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(c, fiber.StatusNotFound, "Master Eval Question Not Found", err.Error())
		}
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch master eval question", err.Error())
	}

	if err := database.DB.Delete(&question).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not delete master eval question", err.Error())
	}

	return response.Success(c, "Master Eval Question Delete Successfully", question)
}

const maxQuestionTreeDepth = 10

func masterEvalQuestionIDs(list []models.MasterEvalQuestion) []uuid.UUID {
	ids := make([]uuid.UUID, 0, len(list))
	for _, q := range list {
		ids = append(ids, q.ID)
	}
	return ids
}

func buildMasterEvalQuestionTree(questions []models.MasterEvalQuestion) []*dto.CategoryEvalQuestion {
	categories := make(map[uuid.UUID]*dto.CategoryEvalQuestion)
	nodes := make(map[uuid.UUID]*dto.MasterEvalQuestion)

	// Build category dan question nodes
	for _, q := range questions {
		if _, exists := categories[q.CategoryID]; !exists {
			categories[q.CategoryID] = &dto.CategoryEvalQuestion{
				CategoryID:    q.CategoryID,
				CategoryTitle: q.Category.Title,
				Questions:     []*dto.MasterEvalQuestion{},
			}
		}

		node := &dto.MasterEvalQuestion{
			ID:            q.ID,
			Title:         q.Title,
			CategoryID:    q.CategoryID,
			ParentID:      q.ParentID,
			CategoryTitle: q.Category.Title,
			Children:      []*dto.MasterEvalQuestion{},
		}

		nodes[q.ID] = node
	}

	// Build parent-child relationship
	for _, q := range questions {
		node := nodes[q.ID]

		if q.ParentID != nil {
			if parent, exists := nodes[*q.ParentID]; exists {
				parent.Children = append(parent.Children, node)
				continue
			}
		}

		categories[q.CategoryID].Questions = append(
			categories[q.CategoryID].Questions,
			node,
		)
	}

	// Convert map to slice
	result := make([]*dto.CategoryEvalQuestion, 0, len(categories))

	for _, category := range categories {
		result = append(result, category)
	}

	return result
}

func MasterEvalQuestion(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 10)
	search := c.Query("search", "")
	category_id := c.Query("category_id", "")

	if page < 1 {
		page = 1
	}

	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	query := database.DB.Model(&models.MasterEvalQuestion{}).Where("parent_id IS NULL")

	if search != "" {
		query = query.Where("title ILIKE ?", "%"+search+"%")
	}

	if category_id != "" {
		CategoryID, err := uuid.Parse(category_id)
		if err != nil {
			return response.Error(c, fiber.StatusBadRequest, "Invalid Category ID", err.Error())
		}
		query = query.Where("category_id = ?", CategoryID)
	}

	var totalItems int64
	if err := query.Count(&totalItems).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch data", err.Error())
	}

	// root
	var roots []models.MasterEvalQuestion
	if err := query.
		Joins("Category").
		Order(`"Category"."order" ASC`).
		Order("created_at ASC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&roots).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not fect master eval question", err.Error())
	}

	// child
	all := append([]models.MasterEvalQuestion{}, roots...)
	parentIDs := masterEvalQuestionIDs(roots)

	for depth := 0; depth < maxQuestionTreeDepth && len(parentIDs) > 0; depth++ {
		var children []models.MasterEvalQuestion
		if err := database.DB.
			Joins("Category").
			Where("parent_id IN ?", parentIDs).
			Order(`"Category"."order" ASC`).
			Order("created_at ASC").
			Find(&children).Error; err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Could not fetch data", err.Error())
		}
		if len(children) == 0 {
			break
		}

		all = append(all, children...)
		parentIDs = masterEvalQuestionIDs(children)
	}

	tree := buildMasterEvalQuestionTree(all)

	return response.PaginatedSuccess(c, "Master eval question retrieved successfully", tree, page, pageSize, totalItems)
}

// Master Eval Answer
func CreateMasterEvalAnswer(c *fiber.Ctx) error {
	var req dto.MasterEvalAnswerCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := validate.Struct(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request payload", err.Error())
	}

	answer := models.MasterEvalAnswer{
		Scale:     req.Scale,
		Predicate: req.Predicate,
		Score:     req.Score,
	}

	err := database.DB.WithContext(c.Context()).Transaction(func(tx *gorm.DB) error {
		return tx.Create(&answer).Error
	})

	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Could not create master eval answer", err.Error())
	}

	return response.Success(c, "master eval answer created successfully", answer)
}

func UpdateMasterEvalAnswer(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid Master Answer ID", err.Error())
	}

	var req dto.MasterEvalAnswerUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := validate.Struct(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request payload", err.Error())
	}

	var answer models.MasterEvalAnswer
	if err := database.DB.First(&answer, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(c, fiber.StatusNotFound, "Answer not found", err.Error())
		}
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch master answer", err.Error())
	}

	if req.Scale != nil {
		answer.Scale = *req.Scale
	}

	if req.Predicate != nil {
		answer.Predicate = *req.Predicate
	}

	if req.Score != nil {
		answer.Score = *req.Score
	}

	if err := database.DB.Save(&answer).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not update master answer", err.Error())
	}

	return response.Success(c, "Master Answer updated successfully", answer)
}

func DeleteMasterEvalAnswer(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid Answer ID", err.Error())
	}

	var answer models.MasterEvalAnswer
	if err := database.DB.First(&answer, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(c, fiber.StatusNotFound, "Master Answer not found", err.Error())
		}
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch master answer", err.Error())
	}

	if err := database.DB.Delete(&answer).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not delete master answer", err.Error())
	}

	return response.Success(c, "Master answer delete successfully", answer)

}

func toMasterEvalAnswer(m models.MasterEvalAnswer) dto.MasterEvalAnswer {
	return dto.MasterEvalAnswer{
		ID:        m.ID,
		Scale:     m.Scale,
		Predicate: m.Predicate,
		Score:     m.Score,
	}
}

func GetAllMasterEvalAnswer(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 10)
	search := c.Query("search", "")

	if page < 1 {
		page = 1
	}

	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	var master_eval_answer []models.MasterEvalAnswer
	var totalItems int64

	query := database.DB.Model(&models.MasterEvalAnswer{})

	if search != "" {
		searchTerm := "%" + search + "%"
		query = query.Where("predicate ILIKE ?", searchTerm)
	}

	query.Count(&totalItems)

	if err := query.
		Order("created_at ASC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&master_eval_answer).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch source", err.Error())
	}

	result := make([]dto.MasterEvalAnswer, 0, len(master_eval_answer))
	for _, master_eval_answer := range master_eval_answer {
		result = append(result, toMasterEvalAnswer(master_eval_answer))
	}

	return response.PaginatedSuccess(c, "Master Answer retrieved successfully", result, page, pageSize, totalItems)
}

func uniqueAnswerUUIDs(ids []uuid.UUID) []uuid.UUID {
	seen := make(map[uuid.UUID]struct{}, len(ids))
	out := make([]uuid.UUID, 0, len(ids))
	for _, id := range ids {
		if _, ok := seen[id]; !ok {
			seen[id] = struct{}{}
			out = append(out, id)
		}
	}
	return out
}

func CreateEvaluation(c *fiber.Ctx) error {
	var req dto.EvaluationCourseCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request format", err.Error())
	}

	if err := validate.Struct(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request payload", err.Error())
	}

	questionStrs := make([]string, 0, len(req.Answer))
	answerStrs := make([]string, 0, len(req.Answer))

	for _, item := range req.Answer {
		questionStrs = append(questionStrs, item.QuestionID.String())
		answerStrs = append(answerStrs, item.AnswerID.String())
	}

	questionIDs, err := utils.ParseIDs(questionStrs)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid Question ID", err.Error())
	}

	answerIDs, err := utils.ParseIDs(answerStrs)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid Answer ID", err.Error())
	}

	var questionCount int64
	if err := database.DB.Model(&models.MasterEvalQuestion{}).
		Where("id IN ?", questionIDs).
		Count(&questionCount).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not validate question", err.Error())
	}

	if int(questionCount) != len(questionIDs) {
		return response.Error(c, fiber.StatusBadRequest, "One or more Question not found", "")
	}

	uniqueAnswerIDs := uniqueAnswerUUIDs(answerIDs)

	var answerCount int64
	if err := database.DB.Model(&models.MasterEvalAnswer{}).
		Where("id IN ?", uniqueAnswerIDs).
		Count(&answerCount).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not validate answer", err.Error())
	}

	if int(answerCount) != len(uniqueAnswerIDs) {
		return response.Error(c, fiber.StatusBadRequest, "One or more Answer not found", "invalid answer_id")
	}

	suggestion := ""
	if req.Suggestion != nil {
		suggestion = *req.Suggestion
	}

	var evaluationCourse models.EvaluationCourse

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		evaluationCourse = models.EvaluationCourse{
			Suggestion: suggestion,
		}

		if err := tx.Create(&evaluationCourse).Error; err != nil {
			return err
		}

		evaluationAnswers := make([]models.EvaluationAnswer, 0, len(req.Answer))

		for i := range req.Answer {
			evaluationAnswers = append(evaluationAnswers, models.EvaluationAnswer{
				QuestionID:         questionIDs[i],
				AnswerID:           answerIDs[i],
				EvaluationCourseID: evaluationCourse.ID,
			})
		}

		return tx.Create(&evaluationAnswers).Error
	})

	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not create evaluation", err.Error())
	}

	return response.Success(c, "Evaluation created successfully", evaluationCourse)
}

func GetListEvaluationAnswer(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid Evaluation ID", err.Error())
	}

	// Evaluation Course
	var evaluation models.EvaluationCourse

	if err := database.DB.First(&evaluation, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(c, fiber.StatusNotFound, "Evaluation not found", err.Error())
		}

		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch evaluation", err.Error())
	}

	// Evaluation Answers
	var evaluationAnswers []models.EvaluationAnswer

	if err := database.DB.
		Preload("Question.Category").
		Preload("Answer").
		Where("evaluation_course_id = ?", evaluation.ID).
		Find(&evaluationAnswers).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Could not fetch Evaluation Answer", err.Error())
	}

	// Group answers by category
	categoryMap := make(map[uuid.UUID]int)
	categories := make([]dto.EvaluationCategoryList, 0)

	for _, item := range evaluationAnswers {
		if item.Question == nil ||
			item.Question.Category == nil ||
			item.Answer == nil {
			continue
		}

		category := item.Question.Category
		categoryIndex, exists := categoryMap[category.ID]

		if !exists {
			categoryIndex = len(categories)
			categoryMap[category.ID] = categoryIndex

			categories = append(categories, dto.EvaluationCategoryList{
				Category:  category.Title,
				Questions: make([]dto.EvaluationQuestionList, 0),
			})
		}

		categories[categoryIndex].Questions = append(
			categories[categoryIndex].Questions,
			dto.EvaluationQuestionList{
				Question: item.Question.Title,
				Answer: dto.EvaluationAnswerList{
					Scale:     item.Answer.Scale,
					Predicate: item.Answer.Predicate,
					Score:     item.Answer.Score,
				},
			},
		)
	}

	result := dto.EvaluationCourseList{
		Categories: categories,
		Suggestion: evaluation.Suggestion,
	}

	return response.Success(c, "Evaluation Retrieved Successfully", result)
}
