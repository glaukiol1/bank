
const express = require('express') 
const app = express() 

const bodyParser = require('body-parser') 
const cookieParser = require('cookie-parser')

app.use(cookieParser())
app.use(bodyParser.urlencoded({
	extended: true
}));
app.use(express.json()) 
app.set('view engine', 'pug')
app.use(express.static('public')) 


app.get('/', (req,res) => { 
	res.render('index.pug') 
})

app.listen(3000, ()=>{console.log('Server Started on 127.0.0.1:3000!')})

