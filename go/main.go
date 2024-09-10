package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DB"),
	))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT NOW()")
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		defer rows.Close()

		var now string
		if rows.Next() {
			err := rows.Scan(&now)
			if err != nil {
				http.Error(w, "Scan error", http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
		}

		fmt.Fprintf(w, "Current time from MySQL: %s", now)
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		usersRows, err := db.Query("SELECT id, name FROM users")
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		defer usersRows.Close()

		var result string
		for usersRows.Next() {
			var id int
			var name string
			err := usersRows.Scan(&id, &name)
			if err != nil {
				http.Error(w, "Scan error", http.StatusInternalServerError)
				fmt.Println(err)
				return
			}

			// N+1クエリ: 各ユーザーの投稿を取得
			postsRows, err := db.Query("SELECT title FROM posts WHERE user_id = ?", id)
			if err != nil {
				http.Error(w, "Database error", http.StatusInternalServerError)
				fmt.Println(err)
				return
			}

			var posts []string
			for postsRows.Next() {
				var title string
				err := postsRows.Scan(&title)
				if err != nil {
					http.Error(w, "Scan error", http.StatusInternalServerError)
					fmt.Println(err)
					return
				}
				posts = append(posts, title)
			}
			postsRows.Close()

			result += fmt.Sprintf("User: %s\nPosts: %v\n", name, posts)
		}

		fmt.Fprintf(w, result)
	})

	http.HandleFunc("/complex-query", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
			SELECT u.name, COUNT(p.id) AS post_count
			FROM users u
			LEFT JOIN posts p ON u.id = p.user_id
			GROUP BY u.name
		`)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		defer rows.Close()

		var result string
		for rows.Next() {
			var name string
			var postCount int
			err := rows.Scan(&name, &postCount)
			if err != nil {
				http.Error(w, "Scan error", http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
			result += fmt.Sprintf("User: %s, Posts: %d\n", name, postCount)
		}

		fmt.Fprintf(w, result)
	})

	http.HandleFunc("/some_table", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, description FROM some_table")
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		defer rows.Close()

		var result string
		for rows.Next() {
			var id int
			var description string
			err := rows.Scan(&id, &description)
			if err != nil {
				http.Error(w, "Scan error", http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
			result += fmt.Sprintf("ID: %d, Description: %s\n", id, description)
		}

		fmt.Fprintf(w, result)
	})

	http.HandleFunc("/large_table", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, LENGTH(data) FROM large_table")
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		defer rows.Close()

		var result string
		for rows.Next() {
			var id int
			var dataLength int
			err := rows.Scan(&id, &dataLength)
			if err != nil {
				http.Error(w, "Scan error", http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
			result += fmt.Sprintf("ID: %d, Data Length: %d bytes\n", id, dataLength)
		}

		fmt.Fprintf(w, result)
	})

	http.ListenAndServe(":8080", nil)
}
