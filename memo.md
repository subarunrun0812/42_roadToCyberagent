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
[RESTful APIとは](https://qiita.com/NagaokaKenichi/items/0647c30ef596cedf4bf2)
swagger-uiのweb上で、`/api/games`が取得できない問題。
→swagger-uiコンテナ(localhost:3000)がgoコンテナ(localhost:8000)にアクセスすることを許可したら解決した。keyword: CORS（Cross-Origin Resource Sharing）
swagger-uiで[yml syntax](https://docs.ansible.com/ansible/latest/reference_appendices/YAMLSyntax.html)の勉強をした

ER図の `||--o{` という表記は 「1対多の関係 (1:N)」 を示しています。
それぞれの記号の意味を解説すると：

`||` → 1（one） を表す（必須の関係）
`o{` → 多（many） を表す（0個以上の関係）


**apiの検証のためにgoで定義する**
GoでAPIエンドポイントを実装した。
ちなみに、エントリーポイントとエンドポイントは同じ意味である。

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
### Go
標準パッケージを用いて、API serverを実装する。GinなどのWeb FrameWorkは使用禁止




/pkg/server/handler の実装
→ POST /user/create を処理するハンドラーを作成

/pkg/server/server.go へのルーティングの追加
→ POST /user/create のリクエストが適切なハンドラーにルーティングされるように設定

Token の生成処理の実装
→ 認証用のトークン（ランダムな文字列や UUID など）の生成を行う

MySQL テーブルへのアクセス処理（DAO）の実装
→ ユーザー情報を MySQL に保存・取得するためのデータアクセス処理（DAO: Data Access Object）を実装