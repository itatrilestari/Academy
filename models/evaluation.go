package models

import "github.com/google/uuid"

type MasterEvalCategory struct {
	Base

	Title		string		`json:"title" gorm:"not null"`
	Order		int64		`json:"order"`

	Question []MasterEvalQuestion `json:"question,omitempty" gorm:"foreignkey:CategoryID"`
}

type MasterEvalQuestion struct {
	Base

	Title		string				`json:"title" gorm:"not null"`
	ParentID	*uuid.UUID			`json:"parent_id" gorm:"type:uuid;index"`
	Parent		*MasterEvalQuestion	`json:"parent,omitempty" gorm:"foreignkey:ParentID"`

	Children []MasterEvalQuestion	`json:"children,omitempty" gorm:"foreignkey:ParentID"`
	
	CategoryID	uuid.UUID			`json:"category_id" gorm:"type:uuid;not null;index"`
	Category	*MasterEvalCategory	`json:"category,omitempty" gorm:"foreignkey:CategoryID"`
	
	EvaluationAnswer []EvaluationAnswer	`json:"evaluation_answer,omitempty" gorm:"foreignkey:QuestionID"`
}

type MasterEvalAnswer struct {
	Base

	Scale		int64		`json:"scale"`
	Predicate	string		`json:"predicate"`
	Score		string		`json:"score"`
}

type EvaluationAnswer struct {
	Base

	QuestionID	uuid.UUID			`json:"question_id" gorm:"type:uuid;not null;index"`
	Question	*MasterEvalQuestion	`json:"question,omitempty" gorm:"foreignkey:QuestionID"`

	AnswerID	uuid.UUID			`json:"answer_id" gorm:"type:uuid;not null;index"`
	Answer		*MasterEvalAnswer	`json:"answer,omitempty" gorm:"foreignkey:AnswerID"`

	EvaluationCourseID	uuid.UUID			`json:"evaluation_course_id" gorm:"type:uuid;not null;index"`
	EvaluationCourse	*EvaluationCourse	`json:"evaluation_course,omitempty" gorm:"foreignkey:EvaluationCourseID"`
}

type EvaluationCourse struct {
	Base

	Suggestion				string				`json:"suggestion"`
}