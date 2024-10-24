# Проектирование базы данных

### Цель
Спроектировать базу данных существующего проекта (Aviasales — сервис для покупки дешевых авиабилетов) и реализовать REST API для работы с данными.

### Задачи
    1. Определить основные сущности
    2. Составить модель базы данных с использованием UML
    3. Развернуть локальный сервер PostgreSQL
    4. Спроектировать БД
    5. Реализовать API

## Реализация

### 1. Определение сущностей

Моделирование БД состоит из множества этапов, первый из которых — это определить сущности, исходя из бизнес-требований заказчика. Это буквально диалог с клиентом, во время которого изучаются технические и логические аспекты отрасли. Если требуется разработать продукт для финансовой компании, то сущности будут примерно такими:

1. Счет
2. Транзакция
3. Заем
4. и т.п.

Рассмотрим этот этап проектирования в контексте нашей задачи:

![alt text](/img/Блоки.png)

При поиске авиабилетов видим блоки с информацией об авиарейсах. Сымитируем упрощённые требования заказчика:

Результат поиска авиабилетов должен быть оформлен в виде блоков.
В каждом блоке должны быть указаны:

```bash
    Город вылета
    Город назначения
    Время вылета
    Время прилета в пункт назначения
    Стоимость авиабилета
    Расчетное время полета
    Авиакомпания, обслуживающая авиаперелет
```

Пока остановимся на этом и выделим сущности из требований:

#### 1. Аэропорт

    a_id      - уникальный идентификатор аэропорта
    a_name    - название аэропорта
    a_city_id - город базирования аэропорта

#### 2. Город

    c_id   - уникальный идентификатор города
    c_name - название города
    c_time_zone - часовой пояс (важно при расчете времени прибытия)

#### 3. Авиакомпания

    a_id   - уникальный идентификатор авиакомпании
    a_name - название авиакомпании

#### 4. Маршрут

    r_id             - уникальный идентификатор машрута
    r_airport_depart - город отправления
    r_airport_dest   - город назначения
    r_way            - время в пути

#### 5. Рейс

    f_id                - уникальный идентификатор рейса
    f_number            - номер рейса (например SU 1111)
    f_airline_id        - id авиакомпании, которая обслуживает рейс
    f_route             - маршрут следования
    f_time              - дата и время начала рейса
    f_cost              - стоимость рейса на 1 эконом место

Список можно продолжать бесконечно. Конечно хочется учесть все параметры, но на данном этапе ключевая задача понять принцип проектирования БД.

### 2. Составление модели

#### 2.1 Теория

Литература по проектированию баз данных учит разделять моделирование на уровни (суперпростыми словами):

1. Составление бизнес-модели

Заказчик говорит, что он хочет, а разработчик составляет базовую модель на основе требований. Чем лучше разработчик поймёт технические аспекты отрасли заказчика, тем лучше.

2. Инфологический уровень

Концептуальный уровень, на котором модель не привязана к конкретной СУБД. Состоит из сущностей, некоторых связей и атрибутов.

3. Даталогический уровень

Логический уровень, на котором повышена детализация и может быть привязка к виду базы данных. Помимо предыдущих пунктов добавляются индексы, ключи и типы.

4. Физический уровень

Уровень модели базы данных с наивысшей детализацией и привязкой к конкретной СУБД. На данном уровне прорабатываются все мельчайшие технические аспекты.

Естественно, никто не проектирует базы данных по строгим инструкциям, и поэтому модель затрагивает сразу несколько уровней.

#### 2.2 Модель базы данных

На основе сущностей составим UML диаграмму:

![alt text](/img/БД.png)

Связующие ключи, индексы и уникальность полей в объектах UML-диаграммы принято отображать как методы, потому что в контексте СУБД определить, к примеру, ключ — это действие. Также на данном этапе моделирования пока пропустим определение индексов и ограничений. Соответственно, модель находится между инфологическим и даталогическим уровнями.

### 3. Локальный сервер PostgreSQL

Развернем локальный сервер PostgreSQL в Docker контейнере. Для этого воспользуемся надстройкой docker-compose и сформируем compose-file.

