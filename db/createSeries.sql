USE mh
DROP TABLE IF EXISTS series;
CREATE TABLE IF NOT EXISTS series (
    id int not null,
    name varchar(10) not null
)
