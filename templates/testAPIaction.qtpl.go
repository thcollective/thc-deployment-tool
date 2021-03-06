// Code generated by qtc from "testAPIaction.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line testAPIaction.qtpl:1
package templates

//line testAPIaction.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line testAPIaction.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line testAPIaction.qtpl:1
func StreamTestAPIaction(qw422016 *qt422016.Writer, testBranch string, ApiFolderSelected string, ApiFilesSelected string, ApiEnvSelected string) {
//line testAPIaction.qtpl:1
	qw422016.N().S(`
# Copyright 2021 The Hacker Collective, LLC.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

name: test-api-collection
on:
  push:
    branches: [ `)
//line testAPIaction.qtpl:19
	qw422016.E().S(testBranch)
//line testAPIaction.qtpl:19
	qw422016.N().S(` ]
jobs:
  test-api:
    runs-on: ubuntu-latest
    steps:
    # Checks-out your repository under $GITHUB_WORKSPACE, so your job can access it
    - uses: actions/checkout@v2
      
    # INstall Node on the runner
    - name: Install Node
      uses: actions/setup-node@v1
      with: 
        node-version: '12.x'
    
    # Install the newman command line utility and also install the html extra reporter
    - name: Install newman
      run: |
       npm install -g newman
       npm install -g newman-reporter-htmlextra
    # Make directory to upload the test results
    - name: Make Directory for results
      run: mkdir -p testResults

    # Run the API collection
    - name: Run API collection
      run: |
       newman run ./`)
//line testAPIaction.qtpl:45
	qw422016.E().S(ApiFolderSelected)
//line testAPIaction.qtpl:45
	qw422016.N().S(`/`)
//line testAPIaction.qtpl:45
	qw422016.E().S(ApiFilesSelected)
//line testAPIaction.qtpl:45
	qw422016.N().S(` -e ./`)
//line testAPIaction.qtpl:45
	qw422016.E().S(ApiFolderSelected)
//line testAPIaction.qtpl:45
	qw422016.N().S(`/`)
//line testAPIaction.qtpl:45
	qw422016.E().S(ApiEnvSelected)
//line testAPIaction.qtpl:45
	qw422016.N().S(` + `)
//line testAPIaction.qtpl:45
	qw422016.N().S("`")
//line testAPIaction.qtpl:45
	qw422016.N().S(` -r htmlextra --reporter-hmlextra-export testResults/htmlreport.html --reporter-htmlextra-darkTheme  > testResults/runreport1.html
    # Upload the contents of Test Results directory to workspace
    - name: Output the run Details
      uses: actions/upload-artifact@v2
      with: 
       name: RunReports
       path: testResults
`)
//line testAPIaction.qtpl:52
}

//line testAPIaction.qtpl:52
func WriteTestAPIaction(qq422016 qtio422016.Writer, testBranch string, ApiFolderSelected string, ApiFilesSelected string, ApiEnvSelected string) {
//line testAPIaction.qtpl:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line testAPIaction.qtpl:52
	StreamTestAPIaction(qw422016, testBranch, ApiFolderSelected, ApiFilesSelected, ApiEnvSelected)
//line testAPIaction.qtpl:52
	qt422016.ReleaseWriter(qw422016)
//line testAPIaction.qtpl:52
}

//line testAPIaction.qtpl:52
func TestAPIaction(testBranch string, ApiFolderSelected string, ApiFilesSelected string, ApiEnvSelected string) string {
//line testAPIaction.qtpl:52
	qb422016 := qt422016.AcquireByteBuffer()
//line testAPIaction.qtpl:52
	WriteTestAPIaction(qb422016, testBranch, ApiFolderSelected, ApiFilesSelected, ApiEnvSelected)
//line testAPIaction.qtpl:52
	qs422016 := string(qb422016.B)
//line testAPIaction.qtpl:52
	qt422016.ReleaseByteBuffer(qb422016)
//line testAPIaction.qtpl:52
	return qs422016
//line testAPIaction.qtpl:52
}
