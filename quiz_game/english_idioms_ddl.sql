drop table if exists quiz.english_idioms;
create table quiz.english_idioms
(question varchar(2000),
answer varchar(2000),
times integer,
last_date timestamp without time zone,
next_date timestamp without time zone);

--backup code
/*
pg_dump --host "localhost" --column-inserts --data-only --port "5432" --username "postgres" --format=p --encoding "WIN1251" --table "public.english_idioms" "quizes" > ./quiz_game/english_idioms_dump.sql
*/