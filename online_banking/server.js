
const express = require('express') // Require ExpressJS

const app = express() // Init the app AKA server

const bodyParser = require('body-parser') // Require Body Parser Module
const cookieParser = require('cookie-parser') // Require Cookie Parser Module

app.use(cookieParser()) // Use cookie parser module in our server
app.use(bodyParser.urlencoded({
	extended: true
})); // Use body parser module in our server
app.use(express.json()) // Use JSON in our server
app.set('view engine', 'pug') // Use the pug langauge to default view
app.use(express.static('public')) // Set the public folder to be static in our server


app.get('/', (req,res) => { // Setup GET request on the / route
	res.render('index.pug') // Render the index.pug file to the client
})

app.listen(3000, ()=>{console.log('Server Started on 127.0.0.1:3000!')}) // Listen on port 3000

