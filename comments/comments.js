const express = require("express");
let router = express.Router();
const Sequelize = require("sequelize");
const sequelize = require('../db/db.js');

var Comment = sequelize.define("Comment", {
  body: Sequelize.STRING,
  author: Sequelize.STRING,
  slug: Sequelize.STRING,
  date: Sequelize.DATE
});

router.get("/comments", (req, res) => {
  Comment.findAll()
    .then(comments => {
      res.json({ comments: comments });
    })
    .catch(err => {
      res.json({ error: err });
    });
});

router.get("/comments/:slug", (req, res) => {
  Comment.findAll({
    where: {
      slug: req.params["slug"]
    }
  }).then(comments => {
    res.json(comments);
  });
});

router.post("/comments/:slug", (req, res) => {
  console.log(req);
  Comment.create({
    body: "Hello",
    author: "Elliot Forbes",
    slug: req.params.slug,
    date: new Date().getTime()
  }).then(comment => {
    res.json({ status: 200, comment: comment });
  });
});

router.delete("/comments/:slug/:id", (req, res) => {
  res.send("Dis also works");
});

module.exports = router;
