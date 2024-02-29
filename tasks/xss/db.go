package xss

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "modernc.org/sqlite"
)

// This task is a XSS task ,it contains a in-memory sqlite database and a page where you can make comments.
// The databse should contain tables for a comment section.

func SetupDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err = CreateTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}

func CreateTables(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE comments (id INTEGER PRIMARY KEY, comment TEXT, user TEXT, date TEXT DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		return fmt.Errorf("failed to create comments table: %w", err)
	}

	_, err = db.Exec("INSERT INTO comments (comment, user) VALUES ('Fet Blogg Bro!', 'Susanne')")
	if err != nil {
		return fmt.Errorf("failed to insert admin user: %w", err)
	}

	_, err = db.Exec("INSERT INTO comments (comment, user) VALUES ('Takk! Jeg har satt opp varsler på nye innlegg. Så jeg kommer og leser hvis noe blir lagt inn', 'admin')")
	if err != nil {
		return fmt.Errorf("failed to insert admin user: %w", err)
	}

	return nil
}

func CreateComment(db *sql.DB, comment, user string) error {
	_, err := db.Exec("INSERT INTO comments (comment, user) VALUES (?, ?)", SurroundLinkWithTag(comment), user)
	if err != nil {
		return fmt.Errorf("failed to insert comment: %w", err)
	}
	sseChan <- "new comment"
	return nil
}

func GetCommentCount(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM comments").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get comment count: %w", err)
	}
	return count, nil
}

var linkRegex = regexp.MustCompile(`https?://[^\s]+`)

func SurroundLinkWithTag(comment string) string {

	match := linkRegex.FindAllString(comment, -1)
	for _, m := range match {
		comment = linkRegex.ReplaceAllString(comment, fmt.Sprintf("<a href=\"%s\">%s</a>", m, m))
	}
	return comment

}

func GetComments(db *sql.DB) ([]Comments, error) {
	comments := []Comments{}
	rows, err := db.Query("SELECT comment, user,date FROM comments")
	if err != nil {
		return nil, fmt.Errorf("failed to get comments: %w", err)
	}
	for rows.Next() {
		var comment Comments
		err = rows.Scan(&comment.Comment, &comment.User, &comment.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
