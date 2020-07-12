package main

import "fmt"

func main(){
	var p Phone
	p = &Iphone{"iphone11", "chen"}
	fmt.Println(p.call())

	p = &Huawei{"pro40", "lin"}
	fmt.Println(p.call())
}


type Phone interface{
	call() string
}

type Iphone struct{
	name string
	owner string
}

func (i *Iphone) call() string{
	return fmt.Sprintf("Name: %s, phone type: %s", i.owner, i.name)
}

type Huawei struct{
	name string
	owner string
}

func (h *Huawei) call() string{
	return fmt.Sprintf("Name: %s, phone type: %s", h.owner, h.name)
}




