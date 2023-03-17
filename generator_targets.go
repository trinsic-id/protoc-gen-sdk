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
		ModuleBase:  &pgs.ModuleBase{},
		ServiceTpl:  template.Must(template.New("dartService").Funcs(funcs).Parse(lang_types.DartServiceTpl)),
		SampleTpl:   template.Must(template.New("dartDocSample").Funcs(funcs).Parse(lang_types.DartDocTpl)),
		FileCase:    pgs.Name.LowerSnakeCase,
		FileExt:     "dart",
		TargetName:  "dart_path",
		FileSuffix:  "_service",
		DocFilePath: "../../example",
	}
}

func TrinsicPython() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		ServiceTpl: template.Must(template.New("pythonService").Funcs(funcs).Parse(lang_types.PythonServiceTpl)),
		FileCase:   pgs.Name.LowerSnakeCase,
		FileExt:    "py",
		TargetName: "python_path",
		FileSuffix: "_service",
	}
}

func TrinsicGolangInterface() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		ServiceTpl: template.Must(template.New("golangServiceInterface").Funcs(funcs).Parse(lang_types.GoServiceInterfaceTpl)),
		FileCase:   pgs.Name.LowerSnakeCase,
		FileExt:    "go",
		TargetName: "golang_path",
		FileSuffix: "_service",
	}
}

func TrinsicGolangImplementation() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		ServiceTpl: template.Must(template.New("golangServiceImplementation").Funcs(funcs).Parse(lang_types.GoServiceImplTpl)),
		FileCase:   pgs.Name.LowerSnakeCase,
		FileExt:    "go",
		TargetName: "golang_path",
		FileSuffix: "_service",
	}
}

func TrinsicDotnet() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		ServiceTpl: template.Must(template.New("dotnetService").Funcs(funcs).Parse(lang_types.DotnetServiceTpl)),
		FileCase:   pgs.Name.UpperCamelCase,
		FileExt:    "cs",
		TargetName: "dotnet_path",
		FileSuffix: "Service",
	}
}

func TrinsicTypescript() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		ServiceTpl: template.Must(template.New("typescriptService").Funcs(funcs).Parse(lang_types.TypescriptServiceTpl)),
		FileCase:   pgs.Name.UpperCamelCase,
		FileExt:    "ts",
		TargetName: "typescript_path",
		FileSuffix: "Service",
	}
}

func TrinsicJava() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		ServiceTpl: template.Must(template.New("javaService").Funcs(funcs).Parse(lang_types.JavaServiceTpl)),
		FileCase:   pgs.Name.UpperCamelCase,
		FileExt:    "java",
		TargetName: "javakotlin_path",
		FileSuffix: "Service",
	}
}

func TrinsicKotlin() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		ServiceTpl: template.Must(template.New("kotlinService").Funcs(funcs).Parse(lang_types.KotlinServiceTpl)),
		FileCase:   pgs.Name.UpperCamelCase,
		FileExt:    "kt",
		TargetName: "javakotlin_path",
		FileSuffix: "ServiceKt",
	}
}

func TrinsicRuby() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		ServiceTpl: template.Must(template.New("rubyService").Funcs(funcs).Parse(lang_types.RubyServiceTpl)),
		FileCase:   pgs.Name.LowerSnakeCase,
		FileExt:    "rb",
		TargetName: "ruby_path",
		FileSuffix: "_service",
	}
}

func TrinsicSwift() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		ServiceTpl: template.Must(template.New("swiftService").Funcs(funcs).Parse(lang_types.SwiftServiceTpl)),
		FileCase:   pgs.Name.UpperCamelCase,
		FileExt:    "swift",
		TargetName: "swift_path",
		FileSuffix: "Service",
	}
}

func TrinsicDashboardBff() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		ServiceTpl: template.Must(template.New("dashboardBffService").Funcs(funcs).Parse(lang_types.DashboardBFFServiceTpl)),
		FileCase:   pgs.Name.UpperCamelCase,
		FileExt:    "cs",
		TargetName: "dashboardbff_path",
		FileSuffix: "Service",
	}
}

func TrinsicDashboardFrontend() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		ServiceTpl: template.Must(template.New("dashboardFrontendService").Funcs(funcs).Parse(lang_types.DashboardFrontendServiceTpl)),
		FileCase:   pgs.Name.UpperCamelCase,
		FileExt:    "ts",
		TargetName: "dashboardfrontend_path",
		FileSuffix: "Service",
	}
}

func TrinsicDocs() *tpp.TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &tpp.TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		ServiceTpl: template.Must(template.New("docsSamples").Funcs(funcs).Parse(lang_types.DocsSampleTpl)),
		FileCase:   pgs.Name.UpperCamelCase,
		FileExt:    "md",
		TargetName: "docs_path",
		FileSuffix: "-service-gen",
	}
}
