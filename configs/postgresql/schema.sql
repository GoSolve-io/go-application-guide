CREATE TABLE bikes (
	id uuid NOT NULL,
	model_name varchar NULL,
	weight numeric NULL,
	price_per_h integer NULL,
	CONSTRAINT bikes_pk PRIMARY KEY (id)
);
CREATE INDEX bikes_model_name_idx ON public.bikes USING btree (model_name);

CREATE TYPE customer_type AS ENUM (
	'business',
	'individual'
);

CREATE TABLE customers (
	id uuid NOT NULL,
	first_name varchar NULL,
	surname varchar NULL,
	email varchar NOT NULL,
	"type" customer_type NOT NULL,
	CONSTRAINT customers_pk PRIMARY KEY (id)
);

CREATE TYPE reservation_status AS ENUM (
	'approved',
	'canceled'
);

CREATE TABLE reservations (
	id uuid NOT NULL,
	"status" reservation_status NOT NULL,
	start_time timestamptz(0) NOT NULL,
	end_time timestamptz(0) NOT NULL,
	bike_id uuid NOT NULL,
	customer_id uuid NOT NULL,
	total_value integer NOT NULL,
	applied_discount integer NOT NULL,
	CONSTRAINT reservations_pk PRIMARY KEY (id),
	CONSTRAINT bikes_fk FOREIGN KEY (bike_id) REFERENCES bikes(id) ON UPDATE CASCADE ON DELETE RESTRICT,
	CONSTRAINT reservations_fk FOREIGN KEY (customer_id) REFERENCES customers(id) ON UPDATE CASCADE ON DELETE RESTRICT DEFERRABLE
);
CREATE INDEX reservations_bike_timerange_idx ON public.reservations USING btree (bike_id, start_time, end_time);
