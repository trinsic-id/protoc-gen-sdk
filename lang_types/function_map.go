package lang_types

import "reflect"

func GetTemplateFuncs() map[string]interface{} {
	// TODO - Is there a way to reflection get all LanguageGenerators?
	langGenerators := map[string]LanguageGenerator{
		"Dart":   GetDartGenerator(),
		"Golang": GetGolangGenerator(),
	}

	funcs := map[string]interface{}{
		"BuildMetadata":       BuildMetadata,
		"MethodIsStreaming":   MethodIsStreaming,
		"MethodParamType":     MethodParamType,
		"SdkAnonymous":        SdkAnonymous,
		"SdkTemplateGenerate": SdkTemplateGenerate,

		"DocsCreateService":         DocsCreateService,
		"DocCreateServiceInjection": DocCreateServiceInjection,
		"DocsDocComment":            DocsDocComment,
		"DocsFileNameService":       DocsFileNameService,
		"DocsFileName":              DocsFileName,
		"DocMethodInjection":        DocMethodInjection,
		"DocMethodInjectionEnd":     DocMethodInjectionEnd,
		"DocsMethodName":            DocsMethodName,
		"DocsMethodTab":             DocsMethodTab,
		"DocsServiceTab":            DocsServiceTab,

		"DotnetMethodReturnType":     DotnetMethodReturnType,
		"DotnetDocComment":           DotnetDocComment,
		"DotnetMethodParamType":      DotnetMethodParamType,
		"DotnetMethodArguments":      DotnetMethodArguments,
		"DotnetDefaultRequestObject": DotnetDefaultRequestObject,
		"DotnetAnnotations":          DotnetAnnotations,

		"DashboardBffMethodArguments":       DashboardBffMethodArguments,
		"DashboardFrontendClassDefinition":  DashboardFrontendClassDefinition,
		"DashboardFrontendServicePath":      DashboardFrontendServicePath,
		"DashboardFrontendMethodReturnType": DashboardFrontendMethodReturnType,
		"DashboardFrontendMethodArguments":  DashboardFrontendMethodArguments,
		"DashboardFrontendRequestConstruct": DashboardFrontendRequestConstruct,
		"DashboardFrontendMethodName":       DashboardFrontendMethodName,

		"JavaMethodReturnType":       JavaMethodReturnType,
		"JavaDocComment":             JavaDocComment,
		"JavaAsync":                  JavaAsync,
		"JavaAwait":                  JavaAwait,
		"JavaStreamStub":             JavaStreamStub,
		"JavaMethodParamType":        JavaMethodParamType,
		"JavaMethodArguments":        JavaMethodArguments,
		"JavaDefaultRequestObject":   JavaDefaultRequestObject,
		"JavaAnnotations":            JavaAnnotations,
		"KotlinMethodReturnType":     KotlinMethodReturnType,
		"KotlinDocComment":           KotlinDocComment,
		"KotlinAsync":                KotlinAsync,
		"KotlinAwait":                KotlinAwait,
		"KotlinAnnotations":          KotlinAnnotations,
		"KotlinMethodArguments":      KotlinMethodArguments,
		"KotlinDefaultRequestObject": KotlinDefaultRequestObject,

		"PythonDocComment":           PythonDocComment,
		"PythonMethodReturnType":     PythonMethodReturnType,
		"PythonBuildMetadata":        PythonBuildMetadata,
		"PythonMethodArguments":      PythonMethodArguments,
		"PythonDefaultRequestObject": PythonDefaultRequestObject,
		"PythonAnnotations":          PythonAnnotations,

		"PythonDocCreateServiceInjection": PythonDocCreateServiceInjection,
		"PythonDocMethodInjection":        PythonDocMethodInjection,

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
		"SwiftAnnotations":      SwiftAnnotations,

		"TypescriptMethodReturnType":     TypescriptMethodReturnType,
		"TypescriptMethodParamType":      TypescriptMethodParamType,
		"TypescriptDocComment":           TypescriptDocComment,
		"TypescriptAsync":                TypescriptAsync,
		"TypescriptAwait":                TypescriptAwait,
		"TypescriptBuildMetadata":        TypescriptBuildMetadata,
		"TypescriptMethodArguments":      TypescriptMethodArguments,
		"TypescriptDefaultRequestObject": TypescriptDefaultRequestObject,
	}
	for langName, generator := range langGenerators {
		dt := reflect.TypeOf(generator)
		for i := 0; i < dt.NumMethod(); i++ {
			funcs[langName+dt.Method(i).Name] = dt.Method(i).Func
		}
	}
	return funcs
}
