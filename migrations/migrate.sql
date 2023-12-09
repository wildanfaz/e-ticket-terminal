CREATE TABLE IF NOT EXISTS users (
    id varchar(36) NOT NULL DEFAULT (uuid()),
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    balance bigint NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS locations (
    id bigint NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS terminals (
    id bigint NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    location_id bigint NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (location_id) REFERENCES locations(id)
);

CREATE TABLE IF NOT EXISTS routes (
    id bigint NOT NULL AUTO_INCREMENT,
    from_terminal_id bigint NOT NULL,
    to_terminal_id bigint NOT NULL,
    price bigint NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (from_terminal_id) REFERENCES terminals(id),
    FOREIGN KEY (to_terminal_id) REFERENCES terminals(id)
);

CREATE TABLE IF NOT EXISTS transactions (
    id bigint NOT NULL AUTO_INCREMENT,
    user_id varchar(36) NOT NULL,
    from_terminal_id bigint NOT NULL,
    to_terminal_id bigint,
    is_success bool NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (from_terminal_id) REFERENCES terminals(id),
    FOREIGN KEY (to_terminal_id) REFERENCES terminals(id)
);

INSERT INTO locations (name) VALUES 
('Jakarta'), 
('Tegal'),
('Pekalongan'),
('Semarang'), 
('Yogyakarta');

INSERT INTO terminals (name, location_id) VALUES
('Terminal Jakarta', 1), 
('Terminal Tegal', 2), 
('Terminal Pekalongan', 3),
('Terminal Semarang', 4), 
('Terminal Yogyakarta', 5);

INSERT INTO routes (from_terminal_id, to_terminal_id, price) VALUES
(1, 2, 100000), (1, 3, 200000), (1, 4, 300000), (1, 5, 400000),
(2, 1, 100000), (2, 3, 100000), (2, 4, 200000), (2, 5, 300000),
(3, 1, 200000), (3, 2, 100000), (3, 4, 100000), (3, 5, 200000),
(4, 1, 300000), (4, 2, 200000), (4, 3, 100000), (4, 5, 100000),
(5, 1, 400000), (5, 2, 300000), (5, 3, 200000), (5, 4, 100000);