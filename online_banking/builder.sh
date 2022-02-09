#/bin/bash
echo "Making Project"
mkdir public
mkdir public/img
mkdir public/js
mkdir public/css
mkdir views
echo "Made Directorys & Files"
echo " 
{
	\"name\": \"expressjs-template\",
	\"version\": \"0.0.1\",
	\"description\": \"Made by @glaukiol1 on github\",
	\"main\": \"server.js\",
	\"scripts\": {
		\"start\": \"node server.js\"
	},
	\"author\": \"@glaukiol1\",
	\"license\": \"MIT\",
	\"dependencies\": {
		\"express\": \"*\",
		\"body-parser\": \"*\",
		\"cookie-parser\": \"*\",
		\"pug\": \"*\"
	}
}
" >> package.json
echo "Wrote package.json! \n"
echo "
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
" >> server.js
echo "Wrote server.js! \n"
echo "
html
	head
		title Test Server
		style
			include ../public/css/index.css
	body
		h1 Hello!
		p This is a test server! In the console you can see the javascript file output, and we also have styling!
		a(href=\"https://github.com/glaukiol1\") GitHub
		script(src=\"/js/index.js\")
" >> views/index.pug
echo "Wrote views/index.pug! \n"
echo "
console.log(\"Hello From JavaScript File!\")
" >> public/js/index.js
echo "Wrote public/js/index.js! \n"
echo "
backgroud-color: lightblue;
" >> public/css/index.css
echo "Wrote public/css/index.css! \n"
echo "\n\n\nFinished writing files! \n"
echo "Installing node modules! \n\n"
sudo npm i