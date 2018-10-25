# status-code-app
An app that responses with arbitrary status code

## Run locally
```
$ go run main.go
```

and access http://localhost:3000/418

## Run on Heroku
[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

```
$ heroku create
$ git push heroku master
$ heroku open
```

and add your favorite status code in the address bar on your browser.
