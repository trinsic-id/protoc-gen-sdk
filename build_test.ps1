Set-Location $PSScriptRoot
go build
protoc --proto_path=testdata `
       --trinsic-sdk_out="trust-registry=trustregistry,universal-wallet=wallet,verifiable-credentials=credential,templates=template,python_path=C?/work/sdk/python/trinsic,dart_path=C?/work/sdk/dart/lib/src,golang_path=C?/work/sdk/dart/lib/src:" `
       --plugin="protoc-gen-trinsic-sdk=${PSScriptRoot}/protoc-gen-sdk.exe" `
        .\testdata\services\account\v1\*.proto .\testdata\services\provider\v1\*.proto .\testdata\services\verifiable-credentials\v1\*.proto .\testdata\services\verifiable-credentials\templates\v1\*.proto .\testdata\services\trust-registry\v1\*.proto .\testdata\services\universal-wallet\v1\*.proto
