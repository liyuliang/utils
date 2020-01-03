package format

import "testing"

func TestSliceDelDuplicate(t *testing.T) {

	input := []string{"hello", "world", "hello", "golang", "hello", "ruby", "php", "java"}
	expect := []string{"hello", "world", "golang", "ruby", "php", "java"}

	result := SliceDelDuplicate(input)

	for i, value := range result {
		if value != expect[i] {
			t.Error("SliceDelDuplicate failed, expect :" + expect[i] + " ,but get " + value)
		}
	}
}

func TestSliceChunk(t *testing.T) {
	var input []interface{}

	for i := 1; i < 10; i++ {
		input = append(input, i)
	}

	result := SliceChunk(input, 4)

	if len(result) != 3 {
		t.Error("Slice Chunk cut to group failed, expect group size is 3, but get ", len(result))
	}

	if result[0][0].(int) != 1 {
		t.Error("Slice Chunk cut to group failed, expect group size is 3, but get ", len(result))
	}

	if result[1][0].(int) != 5 {
		t.Error("Slice Chunk cut to group failed, expect group size is 3, but get ", len(result))
	}
	if result[2][0].(int) != 9 {
		t.Error("Slice Chunk cut to group failed, expect group size is 3, but get ", len(result))
	}

}
