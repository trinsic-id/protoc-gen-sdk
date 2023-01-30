package lang_types

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pgs "github.com/lyft/protoc-gen-star"
	"github.com/trinsic-id/protoc-gen-sdk/services/options"
	"strings"
)

// FieldType intersection between pgs.FieldType and pgs.FieldTypeElem
type FieldType interface {
	ProtoType() pgs.ProtoType
	IsEmbed() bool
	IsEnum() bool
	Imports() []pgs.File
	Enum() pgs.Enum
	Embed() pgs.Message
}

// EntityWithParent intersection between pgs.Message and pgs.Enum
type EntityWithParent interface {
	pgs.Entity
	Parent() pgs.ParentEntity
}

func MessageType(entity EntityWithParent) string {
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

func MethodParamType(method pgs.Method) string {
	return MessageType(method.Input())
}

func SdkTemplateGenerate(method pgs.Method) bool {
	optValue, _ := proto.GetExtension(method.Descriptor().GetOptions(), options.E_SdkTemplateOption)
	if optValue != nil {
		templateOption := optValue.(*options.SdkTemplateOption)
		return !templateOption.GetIgnore()
	}
	return true
}

func SdkAnonymous(method pgs.Method) bool {
	optValue, _ := proto.GetExtension(method.Descriptor().GetOptions(), options.E_SdkTemplateOption)
	if optValue != nil {
		templateOption := optValue.(*options.SdkTemplateOption)
		return templateOption.GetAnonymous()
	}
	return false
}

func SdkExperimental(method pgs.Method) bool {
	optValue, _ := proto.GetExtension(method.Descriptor().GetOptions(), options.E_SdkTemplateOption)
	if optValue != nil {
		templateOption := optValue.(*options.SdkTemplateOption)
		return templateOption.GetExperimental()
	}
	return false
}

func SdkDeprecated(method pgs.Method) bool {
	optValue, _ := proto.GetExtension(method.Descriptor().GetOptions(), options.E_SdkTemplateOption)
	if optValue != nil {
		templateOption := optValue.(*options.SdkTemplateOption)
		return templateOption.GetDeprecated()
	}
	return false
}

func SdkNoArguments(method pgs.Method) bool {
	optValue, _ := proto.GetExtension(method.Descriptor().GetOptions(), options.E_SdkTemplateOption)
	if optValue != nil {
		templateOption := optValue.(*options.SdkTemplateOption)
		return templateOption.GetNoArguments()
	}
	return false
}

func BuildMetadata(method pgs.Method, async bool) string {
	s := "(request)"
	if SdkAnonymous(method) {
		s = "()"
	}
	if async {
		return "await BuildMetadataAsync" + s
	}
	return "BuildMetadata" + s
}

func GetAnnotatedComment(method pgs.Method) []string {
	var annotationComments []string
	if SdkExperimental(method) {
		annotationComments = append(annotationComments, "This method is EXPERIMENTAL")
	}
	if SdkDeprecated(method) {
		annotationComments = append(annotationComments, "This method is DEPRECATED")
	}
	return annotationComments
}
