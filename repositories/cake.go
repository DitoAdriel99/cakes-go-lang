package repositories

import (
	"database/sql"
	"fmt"
	"go-learn/config"
	"go-learn/entities"
)

type CakesRepo struct {
	conn *sql.DB
}

func NewCakesRepositories() *CakesRepo {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}

	return &CakesRepo{
		conn: conn,
	}
}

func (c *CakesRepo) Create(cake *entities.Cake) error {
	query := `INSERT INTO cakes (title, description, rating, image, created_at, updated_at) VALUES (?,?,?,?,?,?)`

	res, err := c.conn.Exec(query, cake.Title, cake.Description, cake.Rating, cake.Image, cake.CreatedAt, cake.UpdatedAt)

	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}

	cake.ID, err = res.LastInsertId()
	if err != nil {
		err = fmt.Errorf("getting last inserted id: %w", err)
		return err
	}

	return nil
}

func (c *CakesRepo) GetAll() ([]entities.Cake, error) {

	rows, err := c.conn.Query(`SELECT id, title, description, rating, image, created_at, updated_at FROM cakes`)
	if err != nil {
		err = fmt.Errorf("error querying: %w", err)
		return nil, err
	}
	var cakeObjects []entities.Cake

	for rows.Next() {
		var cakeObject entities.Cake

		err = rows.Scan(
			&cakeObject.ID,
			&cakeObject.Title,
			&cakeObject.Description,
			&cakeObject.Rating,
			&cakeObject.Image,
			&cakeObject.CreatedAt,
			&cakeObject.UpdatedAt,
		)

		if err != nil {
			err = fmt.Errorf("error scanning: %w", err)
			return nil, err
		}
		cakeObjects = append(cakeObjects, cakeObject)
	}
	return cakeObjects, nil
}

func (c *CakesRepo) FindById(id int64) (*entities.Cake, error) {
	query := `SELECT * FROM cakes WHERE id = ?`

	var cakeObject entities.Cake

	err := c.conn.QueryRow(query, id).Scan(
		&cakeObject.ID,
		&cakeObject.Title,
		&cakeObject.Description,
		&cakeObject.Rating,
		&cakeObject.Image,
		&cakeObject.CreatedAt,
		&cakeObject.UpdatedAt,
	)

	if err != nil {
		err = fmt.Errorf("scanning activity objects: %w", err)

		return nil, err
	}

	return &cakeObject, nil
}

func (c *CakesRepo) Update(cake *entities.Cake) (*entities.Cake, error) {
	query := `
		UPDATE cakes
		SET 
			title = ?,
			description = ?,
			rating = ?,
			image = ?,
			created_at = ?,
			updated_at = ?
		WHERE id = ?	
	`

	_, err := c.conn.Exec(query,
		cake.Title,
		cake.Description,
		cake.Rating,
		cake.Image,
		cake.CreatedAt,
		cake.UpdatedAt,
		cake.ID,
	)
	if err != nil {
		err = fmt.Errorf("executing query update: %w", err)

		return nil, err
	}

	return cake, nil
}

func (c *CakesRepo) Delete(id int64) error {
	query := `DELETE FROM cakes WHERE id = ?`

	_, err := c.conn.Exec(query, id)
	if err != nil {
		err = fmt.Errorf("executing query update: %w", err)
		return err
	}
	return nil
}
