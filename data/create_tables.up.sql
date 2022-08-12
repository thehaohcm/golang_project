CREATE TABLE IF NOT EXISTS friends(user_email varchar not null, friend_email varchar not null, blocked bool default false);
ALTER TABLE public.friends ADD CONSTRAINT friends_un UNIQUE (user_email, friend_email);
CREATE TABLE IF NOT EXISTS subscribers(requestor varchar not null, target varchar not null, blocked bool default false);
ALTER TABLE public.subscribers ADD CONSTRAINT subscribers_un UNIQUE (requestor, target);
INSERT INTO subscribers VALUES('thehaohcm@yahoo.com.vn','hao.nguyen@s3corp.com.vn',false);
INSERT INTO subscribers VALUES('thehaohcm@yahoo.com.vn','thehaohcm@gmail.com',false);

INSERT INTO friends VALUES('thehaohcm@yahoo.com.vn','hao.nguyen@s3corp.com.vn',false);
INSERT INTO friends VALUES('thehaohcm@gmail.com','hao.nguyen@s3corp.com.vn',false);
INSERT INTO friends VALUES('hao.nguyen@s3corp.com.vn','chinh.nguyen@s3corp.com.vn',false);
INSERT INTO friends VALUES('chinh.nguyen@s3corp.com.vn','hao.nguyen@s3corp.com.vn',false);
INSERT INTO friends VALUES('thehaohcm@yahoo.com.vn','chinh.nguyen@s3corp.com.vn',true);