package lang_types

import (
	"fmt"
	"strings"

	pgs "github.com/lyft/protoc-gen-star"
)

func PythonMessageType(entity EntityWithParent) string {
	names := make([]string, 0)
	outer := entity
	ok := true
	for ok {
		name := outer.Name().String()
		names = append([]string{strings.Title(name)}, names...)
		outer, ok = outer.Parent().(pgs.Message)
	}
	return fmt.Sprintf("%s", strings.Join(names, "."))
}

func PythonMethodParamType(method pgs.Method) string {
	return pythonMethodType(method.Input(), method.ClientStreaming())
}

func PythonMethodReturnType(method pgs.Method) string {
	return pythonMethodType(method.Output(), method.ServerStreaming())
}

func EntityDocComment(service pgs.Entity) string {
	return strings.Trim(service.SourceCodeInfo().LeadingComments(), "\n\r\f \t")
}

func pythonMethodType(message pgs.Message, streaming bool) string {
	t := PythonMessageType(message)
	if streaming {
		return fmt.Sprintf("TODO.Streaming[%s]", t)
	}
	return t
}
