package course

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	ID                int64           `json:"id"`
	Uid               int             `json:"uid"`
	CourseSection     []CourseSection `json:"sections"`
	Name              string          `json:"name"`
	Brief             string          `json:"brief"`
	TeacherID         string          `json:"teacher_id"`
	PreviewFirstField string          `json:"preview_first_field"` // 课程概述
	Price             float32         `json:"price"`
	PriceTag          string          `json:"price_tag"`
	Discounts         string          `json:"discounts"` // 课程优惠价格
	Cover             string          `json:"cover"`
	ShareTitle        string          `json:"share_title"`
	ShareDerscription string          `json:"share_description"`
	Description       string          `json:"description"` // 课程描述
	SortNum           int             `json:"sout_num"`
	Status            int             `json:"status"` // 课程状态： 1 上架 0 下架
	IsDel             bool            `json:"is_del"`
}

type CourseSection struct {
	gorm.Model
	CourseID     uint           `json:"course_id"`
	CourseLesson []CourseLesson `json:"lessons"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	OrderNum     int            `json:"order_num"`
	Status       int            `json:"status"` // 0 隐藏 1 待更新 2 已发布
	IsDel        bool           `json:"is_del"`
}

type CourseLesson struct {
	gorm.Model
	CourseSectionID uint        `json:"section_id"`
	Theme           string      `json:"theme"`
	Duration        time.Time   `json:"duration"`
	IsFree          bool        `json:"is_free"`
	OrderNum        int         `json:"order_num"`
	Status          int         `json:"status"` // 0 隐藏 1 待更新 2 已发布
	IsDel           bool        `json:"is_del"`
	CourseMediaID   string      `json:"media_id"`
	CourseMedia     CourseMedia `json:"media"`
}

type CourseMedia struct {
	gorm.Model
	CourseLessonID uint   `json:"lesson_id"`
	CoverImageUrl  string `json:"cover_image_url"`
	MediaType      string `json:"media_type"` // 媒体类型：0 视频、1 音频 2文档
	Status         int    `json:"status"`     // 0 隐藏 1 待更新 2 已发布
	IsDel          bool   `json:"is_del"`
}
