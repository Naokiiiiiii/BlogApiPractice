create table if not exists users (
  user_id integer unsigned auto_increment primary key,
  google_id varchar(255),
  username varchar(255) not null,
  email varchar(255) not null UNIQUE,
  created_at datetime
);

create table if not exists articles (
  article_id integer unsigned auto_increment primary key,
  title varchar(100) not null,
  contents text not null,
  user_id integer unsigned not null,
  created_at datetime,
  foreign key (user_id) references users(user_id)
);

create table if not exists comments (
  comment_id integer unsigned auto_increment primary key,
  article_id integer unsigned not null,
  user_id integer unsigned not null,
  message text not null,
  created_at datetime,
  foreign key (article_id) references articles(article_id),
  foreign key (user_id) references users(user_id)
);

create table if not exists nice (
  nice_id integer unsigned auto_increment primary key,
  article_id integer unsigned not null,
  user_id integer unsigned not null,
  created_at datetime,
  foreign key (article_id) references articles(article_id),
  foreign key (user_id) references users(user_id)
);

-- データの挿入
insert into users (username, email, created_at) values ('naoki', 'exsample@gmail.com', now());

insert into articles (title, contents, user_id, created_at) values ('firstPost', 'This is my first blog', 1, now()); 

insert into articles (title, contents, user_id, created_at) values ('secondPost', 'This is my second blog', 1, now()); 

insert into comments (article_id, user_id, message, created_at) values (1, 1, '1st comment yeah', now());

insert into comments (article_id, user_id, message, created_at) values (1, 1, '2nd comment yeah', now());

insert into nice (article_id, user_id, created_at) values (1, 1, now());