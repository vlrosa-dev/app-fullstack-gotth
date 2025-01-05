CREATE TABLE IF NOT EXISTS users (
    id              VARCHAR(36) NOT NULL PRIMARY KEY,
    first_name      VARCHAR(255) NOT NULL,
    last_name       VARCHAR(255) NOT NULL,
    email           VARCHAR(255) NOT NULL UNIQUE,
    password        VARCHAR(255) NOT NULL,
    created_at      TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP(3) NOT NULL
);

CREATE TABLE IF NOT EXISTS verifications (
    id              SERIAL,
    email           VARCHAR(255) NOT NULL,
    code  		   	VARCHAR(10) NOT NULL,
    expires_at 	   	TIMESTAMP NOT NULL,
    type           	VARCHAR(10) NOT NULL,
	PRIMARY KEY (id),
    FOREIGN KEY (email) REFERENCES users(email)
);
