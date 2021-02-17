

CREATE TABLE bikes (
	id uuid NOT NULL,
	model_name varchar NULL,
	weight numeric NULL,
	price_per_h numeric NULL,
	CONSTRAINT bikes_pk PRIMARY KEY (id)
);
CREATE INDEX bikes_model_name_idx ON public.bikes USING btree (model_name);

CREATE TYPE customer_type AS ENUM ('business', 'individual');
CREATE TABLE customers (
	id uuid NOT NULL,
	"type" customer_type NOT NULL,
	first_name varchar NULL,
	surname varchar NULL,
	email varchar NOT NULL,
	CONSTRAINT customers_pk PRIMARY KEY (id)
);

CREATE TABLE reservations (
	id uuid NOT NULL,
	"from" timestamptz(0) NOT NULL,
	"to" timestamptz(0) NOT NULL,
	bike_id uuid NOT NULL,
	customer_id uuid NOT NULL,
	CONSTRAINT reservations_pk PRIMARY KEY (id),
	CONSTRAINT bikes_fk FOREIGN KEY (bike_id) REFERENCES bikes(id) ON UPDATE CASCADE ON DELETE RESTRICT,
	CONSTRAINT customers_fk FOREIGN KEY (customer_id) REFERENCES customers(id) ON UPDATE CASCADE ON DELETE RESTRICT
);
