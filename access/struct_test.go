package access

import (
	"testing"
)

func TestSet(t *testing.T) {

	type Student struct {
		Name    string
		Address string
	}

	data := make(map[string]interface{})
	data["Name"] = "liang"
	data["Address"] = "China"
	data["age"] = 18

	s := new(Student)
	Set(s, data)

	if s.Name != "liang" && s.Address != "China" {
		t.Error("method Set faild")
	}
}
