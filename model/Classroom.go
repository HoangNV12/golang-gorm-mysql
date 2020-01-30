package model

type Classroom struct {
	//Trong Gorm, trường ID được tự động đặt thành Primary key với thuộc tính auto increment (neu la kieu int, con kieu string thi khong).
	ID int
	Name string
	////Gorm đưa ra định nghĩa model bao gồm các trường như Id, CreatedAt, UpdatedAt, DeletedAt
	//gorm.Model
	//Name string
}
