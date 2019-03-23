const express = require('express');
let router = express.Router();

router.get('/comments/:id', (req, res) => {
    res.send("All Comments for a Post");
});

router.post('/comments/:slug', (req, res) => {
    res.send("Dis works");
})

router.delete('/comments/:slug/:id', (req, res) => {
    res.send("Dis also works");
})