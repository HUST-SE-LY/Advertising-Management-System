use ams;
CREATE TABLE IF NOT EXISTS Admin
(
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Customer(
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    manager_name VARCHAR(255),
    address VARCHAR(255),
    manager_tel VARCHAR(50),
    business_license_number VARCHAR(255)
);