param(
    [string]$RenamePairs,
    [string]$PythonPath,
    [string]$DotnetPath,
    [string]$DartPath,
    [string]$GolangPath,
    [string]$TypescriptPath,
    [string]$JavaKotlinPath,
    [string]$RubyPath,
    [string]$SwiftPath
)

Set-Location $PSScriptRoot

go build

$PythonArg = "python_path=${PythonPath}"
$DotnetArg = "dotnet_path=${DotnetPath}"
$DartArg = "dart_path=${DartPath}"
$GolangArg = "golang_path=${GolangPath}"
$TypescriptArg = "typescript_path=${TypescriptPath}"
$JavaKotlinArg = "javakotlin_path=${JavaKotlinPath}"
$RubyArg = "ruby_path=${RubyPath}"
$SwiftArg = "swift_path=${SwiftPath}"

protoc --proto_path = ..\sdk\proto `
       --trinsic-sdk_out = "${RenamePairs},${DartArg},${PythonArg},${GolangArg},${TypescriptArg},${DotnetArg},${JavaKotlinArg},${RubyArg},${SwiftArg}:" `
       --plugin = "protoc-gen-trinsic-sdk=${PSScriptRoot}/protoc-gen-sdk.exe" `
        ..\sdk\proto\services\account\v1\*.proto `
        ..\sdk\proto\services\options\*.proto `
        ..\sdk\proto\services\provider\v1\*.proto `
        ..\sdk\proto\services\trust-registry\v1\*.proto `
        ..\sdk\proto\services\universal-wallet\v1\*.proto `
        ..\sdk\proto\services\verifiable-credentials\v1\*.proto `
        ..\sdk\proto\services\verifiable-credentials\templates\v1\*.proto
