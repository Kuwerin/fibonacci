# Сервис, возвращающий срез ряда Фибоначчи.

Не забудьте добавить файл .env с примера env.local или env.docker

### Для запуска приложения с  docker:

```
docker-compose up -d --build
```


### Для того, чтобы получить срез чисел Фибоначчи по http, сделайте POST-запрос: 
```
POST localhost:8000/fibonacci

В body(json):

{
    "x": число1,
    "y": число2
}
```

###  Для того, чтобы получить срез числе Фибоначчи по GRPC, вам потребуется с клиентской части (например, evans) подключиться к  порту 9000 и вызвать процедуру GetFibSlice:

```
evans ./proto/fib.proto -p 9000

call GetFibSlice
x (TYPE_UINT64) => 2
y (TYPE_UINT64) => 4
```