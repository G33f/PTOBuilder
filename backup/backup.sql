CREATE DATABASE PTOBuilder

CREATE TABLE PTOBuilder.role(
    id int primary key,
    name text not null,
    creat_time timestamp,
    update_time timestamp
)

CREATE TABLE PTOBuilder.hero(
    id int primary key,
    name text not null,
    img_url text not null,
    role_id int foreign key references PTOBuilder.role(id),
    description text not null,
    scale_phys_dmg float,
    scale_mag_dmg float,
    scale_phys_armor float,
    scale_mag_armor float,
    scale_health float,
    scale_attack_spd float,
    creat_time timestamp,
    update_time timestamp
)

CREATE TABLE PTOBuilder.stat(
    id int primary key,
    name text not null,
    creat_time timestamp,
    update_time timestamp
)

CREATE TABLE PTOBuilder.hero_stats(
    id int primary key,
    hero_id int foreign key references PTOBuilder.hero(id),
    stat_id int foreign key references PTOBuilder.stat(id),
    value int default 0,
    creat_time timestamp,
    update_time timestamp
)

CREATE TABLE PTOBuilder.hero_skill(
    id int primary key,
    name text not null,
    hero_id int foreign key references PTOBuilder.hero(id),
    img_url text,
    description text,
    button text,
    creat_time timestamp,
    update_time timestamp
)

CREATE TABLE PTOBuilder.hero_skill_stat(
    id int primary key,
    hero_skill_id int foreign key references PTOBuilder.hero_skill(id),
    stat_id int foreign key references PTOBuilder.stat(id),
    value float default 0.0,
    creat_time timestamp,
    update_time timestamp
)

CREATE TABLE PTOBuilder.formula(
    id int primary key,
    format text,
    variables float[]
);