create schema if not exists db_short_link collate latin1_swedish_ci;

create table if not exists tb_customize_short_link
(
	id bigint unsigned auto_increment
		primary key,
	created_at datetime null,
	updated_at datetime null,
	deleted_at datetime null,
	short_id bigint unsigned not null,
	short_code varchar(255) not null,
	long_url varchar(255) null
);

create index idx_long_url
	on tb_customize_short_link (long_url);

create index idx_short_code
	on tb_customize_short_link (short_code);

create index idx_tb_customize_short_link_deleted_at
	on tb_customize_short_link (deleted_at);

create table if not exists tb_short_link
(
	short_id bigint unsigned auto_increment
		primary key,
	created_at datetime null,
	updated_at datetime null,
	deleted_at datetime null,
	short_code varchar(255) not null,
	long_link_md5 varchar(36) default '' not null comment '长连接md5, 用于查询',
	long_url varchar(500) null
);

create index idx_long_link_md5
	on tb_short_link (long_link_md5);

create index idx_short_code
	on tb_short_link (short_code);

