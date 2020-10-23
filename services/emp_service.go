package services

import (
	"github.com/MVC/dao"
	"github.com/MVC/model"
)

func GetEmployee(empId int64) (*model.Employee, error) {
	employee, err := dao.GetEmployee(empId)
	if err != nil {
		return nil, err
	}

	return employee, nil
}
