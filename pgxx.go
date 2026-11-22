package main

//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"log"
//	"time"
//
//	"github.com/jackc/pgerrcode"
//	"github.com/jackc/pgx/v5"
//	"github.com/jackc/pgx/v5/pgconn"
//	"github.com/jackc/pgx/v5/pgxpool"
//	"github.com/labstack/echo/v5"
//)
//
//var (
//	NoRows           = errors.New("no rows")
//	NotEnoughMoney   = errors.New("no money")
//	ErrInvalidref    = errors.New("adress no")
//	ErrALreadyExists = errors.New("already exists")
//)
//
//type Task struct {
//	ID    int
//	Title string
//	Done  bool
//}
//
//type TaskRepo struct {
//	pool *pgxpool.Pool
//}
//
//func NewTaskRepo(pool *pgxpool.Pool) *TaskRepo {
//	return &TaskRepo{pool: pool}
//}
//func (t *TaskRepo) Create(ctx context.Context, title string) (int, error) {
//	var id int
//	err := t.pool.QueryRow(ctx, "INSERT INTO tasks(title) VALUES($1) RETURNING id", title).Scan(&id)
//	if err != nil {
//		return 0, err
//	}
//	return id, err
//}
//func (t *TaskRepo) MarkDone(ctx context.Context, id int) error {
//	tag, err := t.pool.Exec(ctx, "UPDATE tasks SET done=true WHERE id=$1", id)
//	if err != nil {
//		return err
//	}
//	if tag.RowsAffected() == 0 {
//		return errors.New("0 rows")
//	}
//	return nil
//}
//func (t *TaskRepo) GetAll(ctx context.Context) ([]Task, error) {
//	var tasks []Task
//	rows, err := t.pool.Query(ctx, "SELECT id, done, title FROM tasks ")
//	if err != nil {
//		return []Task{}, errors.New("query err")
//	}
//	tasks, err = pgx.CollectRows(rows, pgx.RowToStructByName[Task])
//	if err != nil {
//		return []Task{}, errors.New("geterr err")
//	}
//	return tasks, err
//}
//func (t *TaskRepo) Transaction(ctx context.Context, sender, receiver int, value int64) error {
//	if value <= 0 {
//		return errors.New("amount must be positive")
//	}
//	if sender == receiver {
//		return errors.New("cannot transfer to the same account")
//	}
//
//	return pgx.BeginFunc(ctx, t.pool, func(tx pgx.Tx) error {
//		tag, err := tx.Exec(ctx, "UPDATE accounts SET balance = balance - $2 WHERE id = $1", sender, value)
//		if err != nil {
//			return fmt.Errorf("debit sender: %w", err)
//		}
//		if tag.RowsAffected() == 0 {
//			return fmt.Errorf("sender account %d not found", sender)
//		}
//
//		tag, err = tx.Exec(ctx, "UPDATE accounts SET balance = balance + $2 WHERE id = $1", receiver, value)
//		if err != nil {
//			return fmt.Errorf("credit receiver: %w", err)
//		}
//		if tag.RowsAffected() == 0 {
//			return fmt.Errorf("receiver account %d not found", receiver)
//		}
//
//		_, err = tx.Exec(ctx, "INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES ($1, $2, $3)", sender, receiver, value)
//		if err != nil {
//			return fmt.Errorf("insert transfer: %w", err)
//		}
//
//		return nil
//	})
//}
//func translateerr(err error) error {
//	if err == nil {
//		return nil
//	}
//	if err == pgx.ErrNoRows {
//		return NoRows
//	}
//	var pgerr *pgconn.PgError
//	if errors.As(err, &pgerr) {
//		switch pgerr.Code {
//		case pgerrcode.UniqueViolation:
//			return fmt.Errorf("%w: %s", ErrALreadyExists, pgerr.ConstraintName)
//		case pgerrcode.ForeignKeyViolation:
//			return fmt.Errorf("%w: %s", ErrInvalidref, pgerr.ConstraintName)
//		case pgerrcode.CheckViolation:
//			return fmt.Errorf("нарушено условие (%s): %w", pgerr.ConstraintName, err)
//		}
//	}
//	return nil
//}
//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	config, err := pgxpool.ParseConfig("postgres://postgres:pass@localhost:9998/postgres?sslmode=disable")
//	if err != nil {
//		log.Fatal("conn is no")
//	}
//
//	config.MaxConns = 10
//	config.MinConns = 2
//	config.MaxConnLifetime = time.Hour
//
//	pool, err := pgxpool.NewWithConfig(ctx, config)
//	if err != nil {
//		log.Fatal("conn no")
//	}
//	repo := NewTaskRepo(pool)
//	for range 5 {
//		fmt.Println(repo.Transaction(ctx, 1, 2, 200))
//
//	}
//}
//func Httptranslator(c *echo.Context, err error) error {
//
//}
