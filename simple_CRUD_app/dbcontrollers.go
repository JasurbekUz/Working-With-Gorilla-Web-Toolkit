package main

import (
	"log"
	"fmt"
	"database/sql"
	_"github.com/lib/pq"
)
// DB GET CTRL
func DBGetClassifiedsCtrl () []GetClassified {

	db, err := sql.Open("postgres", DB_CONFIG)

	defer db.Close()

	if err != nil {
		log.Fatalf("db connection error:\n %v", err)
	}

	rows, err := db.Query(SELECT_ALL)

	defer rows.Close()

	if err != nil {
		log.Fatalf("db query error:\n %v", err)
	}

	var classifieds []GetClassified

	for rows.Next() {

		var classified GetClassified

		rows.Scan(
			&classified.ClassifiedId,
			&classified.Title,
			&classified.Price,
			&classified.CreatedAt,
			&classified.Category.CategoryId,
			&classified.Category.Name,
		)

		classifieds = append(classifieds, classified)
	}

	return classifieds

}

// DB GET ONCE CLASSIFIED CTRL
func DBGetClassifiedCtrl (id int) GetClassified {

	db, err := sql.Open("postgres", DB_CONFIG)

	defer db.Close()

	if err != nil {
		log.Fatalf("db connection error:\n %v", err)
	}

	var classified GetClassified

	err = db.QueryRow(
				SELECT_ONCE,
				id,
			).Scan(
				&classified.ClassifiedId,
				&classified.Title,
				&classified.Price,
				&classified.CreatedAt,
				&classified.Category.CategoryId,
				&classified.Category.Name,
		)

	if err != nil {
		//log.Fatalf("db query error:\n %v", err)
	}

	return classified

}
// DB POST CTRL
func DBPostClassifiedCtrl (newClassified PostClassifed) GetClassified {

	db, err := sql.Open("postgres", DB_CONFIG)

	defer db.Close()

	if err != nil {
		log.Fatalf("db connection error:\n %v", err)
	}

	var classified GetClassified

	fmt.Println(newClassified)

	err = db.QueryRow(
				INSERT,
				newClassified.Title,
				newClassified.Price,
				newClassified.CategoryId,
			).Scan(
				&classified.ClassifiedId,
				&classified.Title,
				&classified.Price,
				&classified.CreatedAt,
				&classified.Category.CategoryId,
				//&classified.Category.Name,
		)

	if err != nil {
		log.Fatalf("db query error:\n %v", err)
	}

	return classified

}
// DB PUT CTRL
func DBPutClassifiedCtrl (chClassified PutClassified) GetClassified {

	db, err := sql.Open("postgres", DB_CONFIG)

	defer db.Close()

	if err != nil {
		log.Fatalf("db connection error:\n %v", err)
	}

	var classified GetClassified

	err = db.QueryRow(
				UPDATE,
				chClassified.ClassifiedId,
				chClassified.Title,
				chClassified.Price,
				chClassified.CategoryId,
			).Scan(
				&classified.ClassifiedId,
				&classified.Title,
				&classified.Price,
				&classified.CreatedAt,
				&classified.Category.CategoryId,
				//&classified.Category.Name,
		)

	fmt.Println(classified)

	if err != nil {
		return classified
		log.Fatalf("db query error:\n %v", err)
	}

	return classified

}

// DB DELETE CTRL
func DBDeleteClassifiedCtrl (id uint16) GetClassified {

	db, err := sql.Open("postgres", DB_CONFIG)

	defer db.Close()

	if err != nil {
		log.Fatalf("db connection error:\n %v", err)
	}

	var classified GetClassified

	err = db.QueryRow(
				DELETE,
				id,
			).Scan(
				&classified.ClassifiedId,
				&classified.Title,
				&classified.Price,
				&classified.CreatedAt,
				&classified.Category.CategoryId,
				//&classified.Category.Name,
		)

	if err != nil {
		return classified
		//log.Fatalf("db query error:\n %v", err)
	}

	return classified

}

