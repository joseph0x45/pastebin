package db

import "github.com/joseph0x45/sad"

var migrations = []sad.Migration{
	{
		Version: 1,
		Name:    "pastes",
		SQL: `
      create table pastes (
        id text not null primary key,
        title text not null,
        preview text not null,
        content text not null
      );
    `,
	},
}
