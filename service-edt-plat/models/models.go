package models

// APIResponse 通用API响应结构
type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponse 错误响应结构
type ErrorResponse struct {
	Code    interface{} `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// CreateCourseRequest 创建课程请求
type CreateCourseRequest struct {
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	Detail            string   `json:"detail,omitempty"`
	CoverImage        string   `json:"cover_image,omitempty"`
	Price             float64  `json:"price"`
	OriginalPrice     float64  `json:"original_price,omitempty"`
	CategoryID        int      `json:"category_id"`
	Level             string   `json:"level,omitempty"`
	Duration          int      `json:"duration,omitempty"`
	Status            string   `json:"status,omitempty"`
	StartDate         string   `json:"start_date,omitempty"`
	EndDate           string   `json:"end_date,omitempty"`
	TargetStudents    string   `json:"target_students,omitempty"`
	Requirements      string   `json:"requirements,omitempty"`
	LearningObjectives string   `json:"learning_objectives,omitempty"`
	Certificate       string   `json:"certificate,omitempty"`
	CourseHighlights  []string `json:"course_highlights,omitempty"`
	TeachingMethods   []string `json:"teaching_methods,omitempty"`
	SupportServices   []string `json:"support_services,omitempty"`
	CareerProspects   []string `json:"career_prospects,omitempty"`
}

// Course 课程响应结构
type Course struct {
	ID                int      `json:"id"`
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	Detail            string   `json:"detail"`
	CoverImage        string   `json:"cover_image"`
	Price             float64  `json:"price"`
	OriginalPrice     float64  `json:"original_price"`
	CategoryID        int      `json:"category_id"`
	TeacherID         int      `json:"teacher_id"`
	Level             string   `json:"level"`
	Duration          int      `json:"duration"`
	Status            string   `json:"status"`
	StartDate         string   `json:"start_date"`
	EndDate           string   `json:"end_date"`
	TargetStudents    string   `json:"target_students"`
	Requirements      string   `json:"requirements"`
	LearningObjectives string   `json:"learning_objectives"`
	Certificate       string   `json:"certificate"`
	CourseHighlights  []string `json:"course_highlights"`
	TeachingMethods   []string `json:"teaching_methods"`
	SupportServices   []string `json:"support_services"`
	CareerProspects   []string `json:"career_prospects"`
}

// CreateChapterRequest 创建章节请求
type CreateChapterRequest struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	SortOrder   int    `json:"sort_order,omitempty"`
}

// Chapter 章节响应结构
type Chapter struct {
	ID          int    `json:"id"`
	CourseID    int    `json:"course_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
}

// CreateLessonRequest 创建课时请求
type CreateLessonRequest struct {
	Title        string `json:"title"`
	Description  string `json:"description,omitempty"`
	VideoURL     string `json:"video_url,omitempty"`
	DocumentURL  string `json:"document_url,omitempty"`
	Content      string `json:"content,omitempty"`
	Duration     int    `json:"duration,omitempty"`
	SortOrder    int    `json:"sort_order,omitempty"`
	IsFree       bool   `json:"is_free,omitempty"`
	Type         string `json:"type,omitempty"`
}

// Lesson 课时响应结构
type Lesson struct {
	ID          int    `json:"id"`
	ChapterID   int    `json:"chapter_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	VideoURL    string `json:"video_url"`
	Duration    int    `json:"duration"`
	SortOrder   int    `json:"sort_order"`
	IsFree      bool   `json:"is_free"`
	Type        string `json:"type"`
	DocumentURL string `json:"document_url"`
	Content     string `json:"content"`
}
