# IV.4 Redis
volumeを設定していないので、ランキング機能を作るときなどに追加する必要があるかも

# IV.5 Swagger UI
swagger-uiのweb上で、`/api/games`が取得できない問題。
→swagger-ui(localhost:3000)がgo(localhost:8000)にアクセスすることを許可したら解決した。keyword: CORS（Cross-Origin Resource Sharing）

[yml syntax](https://docs.ansible.com/ansible/latest/reference_appendices/YAMLSyntax.html)