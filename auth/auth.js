const express = require("express");
let router = express.Router();
const passport = require("passport");
const jwt = require("jsonwebtoken");
const fs = require('fs');
const winston = require('../config/winston');

// Perform the login, after login Auth0 will redirect to callback
router.get("/login",
  passport.authenticate("auth0", {
    scope: "openid profile read:groups"
  }),
  (req, res) => { res.redirect("/"); }
);

// The route that performs the act of authenticating a new/existing
// user. They'll be assigned a jwt-token and redirected back to the
// client application which will then handle redirection to their original
// page.
router.get("/callback", function(req, res, next) {
  passport.authenticate("auth0", function(err, user, info) {
    if (err) { return next(err); }
    // if (!user) { return res.redirect("/api/v1/login"); }

    req.logIn(user, function(err) {
      if (err) { return next(err); }
      var privateKey = fs.readFileSync('./private.pem', 'utf8');
      let token = jwt.sign({ user: user }, privateKey, { algorithm: 'HS256'});
      res.redirect(process.env.REDIRECT_URL + "?token=" + token);
    });
  })(req, res, next);
});

// Logs the user out.
router.get("/logout", (req, res) => {
  req.logout();
  res.status(200).json({"status": "logged out"})
});

module.exports = router;