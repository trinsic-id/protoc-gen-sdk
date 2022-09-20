Set-Location $PSScriptRoot

# Compile in the `SdkTemplateOption` in `field-options.proto`
protoc --proto_path="../sdk/proto" --go_out="./" "../sdk/proto/services/options/field-options.proto"
go version
go mod vendor

if ($IsWindows)
{
    go build -o "go-plugin/protoc-gen-sdk-windows-amd64.exe"
}
Elseif ($IsLinux)
{
    go build -o "go-plugin/protoc-gen-sdk-linux-amd64"
}
$RenamePairs = "trust-registry=trustregistry,universal-wallet=wallet,verifiable-credentials=credential,templates=template"
$ProtoPath = "$PSScriptRoot/../sdk/proto"
$PythonPath = "$PSScriptRoot/../sdk/python/trinsic"
$DotnetPath = "$PSScriptRoot/../sdk/dotnet/Trinsic"
$DotnetBffPath = "$PSScriptRoot/../server/dashboard-service/Services"
$DartPath = "$PSScriptRoot/../sdk/dart/lib/src"
$GolangPath = "$PSScriptRoot/../sdk/go/services"
$TypescriptPath = "$PSScriptRoot/../sdk/web/src"
$JavaKotlinPath = "$PSScriptRoot/../sdk/java/src/main/java/trinsic/services"
$RubyPath = "$PSScriptRoot/../sdk/ruby/lib/services"
#$SwiftPath = "$PSScriptRoot/../sdk/swift/Sources/Trinsic"
$SwiftPath = "***SKIP***"

./action.ps1 `
    -ProtoPath $ProtoPath `
    -RenamePairs $RenamePairs `
    -PythonPath $PythonPath `
    -DotnetPath $DotnetPath `
    -DotnetBffPath $DotnetBffPath `
    -DartPath $DartPath `
    -GolangPath $GolangPath `
    -TypescriptPath $TypescriptPath `
    -JavaKotlinPath $JavaKotlinPath `
    -RubyPath $RubyPath `
    -SwiftPath $SwiftPath