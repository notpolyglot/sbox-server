package db

import (
	"context"
	"sbox-backend/structs"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func CreateUser(id string) error {
	_, err := Conn.Exec(context.Background(), "INSERT INTO steam_users (id64) VALUES($1)", id)
	if err != nil {
		return err
	}

	_, err = Conn.Exec(context.Background(), "INSERT INTO money (user_id64, value) VALUES($1, $2)", id, 500)
	if err != nil {
		return err
	}

	return nil
}

func GetUser(id string) (structs.User, error) {
	rows, err := Conn.Query(context.Background(), "SELECT * from steam_users WHERE id64 = $1", id)
	if err != nil {
		return structs.User{}, err
	}

	var user structs.User
	if err := pgxscan.ScanOne(&user, rows); err != nil {
		return structs.User{}, err
	}

	return user, nil
}

func GetMoney(id string) (int, error) {
	row := Conn.QueryRow(context.Background(), "SELECT value from money WHERE user_id64 = $1", id)

	var money int
	err := row.Scan(&money)
	if err != nil {
		return 0, err
	}

	return money, nil
}

func AddMoney(recipient string, amount int) error {
	_, err := Conn.Exec(context.Background(), "UPDATE money SET value = value + $1 WHERE user_id64 = $2", amount, recipient)
	if err != nil {
		return err
	}

	return nil
}

func SubtractMoney(sender string, amount int) error {
	_, err := Conn.Exec(context.Background(), "UPDATE money SET value = value - $1 WHERE user_id64 = $2", amount, sender)
	if err != nil {
		return err
	}

	return nil
}

func LogPayment(sender string, recipient string, amount int) error {
	_, err := Conn.Exec(context.Background(), "INSERT INTO payments (sender, recipient, amount) VALUES($1, $2, $3)", sender, recipient, amount)
	if err != nil {
		return err
	}

	return nil
}
