package format

import "testing"

const u = "https://www.baidu.com/"
const urlEncode = "https%3A%2F%2Fwww.baidu.com%2F"

func TestUrlEncode(t *testing.T) {

	result := UrlEncode(u)

	if result != urlEncode {
		t.Error("u encode wrong")
	}
}

func TestUrlDecode(t *testing.T) {

	result, err := UrlDecode(urlEncode)
	if err != nil {
		t.Error(err.Error())
	} else {
		if result != u {
			t.Error(result)
			t.Error("u decode wrong")
		}
	}
}
