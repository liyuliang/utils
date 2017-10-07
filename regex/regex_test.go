package regex

import (
	"testing"
)

func Test_Get(t *testing.T) {
	url := "/zh-hans/slice-of-life/the-sound-of-your-heart/list?title_no=235&serviceZone=CHINA&page=11"

	title_no := Get(url, `\?title_no=(\d+)`)
	if title_no != "235" {
		t.Error("regexp get wrong , expect 235, but get", title_no)
	}
}

func Test_replace(t *testing.T) {
	url := "/zh-hans/slice-of-life/the-sound-of-your-heart/list?title_no=235&serviceZone=CHINA&page=11"

	url2 := Replace(url, `&page=\d+`, `&page=`+"okokok")
	url3 := Replace(url, `list\?`, "???")

	if url2 != "/zh-hans/slice-of-life/the-sound-of-your-heart/list?title_no=235&serviceZone=CHINA&page=okokok" {
		t.Error("regexp replace wrong")
	}
	if url3 != "/zh-hans/slice-of-life/the-sound-of-your-heart/???title_no=235&serviceZone=CHINA&page=11" {
		t.Error("regexp replace wrong")
	}
}
