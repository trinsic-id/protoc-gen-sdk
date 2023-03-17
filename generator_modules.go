package main

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"github.com/trinsic-id/protoc-gen-sdk/lang_types"
	"text/template"
)

type TrinsicModule struct {
	*pgs.ModuleBase
	ctx        pgsgo.Context
	serviceTpl *template.Template
	fileCase   func(pgs.Name) pgs.Name
	fileExt    string
	targetName string
	fileSuffix string
}

type ITrinsicModule interface {
	OutputFileName(name string) string
}

func (m *TrinsicModule) OutputFileName(baseName string) string {
	// Handle argument renaming
	targetName := pgs.Name(m.Parameters().StrDefault(baseName, baseName))
	return renderFilePath(m.fileCase(targetName).String() + m.fileSuffix + "." + m.fileExt)
}

func trinsicDart() *TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("dartService").Funcs(funcs).Parse(lang_types.DartServiceTpl)),
		fileCase:   pgs.Name.LowerSnakeCase,
		fileExt:    "dart",
		targetName: "dart_path",
		fileSuffix: "_service",
	}
}

func trinsicPython() *TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("pythonService").Funcs(funcs).Parse(lang_types.PythonServiceTpl)),
		fileCase:   pgs.Name.LowerSnakeCase,
		fileExt:    "py",
		targetName: "python_path",
		fileSuffix: "_service",
	}
}

func trinsicGolangInterface() *TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("golangServiceInterface").Funcs(funcs).Parse(lang_types.GoServiceInterfaceTpl)),
		fileCase:   pgs.Name.LowerSnakeCase,
		fileExt:    "go",
		targetName: "golang_path",
		fileSuffix: "_service",
	}
}

func trinsicGolangImplementation() *TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("golangServiceImplementation").Funcs(funcs).Parse(lang_types.GoServiceImplTpl)),
		fileCase:   pgs.Name.LowerSnakeCase,
		fileExt:    "go",
		targetName: "golang_path",
		fileSuffix: "_service",
	}
}

func trinsicDotnet() *TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("dotnetService").Funcs(funcs).Parse(lang_types.DotnetServiceTpl)),
		fileCase:   pgs.Name.UpperCamelCase,
		fileExt:    "cs",
		targetName: "dotnet_path",
		fileSuffix: "Service",
	}
}

func trinsicTypescript() *TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("typescriptService").Funcs(funcs).Parse(lang_types.TypescriptServiceTpl)),
		fileCase:   pgs.Name.UpperCamelCase,
		fileExt:    "ts",
		targetName: "typescript_path",
		fileSuffix: "Service",
	}
}

func trinsicJava() *TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("javaService").Funcs(funcs).Parse(lang_types.JavaServiceTpl)),
		fileCase:   pgs.Name.UpperCamelCase,
		fileExt:    "java",
		targetName: "javakotlin_path",
		fileSuffix: "Service",
	}
}

func trinsicKotlin() *TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("kotlinService").Funcs(funcs).Parse(lang_types.KotlinServiceTpl)),
		fileCase:   pgs.Name.UpperCamelCase,
		fileExt:    "kt",
		targetName: "javakotlin_path",
		fileSuffix: "ServiceKt",
	}
}

func trinsicRuby() *TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("rubyService").Funcs(funcs).Parse(lang_types.RubyServiceTpl)),
		fileCase:   pgs.Name.LowerSnakeCase,
		fileExt:    "rb",
		targetName: "ruby_path",
		fileSuffix: "_service",
	}
}

func trinsicSwift() *TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("swiftService").Funcs(funcs).Parse(lang_types.SwiftServiceTpl)),
		fileCase:   pgs.Name.UpperCamelCase,
		fileExt:    "swift",
		targetName: "swift_path",
		fileSuffix: "Service",
	}
}

func trinsicDashboardBff() *TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("dashboardBffService").Funcs(funcs).Parse(lang_types.DashboardBFFServiceTpl)),
		fileCase:   pgs.Name.UpperCamelCase,
		fileExt:    "cs",
		targetName: "dashboardbff_path",
		fileSuffix: "Service",
	}
}

func trinsicDashboardFrontend() *TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("dashboardFrontendService").Funcs(funcs).Parse(lang_types.DashboardFrontendServiceTpl)),
		fileCase:   pgs.Name.UpperCamelCase,
		fileExt:    "ts",
		targetName: "dashboardfrontend_path",
		fileSuffix: "Service",
	}
}

func trinsicDocs() *TrinsicModule {
	funcs := lang_types.GetTemplateFuncs()
	return &TrinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("docsSamples").Funcs(funcs).Parse(lang_types.DocsSampleTpl)),
		fileCase:   pgs.Name.UpperCamelCase,
		fileExt:    "md",
		targetName: "docs_path",
		fileSuffix: "-service",
	}
}
