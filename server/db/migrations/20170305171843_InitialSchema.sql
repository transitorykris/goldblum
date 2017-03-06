
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `endpoint` (
    `id` INTEGER PRIMARY KEY AUTO_INCREMENT,
    `method` ENUM('GET', 'POST', 'PUT', 'OPTIONS', 'HEAD', 'DELETE', 'PATCH'),
    `path` VARCHAR(255),
    `code` TEXT
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE endpoint;