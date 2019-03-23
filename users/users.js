const express = require('express');
let router = express.Router();
const jwt = require("jsonwebtoken");

router.get('/users', verifyRequest, (req, res) => {
    res.send("All Users");
})

router.get('/user/:id', (req, res) => {
    res.send("A Single User");
})

router.get('/authors', (req, res) => {
    res.send("All Authors");
})


function verifyRequest(req, res, next) {
    const bearerHeader = req.headers['Authorization'];
    if(typeof bearerHeader !== 'undefined') {
        const bearer = bearerHeader.split(' ');
        const bearerToken = bearer[1];
        req.token = bearerToken;
        next();
    } else {
        res.sendStatus(403);
    }
}

module.exports = router;