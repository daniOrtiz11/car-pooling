-- Table: "table-booking-sch"."STATUS"

-- DROP TABLE "table-booking-sch"."STATUS";

CREATE TABLE "table-booking-sch"."STATUS"
(
    id bigint NOT NULL,
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "STATUS_pkey" PRIMARY KEY (id),
    CONSTRAINT name_unique UNIQUE (name)
)

TABLESPACE pg_default;

ALTER TABLE "table-booking-sch"."STATUS"
    OWNER to postgres;

-- Table: "table-booking-sch"."TABLES"

-- DROP TABLE "table-booking-sch"."TABLES";

CREATE TABLE "table-booking-sch"."TABLES"
(
    timestamp_created timestamp with time zone,
    seats_taken integer NOT NULL,
    seats integer NOT NULL,
    id bigint NOT NULL,
    timestamp_last_updated timestamp with time zone,
    status bigint NOT NULL,
    CONSTRAINT "TABLES_pkey" PRIMARY KEY (id),
    CONSTRAINT status_fk FOREIGN KEY (status)
        REFERENCES "table-booking-sch"."STATUS" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE "table-booking-sch"."TABLES"
    OWNER to postgres;

-- Table: "table-booking-sch"."BOOKINGS"

-- DROP TABLE "table-booking-sch"."BOOKINGS";

CREATE TABLE "table-booking-sch"."BOOKINGS"
(
    id bigint NOT NULL,
    people integer NOT NULL,
    status bigint NOT NULL,
    timestamp_created timestamp with time zone,
    timestamp_last_updated time with time zone,
    table bigint,
    CONSTRAINT "BOOKINGS_pkey" PRIMARY KEY (id),
    CONSTRAINT table_fk FOREIGN KEY (table)
        REFERENCES "table-booking-sch"."TABLES" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT status_fk FOREIGN KEY (status)
        REFERENCES "table-booking-sch"."STATUS" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE "table-booking-sch"."BOOKINGS"
    OWNER to postgres;


INSERT INTO "table-booking-sch"."STATUS"
(id, name) VALUES (1,'WAITING');
	
INSERT INTO "table-booking-sch"."STATUS"(
	id, name)
	VALUES (2,'EATING');
	
INSERT INTO "table-booking-sch"."STATUS"(
	id, name)
	VALUES (3,'COMPLETED');