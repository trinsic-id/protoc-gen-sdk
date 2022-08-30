param([Parameter(Mandatory = $true)][string]$ProtoPath, [Parameter(Mandatory = $true)][string]$RenamePairs, [Parameter(Mandatory = $true)][string]$PythonPath, [Parameter(Mandatory = $true)][string]$DotnetPath, [Parameter(Mandatory = $true)][string]$DartPath, [Parameter(Mandatory = $true)][string]$GolangPath, [Parameter(Mandatory = $true)][string]$TypescriptPath, [Parameter(Mandatory = $true)][string]$JavaKotlinPath, [Parameter(Mandatory = $true)][string]$RubyPath, [Parameter(Mandatory = $true)][string]$SwiftPath)

Set-Location $PSScriptRoot

$ProtoPath = (Resolve-Path $ProtoPath).Path
$PythonPath = (Resolve-Path $PythonPath).Path.Replace(":","?")
$DartPath = (Resolve-Path $DartPath).Path.Replace(":","?")
$DotnetPath = (Resolve-Path $DotnetPath).Path.Replace(":","?")
$GolangPath = (Resolve-Path $GolangPath).Path.Replace(":","?")
$TypescriptPath = (Resolve-Path $TypescriptPath).Path.Replace(":","?")
$JavaKotlinPath = (Resolve-Path $JavaKotlinPath).Path.Replace(":","?")
$RubyPath = (Resolve-Path $RubyPath).Path.Replace(":","?")
$SwiftPath = (Resolve-Path $SwiftPath).Path.Replace(":","?")


$PythonArg = "python_path=${PythonPath}"
$DotnetArg = "dotnet_path=${DotnetPath}"
$DartArg = "dart_path=${DartPath}"
$GolangArg = "golang_path=${GolangPath}"
$TypescriptArg = "typescript_path=${TypescriptPath}"
$JavaKotlinArg = "javakotlin_path=${JavaKotlinPath}"
$RubyArg = "ruby_path=${RubyPath}"
$SwiftArg = "swift_path=${SwiftPath}"

# TODO - Support ARM64 identification
$PluginPath = Resolve-Path "${PSScriptRoot}/go-plugin/protoc-gen-sdk-$( If ($IsWindows)
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
} )-amd64$( If ($IsWindows)
{
    '.exe'
}
Else
{
    ''
} )"

Write-Output $PluginPath

foreach ($Item in Get-ChildItem -Path $ProtoPath -Include *.proto -Recurse)
{
    $File = $Item.FullName
    $Expr = "protoc --plugin=protoc-gen-trinsic-sdk=${PluginPath} --trinsic-sdk_out=${RenamePairs},${DartArg},${PythonArg},${GolangArg},${TypescriptArg},${DotnetArg},${JavaKotlinArg},${RubyArg},${SwiftArg}: -I $ProtoPath $File"
    Write-Output $Expr
    Invoke-Expression $Expr
}