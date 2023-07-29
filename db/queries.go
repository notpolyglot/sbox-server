package db

import (
	"context"
	"errors"
	"sbox-backend/structs"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func GetUser(id64 string) (structs.User, error) {
	var rows []structs.User
	err := pgxscan.Select(context.Background(), Conn, &rows, "SELECT * from steam_users WHERE id64 = $1", id64)

	if err != nil {
		return structs.User{}, err
	}

	if len(rows) == 0 {
		return structs.User{}, errors.New("no rows")
	}

	return rows[0], nil
}
