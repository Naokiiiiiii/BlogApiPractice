create table if not exists users (
  user_id integer unsigned auto_increment primary key,
  google_id varchar(255),
  username varchar(255) not null,
  email varchar(255) NOT NULL UNIQUE,
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
  nice_id integer unsigned auto_increment
  article_id integer unsigned not null
  user_id integer unsigned not null
  foreign key (article_id) references articles(article_id)
  foreign key (user_id) references users(user_id)
)