[docker-compose.yml](./postgresql/docker-compose.yml)

Docker Compose позволяет разворачивать одновременно несколько контейнеров и организовывать связь между ними. Чтобы информация, которую мы будем вносить в БД, не была потеряна после перезапуска контейнера, вмонтируем Docker Volume:

```bash
volumes:
  aviadb-data:  aviadb-data:/var/lib/postgresql/data
  pgadmin-data: pgadmin-data:/var/lib/pgadmin
```

Теперь запустим контейнер в фоновом режиме и проверим работу:

```bash
docker-compose up -d
```

В строке браузера введём http://localhost:5050/browser/. Если всё настроено правильно, то откроется веб-клиент pgAdmin, с помощью которого и будет спроектирована база данных.

![alt text](/img/pgAdmin.png)

#### Что произошло?

С помощью Docker Compose был развернут контейнер data_base_container, внутри которого расположены два других контенера:

    postgres_container - СУБД PostgreSQL
    pgadmin_container  - Web клиент pgAdmin

![alt text](/img/container.png)

Проверим работоспособность БД и создадим тестовую таблицу:

```bash
CREATE TABLE TEST(
	id		SERIAL,
	name 	varchar(100)
);
```

![alt text](/img/test_table.png)

Работает! Приступим непосредственно к проектированию БД.

### 4. Проектирование БД

С использованием Web клиента pgAdmin определим таблицы в соответствии с диаграммой:

#### Рейс

```bash
CREATE TABLE flight(
f_id 			SERIAL 		NOT NULL,
f_number 		VARCHAR(50) NOT NULL,
f_airline_id 	INT 		NOT NULL,
f_route			INT			NOT NULL,
f_time 			TIMESTAMP	NOT NULL,
f_cost 			MONEY		NOT NULL,
CONSTRAINT PK_flight PRIMARY KEY (f_id),
FOREIGN KEY (f_route) REFERENCES route(r_id),
FOREIGN KEY (f_airline_id) REFERENCES airline(al_id)
);
```

#### Авиакомпания

```bash
CREATE TABLE airline(
al_id 		SERIAL 		NOT NULL,
al_name 	VARCHAR(50) NOT NULL,
CONSTRAINT PK_airline PRIMARY KEY (al_id)
);
```

#### Маршрут

```bash
CREATE TABLE route(
r_id 				SERIAL 		NOT NULL,
r_airport_depart 	INT			NOT NULL,
r_airport_dest		INT 		NOT NULL,
r_way				SMALLINT	NULL,
CONSTRAINT PK_route PRIMARY KEY (r_id),
FOREIGN KEY (r_airport_depart) REFERENCES airport(ap_id) ON DELETE SETNULL,
FOREIGN KEY (r_airport_dest) REFERENCES airport(ap_id) ON DELETE SET NULL
);
```

#### Аэропорт

```bash
CREATE TABLE airport(
ap_id 		SERIAL 		NOT NULL,
ap_name 	VARCHAR(50) NOT NULL,
ap_city_id 	INT 		NOT NULL,
CONSTRAINT PK_airport PRIMARY KEY (ap_id),
FOREIGN KEY (ap_city_id) REFERENCES city(c_id) ON DELETE SET NULL
);
```

#### Город

```bash
CREATE TABLE city(
	c_id SERIAL NOT NULL,
	c_name VARCHAR(100) NOT NULL,
	c_time_zone INT NOT NULL,

	CONSTRAINT PK_city PRIMARY KEY (c_id)
);
```

В итоге получим схему, аналогичную UML-диаграмме. 

![alt text](/img/БД_pgAdmin.png)

БД будет заполнена синтетическими данными для наглядности.

### Определение API

На базе стандартной библиотеки языка Go реализуем API для поиска доступных авиабилетов.

#### Вывод всех авиабилетов в БД

Для начала реализуем ручку, которая будет выводить вообще все аварейсы, хранящиеся в БД

```bash
URL  - http://localhost:8090/flights
Body - {
            "city_depart"   : "(Оставить пустым)",
            "city_dest"     : "(Оставить пустым)",
            "time"          : "(Оставить пустым)"
}
```