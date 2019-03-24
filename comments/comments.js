const express = require('express');
const Sequelize = require('sequelize');
const fs = require("fs");

const sequelize = new Sequelize('defaultdb', 'doadmin', 'xclffhc3lz55ubuq', {
    dialect: 'postgres',
    dialectOptions: {
        ssl: {
            rejectUnauthorized: false,
            ca : fs.readFileSync("./ca-certificate.crt").toString(),
        }
    },
    host: 'db-postgresql-nyc1-45572-do-user-1074347-0.db.ondigitalocean.com',
    port: 25060
})
let router = express.Router();

var Comment = sequelize.define('Comment', {
    body: Sequelize.STRING,
    author: Sequelize.STRING,
    slug: Sequelize.STRING,
    date: Sequelize.DATE
});

router.get('/comments', (req, res) => {
    sequelize.sync()
        .then(() => Comment.findAll())
        .then(comments => {
            res.json({comments: comments})
        })
        .catch(err => {
            res.json({error: err});
        })
})

router.get('/comments/:slug', (req, res) => {
    sequelize.sync()
        .then(() => Comment.findAll({
            where: {
                slug: req.params['slug']
            }
        }))
        .then(comments => {
            res.json(comments);
        })
});

router.post('/comments/:slug', (req, res) => {

    console.log(req.body);

    sequelize.sync()
        .then(() => Comment.create({
            body: 'Hello',
            author: 'Elliot Forbes',
            slug: req.params.slug,
            date: new Date().getTime()
        }))
        .then(comment => {
            res.json({"status": 200, comment: comment});
        })
})

router.delete('/comments/:slug/:id', (req, res) => {
    res.send("Dis also works");
})

module.exports = router;