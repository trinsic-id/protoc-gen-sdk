Set-Location $PSScriptRoot

# Compile in the `SdkTemplateOption` in `field-options.proto`
protoc --proto_path="../sdk/proto" --go_out="./" "../sdk/proto/services/options/field-options.proto"
go version
go mod vendor

# TODO - Support server work vs sdk work
$BuildTarget = "server"  # "sdk"

$ProcessorArch = [System.Runtime.InteropServices.RuntimeInformation]::OSArchitecture.ToString().ToLower()
$BuildPath = "go-plugin/protoc-gen-sdk-$( If ($IsWindows)
{
    'windows'
}
ElseIf ($IsLinux)
{
    'linux'
}
ElseIf ($IsMacOS)
{
    'darwin'
} )-${ProcessorArch}$( If ($IsWindows)
{
    '.exe'
}
Else
{
    ''
} )"

go build -o $BuildPath


$RenamePairs = "trust-registry=trustregistry,universal-wallet=wallet,verifiable-credentials=credential,templates=template"

# Default to doing nothing
$PythonPath = "***SKIP***"
$DotnetPath = "***SKIP***"
$DashboardBffPath = "***SKIP***"
$DartPath = "***SKIP***"
$GolangPath = "***SKIP***"
$TypescriptPath = "***SKIP***"
$JavaKotlinPath = "***SKIP***"
$RubyPath = "***SKIP***"
$SwiftPath = "***SKIP***"

if ($BuildTarget -eq "sdk")
{
    $ProtoPath = "$PSScriptRoot/../sdk/proto"
    $PythonPath = "$PSScriptRoot/../sdk/python/trinsic"
    $DotnetPath = "$PSScriptRoot/../sdk/dotnet/Trinsic"
    $DartPath = "$PSScriptRoot/../sdk/dart/lib/src"
    $GolangPath = "$PSScriptRoot/../sdk/go/services"
    $TypescriptPath = "$PSScriptRoot/../sdk/web/src"
    $JavaKotlinPath = "$PSScriptRoot/../sdk/java/src/main/java/trinsic/services"
    $RubyPath = "$PSScriptRoot/../sdk/ruby/lib/services"
    $SwiftPath = "$PSScriptRoot/../sdk-swift/Sources/Trinsic"
}
elseif ($BuildTarget -eq "server") {
    $ProtoPath = "$PSScriptRoot/../server/proto"
    $DashboardBffPath = "$PSScriptRoot/../server/dashboard/service/Dashboard/Services"
    $DashboardFrontendPath = "$PSScriptRoot/../server/dashboard/app/src/app/services"
}
else {
    throw "unsupported build target"
}

./action.ps1 `
    -ProtoPath $ProtoPath `
    -RenamePairs $RenamePairs `
    -PythonPath $PythonPath `
    -DotnetPath $DotnetPath `
    -DashboardBffPath $DashboardBffPath `
    -DashboardFrontendPath $DashboardFrontendPath `
    -DartPath $DartPath `
    -GolangPath $GolangPath `
    -TypescriptPath $TypescriptPath `
    -JavaKotlinPath $JavaKotlinPath `
    -RubyPath $RubyPath `
    -SwiftPath $SwiftPath