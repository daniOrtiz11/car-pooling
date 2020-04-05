CREATE SCHEMA "table-booking-sch";

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE "table-booking-sch"."BOOKINGS" (
    id bigint NOT NULL,
    people integer NOT NULL,
    status bigint NOT NULL,
    timestamp_created timestamp with time zone,
    timestamp_last_updated time with time zone,
    "table" bigint
);


CREATE TABLE "table-booking-sch"."STATUS" (
    id bigint NOT NULL,
    name character varying(50) NOT NULL
);

CREATE TABLE "table-booking-sch"."TABLES" (
    timestamp_created timestamp with time zone,
    seats integer NOT NULL,
    id bigint NOT NULL,
    timestamp_last_updated timestamp with time zone,
    status bigint NOT NULL
);

ALTER TABLE ONLY "table-booking-sch"."TABLES"
    ADD CONSTRAINT "CARS_pkey" PRIMARY KEY (id);

ALTER TABLE ONLY "table-booking-sch"."BOOKINGS"
    ADD CONSTRAINT "JOURNEYS_pkey" PRIMARY KEY (id);

ALTER TABLE ONLY "table-booking-sch"."STATUS"
    ADD CONSTRAINT "STATUS_pkey" PRIMARY KEY (id);

ALTER TABLE ONLY "table-booking-sch"."STATUS"
    ADD CONSTRAINT name_unique UNIQUE (name);

ALTER TABLE ONLY "table-booking-sch"."TABLES"
    ADD CONSTRAINT status_fk FOREIGN KEY (status) REFERENCES "table-booking-sch"."STATUS"(id) NOT VALID;

ALTER TABLE ONLY "table-booking-sch"."BOOKINGS"
    ADD CONSTRAINT status_fk FOREIGN KEY (status) REFERENCES "table-booking-sch"."STATUS"(id) NOT VALID;