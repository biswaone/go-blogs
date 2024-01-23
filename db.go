package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func GetDB() *pgx.Conn {
	DATABASE_URL := "postgres://goblogs:goblogs@localhost:5432/goblogs"
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func setupDatabaseTables(conn *pgx.Conn) {
	if err := createUsersTable(conn); err != nil {
		log.Fatal("error creating users table err: ", err)
	}
	if err := createPostsTable(conn); err != nil {
		log.Fatal("error creating blogs table, err: ", err)
	}
	if err := createCommentsTable(conn); err != nil {
		log.Fatal("error creating blogs table, err: ", err)
	}
	if err := createTagsTable(conn); err != nil {
		log.Fatal("error creating blogs table, err: ", err)
	}
	if err := createPostTagsTable(conn); err != nil {
		log.Fatal("error creating blogs table, err: ", err)
	}
}

func createUsersTable(conn *pgx.Conn) error {
	query := `
		CREATE TABLE IF NOT EXISTS Users (
			UserID SERIAL PRIMARY KEY,
			Name VARCHAR(50) NOT NULL,
			Email VARCHAR(100) UNIQUE NOT NULL,
			Password VARCHAR(255) NOT NULL,
			ProfilePicture VARCHAR(255),
			RegistrationDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := conn.Exec(context.Background(), query)
	return err
}

func createPostsTable(conn *pgx.Conn) error {
	query := `
		CREATE TABLE IF NOT EXISTS Post (
			PostID SERIAL PRIMARY KEY,
			Title VARCHAR(255) NOT NULL,
			Content TEXT NOT NULL,
			AuthorID INT,
			CreationDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			LastUpdateDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (AuthorID) REFERENCES Users(UserID)
		);
	`
	_, err := conn.Exec(context.Background(), query)
	return err
}

func createCommentsTable(conn *pgx.Conn) error {
	query := `
		CREATE TABLE IF NOT EXISTS Comment (
			CommentID SERIAL PRIMARY KEY,
			PostID INT,
			UserID INT,
			Content TEXT NOT NULL,
			CommentDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (PostID) REFERENCES Post(PostID),
			FOREIGN KEY (UserID) REFERENCES Users(UserID)
		);
	`
	_, err := conn.Exec(context.Background(), query)
	return err
}

func createTagsTable(conn *pgx.Conn) error {
	query := `
		CREATE TABLE IF NOT EXISTS Tag (
			TagID SERIAL PRIMARY KEY,
			TagName VARCHAR(50) NOT NULL
		);
	`
	_, err := conn.Exec(context.Background(), query)
	return err
}

func createPostTagsTable(conn *pgx.Conn) error {
	query := `
		CREATE TABLE IF NOT EXISTS PostTag (
			PostID INT,
			TagID INT,
			PRIMARY KEY (PostID, TagID),
			FOREIGN KEY (PostID) REFERENCES Post(PostID),
			FOREIGN KEY (TagID) REFERENCES Tag(TagID)
		);
	`
	_, err := conn.Exec(context.Background(), query)
	return err
}
