const express = require("express")
const axios = require("axios")
const bodyParser = require("body-parser")
const cors = require("cors")

const app = express()
app.use(bodyParser.json())
app.use(cors())


const events = [];

app.post("/event-bus/events", (req, resp) => {
  const event = req.body // event will have type, data

  events.push(event)

  // broadcasting event to every service
  // post service
  axios.post("http://localhost:4001/api/v1/event-bus/events/listeners", event)
    .catch( err => console.error(err))

  // comment service
  axios.post("http://localhost:4002/api/v1/event-bus/events/listeners", event)
    .catch( err => console.error(err))
  
  // query service
  axios.post("http://localhost:4003/event-bus/events/listeners", event)
    .catch( err => console.error(err))

    resp.send({})
})

app.get("/event-bus/events", (req, resp) => {
  resp.send(events)
})

app.listen(4005, () => console.log("Event bus started :4005"))