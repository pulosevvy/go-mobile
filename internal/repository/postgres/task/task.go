package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	entity "go-mobile/internal/entitiy"
	taskDto "go-mobile/internal/handler/http/task/dto"
	"go-mobile/package/database/postgres"
	sl "go-mobile/package/logger/slog"
	"log/slog"
)

type taskRepository struct {
	db  *postgres.Postgres
	log *slog.Logger
}

func NewTaskRepository(db *postgres.Postgres, log *slog.Logger) *taskRepository {
	return &taskRepository{db, log}
}

func (tr taskRepository) GetByUserId(ctx context.Context, userId string, dto *taskDto.GetByUser) ([]entity.TaskToResponse, error) {
	const fn = "taskRepository.GetByUserId"

	query := tr.db.Builder.Select("id, name, hours, start_task, end_task, user_id").
		From("tasks").
		Where(squirrel.Eq{"user_id": userId})

	if dto.StartDate != "" {
		query = query.Where(squirrel.GtOrEq{"start_task": dto.StartDateUnix})
	}
	if dto.EndDate != "" {
		query = query.Where(squirrel.LtOrEq{"end_task": dto.EndDateUnix})
	}

	query = query.OrderBy("hours DESC")

	sql, args, err := query.ToSql()
	if err != nil {
		tr.log.Error("TaskRepository - GetByUserId", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	tr.log.Info("SQL", "query", sql)

	rows, err := tr.db.Conn.Query(ctx, sql, args...)
	if err != nil {
		tr.log.Error("TaskRepository - GetByUserId - BUILD SQL", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer rows.Close()

	var tasks []entity.TaskToResponse
	for rows.Next() {
		var task entity.TaskToResponse
		err := rows.Scan(&task.Id, &task.Name, &task.Hours, &task.StartTask, &task.EndTask, &task.UserID)
		if err != nil {
			tr.log.Error("TaskRepository - getByUser - EXEC SQL", sl.Err(err))
			return nil, fmt.Errorf("%s: %w", fn, err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (tr taskRepository) CreateTask(ctx context.Context, dto *taskDto.CreateTaskDto) (*string, error) {
	const fn = "taskRepository.CreateTask"

	query := tr.db.Builder.Insert("tasks").
		Columns("name")

	//For future. One tasks - more users
	values := []interface{}{dto.Name}
	if dto.UserId != "" {
		query = query.Columns("user_id")
		values = append(values, dto.UserId)
	}

	sql, args, err := query.
		Values(values...).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		tr.log.Error("TaskRepository - CreateTask - BUILD SQL", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	tr.log.Info("SQL", "query", sql)

	var id string
	err = tr.db.Conn.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		tr.log.Error("TaskRepository - CreateTask - EXEC SQL", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &id, nil
}

func (tr taskRepository) StartTime(ctx context.Context, taskId string, dto *taskDto.StartTaskDto) error {
	const fn = "taskRepository.StartTime"
	fmt.Println(dto.StartTime)
	sql, args, err := tr.db.Builder.Update("tasks").
		Set("start_task", dto.StartTime).
		Where(squirrel.Eq{"id": taskId}).
		Where(squirrel.Eq{"user_id": dto.UserID}).
		ToSql()

	tr.log.Info("SQL", "query", sql)

	if err != nil {
		tr.log.Error("TaskRepository - StartTime - Build SQL", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	_, err = tr.db.Conn.Exec(ctx, sql, args...)
	if err != nil {
		tr.log.Error("TaskRepository - StartTime - EXEC SQL", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (tr taskRepository) EndTime(ctx context.Context, taskId string, hours float64, dto *taskDto.EndTaskDto) error {
	const fn = "taskRepository.EndTime"

	sql, args, err := tr.db.Builder.Update("tasks").
		Set("end_task", dto.EndTime).
		Set("hours", hours).
		Where(squirrel.Eq{"id": taskId}).
		Where(squirrel.Eq{"user_id": dto.UserID}).
		ToSql()

	tr.log.Info("SQL", "query", sql)

	if err != nil {
		tr.log.Error("TaskRepository - EndTime - Build SQL", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	_, err = tr.db.Conn.Exec(ctx, sql, args...)
	if err != nil {
		tr.log.Error("TaskRepository - EndTime - EXEC SQL", sl.Err(err))
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (tr taskRepository) FindTaskByCustomField(ctx context.Context, field, value string) (*entity.TaskToResponse, error) {
	const fn = "TaskRepository.FindTaskByCustomField"
	var task entity.TaskToResponse

	sql, args, err := tr.db.Builder.Select("id, name, hours, start_task, end_task, user_id").
		From("tasks").
		Where(squirrel.Eq{field: value}).
		ToSql()

	if err != nil {
		tr.log.Error("TaskRepository - FindTaskByCustomField - BUILD SQL", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	tr.log.Info("SQL", "query", sql)

	err = tr.db.Conn.QueryRow(ctx, sql, args...).
		Scan(&task.Id, &task.Name, &task.Hours, &task.StartTask, &task.EndTask, &task.UserID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		tr.log.Error("TaskRepository - FindTaskByCustomField - - EXEC SQL", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &task, nil
}
