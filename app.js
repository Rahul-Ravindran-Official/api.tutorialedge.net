const express = require("express");
const passport = require("passport");
const Auth0Strategy = require("passport-auth0");
const session = require("express-session");
const cookieParser = require('cookie-parser');
const bodyParser = require('body-parser');
const jwt = require("jsonwebtoken");
const authRouter = require('./auth/auth.js');
const userRouter = require('./users/users.js');
const commentsRouter = require('./comments/comments.js');
const cors = require('cors');
const dotenv = require('dotenv').config();

const app = express();
const CONFIG = require('./config.json');
const ENV = process.env.NODE_ENV || 'development';

const strategy = new Auth0Strategy(
  {
    domain: CONFIG.domain,
    clientID: CONFIG[ENV].clientID,
    clientSecret: process.env.CLIENT_SECRET,
    callbackURL: CONFIG[ENV].callbackURL
  },
  function(accessToken, refreshToken, extraParams, profile, done) {
    return done(null, profile);
  }
);

passport.serializeUser(function(user, done) {
  done(null, user);
});

passport.deserializeUser(function(user, done) {
  done(null, user);
});

var sess = {
  secret: "CHANGE THIS SECRET",
  cookie: { maxAge: 60000 },
  resave: false,
  saveUninitialized: true
};

passport.use(strategy);
app.use(cookieParser());
app.use(bodyParser.json())
app.use(session(sess));
app.use(cors());
app.use(passport.initialize());
app.use(passport.session());

app.get("/health", (req, res) => {
  res.json({ status: "up" });
});

app.use('/api/v1', authRouter);
app.use('/api/v1', userRouter);
app.use('/api/v1', commentsRouter);

app.listen(CONFIG.port, () =>
  console.log(`api.tutorialedge.net listening on port ${CONFIG.port}!`)
);
