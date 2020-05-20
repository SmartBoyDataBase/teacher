package model

import (
	"sbdb-teacher/infrastructure"
	"time"
)

type Teacher struct {
	Id       uint64    `json:"id"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Sex      string    `json:"sex"`
}

func Get(id uint64) (Teacher, error) {
	result := Teacher{
		Id: id,
	}
	row := infrastructure.DB.QueryRow(`
	SELECT name, birthday, sex
	FROM teacher
	WHERE user_id=$1;
	`)
	err := row.Scan(&result.Name, &result.Birthday, &result.Sex)
	return result, err
}

func Create(teacher Teacher) (Teacher, error) {
	_, err := infrastructure.DB.Exec(`
	INSERT INTO teacher(user_id, name, birthday, sex) 
	VALUES ($1,$2,$3,$4);
	`, teacher.Id, teacher.Name,
		teacher.Birthday, teacher.Sex)
	if err != nil {
		return teacher, err
	}
	_, err = infrastructure.DB.Exec(`
	INSERT INTO user_role(user_id, role_id) 
	VALUES ($1, 3);
	`, teacher.Id)
	return teacher, err
}

func Put(teacher Teacher) error {
	_, err := infrastructure.DB.Exec(`
	UPDATE teacher
	SET name=$2,
		birthday=$3,
	    sex=$4
	WHERE user_id=$1;
	`, teacher.Id, teacher.Name, teacher.Birthday, teacher.Sex)
	return err
}

func Delete(id uint64) error {
	// todo: maybe drop cascade?
	_, err := infrastructure.DB.Exec(`
	DELETE FROM teacher
	WHERE user_id=$1;`, id)
	return err
}

func All() ([]Teacher, error) {
	var result []Teacher
	rows, err := infrastructure.DB.Query(`
	SELECT user_id, name, birthday, sex
	FROM teacher;
	`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var current Teacher
		err = rows.Scan(&current.Id, &current.Name, &current.Birthday, &current.Sex)
		if err != nil {
			return result, err
		}
		result = append(result, current)
	}
	return result, nil
}
