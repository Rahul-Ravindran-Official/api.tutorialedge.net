const express = require("express");
const passport = require("passport");
const Auth0Strategy = require("passport-auth0");
const session = require("express-session");
const cookieParser = require('cookie-parser');
const jwt = require("jsonwebtoken");
const authRouter = require('./auth/auth.js');
const userRouter = require('./users/users.js');
const cors = require('cors');
const app = express();
const port = process.env.PORT || 3000;
const CONFIG = require('./config.json');

const strategy = new Auth0Strategy(
  {
    domain: CONFIG.domain,
    clientID: CONFIG.clientID,
    clientSecret: CONFIG.clientSecret,
    callbackURL: CONFIG.callbackURL,
    state: true
  },
  function(accessToken, refreshToken, extraParams, profile, done) {
    console.log(profile);
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
app.use(session(sess));
app.use(cors());
app.use(passport.initialize());
app.use(passport.session());

app.get("/", (req, res) => {
    res.json(jwt.verify(req.cookies['jwt-token'], 'shhhh'));
});

app.use('/', authRouter);
app.use('/api/v1', userRouter);

app.listen(port, () =>
  console.log(`api.tutorialedge.net listening on port ${port}!`)
);
