CREATE TABLE IF NOT EXISTS USERS (
    uuid VARCHAR(36) PRIMARY KEY,
    user_name VARCHAR(20) NOT NULL,
    coins INT,
    high_score INT
);

CREATE TABLE IF NOT EXISTS ITEM_MASTER (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    rarity INT
);

CREATE TABLE IF NOT EXISTS GAME_RECORDS (
    user_id VARCHAR(36),
    score INT,
    earned_coins INT,
    FOREIGN KEY (user_id) REFERENCES USERS(uuid)
);

CREATE TABLE IF NOT EXISTS USER_ITEMS (
    user_id VARCHAR(36),
    item_id INT,
    quantity INT,
    PRIMARY KEY (user_id, item_id),
    FOREIGN KEY (user_id) REFERENCES USERS(uuid),
    FOREIGN KEY (item_id) REFERENCES ITEM_MASTER(id)
);

CREATE TABLE IF NOT EXISTS GACHA_MASTER (
    -- id: 異なるガチャを識別するためのID (初心者ガチャ、期間限定ガチャ、など)
    id INT PRIMARY KEY,
    name VARCHAR(255),
    item_id INT,
    probability FLOAT,
    single_cost INT,
    multi_cost INT,
    FOREIGN KEY (item_id) REFERENCES ITEM_MASTER(id)
);