workflow "Main workflow" {
  resolves = ["release windows/amd64"]
  on = "release"
}

action "release windows/amd64" {
  uses = "ngs/go-release.action@v1.0.1"
  env = {
    GOOS = "windows"
    GOARCH = "amd64"
  }
  secrets = ["GITHUB_TOKEN"]
}
