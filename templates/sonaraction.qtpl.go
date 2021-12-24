// Code generated by qtc from "sonaraction.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line sonaraction.qtpl:1
package templates

//line sonaraction.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line sonaraction.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line sonaraction.qtpl:1
func StreamSonaraction(qw422016 *qt422016.Writer, sonarBranch string) {
//line sonaraction.qtpl:1
	qw422016.N().S(`

on:
  # Trigger analysis when pushing in master or pull requests, and when creating
  # a pull request.
  push:
    branches:
      - `)
//line sonaraction.qtpl:8
	qw422016.E().S(sonarBranch)
//line sonaraction.qtpl:8
	qw422016.N().S(`
  pull_request:
      types: [opened, synchronize, reopened]
name: Main Workflow
jobs:
  sonarcloud:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
      with:
        # Disabling shallow clone is recommended for improving relevancy of reporting
        fetch-depth: 0
    - name: SonarCloud Scan
      uses: sonarsource/sonarcloud-github-action@master
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        SONAR_TOKEN: ${{ secrets.SONAR_TOKEN }}

`)
//line sonaraction.qtpl:26
}

//line sonaraction.qtpl:26
func WriteSonaraction(qq422016 qtio422016.Writer, sonarBranch string) {
//line sonaraction.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line sonaraction.qtpl:26
	StreamSonaraction(qw422016, sonarBranch)
//line sonaraction.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line sonaraction.qtpl:26
}

//line sonaraction.qtpl:26
func Sonaraction(sonarBranch string) string {
//line sonaraction.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
//line sonaraction.qtpl:26
	WriteSonaraction(qb422016, sonarBranch)
//line sonaraction.qtpl:26
	qs422016 := string(qb422016.B)
//line sonaraction.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
//line sonaraction.qtpl:26
	return qs422016
//line sonaraction.qtpl:26
}
