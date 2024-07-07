package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	entity "go-mobile/internal/entitiy"
	"go-mobile/package/database/postgres"
	sl "go-mobile/package/logger/slog"
	"log/slog"
)

type userRepository struct {
	db  *postgres.Postgres
	log *slog.Logger
}

func NewUserRepository(db *postgres.Postgres, log *slog.Logger) *userRepository {
	return &userRepository{db, log}
}

func (u userRepository) Create(ctx context.Context, passport, passportSeries, passportNumber string) error {
	const fn = "UserRepository.Create"

	sql, args, err := u.db.Builder.Insert("users").
		Columns("passport, passport_series, passport_number").
		Values(passport, passportSeries, passportNumber).
		ToSql()

	u.log.Info("SQL", "query", sql)
	if err != nil {
		u.log.Error("UserRepository - Create", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	_, err = u.db.Conn.Exec(ctx, sql, args...)

	if err != nil {
		u.log.Error("UserRepository - Create", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (u userRepository) GetUserByPassport(ctx context.Context, passport string) (*entity.UserToResponse, error) {
	const fn = "GetUserByPassport"
	var user entity.UserToResponse

	sql, args, err := u.db.Builder.Select("id, name, surname, patronymic, address, passport").
		From("users").
		Where(squirrel.Eq{"passport": passport}).
		ToSql()

	u.log.Info("SQL", "query", sql)
	if err != nil {
		u.log.Error("UserRepository - GetUserByPassport", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	err = u.db.Conn.QueryRow(ctx, sql, args...).
		Scan(&user.Id, &user.Name, &user.Surname, &user.Patronymic, &user.Address, &user.Passport)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		u.log.Error("UserRepository - GetUserByPassport", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &user, nil
}

func (u userRepository) GetAll(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetById(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Update(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Delete(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
