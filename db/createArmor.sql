USE mh
DROP TABLE IF EXISTS armor;
CREATE TABLE IF NOT EXISTS armor (
    id int not null auto_increment,
    name varchar(30) not null,
    position int not null,
    series int not null,
    defense int not null,
    fire int not null,
    water int not null,
    thunder int not null,
    ice int not null,
    dragon int not null,
    slot1 int not null,
    slot2 int not null,
    slot3 int not null,
    skill1 int not null,
    skill2 int not null
)
