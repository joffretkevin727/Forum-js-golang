const express = require('express');
const path = require('path');
const app = express();

app.use(express.static(path.join(__dirname, 'Frontend')));
app.use(express.static(path.join(__dirname, '..', 'Backend')));

app.get('/home', (req, res) => {
    res.sendFile(path.join(__dirname, 'Frontend/home.html'));
});

app.get('/product', (req, res) => {
    res.sendFile(path.join(__dirname, 'template/product.html'));
});

app.get('/cart', (req, res) => {
    res.sendFile(path.join(__dirname, 'template/cart.html'));
});

app.get('/connexion', (req, res) => {
    res.sendFile(path.join(__dirname, 'template/connexion.html'));
});

app.get('/register', (req, res) => {
    res.sendFile(path.join(__dirname, 'template/register.html'));
});

app.get('/profil', (req, res) => {
    res.sendFile(path.join(__dirname, 'template/profil.html'));
});

app.get('/delivery', (req, res) => {
    res.sendFile(path.join(__dirname, 'template/delivery.html'));
});

app.listen(6969, () => console.log("Serveur : http://localhost:6969/home"));

app.get('/favoris', (req, res) => {
    res.sendFile(path.join(__dirname, 'template/favoris.html'));
});