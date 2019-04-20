FROM node:9-slim
WORKDIR /app
COPY package.json /app
RUN npm install
COPY . /app
RUN NODE_ENV=production
CMD ["node", "app.js"]