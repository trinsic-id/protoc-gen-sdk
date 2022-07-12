package lang_types

import (
	"github.com/golang/protobuf/proto"
	pgs "github.com/lyft/protoc-gen-star"
	"github.com/trinsic-id/protoc-gen-sdk/services/options"
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
