const express = require("express")
const axios = require("axios")
const bodyParser = require("body-parser")
const cors = require("cors")

const app = express()
app.use(bodyParser.json())
app.use(cors())


app.post("/event-bus/events", (req, resp) => {
  const event = resp.body // type, data

  // post service
  axios.post("http://localhost:4001/event-bus/event/listeners", event)
    .catch( err => console.error(err))

  // comment service
  axios.post("http://localhost:4002/event-bus/event/listeners", event)
    .catch( err => console.error(err))
  
  
  // axios.post("http://localhost:")
})