CREATE TABLE "public.PTOBuilder.hero" (
                                          "id" serial NOT NULL,
                                          "name" BINARY NOT NULL,
                                          "img_url" BINARY NOT NULL,
                                          "role_id" BINARY NOT NULL,
                                          "description" BINARY NOT NULL,
                                          "scale_phys_dmg" FLOAT NOT NULL,
                                          "scale_mag_dmg" FLOAT NOT NULL,
                                          "scale_phys_armor" FLOAT NOT NULL,
                                          "scale_mag_armor" FLOAT NOT NULL,
                                          "scale_health" FLOAT NOT NULL,
                                          "scale_attack_spd" FLOAT NOT NULL,
                                          "create_time" TIME NOT NULL,
                                          "update_time" TIME NOT NULL,
                                          CONSTRAINT "PTOBuilder.hero_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "public.PTOBuilder.role" (
                                          "id" serial NOT NULL,
                                          "name" BINARY NOT NULL,
                                          "create_time" DATETIME NOT NULL,
                                          "update_time" DATETIME NOT NULL,
                                          CONSTRAINT "PTOBuilder.role_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "public.PTOBuilder.hero_skill" (
                                                "id" serial NOT NULL,
                                                "name" BINARY NOT NULL,
                                                "hero_id" integer NOT NULL,
                                                "level" integer NOT NULL,
                                                "img_url" BINARY NOT NULL,
                                                "description" BINARY NOT NULL,
                                                "button" BINARY NOT NULL,
                                                "create_time" DATETIME NOT NULL,
                                                "update_time" BINARY NOT NULL,
                                                CONSTRAINT "PTOBuilder.hero_skill_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "public.PTOBuilder.hero_stats" (
                                                "id" serial NOT NULL,
                                                "hero_id" integer NOT NULL,
                                                "stat_id" integer NOT NULL,
                                                "value" FLOAT NOT NULL,
                                                "create_time" DATETIME NOT NULL,
                                                "update_time" DATETIME NOT NULL,
                                                CONSTRAINT "PTOBuilder.hero_stats_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "public.PTOBuilder.stat" (
                                          "id" serial NOT NULL,
                                          "name" BINARY NOT NULL,
                                          "create_time" DATETIME NOT NULL,
                                          "update_time" DATETIME NOT NULL,
                                          CONSTRAINT "PTOBuilder.stat_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "public.PTOBuilder.formula" (
                                             "id" serial NOT NULL,
                                             "hero_skill_id" integer NOT NULL,
                                             "stat_id" integer NOT NULL,
                                             "formula" BINARY NOT NULL,
                                             CONSTRAINT "PTOBuilder.formula_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );



ALTER TABLE "PTOBuilder.hero" ADD CONSTRAINT "PTOBuilder.hero_fk0" FOREIGN KEY ("role_id") REFERENCES "PTOBuilder.role"("id");


ALTER TABLE "PTOBuilder.hero_skill" ADD CONSTRAINT "PTOBuilder.hero_skill_fk0" FOREIGN KEY ("hero_id") REFERENCES "PTOBuilder.hero"("id");

ALTER TABLE "PTOBuilder.hero_stats" ADD CONSTRAINT "PTOBuilder.hero_stats_fk0" FOREIGN KEY ("hero_id") REFERENCES "PTOBuilder.hero"("id");
ALTER TABLE "PTOBuilder.hero_stats" ADD CONSTRAINT "PTOBuilder.hero_stats_fk1" FOREIGN KEY ("stat_id") REFERENCES "PTOBuilder.stat"("id");


ALTER TABLE "PTOBuilder.formula" ADD CONSTRAINT "PTOBuilder.formula_fk0" FOREIGN KEY ("hero_skill_id") REFERENCES "PTOBuilder.hero_skill"("id");
ALTER TABLE "PTOBuilder.formula" ADD CONSTRAINT "PTOBuilder.formula_fk1" FOREIGN KEY ("stat_id") REFERENCES "PTOBuilder.stat"("id");






