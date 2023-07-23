CREATE TABLE IF NOT EXISTS `pix_key` (
    `id` varchar(36) NOT NULL PRIMARY KEY,
    `pix_key_type` varchar(255) NOT NULL,
    `pix_key` varchar(255) NOT NULL,
    `account_type` varchar(255) NOT NULL,
    `agency_number` INT NOT NULL,
    `account_number` INT NOT NULL,
    `account_holder_name` varchar(255) NOT NULL,
    `account_holder_last_name` varchar(255) NOT NULL,
    `created_at` datetime NOT NULL,
    `modified_at` datetime
    )ENGINE=InnoDB DEFAULT CHARSET=UTF8;