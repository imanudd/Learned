==================================

CREATE DATABASE enigmalaundry

==================================

CREATE TABLE customers (
id SERIAL PRIMARY KEY,
name VARCHAR(50) NOT NULL,
phone_number VARCHAR(15) NOT NULL
);

CREATE TABLE services (
id SERIAL PRIMARY KEY,
service_name VARCHAR(50) NOT NULL,
unit VARCHAR(20) NOT NULL,
price INT not null
);

CREATE TABLE cassiers(
id SERIAL PRIMARY KEY,
name VARCHAR(30) NOT NULL
);

CREATE TABLE transactions(
id SERIAL PRIMARY KEY NOT NULL,
id_customer INT NOT NULL,
id_cassier INT NOT NULL,
start_date DATE,
end_date DATE,

CONSTRAINT fk_customer 
	FOREIGN KEY (id_customer)
	REFERENCES customers(id),

CONSTRAINT fk_cassier 
	FOREIGN KEY (id_cassier)
	REFERENCES cassiers(id)
);

CREATE TABLE detail_transaction(
id SERIAL PRIMARY KEY NOT NULL,
id_transaction INT NOT NULL,
id_service INT NOT NULL,
qty INT NOT NULL,
final_price INT NOT NULL,
	
CONSTRAINT fk_transaction 
	FOREIGN KEY (id_transaction)
	REFERENCES transactions(id),

CONSTRAINT fk_service
	FOREIGN KEY (id_service)
	REFERENCES services(id)
);


