INSERT INTO ITEM_MASTER (id, name, rarity) VALUES
    (1, 'star one', 1),
    (2, 'star two', 2),
    (3, 'star three', 3);

INSERT INTO GACHA_MASTER (id, name, single_cost, multi_cost) VALUES
    (1, 'normal gacha', 100, 1000),
    (2, 'season gacha', 200, 2000);

/* normal gacha */
INSERT INTO GACHA_PROBABILITY (gacha_id, item_id, item_probability) VALUES
    (1, 1, 5.0),
    (1, 2, 30.0),
    (1, 3, 65.0);

/* season gacha */
INSERT INTO GACHA_PROBABILITY (gacha_id, item_id, item_probability) VALUES
    (2, 1, 10.0),
    (2, 2, 40.0),
    (2, 3, 50.0);