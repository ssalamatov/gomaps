package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ssalamatov/gomaps/internal/city"
	"github.com/ssalamatov/gomaps/internal/config"
	"github.com/ssalamatov/gomaps/internal/country"
)

func NewPgClient(ctx context.Context, config *config.Config) (pool *pgxpool.Pool, err error) {
	dsn := config.GetDbDsn()
	pool, err = pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("%v %w", err, ErrConnect)
	}
	return
}

func GetCountries(ctx context.Context, pool *pgxpool.Pool) (*[]country.GetCountryDTO, error) {
	rows, err := pool.Query(ctx, `SELECT id, name FROM country`)
	if err != nil {
		return nil, fmt.Errorf("%v %w", err, ErrSqlQueryFailed)
	}

	countries := make([]country.GetCountryDTO, 0)
	for rows.Next() {
		var country country.GetCountryDTO
		if err := rows.Scan(&country.Id, &country.Name); err != nil {
			return nil, fmt.Errorf("%v %w", err, ErrSqlScanFailed)
		}
		countries = append(countries, country)
	}
	return &countries, nil
}

func GetCountryById(ctx context.Context, pool *pgxpool.Pool, id int) (*country.GetCountryDTO, error) {
	rows, err := pool.Query(ctx, `SELECT id, name FROM country WHERE id=$1`, id)
	if err != nil {
		return nil, fmt.Errorf("%v %w", err, ErrSqlQueryFailed)
	}

	if rows.Next() {
		country := new(country.GetCountryDTO)
		if err := rows.Scan(&country.Id, &country.Name); err != nil {
			return nil, fmt.Errorf("%v %w", err, ErrSqlScanFailed)
		}
		return country, nil
	}
	return nil, fmt.Errorf("%w", ErrNotFound)
}

func GetCountryByName(ctx context.Context, pool *pgxpool.Pool, name string) (*country.GetCountryDTO, error) {
	rows, err := pool.Query(ctx, `SELECT id, name FROM country WHERE name=$1`, name)
	if err != nil {
		return nil, fmt.Errorf("%v %w", err, ErrSqlQueryFailed)
	}

	if rows.Next() {
		country := new(country.GetCountryDTO)
		if err := rows.Scan(&country.Id, &country.Name); err != nil {
			return nil, fmt.Errorf("%v %w", err, ErrSqlScanFailed)
		}
		return country, nil
	}
	return nil, fmt.Errorf("%w", ErrNotFound)
}

func GetCities(ctx context.Context, pool *pgxpool.Pool) (*[]city.GetCityDTO, error) {
	rows, err := pool.Query(ctx, `SELECT id, name, city.is_capital, city.found_at, city.population FROM city`)
	if err != nil {
		return nil, fmt.Errorf("%v %w", err, ErrSqlQueryFailed)
	}

	cities := make([]city.GetCityDTO, 0)
	for rows.Next() {
		var city city.GetCityDTO
		if err := rows.Scan(&city.Id, &city.Name, &city.IsCapital, &city.FoundAt, &city.Population); err != nil {
			return nil, fmt.Errorf("%v %w", err, ErrSqlScanFailed)
		}
		cities = append(cities, city)
	}
	return &cities, nil
}

func GetCityInfo(ctx context.Context, pool *pgxpool.Pool, name string) (*[]city.GetCityInfoDTO, error) {
	rows, err := pool.Query(
		ctx, `SELECT city.id, city.name, city.is_capital, city.found_at, city.population, country.name AS country FROM
			(SELECT * FROM city WHERE name=$1) AS city
				INNER JOIN country ON city.country_id=country.id`, name)
	if err != nil {
		return nil, fmt.Errorf("%v %w", err, ErrSqlQueryFailed)
	}

	cities := make([]city.GetCityInfoDTO, 0)
	for rows.Next() {
		var city city.GetCityInfoDTO
		if err := rows.Scan(&city.Id, &city.Name, &city.IsCapital, &city.FoundAt, &city.Population, &city.Country); err != nil {
			return nil, fmt.Errorf("%v %w", err, ErrSqlScanFailed)
		}
		cities = append(cities, city)
	}
	return &cities, nil
}

func RemoveCity(ctx context.Context, pool *pgxpool.Pool, id int) error {
	ctag, err := pool.Exec(ctx, `DELETE FROM city WHERE id=$1`, id)

	if err != nil {
		return fmt.Errorf("%v %w", err, ErrSqlQueryFailed)
	}
	if ctag.RowsAffected() != 1 {
		return fmt.Errorf("no affected rows %w", ErrNotFound)
	}
	return nil
}

func RemoveCountry(ctx context.Context, pool *pgxpool.Pool, id int) error {
	ctag, err := pool.Exec(ctx, `DELETE FROM country WHERE id=$1`, id)
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrSqlQueryFailed)
	}
	if ctag.RowsAffected() != 1 {
		return fmt.Errorf("no affected rows %w", ErrNotFound)
	}
	return nil
}

func CreateCity(ctx context.Context, pool *pgxpool.Pool, city *city.CreateCityDTO) error {
	country, err := GetCountryByName(ctx, pool, city.Country)
	if err != nil {
		return err
	}

	ctag, err := pool.Exec(
		ctx,
		`INSERT INTO city (name, country_id, is_capital, found_at, population) VALUES ($1, $2, $3, $4, $5)`,
		city.Name, country.Id, city.IsCapital, city.FoundAt, city.Population)

	if err != nil {
		return fmt.Errorf("%v %w", err, ErrSqlQueryFailed)
	}
	if ctag.RowsAffected() != 1 {
		return fmt.Errorf("no affected rows %w", ErrNotFound)
	}
	return nil
}
