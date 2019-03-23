const express = require("express");
let router = express.Router();
const passport = require("passport");
const jwt = require("jsonwebtoken");

// Perform the login, after login Auth0 will redirect to callback
router.get("/login", setPreviousPageCookie,
  passport.authenticate("auth0", {
    scope: "openid email profile"
  }), 
  (req, res) => {
    console.log(req.headers('Referer'))
    res.redirect("/");
  }
);

router.get("/callback", function(req, res, next) {
  passport.authenticate("auth0", function(err, user, info) {
    if (err) { return next(err); }
    if (!user) { return res.redirect("/login"); }
    
    req.logIn(user, function(err) {
      if (err) {
        return next(err);
      }
      let token = jwt.sign({ user: user }, "shhhh");
      res.cookie("jwt-token", token).send();
      res.redirect("http://localhost:1313/redirect/");
    });
  })(req, res, next);
});

router.get("/logout", (req, res) => {
  req.logout();
  res.redirect("/");
});

function setPreviousPageCookie(req, res, next) {
  // res.cookie("redirectUrl", req.header('Referer')).send();
  next();
}

module.exports = router;
