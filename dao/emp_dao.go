package dao

import (
	"errors"
	"fmt"

	"github.com/MVC/model"
)

var (
	employees = map[int64]*model.Employee{
		100: {EmpId: 101, EmpName: "Carlos", EmpEmail: "carlos@email.com"},
	}
)

func GetEmployee(empId int64) (*model.Employee, error) {
	if employee := employees[empId]; employee != nil {
		return employee, nil
	}
	return nil, errors.New(fmt.Sprintf("Employee %v was not found", empId))
}
