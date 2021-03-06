// Code generated by qtc from "redpanda_sub.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line redpanda_sub.qtpl:1
package templates

//line redpanda_sub.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line redpanda_sub.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line redpanda_sub.qtpl:1
func StreamRedpandaSubscribe(qw422016 *qt422016.Writer, appName string, redpandaTopic string) {
//line redpanda_sub.qtpl:1
	qw422016.N().S(`
const { Kafka } = require('kafkajs')
const { v4 } = require('uuid')
const { SchemaRegistry } = require('@kafkajs/confluent-schema-registry')
require("dotenv").config()

const kafka = new Kafka({
  clientId: '`)
//line redpanda_sub.qtpl:8
	qw422016.E().S(appName)
//line redpanda_sub.qtpl:8
	qw422016.N().S(`',
  brokers: [process.env.BROKER_URL1, process.env.BROKER_URL2, process.env.BROKER_URL3],

})

const registry = new SchemaRegistry({ host: process.env.SCHEMA_URL1 })

async function sub() {

  const consumer = kafka.consumer({ groupId: v4() })

  await consumer.connect()

  await consumer.subscribe({ topic: '`)
//line redpanda_sub.qtpl:21
	qw422016.E().S(redpandaTopic)
//line redpanda_sub.qtpl:21
	qw422016.N().S(`', fromBeginning: true })

  let msgCount = 0

  await consumer.run({
    eachMessage: async ({ topic, partition, message }) => {
      const decodedValue = await registry.decode(message.value)
      const payload =
      {
        topic: topic,
        payload: decodedValue,
        partition: partition,
        offset: message.offset,
        timestamp: message.timestamp,
      }

      console.log(payload)

    },
  })

}

sub()

`)
//line redpanda_sub.qtpl:46
}

//line redpanda_sub.qtpl:46
func WriteRedpandaSubscribe(qq422016 qtio422016.Writer, appName string, redpandaTopic string) {
//line redpanda_sub.qtpl:46
	qw422016 := qt422016.AcquireWriter(qq422016)
//line redpanda_sub.qtpl:46
	StreamRedpandaSubscribe(qw422016, appName, redpandaTopic)
//line redpanda_sub.qtpl:46
	qt422016.ReleaseWriter(qw422016)
//line redpanda_sub.qtpl:46
}

//line redpanda_sub.qtpl:46
func RedpandaSubscribe(appName string, redpandaTopic string) string {
//line redpanda_sub.qtpl:46
	qb422016 := qt422016.AcquireByteBuffer()
//line redpanda_sub.qtpl:46
	WriteRedpandaSubscribe(qb422016, appName, redpandaTopic)
//line redpanda_sub.qtpl:46
	qs422016 := string(qb422016.B)
//line redpanda_sub.qtpl:46
	qt422016.ReleaseByteBuffer(qb422016)
//line redpanda_sub.qtpl:46
	return qs422016
//line redpanda_sub.qtpl:46
}
