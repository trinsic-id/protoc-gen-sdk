package postprocessors

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"os"
	"strings"
	"text/template"
)

type TrinsicModule struct {
	*pgs.ModuleBase
	ctx               pgsgo.Context
	ServiceTpl        *template.Template
	SampleTpl         *template.Template
	FileCase          func(pgs.Name) pgs.Name
	FileExt           string
	TargetName        string
	ServiceFileSuffix string
	SampleFileSuffix  string
	SampleFilePath    string
}

type ITrinsicModule interface {
	OutputFileName(name string) string
	OutputDocPath(name string) string
}

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

func (m *TrinsicModule) generateServices(f pgs.File) {
	targetPath := m.Parameters().Str(m.TargetName)
	m.Debugf("Generate - %v: %v\n", m.TargetName, targetPath)
	if strings.Contains(targetPath, "***SKIP***") || len(targetPath) == 0 {
		return
	}
	targetPath = RenderFilePath(targetPath)

	buildTarget := m.Parameters().StrDefault("build_target", "sdk")

	for _, service := range f.Services() {
		var sdkData = NewSdkArtifact(f, m, service)
		var docData = NewDocArtifact(f, m, service)
		if m.ServiceTpl != nil && buildTarget == "sdk" {
			m.AddCustomTemplateFile(sdkData.TargetPath(), m.ServiceTpl, sdkData, os.ModePerm)
		}
		if m.ServiceTpl != nil && buildTarget == "server" {
			m.AddCustomTemplateFile(sdkData.TargetPath(), m.ServiceTpl, sdkData, os.ModePerm)
		}
		if m.SampleTpl != nil && buildTarget == "docs" {
			m.AddCustomTemplateFile(docData.TargetPath(), m.SampleTpl, docData, os.ModePerm)
		}
	}
}

func (m *TrinsicModule) OutputFileName(baseName string) string {
	// Handle argument renaming
	targetName := pgs.Name(m.Parameters().StrDefault(baseName, baseName))
	return RenderFilePath(m.FileCase(targetName).String() + m.ServiceFileSuffix + "." + m.FileExt)
}

func (m *TrinsicModule) OutputDocPath(baseName pgs.Name) string {
	targetName := strings.Replace(baseName.LowerSnakeCase().String(), "_", "", -1)
	targetName = strings.ToLower(targetName)
	return RenderFilePath(targetName + "_service_examples" + m.SampleFileSuffix + "." + m.FileExt)
}
