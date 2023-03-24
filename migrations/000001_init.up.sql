CREATE TABLE customer
(
    id VARCHAR(255) CONSTRAINT users_pk PRIMARY KEY,
    first_query date
);

CREATE TABLE weather
(
    id SERIAL CONSTRAINT weather_pk PRIMARY KEY,
    city VARCHAR(255),
    weather_description VARCHAR(255),
    temp INTEGER,
    feels_like INTEGER,
    temp_min INTEGER,
    temp_max INTEGER,
    wind_speed INTEGER,
    clouds INTEGER
);


CREATE TABLE statistic
(
    id SERIAL CONSTRAINT statistics_pk PRIMARY KEY,
    user_id VARCHAR(255) REFERENCES customer,
    query VARCHAR(255),
    result INTEGER REFERENCES weather
);