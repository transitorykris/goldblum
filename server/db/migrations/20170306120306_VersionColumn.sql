
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE `endpoint` ADD COLUMN `version` INTEGER DEFAULT 0 AFTER `id`;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE `endpoint` DROP `version`;