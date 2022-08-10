CREATE TABLE IF NOT EXISTS users(email varchar not null primary key);
CREATE TABLE IF NOT EXISTS friends(user_email varchar not null, friend_email varchar not null, blocked bool default false,
foreign key (user_email) references users(email),
foreign key (friend_email) references users(email));
CREATE TABLE IF NOT EXISTS subscribers(requestor varchar not null, target varchar not null, blocked bool default false,
primary key (requestor, target),
foreign key (requestor) references users(email),
foreign key (target) references users(email));


INSERT INTO users VALUES('thehaohcm@yahoo.com.vn');
INSERT INTO users VALUES('hao.nguyen@s3corp.com.vn');
INSERT INTO users VALUES('thehaohcm@gmail.com');
INSERT INTO users VALUES('chinh.nguyen@s3corp.com.vn');
INSERT INTO users VALUES('hung.tong@s3corp.com.vn');
INSERT INTO users VALUES('abc@def.com');
INSERT INTO users VALUES('abc1#@def.com');
INSERT INTO users VALUES('abc1@def.com');

INSERT INTO subscribers VALUES('thehaohcm@yahoo.com.vn','hao.nguyen@s3corp.com.vn',false);
INSERT INTO subscribers VALUES('thehaohcm@yahoo.com.vn','thehaohcm@gmail.com',false);

INSERT INTO friends VALUES('thehaohcm@yahoo.com.vn','hao.nguyen@s3corp.com.vn',false);
INSERT INTO friends VALUES('thehaohcm@gmail.com','hao.nguyen@s3corp.com.vn',false);
INSERT INTO friends VALUES('hao.nguyen@s3corp.com.vn','chinh.nguyen@s3corp.com.vn',false);
INSERT INTO friends VALUES('chinh.nguyen@s3corp.com.vn','hao.nguyen@s3corp.com.vn',false);
INSERT INTO friends VALUES('thehaohcm@yahoo.com.vn','chinh.nguyen@s3corp.com.vn',true);