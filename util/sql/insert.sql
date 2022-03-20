insert into country (name) values
    ('USA'),
    ('Russia'),
    ('Italy'),
    ('Spain'),
    ('Finland'),
    ('Germany'),
    ('Canada'),
    ('Brazil'),
    ('Belarus'),
    ('Armenia');


insert into city (country_id, is_capital, found_at, name, population) values
    (2, true, now(), 'Moscow', 5),
    (1, true, now(), 'Washington', 2),
    (2, false, now(), 'Ufa', 3),
    (2, false, now(), 'Samara', 5),
    (2, false, now(), 'Perm', 1);
