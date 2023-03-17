package postprocessors

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
)

type TrinsicArtifact interface {
	TargetPath() string
}

type TrinsicSdkTemplateArtifact struct {
	File    pgs.File
	Module  *TrinsicModule
	Service pgs.Service
}

type TrinsicDocTemplateArtifact struct {
	File    pgs.File
	Module  *TrinsicModule
	Service pgs.Service
}

func NewDocArtifact(f pgs.File, m *TrinsicModule, service pgs.Service) *TrinsicDocTemplateArtifact {
	return &TrinsicDocTemplateArtifact{
		File:    f,
		Module:  m,
		Service: service,
	}
}

func (t *TrinsicDocTemplateArtifact) TargetPath() string {
	targetPath := t.Module.JoinPath(t.Module.Parameters().StrDefault(t.Module.TargetName, ""), t.Module.SampleFilePath, t.Module.OutputDocPath(t.Service.Name()))

	t.Module.ModuleBase.Debugf("Samples Target Path=%s\n", targetPath)
	// Handle case-insensitive target file
	return FindTargetFileCaseInsensitive(targetPath)
}

func NewSdkArtifact(f pgs.File, m *TrinsicModule, service pgs.Service) *TrinsicSdkTemplateArtifact {
	return &TrinsicSdkTemplateArtifact{
		File:    f,
		Module:  m,
		Service: service,
	}
}

func (t *TrinsicSdkTemplateArtifact) TargetPath() string {
	baseName := t.File.InputPath().BaseName()
	targetPath := t.Module.JoinPath(t.Module.Parameters().StrDefault(t.Module.TargetName, ""), t.Module.OutputFileName(baseName))

	t.Module.ModuleBase.Debugf("Target Path=%s\n", targetPath)
	// Handle case-insensitive target file
	return FindTargetFileCaseInsensitive(targetPath)
}
