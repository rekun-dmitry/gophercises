drop table if exists english_idioms;
create table quiz.english_idioms
(question varchar(2000),
answer varchar(2000),
times integer,
last_date timestamp without time zone,
next_date timestamp without time zone);

