package pssql

import (
	"context"
	"fmt"
	. "go-xml-parser/internal/models"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const AddressObjectQuery = "insert into addresses(id, object_guid) values($1, $2)"
const ParamsQuery = "insert into params(id, object_id,updatedate,startdate,enddate) values($1, $2, $3,$4,$5)"

type Db struct {
	Conn  *pgxpool.Pool
	Batch *pgx.Batch
}

func (db *Db) Init(cfg Config) error {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Server.Host, cfg.Server.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Dbname)
	conn, err := pgxpool.New(context.Background(), psqlconn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err
	}
	db.Conn = conn
	db.Batch = &pgx.Batch{}
	return nil
}

func (db *Db) QueryBatch(query string, arguments ...any) {
	db.Batch.Queue(query, arguments...)
}

func (db *Db) ExecBatch() {
	inbr := db.Conn.SendBatch(context.Background(), db.Batch)
	inbr.Close()
	db.Batch = &pgx.Batch{}
}
func (db *Db) ExecQuery(query string) {
	_, err := db.Conn.Exec(context.Background(), query)
	if err != nil {
		fmt.Println(err)
	}
}
func (db *Db) Close() {
	db.Conn.Close()
}
