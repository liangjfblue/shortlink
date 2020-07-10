create schema if not exists db_short_link collate latin1_swedish_ci;

create table if not exists tb_short_link
(
	id bigint unsigned auto_increment
		primary key,
	short_id bigint unsigned not null comment '短连接id',
	created_at datetime null comment '生成时间',
	updated_at datetime null comment '更新时间',
	deleted_at datetime null comment '删除时间',
	short_code varchar(255) not null comment '短码',
	long_link_md5 varchar(36) default '' not null comment '长连接md5, 用于查询',
	long_url varchar(500) null comment '长连接',
	type tinyint default 0 null comment '类型 0-系统 1-用户自定义',
	constraint tb_short_link_short_id_uindex
		unique (short_id)
);

create index idx_long_link_md5
	on tb_short_link (long_link_md5);

create index idx_short_code
	on tb_short_link (short_code);