package lang_types

import (
	"fmt"
	"strings"

	pgs "github.com/lyft/protoc-gen-star"
)

func GoMessageType(entity EntityWithParent) string {
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

func GoMethodReturnType(method pgs.Method) string {
	t := GoMessageType(method.Output())
	t = fmt.Sprintf("*%s.%s", lastPackageNamePart(method), t)
	if method.ServerStreaming() {
		return fmt.Sprintf("%s.%s_%sClient", lastPackageNamePart(method), method.Service().Name().UpperCamelCase(), method.Name().UpperCamelCase())
	}
	return t
}

func lastPackageNamePart(method pgs.Method) string {
	nameParts := strings.Split(method.File().Descriptor().GetOptions().GetGoPackage(), "/")
	return strings.ToLower(nameParts[len(nameParts)-1])
}

func GoMethodParamType(method pgs.Method) string {
	t := GoMessageType(method.Input())
	return fmt.Sprintf("%s.%s", lastPackageNamePart(method), t)
}

func GoDocComment(method pgs.Method) string {
	commentLines := deleteEmpty(strings.Split(method.SourceCodeInfo().LeadingComments(), "\n"))
	commentLines = append(GetAnnotatedComment(method), commentLines...)
	if len(commentLines) == 0 {
		return ""
	}
	return fmt.Sprintf("// %s %s", method.Name().UpperCamelCase(), strings.Join(commentLines, "\n//"))
}

func GoStructPointer(method pgs.Method) string {
	serviceLowerName := method.Service().Name().LowerCamelCase().String()
	return fmt.Sprintf("%s *%sBase", GolangStructPointerVar(method), serviceLowerName)
}

func GolangStructPointerVar(method pgs.Method) string {
	serviceLowerName := lastPackageNamePart(method)
	return serviceLowerName[0:1]
}

func GolangBuildMetadata(method pgs.Method) string {
	if SdkAnonymous(method) {
		return "GetMetadataContext(userContext, nil)"
	}
	return "GetMetadataContext(userContext, request)"
}

func GolangMethodArguments(method pgs.Method) string {
	if SdkNoArguments(method) {
		return ""
	} else {
		return fmt.Sprintf(", request *%s", GoMethodParamType(method))
	}
}

func GolangDefaultRequestObject(method pgs.Method) string {
	if SdkNoArguments(method) {
		return fmt.Sprintf("request := &%s{};\n", GoMethodParamType(method))
	}
	return ""
}

const GoServiceInterfaceTpl = `// BEGIN Interface Code generated by protoc-gen-trinsic. DO NOT EDIT.
// target: {{ .TargetPath }}
{{ range .File.Services }}{{ range .Methods }}{{ if SdkTemplateGenerate . }}
    {{ GolangDocComment . }} 
    {{ .Name.UpperCamelCase }}(userContext context.Context{{ GolangMethodArguments . }}) ({{ GolangMethodReturnType . }}, error){{ end }}{{ end }}
{{ end }}
// END Interface Code generated by protoc-gen-trinsic. DO NOT EDIT.`

const GoServiceImplTpl = `// BEGIN Implementation Code generated by protoc-gen-trinsic. DO NOT EDIT.
// target: {{ .TargetPath }}
{{ range .File.Services }}{{ range .Methods }}{{ if SdkTemplateGenerate . }}
{{ GolangDocComment . }}
func ({{ GolangStructPointer . }}) {{ .Name.UpperCamelCase }}(userContext context.Context{{ GolangMethodArguments . }}) ({{ GolangMethodReturnType . }}, error) {
    {{ GolangDefaultRequestObject . }}md, err := {{ GolangStructPointerVar . }}.{{ GolangBuildMetadata . }}
    if err != nil {
        return nil, err
    }
    response, err := {{ GolangStructPointerVar . }}.client.{{ .Name.UpperCamelCase }}(md, request)
    if err != nil {
		return nil, err
	}
	return response, nil
}{{ end }}{{ end }}{{ end }}
// END Implementation Code generated by protoc-gen-trinsic. DO NOT EDIT.`
