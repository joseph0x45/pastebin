package db

import (
	"database/sql"
	"errors"
	"fmt"
	"pastebin/models"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jmoiron/sqlx"
)

func GetDBConnection() *sqlx.DB {
	conn, err := sqlx.Connect("sqlite3", "pastebin.db")
	if err != nil {
		panic(err)
	}
	const initQuery = `
    create table if not exists pastes (
      id text not null primary key,
      title text not null,
      content text not null
    );
  `
	_, err = conn.Exec(initQuery)
	if err != nil {
		conn.Close()
		panic(err)
	}
	return conn
}

func InsertPaste(conn *sqlx.DB, paste models.Paste) error {
	const query = `
    insert into pastes (
      id, title, content
    )
    values (
      :id, :title, :content
    )
  `
	_, err := conn.NamedExec(query, paste)
	if err != nil {
		return fmt.Errorf("Error while inserting new paste: %w", err)
	}
	return nil
}

func GetAllPastes(conn *sqlx.DB) ([]models.Paste, error) {
	pastes := make([]models.Paste, 0)
	const query = "select * from pastes"
	err := conn.Select(&pastes, query)
	if err != nil {
		return nil, fmt.Errorf("Error while getting all pastes: %w", err)
	}
	return pastes, nil
}

func DeletePaste(conn *sqlx.DB, id string) error {
	const query = `
    delete from pastes where id=$1
  `
	_, err := conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error while deleting paste with id %s: %w", id, err)
	}
	return nil
}

func GetPasteByID(conn *sqlx.DB, id string) (*models.Paste, error) {
	paste := &models.Paste{}
	const query = `
    select * from pastes where id=$1
  `
	err := conn.Get(paste, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("Error while getting paste by id %s: %w", id, err)
	}
	return paste, nil
}
