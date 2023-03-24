package main

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	"github.com/trinsic-id/protoc-gen-sdk/lang_types"
	tpp "github.com/trinsic-id/protoc-gen-sdk/postprocessors"
	"text/template"
)

func TrinsicDart() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase:        &pgs.ModuleBase{},
		ServiceTpl:        template.Must(template.New("dartService").Funcs(funcs).Parse(lang_types.DartServiceTpl)),
		SampleTpl:         template.Must(template.New("dartDocSample").Funcs(funcs).Parse(lang_types.DartDocTpl)),
		FileCase:          pgs.Name.LowerSnakeCase,
		FileExt:           "dart",
		TargetName:        "dart_path",
		ServiceFileSuffix: "_service",
		SampleFilePath:    "../../example",
	}
}

func TrinsicPython() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase:        &pgs.ModuleBase{},
		ServiceTpl:        template.Must(template.New("pythonService").Funcs(funcs).Parse(lang_types.PythonServiceTpl)),
		SampleTpl:         template.Must(template.New("pythonDocSample").Funcs(funcs).Parse(lang_types.PythonDocTpl)),
		FileCase:          pgs.Name.LowerSnakeCase,
		FileExt:           "py",
		TargetName:        "python_path",
		ServiceFileSuffix: "_service",
		SampleFilePath:    "../samples",
	}
}

func TrinsicGolangInterface() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase:        &pgs.ModuleBase{},
		ServiceTpl:        template.Must(template.New("golangServiceInterface").Funcs(funcs).Parse(lang_types.GoServiceInterfaceTpl)),
		FileCase:          pgs.Name.LowerSnakeCase,
		FileExt:           "go",
		TargetName:        "golang_path",
		ServiceFileSuffix: "_service",
	}
}

func TrinsicGolangImplementation() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase:        &pgs.ModuleBase{},
		ServiceTpl:        template.Must(template.New("golangServiceImplementation").Funcs(funcs).Parse(lang_types.GoServiceImplTpl)),
		SampleTpl:         template.Must(template.New("golangDocSample").Funcs(funcs).Parse(lang_types.GoDocTpl)),
		FileCase:          pgs.Name.LowerSnakeCase,
		FileExt:           "go",
		TargetName:        "golang_path",
		ServiceFileSuffix: "_service",
		SampleFilePath:    "../examples",
		SampleFileSuffix:  "_test",
	}
}

func TrinsicDotnet() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase:        &pgs.ModuleBase{},
		ServiceTpl:        template.Must(template.New("dotnetService").Funcs(funcs).Parse(lang_types.DotnetServiceTpl)),
		SampleTpl:         template.Must(template.New("dotnetDocSample").Funcs(funcs).Parse(lang_types.DotnetDocTpl)),
		FileCase:          pgs.Name.UpperCamelCase,
		FileExt:           "cs",
		TargetName:        "dotnet_path",
		ServiceFileSuffix: "Service",
		SampleFilePath:    "../Tests",
	}
}

func TrinsicTypescript() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase:        &pgs.ModuleBase{},
		ServiceTpl:        template.Must(template.New("typescriptService").Funcs(funcs).Parse(lang_types.TypescriptServiceTpl)),
		SampleTpl:         template.Must(template.New("golangDocSample").Funcs(funcs).Parse(lang_types.TypescriptDocTpl)),
		FileCase:          pgs.Name.UpperCamelCase,
		FileExt:           "ts",
		TargetName:        "typescript_path",
		ServiceFileSuffix: "Service",
		SampleFilePath:    "../test",
	}
}

func TrinsicJava() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase:        &pgs.ModuleBase{},
		ServiceTpl:        template.Must(template.New("javaService").Funcs(funcs).Parse(lang_types.JavaServiceTpl)),
		SampleTpl:         template.Must(template.New("golangDocSample").Funcs(funcs).Parse(lang_types.JavaDocTpl)),
		FileCase:          pgs.Name.UpperCamelCase,
		FileExt:           "java",
		TargetName:        "javakotlin_path",
		ServiceFileSuffix: "Service",
		SampleFilePath:    "../../../../test/java/trinsic",
	}
}

func TrinsicKotlin() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase:        &pgs.ModuleBase{},
		ServiceTpl:        template.Must(template.New("kotlinService").Funcs(funcs).Parse(lang_types.KotlinServiceTpl)),
		FileCase:          pgs.Name.UpperCamelCase,
		FileExt:           "kt",
		TargetName:        "javakotlin_path",
		ServiceFileSuffix: "ServiceKt",
	}
}

func TrinsicRuby() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase:        &pgs.ModuleBase{},
		ServiceTpl:        template.Must(template.New("rubyService").Funcs(funcs).Parse(lang_types.RubyServiceTpl)),
		FileCase:          pgs.Name.LowerSnakeCase,
		FileExt:           "rb",
		TargetName:        "ruby_path",
		ServiceFileSuffix: "_service",
	}
}

func TrinsicSwift() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase:        &pgs.ModuleBase{},
		ServiceTpl:        template.Must(template.New("swiftService").Funcs(funcs).Parse(lang_types.SwiftServiceTpl)),
		SampleTpl:         template.Must(template.New("swiftDocSample").Funcs(funcs).Parse(lang_types.SwiftDocTpl)),
		FileCase:          pgs.Name.UpperCamelCase,
		FileExt:           "swift",
		TargetName:        "swift_path",
		ServiceFileSuffix: "Service",
		SampleFilePath:    "../../Tests/TrinsicTests",
	}
}

func TrinsicDashboardBff() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase:        &pgs.ModuleBase{},
		ServiceTpl:        template.Must(template.New("dashboardBffService").Funcs(funcs).Parse(lang_types.DashboardBFFServiceTpl)),
		FileCase:          pgs.Name.UpperCamelCase,
		FileExt:           "cs",
		TargetName:        "dashboardbff_path",
		ServiceFileSuffix: "Service",
	}
}

func TrinsicDashboardFrontend() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase:        &pgs.ModuleBase{},
		ServiceTpl:        template.Must(template.New("dashboardFrontendService").Funcs(funcs).Parse(lang_types.DashboardFrontendServiceTpl)),
		FileCase:          pgs.Name.UpperCamelCase,
		FileExt:           "ts",
		TargetName:        "dashboardfrontend_path",
		ServiceFileSuffix: "Service",
	}
}

func TrinsicDocs() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase:        &pgs.ModuleBase{},
		SampleTpl:         template.Must(template.New("docsSamples").Funcs(funcs).Parse(lang_types.DocsSampleTpl)),
		FileCase:          pgs.Name.UpperCamelCase,
		FileExt:           "md",
		TargetName:        "docs_path",
		ServiceFileSuffix: "-service-gen",
	}
}
