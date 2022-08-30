param(
    [Parameter(Mandatory=$true)][string]$ProtoPath,
    [Parameter(Mandatory=$true)][string]$RenamePairs,
    [Parameter(Mandatory=$true)][string]$PythonPath,
    [Parameter(Mandatory=$true)][string]$DotnetPath,
    [Parameter(Mandatory=$true)][string]$DartPath,
    [Parameter(Mandatory=$true)][string]$GolangPath,
    [Parameter(Mandatory=$true)][string]$TypescriptPath,
    [Parameter(Mandatory=$true)][string]$JavaKotlinPath,
    [Parameter(Mandatory=$true)][string]$RubyPath,
    [Parameter(Mandatory=$true)][string]$SwiftPath
)

go version
go mod vendor
go build

Set-Location $PSScriptRoot

$PROTO_DIR = Resolve-Path "$PSScriptRoot/$ProtoPath"

$PythonArg = "python_path=${PythonPath}"
$DotnetArg = "dotnet_path=${DotnetPath}"
$DartArg = "dart_path=${DartPath}"
$GolangArg = "golang_path=${GolangPath}"
$TypescriptArg = "typescript_path=${TypescriptPath}"
$JavaKotlinArg = "javakotlin_path=${JavaKotlinPath}"
$RubyArg = "ruby_path=${RubyPath}"
$SwiftArg = "swift_path=${SwiftPath}"

$PluginPath = "${PSScriptRoot}/protoc-gen-sdk$(If ($IsWindows) {'.exe'} Else {''})"

foreach ($Item in Get-ChildItem -Path $PROTO_DIR -Include *.proto -Recurse)
{
    $File = $Item.FullName
    $Expr = "protoc --plugin=protoc-gen-trinsic-sdk=${PluginPath} --trinsic-sdk_out=${RenamePairs},${DartArg},${PythonArg},${GolangArg},${TypescriptArg},${DotnetArg},${JavaKotlinArg},${RubyArg},${SwiftArg}: -I $PROTO_DIR $File"
    Write-Output $Expr
    Invoke-Expression $Expr
}