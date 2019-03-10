const express = require("express");
const passport = require("passport");
const Auth0Strategy = require("passport-auth0");
const session = require("express-session");
const cookieParser = require('cookie-parser');

const jwt = require("jsonwebtoken");


const CONFIG = require('./config.json');
const authRouter = require('./auth/auth.js');

const app = express();
const port = 3000;

const strategy = new Auth0Strategy(
  {
    domain: CONFIG.domain,
    clientID: CONFIG.clientID,
    clientSecret: CONFIG.clientSecret,
    callbackURL: CONFIG.callbackURL
  },
  function(accessToken, refreshToken, extraParams, profile, done) {
    return done(null, profile);
  }
);

passport.use(strategy);
// You can use this section to keep a smaller payload
passport.serializeUser(function(user, done) {
  done(null, user);
});

passport.deserializeUser(function(user, done) {
  done(null, user);
});

var sess = {
  secret: "CHANGE THIS SECRET",
  cookie: {},
  resave: false,
  saveUninitialized: true
};

if (app.get("env") === "production") {
  sess.cookie.secure = true; // serve secure cookies, requires https
}

app.use(cookieParser());
app.use(session(sess));
app.use(passport.initialize());
app.use(passport.session());

app.get("/", (req, res) => {
  console.log("Home API Endpoint hit");
  try {
    console.log(req.cookies['jwt-token']);
    console.log(jwt.verify(req.cookies['jwt-token'], 'shhhh'));
    res.send("hello world");
  } catch (err) {
    res.send(err)
  }
});

app.use('/', authRouter);


app.listen(port, () =>
  console.log(`api.tutorialedge.net listening on port ${port}!`)
);
