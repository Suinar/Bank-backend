param(
    [string]$BaseUrl = "http://localhost:18080"
)

$ErrorActionPreference = "Stop"
$collection = Join-Path $PSScriptRoot "bank-backend.postman_collection.json"
$environment = Join-Path $PSScriptRoot "local.postman_environment.json"

npx --yes newman@6.2.1 run $collection `
    --environment $environment `
    --env-var "base_url=$BaseUrl" `
    --reporters cli

if ($LASTEXITCODE -ne 0) {
    exit $LASTEXITCODE
}
