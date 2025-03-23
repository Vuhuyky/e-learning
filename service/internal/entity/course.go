package entity

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name        string  `json:"name"`
	Code        string  `json:"code" gorm:"unique"`
	Introduce   string  `json:"introduce"`
	Description string  `json:"description"`
	Thumnail    string  `json:"thumnail"`
	MultiLogin  *bool   `json:"multiLogin"`
	Active      *bool   `json:"active"`
	Value       float64 `json:"value"`

	CreateId uint     `json:"createId"`
	Create   *Profile `json:"create" gorm:"foreignKey:CreateId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CourseCategorys []CourseCategory `json:"courseCategorys" gorm:"foreignKey:CourseId;"`
	Chapters        []Chapter        `json:"chapters" gorm:"foreignKey:CourseId;"`
	SaleCourses     []SaleCourse     `json:"saleCourses" gorm:"foreignKey:CourseId;"`
}
