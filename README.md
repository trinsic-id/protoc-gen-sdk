# protoc-gen-sdk
Protobuf compiler plugin that generates [Trinsic SDK Wrappers](https://github.com/trinsic-id/sdk) `.rbi` "Ruby Interface" files.

### Installation

```
go get github.com/trinsic-id/protoc-gen-sdk
```

### Usage

```
protoc --trinsicsdk_out=. example.proto
```