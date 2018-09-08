package access

import (
	"testing"
)

func TestSet(t *testing.T) {

	type Student struct {
		Name        string
		Address     string
		EnglishName string
	}

	data := make(map[string]interface{})
	data["Name"] = "liang"
	data["Englishname"] = "liangaaaa"
	data["Address"] = "China"
	data["age"] = 18

	s := new(Student)
	SetMap(s, data)

	if s.Name != "liang" {
		t.Error("method Set Name faild")
	}
	if s.Address != "China" {
		t.Error("method Set Address faild")
	}
	if s.EnglishName != "liangaaaa" {
		t.Error("method Set EnglishName faild")
	}
}

func TestFill(t *testing.T) {

	type Student struct {
		Name        string
		Age         string
		Address     string
		EnglishName string
	}

	type People struct {
		Name        string
		Age         int
		Gender      string
		Address     string
		EnglishName string
	}

	p := new(People)
	p.Name = "liyuliang"
	p.Age = 25
	p.Gender = "male"
	p.Address = "China"
	p.EnglishName = "liang"

	s := new(Student)
	Set(s,p)

	println("--------")
	println(s.Name)
	println(s.Address)
	println(s.EnglishName)
	println("--------")

	//if s.Name != "liyuliang" {
	//	t.Error("method Fill Name faild")
	//}
	//if s.Address != "China" {
	//	t.Error("method Fill Address faild")
	//}
	//if s.EnglishName != "liang" {
	//	t.Error("method Fill EnglishName faild")
	//}
}
