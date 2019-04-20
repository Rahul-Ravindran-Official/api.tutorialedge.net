const express = require("express");
const passport = require("passport");
const Auth0Strategy = require("passport-auth0");
const session = require("express-session");
const cookieParser = require('cookie-parser');
const bodyParser = require('body-parser');
const authRouter = require('./auth/auth.js');
const userRouter = require('./users/users.js');
const commentsRouter = require('./comments/comments.js');
const cors = require('cors');
const morgan = require('morgan');
const winston = require('./config/winston');
const dotenv = require('dotenv').config();

const app = express();
const strategy = new Auth0Strategy({
    domain: process.env.DOMAIN,
    clientID: process.env.CLIENT_ID,
    clientSecret: process.env.CLIENT_SECRET,
    callbackURL: process.env.CALLBACK_URL
  }, function(accessToken, refreshToken, extraParams, profile, done) {
    return done(null, profile);
});

passport.serializeUser(function(user, done) {
  done(null, user);
});

passport.deserializeUser(function(user, done) {
  done(null, user);
});

passport.use(strategy);
app.use(morgan('combined', { stream: winston.stream }));
app.use(cookieParser());
app.use(bodyParser.json())
app.use(session({ secret: "superstrongsecret", resave: true }));
app.use(cors());
app.use(passport.initialize());
app.use(passport.session());

app.get("/health", (req, res) => {
  res.status(200).json({ status: "up" });
});

app.get("/liveliness", (req, res) => {
  res.status(200).json({ status: "up" })
})

app.use('/api/v1', authRouter);
app.use('/api/v1', userRouter);
app.use('/api/v1', commentsRouter);

app.listen(process.env.PORT, () =>
  console.log(`api.tutorialedge.net listening on PORT ${process.env.PORT}!`)
);
