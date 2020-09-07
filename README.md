# highloadcup2018

## 開発方法

1. Build Docker image

```
$ docker build -t app .
```

2. Run Docker container
```
$ docker run --rm --name app -p 8080:80 -v $PWD:/go/src/github.com/habara-k/highloadcup2018 app
```

3. Access http://localhost:8080 from browser

4. Edit main.go (updated automatically by [realize](https://github.com/oxequa/realize))
