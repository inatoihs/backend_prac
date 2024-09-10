-- Users テーブルを作成
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Posts テーブルを作成
CREATE TABLE IF NOT EXISTS posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    title VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- SomeTable テーブルを作成
CREATE TABLE IF NOT EXISTS some_table (
    id INT AUTO_INCREMENT PRIMARY KEY,
    description TEXT
);

-- LargeTable テーブルを作成
CREATE TABLE IF NOT EXISTS large_table (
    id INT AUTO_INCREMENT PRIMARY KEY,
    data BLOB
);

-- ユーザーと投稿データの挿入
INSERT INTO
    users (name)
VALUES
    ('Alice'),
    ('Bob'),
    ('Charlie');

INSERT INTO
    posts (user_id, title)
VALUES
    (1, 'Alice\'s First Post'),
    (1, 'Alice\'s Second Post'),
    (2, 'Bob\'s First Post'),
    (3, 'Charlie\'s Only Post');

-- SomeTable に大量のデータを挿入
INSERT INTO
    some_table (description)
VALUES
    ('Short description 1'),
    ('Short description 2'),
    ('Short description 3');

-- さらに大量のデータを挿入（ここではサンプルとして 1000 行を追加）
INSERT INTO
    some_table (description)
SELECT
    REPEAT('A long description text ', 20)
FROM
    information_schema.tables
LIMIT
    1000;

-- LargeTable に大きなデータを挿入
INSERT INTO
    large_table (data)
VALUES
    (REPEAT('A large binary blob data ', 1000)),
    (REPEAT('Another large binary blob data ', 1000));