package main

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"google.golang.org/protobuf/types/pluginpb"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func (m *TrinsicModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *TrinsicModule) Name() string { return "trinsic-sdk" }

func (m *TrinsicModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, t := range targets {
		m.generateServices(t)
	}
	return m.Artifacts()
}

type TrinsicArtifact interface {
	File() pgs.File
	TargetPath() string
	Module() *TrinsicModule
}

type TrinsicSdkTemplateArtifact struct {
	file    pgs.File
	module  *TrinsicModule
	service pgs.Service
}

func (t *TrinsicSdkTemplateArtifact) File() pgs.File {
	return t.file
}

func renderFilePath(targetPath string) string {
	// prepend a "/" on linux
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		targetPath = "/" + targetPath
	}
	// Handle ":" drive on windows
	targetPath = strings.Replace(targetPath, "?", ":", 1)
	return targetPath
}

func findTargetFileCaseInsensitive(targetPath string) string {
	targetPath = renderFilePath(targetPath)
	targetFolder, targetFile := filepath.Split(targetPath)
	fileInfos, err := os.ReadDir(targetFolder)
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
	return targetPath
}

func (t *TrinsicSdkTemplateArtifact) TargetPath() string {
	baseName := t.File().InputPath().BaseName()
	targetPath := t.module.JoinPath(t.module.Parameters().StrDefault(t.module.targetName, ""), t.module.OutputFileName(baseName))

	// Handle case-insensitive target file
	return findTargetFileCaseInsensitive(targetPath)
}

func (t *TrinsicSdkTemplateArtifact) Module() *TrinsicModule {
	return t.module
}

func (m *TrinsicModule) generateServices(f pgs.File) {
	targetPath := m.Parameters().Str(m.targetName)
	m.Debugf("Generate - %v: %v\n", m.targetName, targetPath)
	if strings.Contains(targetPath, "***SKIP***") || len(targetPath) == 0 {
		return
	}
	targetPath = renderFilePath(targetPath)

	for _, service := range f.Services() {
		var sdkInterface TrinsicArtifact = &TrinsicSdkTemplateArtifact{
			file:    f,
			module:  m,
			service: service,
		}
		m.Debugf("Target file path: %s\n", sdkInterface.TargetPath())
		m.AddCustomTemplateFile(sdkInterface.TargetPath(), m.serviceTpl, sdkInterface, os.ModePerm)
	}
}

func main() {
	supportOptional := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	pgs.Init(pgs.DebugEnv("DEBUG"), pgs.SupportedFeatures(&supportOptional)).
		RegisterModule(trinsicDart()).
		RegisterModule(trinsicDotnet()).
		RegisterModule(trinsicDashboardBff()).
		RegisterModule(trinsicDashboardFrontend()).
		RegisterModule(trinsicGolangInterface()).
		RegisterModule(trinsicGolangImplementation()).
		RegisterModule(trinsicJava()).
		RegisterModule(trinsicKotlin()).
		RegisterModule(trinsicPython()).
		RegisterModule(trinsicRuby()).
		RegisterModule(trinsicSwift()).
		RegisterModule(trinsicTypescript()).
		RegisterModule(trinsicDocs()).
		RegisterPostProcessor(applyTemplateFiles()).
		Render()
}
