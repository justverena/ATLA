package model

import (
	//"context"
	"database/sql"
	"errors"
	"log"
	//"time"
)

type Episode struct {
	ID        int    `json:"id"`
	Title     string `json:"name"`
	Air_Date  string `json:"air_date"`
	CreatedAt string `json:"createdAt"`
	ApdatedAt string `json:"updatedAt"`
}
type EpisodeModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m EpisodeModel) GetAll() ([]*Episode, error) {
	// TODO: implement this method
	return nil, errors.New("not implemented")
}
func (m EpisodeModel) Insert(menu *Episode) error {
	// TODO: implement this method
	return errors.New("not implemented")
}

func (m EpisodeModel) Get(id int) (*Episode, error) {
	// TODO: implement this method
	return nil, errors.New("not implemented")
}

func (m EpisodeModel) Update(menu *Episode) error {
	// TODO: implement this method
	return errors.New("not implemented")
}

func (m EpisodeModel) Delete(id int) error {
	// TODO: implement this method
	return errors.New("not implemented")
}
