const dotenv = require('dotenv');
dotenv.config();

module.exports = {
  // ...other configures
  mailer: {
    naverUser: process.env.MAIL_USER,
    naverPassword: process.env.MAIL_PW
  },
  jwtSecret: process.env.JWT_SECRET,
  port: process.env.PORT,
  db: {
    host: process.env.DB_HOST,
    port: process.env.DB_PORT,
    user: process.env.DB_USER,
    password: process.env.DB_PASS,
    database: process.env.DB_NAME,
},
adminpw: process.env.ADMINPW,

};

