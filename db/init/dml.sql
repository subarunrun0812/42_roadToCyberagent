INSERT INT ITEM_MASTER (id, name, rarity) VALUES
    (1, '1st Item', 1),
    (2, '2nd Item', 2),
    (3, '3rd Item', 3),
    (4, '4th Item', 4),
    (5, '5th Item', 5);

INSERT INT GACHA_MASTER (id, name, item_id, probability, single_cost, multi_cost) VALUES
    (1, '1st Gacha', 1, 0.1, 100, 1000),
    (2, '2nd Gacha', 2, 0.2, 200, 2000),
    (3, '3rd Gacha', 3, 0.3, 300, 3000),
    (4, '4th Gacha', 4, 0.4, 400, 4000),
    (5, '5th Gacha', 5, 0.5, 500, 5000);