const Sequelize = require("sequelize");
const fs = require("fs");
const dotenv = require('dotenv').config();

const sequelize = new Sequelize(
    process.env.DB_TABLE, 
    process.env.DB_USERNAME, 
    process.env.DB_PASSWORD, 
    {
        dialect: "postgres",
        dialectOptions: {
        ssl: {
            rejectUnauthorized: false,
            ca: fs.readFileSync("./ca-certificate.crt").toString()
        }
    },
    host: process.env.DB_HOST,
    port: process.env.DB_PORT
});

module.exports = sequelize;