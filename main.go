package main

import (
	"fmt"
	"github.com/trinsic-id/protoc-gen-sdk/lang_types"
	"google.golang.org/protobuf/types/pluginpb"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type trinsicModule struct {
	*pgs.ModuleBase
	ctx        pgsgo.Context
	serviceTpl *template.Template
	fileCase   func(pgs.Name) pgs.Name
	fileExt    string
	targetName string
	fileSuffix string
	// TODO - Support generating the file path as a function
}

func trinsicDart() *trinsicModule {
	funcs := getTemplateFuncs()
	return &trinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("dartService").Funcs(funcs).Parse(lang_types.DartServiceTpl)),
		fileCase:   pgs.Name.LowerSnakeCase,
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
		fileCase:   pgs.Name.LowerSnakeCase,
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
		fileCase:   pgs.Name.LowerSnakeCase,
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
		fileCase:   pgs.Name.LowerSnakeCase,
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
		fileCase:   pgs.Name.UpperCamelCase,
		fileExt:    "cs",
		targetName: "dotnet_path",
		fileSuffix: "Service",
	}
}

func trinsicDotnetBff() *trinsicModule {
	funcs := getTemplateFuncs()
	return &trinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("dotnetBffService").Funcs(funcs).Parse(lang_types.DotnetBFFServiceTpl)),
		fileCase:   pgs.Name.UpperCamelCase,
		fileExt:    "cs",
		targetName: "dotnetbff_path",
		fileSuffix: "Service",
	}
}

func trinsicTypescript() *trinsicModule {
	funcs := getTemplateFuncs()
	return &trinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("typescriptService").Funcs(funcs).Parse(lang_types.TypescriptServiceTpl)),
		fileCase:   pgs.Name.UpperCamelCase,
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
		fileCase:   pgs.Name.UpperCamelCase,
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
		fileCase:   pgs.Name.UpperCamelCase,
		fileExt:    "kt",
		targetName: "javakotlin_path",
		fileSuffix: "ServiceKt",
	}
}

func trinsicRuby() *trinsicModule {
	funcs := getTemplateFuncs()
	return &trinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("rubyService").Funcs(funcs).Parse(lang_types.RubyServiceTpl)),
		fileCase:   pgs.Name.LowerSnakeCase,
		fileExt:    "rb",
		targetName: "ruby_path",
		fileSuffix: "_service",
	}
}

func trinsicSwift() *trinsicModule {
	funcs := getTemplateFuncs()
	return &trinsicModule{
		ModuleBase: &pgs.ModuleBase{},
		serviceTpl: template.Must(template.New("swiftService").Funcs(funcs).Parse(lang_types.SwiftServiceTpl)),
		fileCase:   pgs.Name.UpperCamelCase,
		fileExt:    "swift",
		targetName: "swift_path",
		fileSuffix: "Service",
	}
}

