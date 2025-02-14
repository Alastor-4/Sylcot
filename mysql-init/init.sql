CREATE DATABASE IF NOT EXISTS sylcot;
ALTER USER 'sylcot' @'%' IDENTIFIED WITH mysql_native_password BY 'sylcot';
GRANT ALL PRIVILEGES ON sylcot.* TO 'sylcot' @'%';
FLUSH PRIVILEGES;