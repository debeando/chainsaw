DROP DATABASE IF EXISTS demo;
CREATE DATABASE IF NOT EXISTS demo;

USE demo;
DROP TABLE IF EXISTS foo;
CREATE TABLE IF NOT EXISTS foo (
  id BIGINT(20) unsigned NOT NULL AUTO_INCREMENT,
  bar_id BIGINT(20) unsigned NOT NULL,
  value VARCHAR(32) NOT NULL,
  status SMALLINT(4) DEFAULT 0 NOT NULL,
  created_at TIMESTAMP NULL DEFAULT NULL,
  modified_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY bar_idx (bar_id),
  KEY status_idx (status),
  KEY created_at_modified_atx (created_at, modified_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DELIMITER $$
DROP PROCEDURE IF EXISTS sp_fill_foo$$
CREATE PROCEDURE sp_fill_foo()
BEGIN
  DECLARE i INT DEFAULT 0;
  WHILE i < 100000  DO
    INSERT INTO foo (
      bar_id,
      value,
      status,
      created_at
    )
    VALUES (
      LPAD(FLOOR(RAND() * 999999.99), 6, '0'),
      LPAD(CONV(FLOOR(RAND()*POW(36,12)), 10, 36), 12, 0),
      LPAD(FLOOR(RAND() * 7), 1, '0'),
      NOW() - INTERVAL FLOOR(RAND() * 30) DAY
    );
    SET i = i + 1;
  END WHILE;
END$$
DELIMITER ;

CALL sp_fill_foo();
