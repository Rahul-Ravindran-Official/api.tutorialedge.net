const express = require("express");
let router = express.Router();
const Sequelize = require("sequelize");
const sequelize = require("../db/db.js");
const jwt = require("jsonwebtoken");
const fs = require("fs");
const winston = require('./../config/winston');

var Comment = sequelize.define("Comment", {
  body: Sequelize.STRING,
  author: Sequelize.STRING,
  slug: Sequelize.STRING,
  date: Sequelize.DATE,
  picture: Sequelize.STRING,
  userID: Sequelize.STRING,
  path: Sequelize.STRING,
});

function verifyJWT(req, res, next) {
  if (typeof req.headers.authorization !== "undefined") {
    let token = req.headers.authorization.split(" ")[1];
    var privateKey = fs.readFileSync("./private.pem", "utf8");
    jwt.verify(token, privateKey, { algorithm: "HS256" }, (err, user) => {
      if (err) {
        console.log(err);
        res.status(500).json({ error: "Not Authorized" });
        throw new Error("Not Authorized");
      }
      console.log(user);
      return next();
    });
  } else {
    throw new Error("Not Authorized");
  }
}

function isAdmin(req, res, next) {
  let token = req.headers.authorization.split(" ")[1];
  var privateKey = fs.readFileSync("./private.pem", "utf8");
  jwt.verify(token, privateKey, { algorithm: "HS256" }, (err, user) => {
    if (err) {
      console.log(err);
      res.status(500).json({ error: "Not Authorized" });
      throw new Error("Not Authorized");
    }
    console.log(user);

    return next();
  });
}

router.get("/comments", (req, res) => {
  winston.log("info", "Comments endpoint hit");
  Comment.findAll()
    .then(comments => {
      res.json({ comments: comments });
    })
    .catch(err => {
      res.json({ error: err });
    });
});

router.get("/comments/:slug", (req, res) => {
  console.log("Request for comments for Article with ID: ", req.params['slug']);
  Comment.findAll({
    where: {
      slug: req.params["slug"]
    }
  }).then(comments => {
    res.status(200).json(comments);
  });
});

router.post("/comments/:slug", verifyJWT, (req, res) => {
  Comment.create({
    body: req.body.body,
    author: req.body.author,
    path: req.body.path,
    slug: req.params.slug,
    date: new Date().getTime(),
    picture: req.body.user.user.picture,
    userID: req.body.user.user.id
  }).then(comment => {
    res.json({ status: 200, comment: comment });
  });
});

router.delete("/comments/:slug/:id", verifyJWT, isAdmin, (req, res) => {
  // Comment.destroy({
  //     where: { id: req.params.id }
  // })
  //     .then((resp) => {
  //         console.log("The Deed Is Done");
  //         res.status(200).json({"status": "success"});
  //     });
  res.status(500).json({ status: "not implemented" });
});

module.exports = router;
