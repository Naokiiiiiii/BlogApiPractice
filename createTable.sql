create table if not exists users (
  user_id integer unsigned auto_increment primary key,
  google_id integer,
  username varchar(255) not null,
  email varchar(255) not null UNIQUE,
  created_at datetime,
  updated_at datetime
);

create table if not exists articles (
  article_id integer unsigned auto_increment primary key,
  title varchar(100) not null,
  contents text not null,
  user_id integer unsigned not null,
  created_at datetime,
  updated_at datetime,
  foreign key (user_id) references users(user_id)
);

create table if not exists comments (
  comment_id integer unsigned auto_increment primary key,
  article_id integer unsigned not null,
  user_id integer unsigned not null,
  message text not null,
  created_at datetime,
  updated_at datetime,
  foreign key (article_id) references articles(article_id),
  foreign key (user_id) references users(user_id)
);

create table if not exists nices (
  nice_id integer unsigned auto_increment primary key,
  article_id integer unsigned not null,
  user_id integer unsigned not null,
  created_at datetime,
  foreign key (article_id) references articles(article_id),
  foreign key (user_id) references users(user_id)
);

-- データの挿入
insert into users (username, email, created_at, updated_at) values ('naoki', 'exsample@gmail.com', now(), now());

insert into articles (title, contents, user_id, created_at, updated_at) values ('firstPost', 'This is my first blog', 1, now(), now()); 

insert into articles (title, contents, user_id, created_at, updated_at) values ('secondPost', 'This is my second blog', 1, now(), now()); 

insert into comments (article_id, user_id, message, created_at, updated_at) values (1, 1, '1st comment yeah', now(), now());

insert into comments (article_id, user_id, message, created_at, updated_at) values (1, 1, '2nd comment yeah', now(), now());

insert into nices (article_id, user_id, created_at) values (1, 1, now());