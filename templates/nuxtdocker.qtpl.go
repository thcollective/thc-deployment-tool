// Code generated by qtc from "nuxtdocker.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line nuxtdocker.qtpl:1
package templates

//line nuxtdocker.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line nuxtdocker.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line nuxtdocker.qtpl:1
func StreamNuxtdocker(qw422016 *qt422016.Writer, portSelected string) {
//line nuxtdocker.qtpl:1
	qw422016.N().S(`

# build stage
FROM node:16-alpine3.14 as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build


ENV HOST 0.0.0.0
EXPOSE `)
//line nuxtdocker.qtpl:13
	qw422016.E().S(portSelected)
//line nuxtdocker.qtpl:13
	qw422016.N().S(`

# Start command
CMD [ "npm", "start" ]

`)
//line nuxtdocker.qtpl:18
}

//line nuxtdocker.qtpl:18
func WriteNuxtdocker(qq422016 qtio422016.Writer, portSelected string) {
//line nuxtdocker.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line nuxtdocker.qtpl:18
	StreamNuxtdocker(qw422016, portSelected)
//line nuxtdocker.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line nuxtdocker.qtpl:18
}

//line nuxtdocker.qtpl:18
func Nuxtdocker(portSelected string) string {
//line nuxtdocker.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
//line nuxtdocker.qtpl:18
	WriteNuxtdocker(qb422016, portSelected)
//line nuxtdocker.qtpl:18
	qs422016 := string(qb422016.B)
//line nuxtdocker.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
//line nuxtdocker.qtpl:18
	return qs422016
//line nuxtdocker.qtpl:18
}