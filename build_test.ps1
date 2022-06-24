Set-Location $PSScriptRoot
go build

$RenamePairs = "trust-registry=trustregistry,universal-wallet=wallet,verifiable-credentials=credential,templates=template"
$PythonPath = "python_path=../sdk/python/trinsic"
$DotnetPath = "dotnet_path=../sdk/dotnet/Trinsic"
$DartPath ="dart_path=../sdk/dart/lib/src"
$GolangPath = "golang_path=../sdk/go/services"
$TypescriptPath = "typescript_path=../sdk/web/src"
$JavaKotlinPath = "javakotlin_path=../sdk/java/src/main/java/trinsic/services"

protoc --proto_path=..\sdk\proto `
       --trinsic-sdk_out="${RenamePairs},${DartPath},${PythonPath},${GolangPath},${TypescriptPath},${DotnetPath},${JavaKotlinPath}:" `
       --plugin="protoc-gen-trinsic-sdk=${PSScriptRoot}/protoc-gen-sdk.exe" `
        ..\sdk\proto\services\account\v1\*.proto ..\sdk\proto\services\provider\v1\*.proto ..\sdk\proto\services\verifiable-credentials\v1\*.proto ..\sdk\proto\services\verifiable-credentials\templates\v1\*.proto ..\sdk\proto\services\trust-registry\v1\*.proto ..\sdk\proto\services\universal-wallet\v1\*.proto
