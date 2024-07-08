package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	entity "go-mobile/internal/entitiy"
	"go-mobile/internal/handler/http/user/dto"
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

func (u userRepository) Create(ctx context.Context, passport, passportSeries, passportNumber string) (*string, error) {
	const fn = "UserRepository.Create"

	sql, args, err := u.db.Builder.Insert("users").
		Columns("passport, passport_series, passport_number").
		Values(passport, passportSeries, passportNumber).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		u.log.Error("UserRepository - Create - BUILD SQL", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	u.log.Info("SQL", "query", sql)

	var id string
	err = u.db.Conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		u.log.Error("UserRepository - Create - EXEC SQL", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &id, nil
}

func (u userRepository) FindUserByCustomField(ctx context.Context, field, value string) (*entity.UserToResponse, error) {
	const fn = "UserRepository.FindUserByCustomField"
	var user entity.UserToResponse

	sql, args, err := u.db.Builder.Select("id, name, surname, patronymic, address, passport").
		From("users").
		Where(squirrel.Eq{field: value}).
		ToSql()

	if err != nil {
		u.log.Error("UserRepository - FindUserByCustomField - BUILD SQL", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	u.log.Info("SQL", "query", sql)

	err = u.db.Conn.QueryRow(ctx, sql, args...).
		Scan(&user.Id, &user.Name, &user.Surname, &user.Patronymic, &user.Address, &user.Passport)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		u.log.Error("UserRepository - FindUserByCustomField - - EXEC SQL", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &user, nil
}

func (u userRepository) GetAll(ctx context.Context, params *dto.GetAllParams) (*entity.UserListResponse, error) {
	const fn = "UserRepository.GetAll"

	query := u.db.Builder.Select("id, name, surname, patronymic, address, passport").From("users")

	conditions := make([]squirrel.Sqlizer, 0)
	if params.Name != "" {
		conditions = append(conditions, squirrel.Eq{"name": params.Name})
	}
	if params.Surname != "" {
		conditions = append(conditions, squirrel.Eq{"surname": params.Surname})
	}
	if params.Patronymic != "" {
		conditions = append(conditions, squirrel.Eq{"patronymic": params.Patronymic})
	}
	if params.Address != "" {
		conditions = append(conditions, squirrel.Eq{"address": params.Address})
	}
	if params.Passport != "" {
		conditions = append(conditions, squirrel.Eq{"passport": params.Passport})
	}
	if params.PassportSeries != "" {
		conditions = append(conditions, squirrel.Eq{"passport_series": params.PassportSeries})
	}
	if params.PassportNumber != "" {
		conditions = append(conditions, squirrel.Eq{"passport_number": params.PassportNumber})
	}

	for _, condition := range conditions {
		query = query.Where(condition)
	}

	if params.OrderBy != "" && (params.OrderSort == "asc" || params.OrderSort == "desc") {
		query = query.OrderBy(fmt.Sprintf("%s %s", params.OrderBy, params.OrderSort))
	}

	if params.Limit >= 0 && params.Page >= 1 {
		offset := (params.Page - 1) * params.Limit
		query = query.Limit(uint64(params.Limit)).Offset(uint64(offset))
	}

	sql, args, err := query.ToSql()
	if err != nil {
		u.log.Error("UserRepository - GetAll - BUILD SQL", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	u.log.Info("SQL", "query", sql)

	rows, err := u.db.Conn.Query(ctx, sql, args...)
	if err != nil {
		u.log.Error("UserRepository - GetAll - EXEC SQL", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer rows.Close()

	//TODO: for totalCount make COUNT(*) query
	totalCount := 0
	var users []*entity.UserToResponse
	for rows.Next() {
		var user entity.UserToResponse
		err := rows.Scan(&user.Id, &user.Name, &user.Surname, &user.Patronymic, &user.Address, &user.Passport)
		if err != nil {
			u.log.Error("UserRepository - GetAll - ROWS SCAN", sl.Err(err))
			return nil, fmt.Errorf("%s: %w", fn, err)
		}
		users = append(users, &user)
		totalCount++
	}

	response := &entity.UserListResponse{
		Page:       &params.Page,
		Limit:      &params.Limit,
		TotalCount: &totalCount,
		Users:      users,
	}

	return response, nil
}

func (u userRepository) Update(c context.Context, dto *dto.UpdateUserDto, userId, series, number string) error {
	const fn = "UserRepository.Update"

	sql, args, err := u.db.Builder.Update("users").
		Set("name", dto.Name).
		Set("surname", dto.Surname).
		Set("patronymic", dto.Patronymic).
		Set("address", dto.Address).
		Set("passport", dto.Passport).
		Set("passport_series", series).
		Set("passport_number", number).
		Where(squirrel.Eq{"id": userId}).
		ToSql()

	if err != nil {
		u.log.Error("UserRepository - Update - BUILD SQL", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}
	u.log.Info("SQL", "query", sql)

	_, err = u.db.Conn.Exec(c, sql, args...)
	if err != nil {
		u.log.Error("UserRepository - Update - EXEC SQL", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (u userRepository) Delete(ctx context.Context, userId string) error {
	const fn = "UserRepository.Delete"

	sql, args, err := u.db.Builder.Delete("users").
		Where(squirrel.Eq{"id": userId}).
		ToSql()

	if err != nil {
		u.log.Error("UserRepository - Delete - BUILD SQL", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}
	u.log.Info("SQL", "query", sql)

	_, err = u.db.Conn.Exec(ctx, sql, args...)
	if err != nil {
		u.log.Error("UserRepository - Update - EXEC SQL", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
