NewsAPI implemented in Clean Architecture using Golang
====


## Description

外部サービスから取得したデータを用いて
記事を作成し, 提供するAPIを [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) で実装した例. 

## Requirement

### 外部APIのトークン

[News API](https://NewsAPI.org) で Developer Plan で APIトークンを取得し, `.env` にトークンを記述して使用する必要がある.  

```
NEWS_API_TOKEN={your api token}
```

## License
This Application is open-sourced software licensed under the [MIT](https://opensource.org/licenses/MIT)

## Author

[shirobrak](https://github.com/shirobrak)

## Used external service

- NewsAPI Developer Plan（ Powered by NewsAPI.org ）

---

## 実装方針

今回の実装では, クリーンアーキテクチャの4層は以下のようになった.  
- Entity層：`entities`
- Service層：`services`
- Adapter層 :`adapters`
- Repository層：`repositories`

また, `TopicsAPI` の実装に要した作成ファイルは以下の通り.  
```
.
|
├── adapters
│   └── ArticleGetter.go
|
├── entities
│   └── responses
│       ├── TopicsAPIResponse.go
│       └── contents
│           └── Article.go
├── repositories
│   └── NewsApiRepository.go
├── router.go
└── services
    └── TopicsArticleService.go
```