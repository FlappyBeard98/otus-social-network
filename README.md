# otus-social-network

### postman-коллекция  
[otus-social-network.postman_collection.json](/tests/otus-social-network.postman_collection.json)

### инструкция по запуску

нужно выполнить скрипт, который запускает docker-compose.yml c базой данных и сервисом и запускает запросы из postman-коллекции

```shell
sh run.sh
```

### инструкция по запуску нагрузки

выполнив этот скрипт
```shell
sh run_load_test.sh 10
```

запустится сценарий с 10 потоками по 1000 итераций и поднимется [dashboard](http://127.0.0.1:5665/), если не указывать параметр то запустится сценарий с 1 потоком(допустимые параметры 10,100,1000,max,over). отчет с цифрами сохраняется в файл [summary.html](./summary.html)