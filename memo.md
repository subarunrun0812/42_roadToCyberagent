# Chapter4
### IV .2 MySQL
mysqlへのアクセス方法
```
❯ docker exec -it mysql /bin/bash
bash-5.1# mysql -u root -p
```
パスワードはdocker-compose.ymlの`MYSQL_PASSWORD`で定義している 

### IV.4 Redis
まだvolumeを設定していないので、ランキング機能を作るときなどに追加する必要があるかも

### IV.5 Swagger UI
swagger-uiのweb上で、`/api/games`が取得できない問題。
→swagger-uioコンテナ(localhost:3000)がgoコンテナ(localhost:8000)にアクセスすることを許可したら解決した。keyword: CORS（Cross-Origin Resource Sharing）
swagger-uiで[yml syntax](https://docs.ansible.com/ansible/latest/reference_appendices/YAMLSyntax.html)の勉強をした

ER図の `||--o{` という表記は 「1対多の関係 (1:N)」 を示しています。
それぞれの記号の意味を解説すると：

`||` → 1（one） を表す（必須の関係）
`o{` → 多（many） を表す（0個以上の関係）


# Exercise 1
・**DDL (Data Definition Language)**  
データベースの構造を定義するファイル。これは、データの型やテーブルの枠組みを作るイメージ。
実際のデータは挿入しない   
(コマンドの例)  
`CREATE`,`DROP`,`ALTER`

・**DML (Data Manipulation Language)**  
DDLで定義したテーブルに、実際にデータを操作(削除や追加)を行うファイルです。
実際のデータを操作する
(コマンドの例)   
`INSERT`,`UPDATE`,`DELETE`

### 4
> コレクションアイテムやガチャの確率定義など、ゲームをプレイするユーザーによって内容が書き換わらないレコードをDMLファイルとして作成

あらかじめ挿入するデータ
・アイテム
アイテムの確率(ガチャの出現率)
アイテム名
レアリティ
・ガチャ
単発ガチャの1回あたりの消費コイン
10ガチャの1回あたりの消費コイン
ガチャ名


dmlで追加したデータの確認方法
```
mysql> use 42tokyo-road-to-dojo-go_db-data;
mysql> show tables;
mysql> select * from <テーブル名>
```

# Exercise 2