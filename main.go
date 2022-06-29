package main

import (
	"github.com/trinsic-id/protoc-gen-sdk/lang_types"
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type trinsicModule struct {
	*pgs.ModuleBase
	ctx        pgsgo.Context
	serviceTpl *template.Template
	fileExt    string
	targetName string
	fileSuffix string
}

func trinsicDart() *trinsicModule {
	funcs := getTemplateFuncs()
	return &trinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("dartService").Funcs(funcs).Parse(lang_types.DartServiceTpl)),
		fileExt:    "dart",
		targetName: "dart_path",
		fileSuffix: "_service",
	}
}

func trinsicPython() *trinsicModule {
	funcs := getTemplateFuncs()
	return &trinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("pythonService").Funcs(funcs).Parse(lang_types.PythonServiceTpl)),
		fileExt:    "py",
		targetName: "python_path",
		fileSuffix: "_service",
	}
}

func trinsicGolangInterface() *trinsicModule {
	funcs := getTemplateFuncs()
	return &trinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("golangServiceInterface").Funcs(funcs).Parse(lang_types.GoServiceInterfaceTpl)),
		fileExt:    "go",
		targetName: "golang_path",
		fileSuffix: "_service",
	}
}

func trinsicGolangImplementation() *trinsicModule {
	funcs := getTemplateFuncs()
	return &trinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("golangServiceImplementation").Funcs(funcs).Parse(lang_types.GoServiceImplTpl)),
		fileExt:    "go",
		targetName: "golang_path",
		fileSuffix: "_service",
	}
}

func trinsicDotnet() *trinsicModule {
	funcs := getTemplateFuncs()
	return &trinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("dotnetService").Funcs(funcs).Parse(lang_types.DotnetServiceTpl)),
		fileExt:    "cs",
		targetName: "dotnet_path",
		fileSuffix: "Service",
	}
}

func trinsicTypescript() *trinsicModule {
	funcs := getTemplateFuncs()
	return &trinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("typescriptService").Funcs(funcs).Parse(lang_types.TypescriptServiceTpl)),
		fileExt:    "ts",
		targetName: "typescript_path",
		fileSuffix: "Service",
	}
}

func trinsicJava() *trinsicModule {
	funcs := getTemplateFuncs()
	return &trinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("javaService").Funcs(funcs).Parse(lang_types.JavaServiceTpl)),
		fileExt:    "java",
		targetName: "javakotlin_path",
		fileSuffix: "Service",
	}
}

func trinsicKotlin() *trinsicModule {
	funcs := getTemplateFuncs()
	return &trinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("kotlinService").Funcs(funcs).Parse(lang_types.KotlinServiceTpl)),
		fileExt:    "kt",
		targetName: "javakotlin_path",
		fileSuffix: "ServiceKt",
	}
}

func getTemplateFuncs() map[string]interface{} {
	funcs := map[string]interface{}{
		"MethodParamType":     lang_types.MethodParamType,
		"MethodIsStreaming":   lang_types.MethodIsStreaming,
		"SdkTemplateGenerate": lang_types.SdkTemplateGenerate,
		"SdkAnonymous":        lang_types.SdkAnonymous,

		"DartMethodReturnType":       lang_types.DartMethodReturnType,
		"DartDocComment":             lang_types.DartDocComment,
		"DartAsync":                  lang_types.DartAsync,
		"DartAwait":                  lang_types.DartAwait,
		"PythonDocComment":           lang_types.PythonDocComment,
		"PythonMethodReturnType":     lang_types.PythonMethodReturnType,
		"GolangDocComment":           lang_types.GoDocComment,
		"GolangMethodReturnType":     lang_types.GoMethodReturnType,
		"GolangMethodParamType":      lang_types.GoMethodParamType,
		"GolangStructPointer":        lang_types.GoStructPointer,
		"GolangStructPointerVar":     lang_types.GolangStructPointerVar,
		"DotnetMethodReturnType":     lang_types.DotnetMethodReturnType,
		"DotnetDocComment":           lang_types.DotnetDocComment,
		"DotnetMethodParamType":      lang_types.DotnetMethodParamType,
		"TypescriptMethodReturnType": lang_types.TypescriptMethodReturnType,
		"TypescriptDocComment":       lang_types.TypescriptDocComment,
		"TypescriptAsync":            lang_types.TypescriptAsync,
		"TypescriptAwait":            lang_types.TypescriptAwait,
		"JavaMethodReturnType":       lang_types.JavaMethodReturnType,
		"JavaDocComment":             lang_types.JavaDocComment,
		"JavaAsync":                  lang_types.JavaAsync,
		"JavaAwait":                  lang_types.JavaAwait,
		"JavaMethodParamType":        lang_types.JavaMethodParamType,
		"KotlinMethodReturnType":     lang_types.KotlinMethodReturnType,
		"KotlinDocComment":           lang_types.KotlinDocComment,
		"KotlinAsync":                lang_types.KotlinAsync,
		"KotlinAwait":                lang_types.KotlinAwait,
	}
	return funcs
}

func (m *trinsicModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *trinsicModule) Name() string { return "trinsic-sdk" }

func (m *trinsicModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, t := range targets {
		m.generateServices(t)
	}
	return m.Artifacts()
}

type TrinsicSdk interface {
	File() pgs.File
	TargetPath() string
	Module() *trinsicModule
}

type trinsicSdk struct {
	file   pgs.File
	module *trinsicModule
}

func (t *trinsicSdk) File() pgs.File {
	return t.file
}

func (t *trinsicSdk) TargetPath() string {
	baseName := t.File().InputPath().BaseName()
	// Handle argument renaming
	targetName := t.module.Parameters().StrDefault(baseName, baseName)
	targetFile := t.module.JoinPath(t.module.Parameters().StrDefault(t.module.targetName, ""), targetName+t.module.fileSuffix+"."+t.module.fileExt)
	// Handle ":" drive on windows
	targetFile = strings.Replace(targetFile, "?", ":", 1)
	return targetFile
}

func (t *trinsicSdk) Module() *trinsicModule {
	return t.module
}

func (m *trinsicModule) generateServices(f pgs.File) {
	baseName := f.InputPath().BaseName()
	templateFile := baseName + m.serviceTpl.Name() + "_service.template_" + m.fileExt

	var sdkInterface TrinsicSdk = &trinsicSdk{
		file:   f,
		module: m,
	}
	if len(f.Services()) > 0 {
		m.AddGeneratorTemplateFile(templateFile, m.serviceTpl, sdkInterface)
	}
}

func main() {
	pgs.Init(pgs.DebugEnv("DEBUG")).
		RegisterModule(trinsicDart()).
		RegisterModule(trinsicPython()).
		RegisterModule(trinsicGolangInterface()).
		RegisterModule(trinsicGolangImplementation()).
		RegisterModule(trinsicDotnet()).
		RegisterModule(trinsicTypescript()).
		RegisterModule(trinsicJava()).
		RegisterModule(trinsicKotlin()).
		RegisterPostProcessor(applyTemplateFiles()).
		Render()
}
