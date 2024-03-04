CREATE TABLE IF NOT EXISTS users (
  ID                         UUID PRIMARY KEY,
	name                       VARCHAR(100) NOT NULL UNIQUE,
	age                        INTEGER NOT NULL,
	gender                     VARCHAR(20) NOT NULL,
	latitude                   VARCHAR(255) NOT NULL,
	longitude                  VARCHAR(255) NOT NULL,
	infected                   BOOLEAN DEFAULT FALSE,
	contamination_notification INTEGER DEFAULT 0,
	created_at                 TIMESTAMP NOT NULL,
	updated_at                 TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS items (
  ID                         UUID PRIMARY KEY,
	description                VARCHAR(100) NOT NULL UNIQUE,
	score                      INTEGER NOT NULL,
	created_at                 TIMESTAMP NOT NULL,
	updated_at                 TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS inventories (
  ID                         UUID PRIMARY KEY,
	user_id                    UUID NOT NULL,
	created_at                 TIMESTAMP NOT NULL,
	updated_at                 TIMESTAMP NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS infecteds (
  user_id_reported           UUID NOT NULL,
	user_id_notified           UUID NOT NULL,
	created_at                 TIMESTAMP NOT NULL,
	updated_at                 TIMESTAMP NOT NULL,
	FOREIGN KEY (user_id_reported) REFERENCES users(id),
	FOREIGN KEY (user_id_notified) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS inventory_items (
  ID                         UUID PRIMARY KEY,
	inventory_id               UUID NOT NULL,
	item_id                    UUID NOT NULL,
	quantity                   INTEGER NOT NULL,
	created_at                 TIMESTAMP NOT NULL,
	updated_at                 TIMESTAMP NOT NULL,
	FOREIGN KEY (inventory_id) REFERENCES inventories(id),
	FOREIGN KEY (item_id) REFERENCES items(id)
);
