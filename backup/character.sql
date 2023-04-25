create table roles
(
    id bigserial primary key,
    name text
);

create table characters
(
    id bigserial primary key,
    role_id bigint,
    name text,
    image_url text,
    description text,
    constraint fk_character_role_id foreign key (role_id) references roles(id)
);

create table skills
(
    id bigserial primary key,
    character_id bigint,
    name text,
    image_url text,
    description text,
    button text,
    constraint fk_skill_character_id foreign key (character_id) references characters(id)
);

create table stats
(
    id bigserial primary key,
    character_id bigint,
    name text,
    value int,
    scaling int,
    constraint fk_stat_character_id foreign key (character_id) references characters(id)
);

create table formulas
(
    id bigserial primary key,
    skill_id bigint,
    level int,
    formula text,
    stats_name text[],
    constraint fk_formula_skill_id foreign key (skill_id) references skills(id)
);