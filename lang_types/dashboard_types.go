package lang_types

import (
	"fmt"
	pgs "github.com/lyft/protoc-gen-star"
	"strings"
)

func DashboardBffMethodArguments(method pgs.Method) string {
	if SdkNoArguments(method) {
		return ""
	} else {
		return fmt.Sprintf("%s.Parser.ParseFrom(request.Data)", DotnetMethodParamType(method))
	}
}

func DashboardFrontendClassDefinition(method pgs.Method) string {
	return fmt.Sprintf("%sDefinition", method.Service().Name().UpperCamelCase())
}

func DashboardFrontendServicePath(method pgs.Method) string {
	classDef := DashboardFrontendClassDefinition(method)
	return fmt.Sprintf("`${proto.%s.fullName}/${proto.%s.methods.%s.name}`", classDef, classDef, method.Name().LowerCamelCase())
}

func DashboardFrontendMethodReturnType(method pgs.Method) string {
	return fmt.Sprintf("proto.%s", MessageType(method.Output()))
}

func DashboardFrontendMethodName(method pgs.Method) string {
	methodName := method.Name().LowerCamelCase().String()
	// TODO - Master list of forbidden names
	if methodName == "delete" {
		methodName += "_"
	}
	return methodName
}

func DashboardFrontendRequestConstruct(method pgs.Method) string {
	if !canDestructure(method) {
		return ""
	}
	var requestFields []string
	for _, field := range method.Input().Fields() {
		fieldName := field.Name().LowerCamelCase()
		requestFields = append(requestFields, fmt.Sprintf("%s: %s", fieldName, fieldName))
	}
	return fmt.Sprintf("\n        const request = {%s};", strings.Join(requestFields, ", "))
}

func DashboardFrontendMethodArguments(method pgs.Method) string {
	if !canDestructure(method) {
		// We can't break this down to primitive types only, don't destructure.
		return TypescriptMethodArguments(method)
	}
	// TODO - Single argument destructuring
	var arguments []string
	for _, field := range method.Input().Fields() {
		arguments = append(arguments, fmt.Sprintf("%s: %s", field.Name().LowerCamelCase(), getTypescriptProtoType(field.Type())))
	}
	return strings.Join(arguments, ", ") //fmt.Sprintf("proto.%s", MessageType(method.Output()))
}

func canDestructure(method pgs.Method) bool {
	for _, field := range method.Input().Fields() {
		if field.Type().ProtoType() == pgs.MessageT {
			// We can't break this down to primitive types only, don't destructure.
			return false
		}
	}
	return true
}

func getTypescriptProtoType(myType pgs.FieldType) string {
	primitiveTypesMap := map[pgs.ProtoType]string{
		pgs.StringT: "string",
		pgs.BoolT:   "boolean",
		pgs.Int32T:  "number",
		pgs.FloatT:  "number",
		pgs.Int64T:  "number",
		pgs.UInt64T: "number",
	}

	arrayPrefix := ""
	baseType, hasType := primitiveTypesMap[myType.ProtoType()]
	if myType.IsRepeated() {
		arrayPrefix = "[]"
	}
	if myType.Field().InOneOf() {
		arrayPrefix += " | undefined"
	}
	if hasType {
		return baseType + arrayPrefix
	}
	return "any"
	// TODO - Make this have smarter imports?
	//fullyQualifedType := myType.Field().Descriptor().GetTypeName()
	//typeParts := strings.Split(fullyQualifedType, ".")
	//return "proto." + typeParts[len(typeParts)-1] + " | undefined"
}

const DashboardBFFServiceTpl = `// BEGIN Code generated by protoc-gen-trinsic. DO NOT EDIT.
// target: {{ .TargetPath }}
{{ range .File.Services }}{{ range .Methods }}{{ if SdkTemplateGenerate . }}
            case nameof(service.{{ .Name.UpperCamelCase }}):
                response = await service.{{ .Name.UpperCamelCase }}Async({{ DashboardBffMethodArguments .}});
                break;
{{ end }}{{ end }}{{ end }}
// END Code generated by protoc-gen-trinsic. DO NOT EDIT.`

const DashboardFrontendServiceTpl = `// BEGIN Code generated by protoc-gen-trinsic. DO NOT EDIT.
// target: {{ .TargetPath }}
{{ range .File.Services }}{{ range .Methods }}{{ if SdkTemplateGenerate . }}
export async function {{ DashboardFrontendMethodName . }}({{ DashboardFrontendMethodArguments . }}): {{ TypescriptMethodReturnType . }} {
    try { {{ DashboardFrontendRequestConstruct . }}
        const encodedRequest = {{TypescriptMethodParamType .}}.encode(request).finish();
        const response = await dashboardClient.aPICall({
            service: {{ DashboardFrontendServicePath . }},
            data: encodedRequest,
        });
        const decodedResponse = {{DashboardFrontendMethodReturnType .}}.decode(
            response.data,
            response.data.length
        );

        return decodedResponse;
    } catch (error) {
        errorHandler(error);
        // TODO: dispatch error
        return {{DashboardFrontendMethodReturnType .}}.fromPartial({});
    }
}
{{ end }}{{ end }}{{ end }}
// END Code generated by protoc-gen-trinsic. DO NOT EDIT.`
