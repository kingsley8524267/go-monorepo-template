param(
    [Parameter(Mandatory=$true)][string]$AppName,
    [string]$Version
)
$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Definition
$AppDir = Join-Path $ScriptDir "..\cmd\$AppName" | Resolve-Path
$VersionFile = Join-Path $ScriptDir ".$AppName"

# 自动递增版本
if (-not $Version) {
    if (Test-Path $VersionFile) {
        $LastVersion = Get-Content $VersionFile
        $BaseVersion = $LastVersion -replace "^v", ""
        $Parts = $BaseVersion.Split(".")
        $Major = [int]$Parts[0]
        $Minor = [int]$Parts[1]
        $Patch = [int]$Parts[2] + 1
        $Version = "v$Major.$Minor.$Patch"
    } else {
        $Version = "v0.1.0"
    }
}

# 获取 go module 名称
$RepoRoot = Resolve-Path "$ScriptDir\.."
$GoMod = Join-Path $RepoRoot "go.mod"
$ModuleName = Select-String -Path $GoMod -Pattern "^module\s+(.+)" | ForEach-Object { $_.Matches[0].Groups[1].Value.Trim() }

# 构建 ldflags 包路径
$LdPackage = "$ModuleName/apps/$AppName"

# 创建输出目录
$BinDir = Join-Path $RepoRoot "bin"
if (-not (Test-Path $BinDir)) {
    New-Item -ItemType Directory -Path $BinDir | Out-Null
}

# 编译
Write-Host "$LdPackage"
Write-Host "Building $AppName version $Version..."
Write-Host "go build -ldflags `"-X ${LdPackage}.Version=$Version`" -o `"$BinDir\$AppName.exe`" `"$AppDir`""
go build -ldflags "-X $LdPackage.version=$Version" -o "$BinDir\$AppName.exe" "$AppDir"

Write-Host "Build complete: $AppName ($Version)"

# 保存版本号
Set-Content -Path $VersionFile -Value $Version
