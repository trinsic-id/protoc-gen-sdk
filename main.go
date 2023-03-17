package main

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	tpp "github.com/trinsic-id/protoc-gen-sdk/postprocessors"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	supportOptional := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	pgs.Init(pgs.DebugMode(), pgs.SupportedFeatures(&supportOptional)).
		RegisterModule(TrinsicDart()).
		//RegisterModule(TrinsicDotnet()).
		//RegisterModule(TrinsicDashboardBff()).
		//RegisterModule(TrinsicDashboardFrontend()).
		//RegisterModule(TrinsicGolangInterface()).
		//RegisterModule(TrinsicGolangImplementation()).
		//RegisterModule(TrinsicJava()).
		//RegisterModule(TrinsicKotlin()).
		//RegisterModule(TrinsicPython()).
		//RegisterModule(TrinsicRuby()).
		//RegisterModule(TrinsicSwift()).
		//RegisterModule(TrinsicTypescript()).
		//RegisterModule(TrinsicDocs()).
		RegisterPostProcessor(tpp.ApplyTemplateFiles()).
		RegisterPostProcessor(tpp.AppendSampleFiles()).
		Render()
}
