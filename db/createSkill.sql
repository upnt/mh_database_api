USE mh
DROP TABLE IF EXISTS skill;
CREATE TABLE IF NOT EXISTS skill (
    id int unsigned not null auto_increment,
    name varchar(10) not null
)
