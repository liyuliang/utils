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
	Set(s, data)

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
