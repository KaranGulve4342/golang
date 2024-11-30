const express = require('express');
const cors = require('cors');
const app = express();


const port = 3000;

app.use(express.json());
app.use(express.urlencoded({ extended: true}));
app.use(cors());

app.get('/', (req, res) => {
    res.status(200).send("Welcome to the server");
})

app.get('/get', (req, res) => {
    res.status(200).send("Get request received");
})

app.post('/post', (req, res) => {
    let myJson = req.body;

    res.status(200).send(myJson);
})

app.post('/postform', (req, res) => {
    
    res.status(200).send(JSON.stringify(req.body));
})

app.listen(port, () => {
    console.log(`Server is running on port: ${port}`);
})
