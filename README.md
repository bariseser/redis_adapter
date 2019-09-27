# Redis Adapter for Go

This module helps to you some redis command which such as set, get and del

### Connect to Redis
````go
redisClient = ConnectRedis("localhost", 6379, database, "passsword")
````

### Set New Key
````go
response, err := redisClient.Set("bariseser", "Bariseser", time.Second*60*60)

if err != nil {
    log.Fatal(err.Error())
}
````

### Get Key
````go
response, err := redisClient.Get("bariseser")

if err != nil {
    log.Fatal(err.Error())
}
````

### Delete Key
````go
err := redisClient.Del("bariseser")

if err != nil {
    log.Fatal(err.Error())
}
````

### Ping Redis
````go
response, err := redisClient.Ping()
if err != nil {
    log.Fatal(err.Error())
}
````