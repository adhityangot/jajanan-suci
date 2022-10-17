CREATE TABLE IF NOT EXISTS products (
	"id" smallserial NOT NULL,
	"name" varchar(255) NOT NULL DEFAULT '',
	"sku" varchar(50) NOT NULL DEFAULT '',
	"price" int NOT NULL DEFAULT 0,
	"added_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
	"updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
	CONSTRAINT products_pk PRIMARY KEY (id)
);