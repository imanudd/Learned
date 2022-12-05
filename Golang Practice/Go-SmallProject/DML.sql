INSERT INTO customers (id,name,phone_number)
VALUES 
(1,'Mirna','0819226810');

INSERT INTO cassiers (id, name)
VALUES 
(1,'iman'),
(2,'jamil');

INSERT INTO services (id,service_name,unit,price)
VALUES 
(1,'Cuci + Strika','KG',7000),
(2,'Laundry Bedcover','buah',50000),
(3,'laundry Boneka','buah',25000);

INSERT INTO transactions (id,id_customer,id_cassier,start_date,end_date)
VALUES
(1,1,1,'2022-01-01','2022-01-03');

INSERT INTO detail_transaction (id_transaction,id_service,qty)
VALUES 
(1,1,2);