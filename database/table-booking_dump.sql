CREATE SCHEMA "docker-booking-sch";
/*
CREATE TABLE BOOKINGS (
    id bigint NOT NULL,
    people integer NOT NULL,
    status bigint NOT NULL,
    timestamp_created timestamp with time zone,
    timestamp_last_updated time with time zone,
    table bigint
);


CREATE TABLE STATUS (
    id bigint NOT NULL,
    name character varying(50) NOT NULL
);

CREATE TABLE TABLES (
    timestamp_created timestamp with time zone,
    seats integer NOT NULL,
    id bigint NOT NULL,
    timestamp_last_updated timestamp with time zone,
    status bigint NOT NULL
);

ALTER TABLE ONLY TABLES
    ADD CONSTRAINT CARS_pkey PRIMARY KEY (id);

ALTER TABLE ONLY BOOKINGS
    ADD CONSTRAINT JOURNEYS_pkey PRIMARY KEY (id);

ALTER TABLE ONLY STATUS
    ADD CONSTRAINT STATUS_pkey PRIMARY KEY (id);

ALTER TABLE ONLY STATUS
    ADD CONSTRAINT name_unique UNIQUE (name);

ALTER TABLE ONLY TABLES
    ADD CONSTRAINT status_fk FOREIGN KEY (status) REFERENCES STATUS(id) NOT VALID;

ALTER TABLE ONLY BOOKINGS
    ADD CONSTRAINT status_fk FOREIGN KEY (status) REFERENCES STATUS(id) NOT VALID;
    */