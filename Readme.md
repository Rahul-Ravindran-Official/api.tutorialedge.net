The Official Repository for the TutorialEdge API
=================================================

Deploying to Production

1. process.env.NODE_ENV needs to be 'production'
2. CLIENT_SECRET needs to be set as an environment variable
3. App should then start

# Running Locally

```
$ nodemon app.js
```

# Docker

```
$ docker build -t api.tutorialedge.net .
$ docker run -it -p 3000:3000 --env-file=.env.ENV api.tutorialedge.net
```