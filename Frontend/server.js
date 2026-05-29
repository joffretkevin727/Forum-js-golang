const express = require('express');
const path = require('path');
const app = express();

app.use(express.static(path.join(__dirname, 'Frontend')));
app.use(express.static(path.join(__dirname, 'Public')));

app.get('/home', (req, res) => {
    res.sendFile(path.join(__dirname, 'Public/home.html'));
});

app.get('/login', (req, res) => {
    res.sendFile(path.join(__dirname, 'Public/login.html'));
});

app.get('/newTopic', (req, res) => {
    res.sendFile(path.join(__dirname, 'Public/create_topic.html'));
});

app.get('/profil', (req, res) => {
    res.sendFile(path.join(__dirname, 'Public/profil.html'));
});

app.listen(6969, () => console.log("Serveur : http://localhost:6969/home"));