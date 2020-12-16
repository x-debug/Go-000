CREATE TABLE `fake_user`
(
    `id`       int(10) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(20)      NOT NULL,
    `password` varchar(20)      NOT NULL,
    `nickname` varchar(20)      NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`)
) ENGINE = InnoDB;

INSERT INTO fake_user values (1, 'bob1', '123456', 'bob1');
INSERT INTO fake_user values (2, 'bob2', '123456', 'bob2');
INSERT INTO fake_user values (3, 'bob3', '123456', 'bob3');
INSERT INTO fake_user values (4, 'bob4', '123456', 'bob4');
INSERT INTO fake_user values (5, 'bob5', '123456', 'bob5');
INSERT INTO fake_user values (6, 'bob6', '123456', 'bob6');
INSERT INTO fake_user values (7, 'bob7', '123456', 'bob7');
INSERT INTO fake_user values (8, 'bob8', '123456', 'bob8');
INSERT INTO fake_user values (9, 'bob9', '123456', 'bob9');
INSERT INTO fake_user values (10, 'bob10', '123456', 'bob10');