func getTemplateFuncs() map[string]interface{} {
	funcs := map[string]interface{}{
		"BuildMetadata":       lang_types.BuildMetadata,
		"MethodIsStreaming":   lang_types.MethodIsStreaming,
		"MethodParamType":     lang_types.MethodParamType,
		"SdkAnonymous":        lang_types.SdkAnonymous,
		"SdkTemplateGenerate": lang_types.SdkTemplateGenerate,

		"DartMethodReturnType":     lang_types.DartMethodReturnType,
		"DartMethodArguments":      lang_types.DartMethodArguments,
		"DartDefaultRequestObject": lang_types.DartDefaultRequestObject,
		"DartDocComment":           lang_types.DartDocComment,
		"DartAsync":                lang_types.DartAsync,
		"DartAwait":                lang_types.DartAwait,
		"DartBuildMetadata":        lang_types.DartBuildMetadata,

		"DotnetMethodReturnType":     lang_types.DotnetMethodReturnType,
		"DotnetDocComment":           lang_types.DotnetDocComment,
		"DotnetMethodParamType":      lang_types.DotnetMethodParamType,
		"DotnetMethodArguments":      lang_types.DotnetMethodArguments,
		"DotnetDefaultRequestObject": lang_types.DotnetDefaultRequestObject,

		"DotnetBffMethodArguments": lang_types.DotnetBffMethodArguments,

		"GolangDocComment":           lang_types.GoDocComment,
		"GolangBuildMetadata":        lang_types.GolangBuildMetadata,
		"GolangMethodReturnType":     lang_types.GoMethodReturnType,
		"GolangMethodParamType":      lang_types.GoMethodParamType,
		"GolangStructPointer":        lang_types.GoStructPointer,
		"GolangStructPointerVar":     lang_types.GolangStructPointerVar,
		"GolangMethodArguments":      lang_types.GolangMethodArguments,
		"GolangDefaultRequestObject": lang_types.GolangDefaultRequestObject,

		"JavaMethodReturnType":       lang_types.JavaMethodReturnType,
		"JavaDocComment":             lang_types.JavaDocComment,
		"JavaAsync":                  lang_types.JavaAsync,
		"JavaAwait":                  lang_types.JavaAwait,
		"JavaStreamStub":             lang_types.JavaStreamStub,
		"JavaMethodParamType":        lang_types.JavaMethodParamType,
		"JavaMethodArguments":        lang_types.JavaMethodArguments,
		"JavaDefaultRequestObject":   lang_types.JavaDefaultRequestObject,
		"KotlinMethodReturnType":     lang_types.KotlinMethodReturnType,
		"KotlinDocComment":           lang_types.KotlinDocComment,
		"KotlinAsync":                lang_types.KotlinAsync,
		"KotlinAwait":                lang_types.KotlinAwait,
		"KotlinMethodArguments":      lang_types.KotlinMethodArguments,
		"KotlinDefaultRequestObject": lang_types.KotlinDefaultRequestObject,

		"PythonDocComment":           lang_types.PythonDocComment,
		"PythonMethodReturnType":     lang_types.PythonMethodReturnType,
		"PythonBuildMetadata":        lang_types.PythonBuildMetadata,
		"PythonMethodArguments":      lang_types.PythonMethodArguments,
		"PythonDefaultRequestObject": lang_types.PythonDefaultRequestObject,

		"RubyMethodReturnType":     lang_types.RubyMethodReturnType,
		"RubyDocComment":           lang_types.RubyDocComment,
		"RubyMethodParamType":      lang_types.RubyMethodParamType,
		"RubyDefaultRequestObject": lang_types.RubyDefaultRequestObject,
		"RubyMethodArguments":      lang_types.RubyMethodArguments,

		"SwiftMethodReturnType": lang_types.SwiftMethodReturnType,
		"SwiftMethodParamType":  lang_types.SwiftMethodParamType,
		"SwiftDocComment":       lang_types.SwiftDocComment,
		"SwiftAsync":            lang_types.SwiftAsync,
		"SwiftAwait":            lang_types.SwiftAwait,
		"SwiftBuildMetadata":    lang_types.SwiftBuildMetadata,

		"TypescriptMethodReturnType":     lang_types.TypescriptMethodReturnType,
		"TypescriptDocComment":           lang_types.TypescriptDocComment,
		"TypescriptAsync":                lang_types.TypescriptAsync,
		"TypescriptAwait":                lang_types.TypescriptAwait,
		"TypescriptBuildMetadata":        lang_types.TypescriptBuildMetadata,
		"TypescriptMethodArguments":      lang_types.TypescriptMethodArguments,
		"TypescriptDefaultRequestObject": lang_types.TypescriptDefaultRequestObject,
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
	targetName := pgs.Name(t.module.Parameters().StrDefault(baseName, baseName))
	targetPath := t.module.JoinPath(t.module.Parameters().StrDefault(t.module.targetName, ""), t.module.fileCase(targetName).String()+t.module.fileSuffix+"."+t.module.fileExt)
	// prepend a "/" on linux
	if runtime.GOOS == "linux" {
		targetPath = "/" + targetPath
	}
	// Handle ":" drive on windows
	targetPath = strings.Replace(targetPath, "?", ":", 1)
	// Handle case-insensitive target file
	targetFolder, targetFile := filepath.Split(targetPath)
	//fmt.Fprintln(os.Stderr, "Target file="+targetFile)
	//fmt.Fprintln(os.Stderr, "Target folder="+targetFolder)
	fileInfos, err := ioutil.ReadDir(targetFolder)
	if err != nil {
		return ""
	}

	for _, info := range fileInfos {
		if strings.ToLower(info.Name()) == strings.ToLower(targetFile) {
			targetFile = info.Name()
			break
		}
	}
	targetPath = filepath.Join(targetFolder, targetFile)
	fmt.Fprintln(os.Stderr, "Target path="+targetPath)
	return targetPath
}

func (t *trinsicSdk) Module() *trinsicModule {
	return t.module
}

func (m *trinsicModule) generateServices(f pgs.File) {
	targetPath := m.Parameters().Str(m.targetName)
	fmt.Fprintf(os.Stderr, "Generate: %v\n", targetPath)
	if strings.Contains(targetPath, "***SKIP***") {
		fmt.Fprintf(os.Stderr, "Skipping: %v\n", m.targetName)
		return
	}
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
	supportOptional := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	pgs.Init(pgs.DebugEnv("DEBUG"), pgs.SupportedFeatures(&supportOptional)).
		RegisterModule(trinsicDart()).
		RegisterModule(trinsicDotnet()).
		RegisterModule(trinsicDotnetBff()).
		RegisterModule(trinsicGolangInterface()).
		RegisterModule(trinsicGolangImplementation()).
		RegisterModule(trinsicJava()).
		RegisterModule(trinsicKotlin()).
		RegisterModule(trinsicPython()).
		RegisterModule(trinsicRuby()).
		RegisterModule(trinsicSwift()).
		RegisterModule(trinsicTypescript()).
		RegisterPostProcessor(applyTemplateFiles()).
		Render()
}
