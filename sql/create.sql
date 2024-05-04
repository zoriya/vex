create table if not exists users(
	id uuid not null primary key,
	name text not null,
	password varchar(100) not null,
	email text not null
);

create table if not exists feeds(
	id uuid not null primary key,
	name text not null,
	link text not null,
	favicon_url text not null,
	tags text[] not null,
	submitter_id uuid not null references users(id)
);

create table if not exists entries(
	id uuid not null primary key,
	feed_id uuid not null references feed(id),
	title text not null,
	link text not null,
	date timestamp with time zone not null,
	content text not null,
	author text
);

create table if not exists entries_users(
	user_id uuid not null references users(id),
	feed_id uuid not null references feeds(id),
	is_read bool not null,
	is_bookmarked bool not null,
	is_read_later bool not null,
	is_ignored bool not null,
	constraint entries_users_pk primary keys(user_id, feed_id)
);

