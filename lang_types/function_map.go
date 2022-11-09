package lang_types

func GetTemplateFuncs() map[string]interface{} {
	funcs := map[string]interface{}{
		"BuildMetadata":       BuildMetadata,
		"MethodIsStreaming":   MethodIsStreaming,
		"MethodParamType":     MethodParamType,
		"SdkAnonymous":        SdkAnonymous,
		"SdkTemplateGenerate": SdkTemplateGenerate,

		"DartMethodReturnType":     DartMethodReturnType,
		"DartMethodArguments":      DartMethodArguments,
		"DartDefaultRequestObject": DartDefaultRequestObject,
		"DartDocComment":           DartDocComment,
		"DartAsync":                DartAsync,
		"DartAwait":                DartAwait,
		"DartBuildMetadata":        DartBuildMetadata,

		"DotnetMethodReturnType":     DotnetMethodReturnType,
		"DotnetDocComment":           DotnetDocComment,
		"DotnetMethodParamType":      DotnetMethodParamType,
		"DotnetMethodArguments":      DotnetMethodArguments,
		"DotnetDefaultRequestObject": DotnetDefaultRequestObject,

		"DashboardBffMethodArguments":       DashboardBffMethodArguments,
		"DashboardFrontendClassDefinition":  DashboardFrontendClassDefinition,
		"DashboardFrontendServicePath":      DashboardFrontendServicePath,
		"DashboardFrontendMethodReturnType": DashboardFrontendMethodReturnType,
		"DashboardFrontendMethodArguments":  DashboardFrontendMethodArguments,
		"DashboardFrontendRequestConstruct": DashboardFrontendRequestConstruct,
		"DashboardFrontendMethodName":       DashboardFrontendMethodName,

		"GolangDocComment":           GoDocComment,
		"GolangBuildMetadata":        GolangBuildMetadata,
		"GolangMethodReturnType":     GoMethodReturnType,
		"GolangMethodParamType":      GoMethodParamType,
		"GolangStructPointer":        GoStructPointer,
		"GolangStructPointerVar":     GolangStructPointerVar,
		"GolangMethodArguments":      GolangMethodArguments,
		"GolangDefaultRequestObject": GolangDefaultRequestObject,

		"JavaMethodReturnType":       JavaMethodReturnType,
		"JavaDocComment":             JavaDocComment,
		"JavaAsync":                  JavaAsync,
		"JavaAwait":                  JavaAwait,
		"JavaStreamStub":             JavaStreamStub,
		"JavaMethodParamType":        JavaMethodParamType,
		"JavaMethodArguments":        JavaMethodArguments,
		"JavaDefaultRequestObject":   JavaDefaultRequestObject,
		"KotlinMethodReturnType":     KotlinMethodReturnType,
		"KotlinDocComment":           KotlinDocComment,
		"KotlinAsync":                KotlinAsync,
		"KotlinAwait":                KotlinAwait,
		"KotlinMethodArguments":      KotlinMethodArguments,
		"KotlinDefaultRequestObject": KotlinDefaultRequestObject,

		"PythonDocComment":           PythonDocComment,
		"PythonMethodReturnType":     PythonMethodReturnType,
		"PythonBuildMetadata":        PythonBuildMetadata,
		"PythonMethodArguments":      PythonMethodArguments,
		"PythonDefaultRequestObject": PythonDefaultRequestObject,

		"RubyMethodReturnType":     RubyMethodReturnType,
		"RubyDocComment":           RubyDocComment,
		"RubyMethodParamType":      RubyMethodParamType,
		"RubyDefaultRequestObject": RubyDefaultRequestObject,
		"RubyMethodArguments":      RubyMethodArguments,

		"SwiftMethodReturnType": SwiftMethodReturnType,
		"SwiftMethodParamType":  SwiftMethodParamType,
		"SwiftDocComment":       SwiftDocComment,
		"SwiftAsync":            SwiftAsync,
		"SwiftAwait":            SwiftAwait,
		"SwiftBuildMetadata":    SwiftBuildMetadata,

		"TypescriptMethodReturnType":     TypescriptMethodReturnType,
		"TypescriptMethodParamType":      TypescriptMethodParamType,
		"TypescriptDocComment":           TypescriptDocComment,
		"TypescriptAsync":                TypescriptAsync,
		"TypescriptAwait":                TypescriptAwait,
		"TypescriptBuildMetadata":        TypescriptBuildMetadata,
		"TypescriptMethodArguments":      TypescriptMethodArguments,
		"TypescriptDefaultRequestObject": TypescriptDefaultRequestObject,
	}
	return funcs
}
