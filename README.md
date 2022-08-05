# protoc-gen-sdk
Protobuf compiler plugin that generates [Trinsic SDK Wrappers](https://github.com/trinsic-id/sdk)

### Installation

```
go get github.com/trinsic-id/protoc-gen-sdk
```

### Usage

See the `build_test.ps1` script for an example of how to build and run the plugin, assuming that [Trinsic SDK](https://github.com/trinsic-id/sdk) is checked out into a parallel directory.
* You must provide relative (or absolute) paths to the various directories to update
* Because `protoc-gen-star`, the major library that this plugin uses, differentiates the final proto output path with `:`, you must escape the absolute windows path with a `?`, eg (`C?\work\sdk` for `C:\work\sdk`)
* Rename pairs are there to map service/file names to the expected output. The language type with handle proper casing.
* To add another language:
  * Add file `lang_types/[LANG NAME]_types.go`
  * Update the corresponding template `const [LANG NAME]ServiceTpl = ...`
  * Create a `trinsicModule` factory function in `main.go` `trinsic[LANG NAME]() *trinsicModule`
  * Register the factory function in `main.go` `func main`
