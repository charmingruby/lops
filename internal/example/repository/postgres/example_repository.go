package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/charmingruby/pack/internal/example/model"
	"github.com/charmingruby/pack/pkg/database/postgres"
	"github.com/jmoiron/sqlx"
)

type ExampleRepo struct {
	db    postgres.Client
	stmts map[string]*sqlx.Stmt
}

func NewExampleRepo(db postgres.Client) (*ExampleRepo, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range exampleQueries() {
		stmt, err := db.Conn.Preparex(statement)
		if err != nil {
			return nil,
				postgres.NewPreparationErr(queryName, "example", err)
		}

		stmts[queryName] = stmt
	}

	return &ExampleRepo{
		db:    db,
		stmts: stmts,
	}, nil
}

func (r *ExampleRepo) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			postgres.NewStatementNotPreparedErr(queryName, "dummy example")
	}

	return stmt, nil
}

func (r *ExampleRepo) FindByID(ctx context.Context, id string) (model.Example, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(findExampleByID)
	if err != nil {
		return model.Example{}, err
	}

	var example model.Example

	if err := stmt.QueryRowContext(ctx, id).Scan(
		&example.ID,
		&example.Name,
	); err != nil {
		if err == sql.ErrNoRows {
			return model.Example{}, nil
		}

		return model.Example{}, err
	}

	return example, nil
}

func (r *ExampleRepo) Create(ctx context.Context, example model.Example) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(createExample)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		example.ID,
		example.Name,
	)

	return err
}

func (r *ExampleRepo) Save(ctx context.Context, example model.Example) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	stmt, err := r.statement(saveExample)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		example.Name,
		example.ID,
	)

	return err
}
