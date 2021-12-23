// Code generated by qtc from "pythons2s.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line pythons2s.qtpl:1
package templates

//line pythons2s.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line pythons2s.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line pythons2s.qtpl:1
func StreamPythonS2S(qw422016 *qt422016.Writer, answer7 string) {
//line pythons2s.qtpl:1
	qw422016.N().S(`

import urllib

import google.auth.transport.requests
import google.oauth2.id_token


def make_authorized_get_request(service_url):
    """
    make_authorized_get_request makes a GET request to the specified HTTP endpoint
    in service_url (must be a complete URL) by authenticating with the
    ID token obtained from the google-auth client library.
    """

    req = urllib.request.Request(service_url)

    auth_req = google.auth.transport.requests.Request()
    id_token = google.oauth2.id_token.fetch_id_token(auth_req, service_url)

    req.add_header("Authorization", f"Bearer {id_token}")
    response = urllib.request.urlopen(req)

    return response.read()

`)
//line pythons2s.qtpl:26
}

//line pythons2s.qtpl:26
func WritePythonS2S(qq422016 qtio422016.Writer, answer7 string) {
//line pythons2s.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pythons2s.qtpl:26
	StreamPythonS2S(qw422016, answer7)
//line pythons2s.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line pythons2s.qtpl:26
}

//line pythons2s.qtpl:26
func PythonS2S(answer7 string) string {
//line pythons2s.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
//line pythons2s.qtpl:26
	WritePythonS2S(qb422016, answer7)
//line pythons2s.qtpl:26
	qs422016 := string(qb422016.B)
//line pythons2s.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
//line pythons2s.qtpl:26
	return qs422016
//line pythons2s.qtpl:26
}
