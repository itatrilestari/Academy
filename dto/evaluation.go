package dto

import (
	// 	"time"

	"github.com/google/uuid"
)

// Master Eval Category
type MasterEvalCategoryCreateRequest struct {
	Title string `json:"title" form:"title" validate:"required,min=1,max=255"`
	Order int64  `json:"order" form:"order" validate:"required"`
}

type MasterEvalCategoryUpdateRequest struct {
	Title *string `json:"title" form:"title" validate:"omitempty,min=1,max=255"`
	Order *int64  `json:"order" form:"order" validate:"omitempty"`
}

type MasterEvalCategory struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Order int64     `json:"order"`
}

// Master Eval Question
type MasterEvalQuestionCreateRequest struct {
	Title      string     `json:"title" form:"title" validate:"required,min=1,max=255"`
	ParentID   *uuid.UUID `json:"parent_id" form:"parent_id" validate:"omitempty"`
	CategoryID uuid.UUID  `json:"category_id" form:"category_id" validate:"required"`
}

type MasterEvalQuestionUpdateRequest struct {
	Title      *string    `json:"title" form:"title" validate:"omitempty,min=1,max=255"`
	ParentID   *uuid.UUID `json:"parent_id" form:"parent_id"`
	CategoryID *uuid.UUID `json:"category_id" form:"category_id" validate:"omitempty"`
}

type CategoryEvalQuestion struct {
	CategoryID    uuid.UUID             `json:"category_id"`
	CategoryTitle string                `json:"category_title"`
	Questions     []*MasterEvalQuestion `json:"questions"`
}

type MasterEvalQuestion struct {
	ID            uuid.UUID             `json:"id"`
	Title         string                `json:"title"`
	CategoryID    uuid.UUID             `json:"category_id"`
	CategoryTitle string                `json:"category_title,omitempty"`
	ParentID      *uuid.UUID            `json:"parent_id"`
	Children      []*MasterEvalQuestion `json:"children"`
}

// Master Eval Answer
type MasterEvalAnswerCreateRequest struct {
	Scale     int64  `json:"scale" form:"scale" validate:"required"`
	Predicate string `json:"predicate" form:"predicate" validate:"required,min=1,max=255"`
	Score     string `json:"score" form:"score" validate:"required,min=1,max=255"`
}

type MasterEvalAnswerUpdateRequest struct {
	Scale     *int64  `json:"scale" form:"scale" validate:"omitempty"`
	Predicate *string `json:"predicate" form:"predicate" validate:"omitempty"`
	Score     *string `json:"score" form:"score" validate:"omitempty"`
}

type MasterEvalAnswer struct {
	ID        uuid.UUID `json:"id"`
	Scale     int64     `json:"scale"`
	Predicate string    `json:"predicate"`
	Score     string    `json:"score"`
}

// Evaluation
type EvaluationAnswerItem struct {
	QuestionID uuid.UUID `json:"question_id" form:"question_id" validate:"required"`
	AnswerID   uuid.UUID `json:"answer_id" form:"answer_id" validate:"required"`
}

type EvaluationCourseCreateRequest struct {
	Answer     []EvaluationAnswerItem `json:"answer" form:"answer" validate:"required,min=1"`
	Suggestion *string                `json:"suggestion" form:"suggestion" validate:"omitempty,max=2000"`
}

type EvaluationCourseList struct {
	Categories []EvaluationCategoryList `json:"categories"`
	Suggestion string                   `json:"suggestion"`
}

type EvaluationCategoryList struct {
	Category  string                   `json:"category"`
	Questions []EvaluationQuestionList `json:"questions"`
}

type EvaluationQuestionList struct {
	Question string               `json:"question"`
	Answer   EvaluationAnswerList `json:"answer"`
}

type EvaluationAnswerList struct {
	Scale     int64  `json:"scale"`
	Predicate string `json:"predicate"`
	Score     string `json:"score"`
}
