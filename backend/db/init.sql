create table measurments(
	id serial PRIMARY KEY,
	device_uuid varchar(255) NULL,
	measured_datetime timestamp NOT NULL,
    measured_value varchar(255)
);