// Code generated by qtc from "vuedocker.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line vuedocker.qtpl:1
package templates

//line vuedocker.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line vuedocker.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line vuedocker.qtpl:1
func StreamVuedocker(qw422016 *qt422016.Writer, portNo string) {
//line vuedocker.qtpl:1
	qw422016.N().S(`
# build stage
FROM node:16-alpine3.14 as build-stage
ENV PORT=`)
//line vuedocker.qtpl:4
	qw422016.E().S(portNo)
//line vuedocker.qtpl:4
	qw422016.N().S(`
ENV HOST 0.0.0.0
RUN apk update && apk add python make g++
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build


# production stage
FROM nginx:stable-alpine as production-stage
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE `)
//line vuedocker.qtpl:18
	qw422016.E().S(portNo)
//line vuedocker.qtpl:18
	qw422016.N().S(`
CMD ["nginx", "-g", "daemon off;"]

`)
//line vuedocker.qtpl:21
}

//line vuedocker.qtpl:21
func WriteVuedocker(qq422016 qtio422016.Writer, portNo string) {
//line vuedocker.qtpl:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line vuedocker.qtpl:21
	StreamVuedocker(qw422016, portNo)
//line vuedocker.qtpl:21
	qt422016.ReleaseWriter(qw422016)
//line vuedocker.qtpl:21
}

//line vuedocker.qtpl:21
func Vuedocker(portNo string) string {
//line vuedocker.qtpl:21
	qb422016 := qt422016.AcquireByteBuffer()
//line vuedocker.qtpl:21
	WriteVuedocker(qb422016, portNo)
//line vuedocker.qtpl:21
	qs422016 := string(qb422016.B)
//line vuedocker.qtpl:21
	qt422016.ReleaseByteBuffer(qb422016)
//line vuedocker.qtpl:21
	return qs422016
//line vuedocker.qtpl:21
}
