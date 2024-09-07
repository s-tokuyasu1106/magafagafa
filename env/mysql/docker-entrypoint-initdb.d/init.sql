CREATE DATABASE IF NOT EXISTS `magafagafa` COLLATE 'utf8mb4_general_ci' ;

GRANT ALL ON `magafagafa`.* TO 'mysqluser'@'%';

USE magafagafa;

create table m_acces_users (
 m_acc_id int AUTO_INCREMENT not null primary key,
 m_user_name text,
 m_user_email text,
 m_user_status varchar(1) not null default '1',
 m_user_regist_date timestamp not null DEFAULT CURRENT_TIMESTAMP,
 m_user_update_date timestamp not null,
 m_user_auth_id int
);




FLUSH PRIVILEGES ;
