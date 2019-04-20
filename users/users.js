const express = require('express');
let router = express.Router();
const jwt = require("jsonwebtoken");
var ManagementClient = require('auth0').ManagementClient;
const dotenv = require('dotenv').config();

let auth0 = new ManagementClient({
    domain: process.env.DOMAIN,
    clientId: process.env.CLIENT_ID,
    clientSecret: process.env.CLIENT_SECRET,
    scope: 'read:users update:users'
});
  

router.get('/users', (req, res) => {
    auth0.getUsers()
        .then(users => {
            res.json({ users: users })
        })
        .catch(err => {
            res.status(500).json({error: err});
        })
});


router.get('/user/:id', (req, res) => {
    res.send("A Single User");
});

module.exports = router;