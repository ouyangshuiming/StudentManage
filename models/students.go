package models

import (
	"StudentManage/pkg/config"

	"database/sql"
	"log"
)

type Student struct {
	ID      int     `json:"id" form:"id"`
	Name    string  `json:"name" form:"name"`
	Class   string  `json:"class" form:"class"`
	Chinese float32 `json:"chinese" form:"chinese"` //语文成绩
	Math    float32 `json:"math" form:"math"`       //数学成绩
	English float32 `json:"english" form:"english"` //英语成绩
}

var db *sql.DB

func init() {
	db = config.GetDB()
}

// 更新某个学生的信息
func UpdateStudent(student Student) (*Student, error) {
	sql := "UPDATE student SET name=?, class=?, chinese=?, math=?, english=? WHERE id=?"
	_, err := db.Exec(sql, student.Name, student.Class, student.Chinese,
		student.Math, student.English, student.ID)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func CreatStudent(student Student) *Student {
	sql := "INSERT INTO student(name, class, chinese, math, english) VALUES (?, ?, ?, ?, ?)"
	result, err := db.Exec(sql, student.Name, student.Class, student.Chinese, student.Math, student.English)
	if err != nil {
		log.Fatalf("Failed inserting into the table: %v", err)
		return nil
	}

	id, err := result.LastInsertId() // 获取插入的新记录的 ID
	if err != nil {
		log.Fatalf("Failed to get last inserted id: %v", err)
		return nil
	}
	student.ID = int(id)

	return &student
}

// 查询所有学生
func GetAllStudent() []Student {
	sql := "SELECT id, name, class, chinese, math, english FROM student"
	rows, err := db.Query(sql) //返回的是结果集
	if err != nil {
		log.Fatalf("Failed to run the query: %v", err)
	}

	defer rows.Close()

	var students []Student
	for rows.Next() { //尝试去结果集的下一行
		var s Student //把每个学生取出来
		rows.Scan(&s.ID, &s.Name, &s.Class, &s.Chinese, &s.Math, &s.English)
		students = append(students, s)
	}

	return students
}

// 根据id获取某个学生的信息
func GetStudent(id int) (*Student, error) {
	sql := "SELECT id, name, class, chinese, math, english FROM student WHERE id = ?"
	row := db.QueryRow(sql, id)

	var s Student
	row.Scan(&s.ID, &s.Name, &s.Class, &s.Chinese, &s.Math, &s.English)
	return &s, nil
}

// 根据id删除某个学生的信息
func DeleteStudent(id int) error {
	sql := "DELETE FROM student WHERE id = ?"
	_, err := db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}
