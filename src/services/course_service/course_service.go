package course_service

import (
	"apis/common"
	"apis/src/models/course"
)

func GetCourseList(page, limit int) (ret []course.Course, err error) {
	err = common.GetGorm().Find(&ret).Error
	if err != nil {
		return
	}
	return
}

func CreateCourse(in course.Course) (err error) {
	w := course.Course{}
	err = common.GetGorm().Where("id = ?", in.ID).First(&w).Error
	if err != nil {
		err = common.GetGorm().Create(&in).Error
		return
	}
	//更新
	err = common.GetGorm().Table("works").Where("id = ?", in.ID).Updates(map[string]interface{}{
		"name":  in.Name,
		"cover": in.Cover,
	}).Error
	if err != nil {
		return
	}
	return
}

func GetUserCourseList() {

}
