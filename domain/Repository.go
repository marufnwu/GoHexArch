package domain

import (

)

type MyRepository struct {
	
}

func (r MyRepository) FindAllUser() ([]User, error)  {
	return []User{
		{"Maruf", 1001, "Male", "Khulna"},
		{"J", 1002, "Female", "Khulna"},
	}, nil
}

func NewRepository() MyRepository{
	// client, err := sql.Open("mysql", "root:@/banking")

	// if err != nil {
	// 	panic(err)
	// }
	// // See "Important settings" section.
	// client.SetConnMaxLifetime(time.Minute * 3)
	// client.SetMaxOpenConns(10)
	// client.SetMaxIdleConns(10)

	c := MyRepository{}

	return c
}