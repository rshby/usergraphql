package link

import (
	"context"
	"database/sql"
	"errors"
	"log"
	database "usergraphql/db"
)

type LinkRepository struct {
	Db *sql.DB
}

func NewLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{
		Db: db,
	}
}

// method Save
func (l *LinkRepository) Save(ctx context.Context, entity *Link) int64 {
	stmt, err := database.Db.PrepareContext(ctx, "INSERT INTO links(Title, Address) VALUES (?, ?)")
	if err != nil {
		log.Fatalf("error prepare statement: %v", err.Error())
	}

	result, err := stmt.ExecContext(ctx, entity.Title, entity.Address)
	if err != nil {
		log.Fatalf("error execute query insert: %v", err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("error get inserted id: %v", err.Error())
	}

	log.Println("row links inserted")
	return id
}

// method get all data links
func (l *LinkRepository) GetAll(ctx context.Context) ([]Link, error) {
	queryString, err := database.Db.PrepareContext(ctx, "select ID, Title, Address from links")
	if err != nil {
		return nil, err
	}
	defer queryString.Close()

	// execute query
	rows, err := queryString.QueryContext(ctx)
	if err != nil {
		return nil, errors.New("cant execute query: " + err.Error())
	}
	defer rows.Close()

	var links []Link

	// scan data
	for rows.Next() {
		var link Link
		err = rows.Scan(&link.Id, &link.Title, &link.Address)
		if err != nil {
			return nil, errors.New("cant scan data: " + err.Error())
		}
		links = append(links, link)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(links) == 0 {
		return nil, errors.New("record links not found")
	}

	// success get data
	return links, nil
}
