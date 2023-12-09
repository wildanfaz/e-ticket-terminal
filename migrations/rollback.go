package migrations

import (
	"context"
	"fmt"

	"github.com/wildanfaz/e-ticket-terminal/configs"
)

func RollbackTables(ctx context.Context) {
	db := configs.InitMySql()

	_, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS transactions")
	if err != nil {
		panic(err)
	}

	_, err = db.ExecContext(ctx, "DROP TABLE IF EXISTS routes")
	if err != nil {
		panic(err)
	}

	_, err = db.ExecContext(ctx, "DROP TABLE IF EXISTS terminals")
	if err != nil {
		panic(err)
	}

	_, err = db.ExecContext(ctx, "DROP TABLE IF EXISTS locations")
	if err != nil {
		panic(err)
	}

	_, err = db.ExecContext(ctx, "DROP TABLE IF EXISTS users")
	if err != nil {
		panic(err)
	}

	fmt.Println("Rollback Success!")
}
