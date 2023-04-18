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

-- select (characters.id, roles.name, characters.description, characters.image_url, characters.name)
--     from characters join roles
--     on roles.id = characters.role_id
--         where characters.id = $1;

-- select stats.id, stats.name, stats.scaling, stats.value from stats
--     where stats.character_id = $1;

-- select skills.id, skills.name, skills.image_url, skills.description, skills.button from skills
--     where skills.character_id = $1
--     order by skills.id;

select formulas.id, formulas.level, formulas.formula, formulas.stats_name from formulas
       where formulas.skill_id = $1;