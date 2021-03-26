const express = require('express')
const serveStatic = require('serve-static')
const path = require('path')
var cors = require('cors')
 

const app = express()
app.use(cors())
//here we are configuring dist to serve app files
app.use(express.static('./static'));

// this * route is to serve project on different page routes except root /
// app.get(/.*/, function (req, res) {
//     res.sendFile(path.join(dirname, 'sign-in.html'))
// })

// app.get('/', function(req, res){
//     res.sendFile(path.join(__dirname, 'sign-in.html'))
// });

const port = process.env.PORT || 8080
app.listen(port)