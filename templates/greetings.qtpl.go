// Code generated by qtc from "greetings.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// Greetings greets up to 42 names.
// It also greets John differently comparing to others.

//line greetings.qtpl:4
package templates

//line greetings.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line greetings.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line greetings.qtpl:4
func StreamGreetings(qw422016 *qt422016.Writer, names []string) {
//line greetings.qtpl:4
	qw422016.N().S(`
	`)
//line greetings.qtpl:5
	if len(names) == 0 {
//line greetings.qtpl:5
		qw422016.N().S(`
		Nobody to greet :(
		`)
//line greetings.qtpl:7
		return
//line greetings.qtpl:8
	}
//line greetings.qtpl:8
	qw422016.N().S(`

	`)
//line greetings.qtpl:10
	for i, name := range names {
//line greetings.qtpl:10
		qw422016.N().S(`
		`)
//line greetings.qtpl:11
		if i == 42 {
//line greetings.qtpl:11
			qw422016.N().S(`
			I'm tired to greet so many people...
			`)
//line greetings.qtpl:13
			break
//line greetings.qtpl:14
		} else if name == "John" {
//line greetings.qtpl:14
			qw422016.N().S(`
			`)
//line greetings.qtpl:15
			streamsayHi(qw422016, "Mr. "+name)
//line greetings.qtpl:15
			qw422016.N().S(`
			`)
//line greetings.qtpl:16
			continue
//line greetings.qtpl:17
		} else {
//line greetings.qtpl:17
			qw422016.N().S(`
			`)
//line greetings.qtpl:18
			StreamHello(qw422016, name)
//line greetings.qtpl:18
			qw422016.N().S(`
		`)
//line greetings.qtpl:19
		}
//line greetings.qtpl:19
		qw422016.N().S(`
	`)
//line greetings.qtpl:20
	}
//line greetings.qtpl:20
	qw422016.N().S(`
`)
//line greetings.qtpl:21
}

//line greetings.qtpl:21
func WriteGreetings(qq422016 qtio422016.Writer, names []string) {
//line greetings.qtpl:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line greetings.qtpl:21
	StreamGreetings(qw422016, names)
//line greetings.qtpl:21
	qt422016.ReleaseWriter(qw422016)
//line greetings.qtpl:21
}

//line greetings.qtpl:21
func Greetings(names []string) string {
//line greetings.qtpl:21
	qb422016 := qt422016.AcquireByteBuffer()
//line greetings.qtpl:21
	WriteGreetings(qb422016, names)
//line greetings.qtpl:21
	qs422016 := string(qb422016.B)
//line greetings.qtpl:21
	qt422016.ReleaseByteBuffer(qb422016)
//line greetings.qtpl:21
	return qs422016
//line greetings.qtpl:21
}

// sayHi is unexported, since it starts with lowercase letter.

//line greetings.qtpl:24
func streamsayHi(qw422016 *qt422016.Writer, name string) {
//line greetings.qtpl:24
	qw422016.N().S(`
	Hi, `)
//line greetings.qtpl:25
	qw422016.E().S(name)
//line greetings.qtpl:25
	qw422016.N().S(`
`)
//line greetings.qtpl:26
}

//line greetings.qtpl:26
func writesayHi(qq422016 qtio422016.Writer, name string) {
//line greetings.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line greetings.qtpl:26
	streamsayHi(qw422016, name)
//line greetings.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line greetings.qtpl:26
}

//line greetings.qtpl:26
func sayHi(name string) string {
//line greetings.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
//line greetings.qtpl:26
	writesayHi(qb422016, name)
//line greetings.qtpl:26
	qs422016 := string(qb422016.B)
//line greetings.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
//line greetings.qtpl:26
	return qs422016
//line greetings.qtpl:26
}

// Note that every template file may contain an arbitrary number
// of template functions. For instance, this file contains Greetings and sayHi
// functions.
//
