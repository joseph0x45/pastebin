package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joseph0x45/goutils"
	"github.com/joseph0x45/pastebin/internal/buildinfo"
	"github.com/joseph0x45/pastebin/internal/models"
	"github.com/joseph0x45/sad"
)

type Conn struct {
	db      *sqlx.DB
	verbose bool
}

func (c *Conn) Close() {
	if c.verbose {
		log.Println("Database connection closed")
	}
	c.db.Close()
}

func GetConn(verbose bool) *Conn {
	dbPath := goutils.Setup()
	if buildinfo.Version == "debug" {
		dbPath = "db.sqlite"
	}
	db, err := sad.OpenDBConnection(sad.DBConnectionOptions{
		Reset:             false,
		EnableForeignKeys: true,
		DatabasePath:      dbPath,
	}, migrations)
	if err != nil {
		panic(err)
	}
	if verbose {
		log.Println("Connected to database file at", dbPath)
	}
	return &Conn{db: db, verbose: verbose}
}

func (c *Conn) InsertPaste(paste *models.Paste) error {
	const query = `
    insert into pastes (
      id, title, content, preview
    )
    values (
      :id, :title, :content, :preview
    )
  `
	if _, err := c.db.NamedExec(query, paste); err != nil {
		return fmt.Errorf("Error while inserting paste %w", err)
	}
	return nil
}

func (c *Conn) GetAllPastes() ([]models.Paste, error) {
	const query = "select * from pastes"
	pastes := []models.Paste{}
	if err := c.db.Select(&pastes, query); err != nil {
		return nil, fmt.Errorf("Error while getting pastes: %w", err)
	}
	return pastes, nil
}

func (c *Conn) DeletePaste(id string) error {
	const query = "delete from pastes where id=?"
	if _, err := c.db.Exec(query, id); err != nil {
		return fmt.Errorf("Error while deleting paste: %w", err)
	}
	return nil
}
