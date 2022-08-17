CREATE TABLE IF NOT EXISTS USER_ACCOUNT(user_email varchar primary key);

CREATE TABLE IF NOT EXISTS RELATIONSHIP(requestor varchar not null ,target varchar not null, is_friend boolean default false, friend_blocked boolean default false, subscribed boolean default false, subscribe_blocked boolean default false,
constraint pk_mail_app_recipients primary key (requestor, target),
CONSTRAINT fk_requestor_user_account FOREIGN KEY(requestor) REFERENCES USER_ACCOUNT(user_email),
CONSTRAINT fk_target_user_account FOREIGN KEY(target) REFERENCES USER_ACCOUNT(user_email),
CONSTRAINT relationship_uniq UNIQUE(requestor, target));

INSERT INTO USER_ACCOUNT(user_email) values('thehaohcm@yahoo.com.vn');
INSERT INTO USER_ACCOUNT(user_email) values('hao.nguyen@s3corp.com.vn');
INSERT INTO USER_ACCOUNT(user_email) values('thehaohcm@gmail.com');
INSERT INTO USER_ACCOUNT(user_email) values('chinh.nguyen@s3corp.com.vn');
INSERT INTO USER_ACCOUNT(user_email) values('son.le@s3corp.com.vn');
INSERT INTO USER_ACCOUNT(user_email) values('hung.tong@s3corp.com.vn');

INSERT INTO RELATIONSHIP(requestor,target, is_friend) values('thehaohcm@yahoo.com.vn','hao.nguyen@s3corp.com.vn',true);
INSERT INTO RELATIONSHIP(requestor,target, is_friend) values('thehaohcm@gmail.com','hao.nguyen@s3corp.com.vn',true);
INSERT INTO RELATIONSHIP(requestor,target, is_friend) values('hao.nguyen@s3corp.com.vn','chinh.nguyen@s3corp.com.vn',true);
INSERT INTO RELATIONSHIP(requestor,target, is_friend) values('chinh.nguyen@s3corp.com.vn','hao.nguyen@s3corp.com.vn',true);
INSERT INTO RELATIONSHIP(requestor,target, is_friend) values('thehaohcm@yahoo.com.vn','chinh.nguyen@s3corp.com.vn',true);
INSERT INTO RELATIONSHIP(requestor,target, is_friend) values('chinh.nguyen@s3corp.com.vn','hung.tong@s3corp.com.vn',true);
INSERT INTO RELATIONSHIP(requestor,target, is_friend) values('hung.tong@s3corp.com.vn','chinh.nguyen@s3corp.com.vn',true);