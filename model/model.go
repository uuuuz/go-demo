package model

type User struct {
	Name   string `db:"name"`
	Age    int    `db:"age"`
	Field1 string `db:"field1"`
}
