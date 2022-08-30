Set-Location $PSScriptRoot

# Compile in the `SdkTemplateOption` in `field-options.proto`
protoc --proto_path=..\sdk\proto --go_out=.\ ..\sdk\proto\services\options\*.proto
go version
go mod vendor
go build

$RenamePairs = "trust-registry=trustregistry,universal-wallet=wallet,verifiable-credentials=credential,templates=template"
$PythonPath = "../sdk/python/trinsic"
$DotnetPath = "../sdk/dotnet/Trinsic"
$DartPath ="../sdk/dart/lib/src"
$GolangPath = "../sdk/go/services"
$TypescriptPath = "../sdk/web/src"
$JavaKotlinPath = "../sdk/java/src/main/java/trinsic/services"
$RubyPath = "../sdk/ruby/lib/services"
$SwiftPath = "../sdk-swift/Sources/Trinsic"

.\action.ps1 `
    -ProtoPath "..\sdk\proto" `
    -RenamePairs $RenamePairs `
    -PythonPath $PythonPath `
    -DotnetPath $DotnetPath `
    -DartPath $DartPath `
    -GolangPath $GolangPath `
    -TypescriptPath $TypescriptPath `
    -JavaKotlinPath $JavaKotlinPath `
    -RubyPath $RubyPath `
    -SwiftPath $SwiftPath