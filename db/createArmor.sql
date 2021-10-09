USE mh
DROP TABLE IF EXISTS armor;
CREATE TABLE IF NOT EXISTS armor (
    id int,
    name varchar(30),
    position int,
    series int,
    defense int,
    fire int,
    water int,
    thunder int,
    ice int,
    dragon int,
    slot1 int,
    slot2 int,
    slot3 int,
    skill1 int,
    skill2 int
)
