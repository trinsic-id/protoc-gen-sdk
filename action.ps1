param(
    [Parameter()][string]$ProtoPath,
    [Parameter()][string]$RenamePairs,
    [Parameter()][string]$PythonPath,
    [Parameter()][string]$DotnetPath,
    [Parameter()][string]$DartPath,
    [Parameter()][string]$GolangPath,
    [Parameter()][string]$TypescriptPath,
    [Parameter()][string]$JavaKotlinPath,
    [Parameter()][string]$SwiftPath,

    [Parameter()][string]$DocsPath,

    [Parameter()][string]$DashboardBffPath,
    [Parameter()][string]$DashboardFrontendPath,

    [Parameter()][string]$BuildTarget = "sdk"
    )

Set-Location $PSScriptRoot

$ProtoPath = (Resolve-Path $ProtoPath).Path
$PythonPath = (Resolve-Path $PythonPath)?.Path?.Replace(":","?") ?? "***SKIP***"
$DartPath = (Resolve-Path $DartPath)?.Path?.Replace(":","?") ?? "***SKIP***"
$DotnetPath = (Resolve-Path $DotnetPath)?.Path?.Replace(":","?") ?? "***SKIP***"
$GolangPath = (Resolve-Path $GolangPath)?.Path?.Replace(":","?") ?? "***SKIP***"
$TypescriptPath = (Resolve-Path $TypescriptPath)?.Path?.Replace(":","?") ?? "***SKIP***"
$JavaKotlinPath = (Resolve-Path $JavaKotlinPath)?.Path?.Replace(":","?") ?? "***SKIP***"
$RubyPath = "***SKIP***"
$SwiftPath = (Resolve-Path $SwiftPath)?.Path?.Replace(":","?")  ?? "***SKIP***"

$DashboardBffPath = (Resolve-Path $DashboardBffPath)?.Path?.Replace(":","?") ?? "***SKIP***"
$DashboardFrontendPath = (Resolve-Path $DashboardFrontendPath)?.Path?.Replace(":","?") ?? "***SKIP***"
$DocsPath = (Resolve-Path $DocsPath)?.Path?.Replace(":","?") ?? "***SKIP***"

$PythonArg = "python_path=${PythonPath}"
$DotnetArg = "dotnet_path=${DotnetPath}"
$DartArg = "dart_path=${DartPath}"
$GolangArg = "golang_path=${GolangPath}"
$TypescriptArg = "typescript_path=${TypescriptPath}"
$JavaKotlinArg = "javakotlin_path=${JavaKotlinPath}"
$RubyArg = "ruby_path=${RubyPath}"
$SwiftArg = "swift_path=${SwiftPath}"

$DocsArg = "docs_path=${DocsPath}"

$DashboardBffArg = "dashboardbff_path=${DashboardBffPath}"
$DashboardFrontendArg = "dashboardfrontend_path=${DashboardFrontendPath}"

$BuildTargetArg = "build_target=${BuildTarget}"

$ProcessorArch = [System.Runtime.InteropServices.RuntimeInformation]::OSArchitecture.ToString().ToLower()
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
} )-${ProcessorArch}$( If ($IsWindows)
{
    '.exe'
}
Else
{
    ''
} )"

#Write-Output $PluginPath

foreach ($Item in Get-ChildItem -Path $ProtoPath -Include *.proto -Recurse)
{
    $File = $Item.FullName
    $Expr = "protoc --plugin=protoc-gen-trinsic-sdk=${PluginPath} --trinsic-sdk_out=${BuildTargetArg},${RenamePairs},${DartArg},${PythonArg},${GolangArg},${TypescriptArg},${DotnetArg},${JavaKotlinArg},${RubyArg},${SwiftArg},${DashboardBffArg},${DashboardFrontendArg},${DocsArg}: -I $ProtoPath $File"
#    Write-Output $Expr
    Invoke-Expression $Expr
}
# Plugin will issue a code-1 warning due to generating hidden "template generator files" that don't (and shouldn't) exist. Ignore this.
Exit 0