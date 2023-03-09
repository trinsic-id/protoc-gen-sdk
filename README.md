# protoc-gen-sdk
Protobuf compiler plugin that generates [Trinsic SDK Wrappers](https://github.com/trinsic-id/sdk)

### Installation
* To install locally:
```bash
go get github.com/trinsic-id/protoc-gen-sdk
```
* For use as a github action:
```yaml
      - uses: trinsic-id/protoc-gen-sdk@v0.1
        id: buildsdkwrappers
```

> The folder is case-sensitive depending on filesystem, the final file is not due to our use case of "TrustRegistryService" mapping to "trustregistry"

> To skip a target output, provide an output path of `***SKIP***`

### Usage

See the `build_test.ps1` script for an example of how to build and run the plugin, assuming that [Trinsic SDK](https://github.com/trinsic-id/sdk) is checked out into a parallel directory (`../sdk`).
* You must provide relative (or absolute) paths to the various directories to update
* Because `protoc-gen-star`, the major library that this plugin uses, differentiates the final proto output path with `:`, you must escape the absolute windows path with a `?`, eg (`C?\work\sdk` for `C:\work\sdk`)
* Rename pairs are there to map service/file names to the expected output. The language type with handle proper casing.
* To add another language:
  * Add file `lang_types/[LANG NAME]_types.go`
  * Update the corresponding template `const [LANG NAME]ServiceTpl = \``\\ BEGIN`
  * The delimiter comment MUST be on the first and last line of the template to prevent recursive generation
  * Create a `trinsicModule` factory function in `main.go` `trinsic[LANG NAME]() *trinsicModule`
  * Register the factory function in `main.go` `func main`

!!! tip "Template"
    