Set-Location $PSScriptRoot
go build
protoc --proto_path=testdata --trinsic-sdk_out=testdata --plugin=protoc-gen-trinsic-sdk=${PSScriptRoot}/protoc-gen-sdk.exe .\testdata\services\account\v1\*.proto