package sql

import (
	"fmt"
	"log"
	"strconv"
)

func (s *SQLdb) SelectQuery(qr string) (string, error) {
	var row struct {
		id      int
		name    string
		surname string
		modtime interface{}
	}
	fmt.Println("qr :", qr)
	err := s.db.QueryRow(qr).Scan(&row.id, &row.name, &row.surname, &row.modtime)
	if err != nil {
		return "", err
	}
	return "ID: " + strconv.Itoa(row.id) + "\n" + "name: " + row.name + "\n" + "surname: " + row.surname, nil
}

func (s *SQLdb) Add(name, surname string) error {
	_, err := s.db.Exec("insert into info (name, surname) values (?, ?)", name, surname)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (s *SQLdb) Delete(id string) error {
	qr := "delete from info where id = " + id
	_, err := s.db.Exec(qr)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
