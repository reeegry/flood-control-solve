```bash   
git clone https://github.com/MaksimovDenis/floodControl.git
```
```bash    
docker pull redis
```
```bash      
docker run --name vk-flood-control -p 6379:6379 -d redis
```
```bash    
cd cmd
```
```bash    
go run .
``` 

# Инструменты:

Начал решение с выбора места для хранилища, остановился на Redis. Плюсом будет то, что Redis хранит свои данные в оперативной памяти, запись\чтение из базы будет быстрой, также у нас не возникнет проблем в случае запуска нескольких экземпляров приложения.

Для конфига выбрал json. Потому что был с ним знаком. И он достаточно прост в использовании.

# Реализация:

В json файле указаны примеры ID пользователей для демонстрации. После создается цикл фиксирующий ежесекундные вызовы функции Check от разных пользователей, каждый из пользователей при первом вызове функции Check записывается в БД со сроком равным N, Если за это время user напишет больше K сообщений, возвращаем false.

Реализация функции Check нахоидтся в файле floodControl.go





Когда завершите задачу, в этом README опишите свой ход мыслей: как вы пришли к решению, какие были варианты и почему выбрали именно этот.

# Что нужно сделать

Реализовать интерфейс с методом для проверки правил флуд-контроля. Если за последние N секунд вызовов метода Check будет больше K, значит, проверка на флуд-контроль не пройдена.

- Интерфейс FloodControl располагается в файле main.go.

- Флуд-контроль может быть запущен на нескольких экземплярах приложения одновременно, поэтому нужно предусмотреть общее хранилище данных. Допустимо использовать любое на ваше усмотрение.

# Необязательно, но было бы круто

Хорошо, если добавите поддержку конфигурации итоговой реализации. Параметры — на ваше усмотрение.
