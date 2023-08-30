CREATE TABLE grocery_list (
	id serial primary key,
	owner varchar(255) not null,
	active boolean not null,
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp default null
);

CREATE TABLE grocery_item ( 
	id serial primary key,
	grocery_list_id int references grocery_list(id),
	name text not null,
	quantity int not null,
	checked boolean not null default false,
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp default null
);