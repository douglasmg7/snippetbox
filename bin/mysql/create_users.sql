CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'localhost';
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';

-- For all hosts (needed for use docker).
CREATE USER 'web'@'%' IDENTIFIED BY 'pass';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'%';