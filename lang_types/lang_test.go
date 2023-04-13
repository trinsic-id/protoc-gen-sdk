package lang_types

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	"testing"
)

func TestFunctionMap(t *testing.T) {
	x := GetTemplateFuncs()
	dartAsyncFunc, ok := x["DartAsyncModifier"]
	if !ok {
		t.Fail()
	}
	s := dartAsyncFunc.(func(method pgs.Method) string)(nil)
	if s != "async" {
		t.Fail()
	}
}
