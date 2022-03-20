package client

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	DB_USER = "docker"
	DB_PASS = "docker"
	DB_NAME = "docker"
	DB_HOST = "localhost"
	DB_PORT = "5432"
)

const TableCreationCityQuery = `CREATE TABLE IF NOT EXISTS city
(	
	id serial primary key,
    country_id int references country(id),
    is_capital boolean default false,
    found_at timestamp not null,
    name varchar(50) not null,
    population int
)`
const TableCreationCountryQuery = `CREATE TABLE IF NOT EXISTS country
(
    id serial primary key,
	name varchar(50) not null unique
)`

func ExecQuery(ctx context.Context, pool *pgxpool.Pool, query string) {
	if _, err := pool.Exec(ctx, query); err != nil {
		log.Fatal(err)
	}
}

func ClearCityTable(ctx context.Context, pool *pgxpool.Pool) {
	ExecQuery(ctx, pool, "DELETE FROM city")
	ExecQuery(ctx, pool, "ALTER SEQUENCE city_id_seq RESTART WITH 1")
}

func ClearCountryTable(ctx context.Context, pool *pgxpool.Pool) {
	ExecQuery(ctx, pool, "DELETE FROM country")
	ExecQuery(ctx, pool, "ALTER SEQUENCE country_id_seq RESTART WITH 1")
}

func Connect() (*pgxpool.Pool, error) {
	ctx := context.Background()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	return pgxpool.Connect(ctx, dsn)

}

func TablesExists(ctx context.Context, pool *pgxpool.Pool) {
	ExecQuery(ctx, pool, TableCreationCityQuery)
	ExecQuery(ctx, pool, TableCreationCountryQuery)
}

func InsertTestData(ctx context.Context, pool *pgxpool.Pool) {
	ExecQuery(ctx, pool, `INSERT INTO country (name) VALUES ('Russia')`)
	ExecQuery(ctx, pool, `INSERT INTO city (country_id, is_capital, found_at, name, population) VALUES (1, true, '2022-03-19T23:36:13.183732Z', 'Moscow', 5)`)
}

func PrepareDB(DoInsert bool) {
	pool, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	TablesExists(ctx, pool)
	ClearCityTable(ctx, pool)
	ClearCountryTable(ctx, pool)

	if DoInsert {
		InsertTestData(ctx, pool)
	}
	pool.Close()
}
