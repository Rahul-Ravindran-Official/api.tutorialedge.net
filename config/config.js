const fs = require('fs');

module.exports = {
  development: {
    username: process.env.DB_USERNAME,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_TABLE,
    host: process.env.DB_HOST,
    dialect: 'postgres',
    port: process.env.DB_PORT,
    dialectOptions: {
      ssl: {
        rejectUnauthorized: false,
        ca: fs.readFileSync(__dirname + '/ca-certificate.crt').toString()
      }
    }
  },
  production: {
    username: process.env.DB_USERNAME,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_NAME,
    host: process.env.DB_HOSTNAME,
    dialect: 'postgres',
    port: process.env.DB_PORT,
    dialectOptions: {
      ssl: {
        rejectUnauthorized: false,
        ca: fs.readFileSync(__dirname + '/ca-certificate.crt').toString()
      }
    }
  }
};