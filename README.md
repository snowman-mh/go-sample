# go-sample

Sample development environment in Golang

## Introduction

https://qiita.com/snowman_mh/items/2e50d7fa8e437411f5da

## Start local server

```bash
cp docker/.env.default docker/.env
make local start
make migrate up
```

## Run test

```bash
make test start
make test run
```
