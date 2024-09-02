create table measurments(
	id int PRIMARY KEY,
	device_uuid varchar(255) NULL,
	measured_datetime timestamp NOT NULL,
    measured_value varchar(255)
);