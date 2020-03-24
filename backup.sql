-- Table: "car-pooling-sch"."STATUS"

-- DROP TABLE "car-pooling-sch"."STATUS";

CREATE TABLE "car-pooling-sch"."STATUS"
(
    id bigint NOT NULL,
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "STATUS_pkey" PRIMARY KEY (id),
    CONSTRAINT name_unique UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE "car-pooling-sch"."STATUS"
    OWNER to postgres;

-- Table: "car-pooling-sch"."CARS"

-- DROP TABLE "car-pooling-sch"."CARS";

CREATE TABLE "car-pooling-sch"."CARS"
(
    timestamp_created timestamp with time zone,
    seats_taken integer NOT NULL,
    seats integer NOT NULL,
    id bigint NOT NULL,
    timestamp_last_updated timestamp with time zone,
    status bigint NOT NULL,
    CONSTRAINT "CARS_pkey" PRIMARY KEY (id),
    CONSTRAINT status_fk FOREIGN KEY (status)
        REFERENCES "car-pooling-sch"."STATUS" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE "car-pooling-sch"."CARS"
    OWNER to postgres;

-- Table: "car-pooling-sch"."JOURNEYS"

-- DROP TABLE "car-pooling-sch"."JOURNEYS";

CREATE TABLE "car-pooling-sch"."JOURNEYS"
(
    id bigint NOT NULL,
    people integer NOT NULL,
    status bigint NOT NULL,
    timestamp_created timestamp with time zone,
    timestamp_last_updated time with time zone,
    car bigint,
    CONSTRAINT "JOURNEYS_pkey" PRIMARY KEY (id),
    CONSTRAINT car_fk FOREIGN KEY (car)
        REFERENCES "car-pooling-sch"."CARS" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT status_fk FOREIGN KEY (status)
        REFERENCES "car-pooling-sch"."STATUS" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE "car-pooling-sch"."JOURNEYS"
    OWNER to postgres;


INSERT INTO "car-pooling-sch"."STATUS"
(id, name) VALUES (1,'WAITING');
	
INSERT INTO "car-pooling-sch"."STATUS"(
	id, name)
	VALUES (2,'JOURNEYING');
	
INSERT INTO "car-pooling-sch"."STATUS"(
	id, name)
	VALUES (3,'COMPLETED');