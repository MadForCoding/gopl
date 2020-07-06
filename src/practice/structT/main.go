package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var chen Employee

func main(){
	chen.ID = 10
	chen.Name = "chen"
	selectById(10).Name = "lin"
	fmt.Println(selectById(10).Name)
}

func selectById(id int) *Employee{
	return &chen
}