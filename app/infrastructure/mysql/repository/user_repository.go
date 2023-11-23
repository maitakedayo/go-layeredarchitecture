package repository

import (
	"errors"
	"strings"
	"time"
	"database/sql"
	"log"

	userDomain "github.com/maitakedayo/go-layeredarchitecture/app/domain/user"
	_ "github.com/lib/pq" // PostgreSQL用サードパーティードライバ
)

// DTO domain領域(User)からAppli領域にDTO変換
type FindUseCaseRepoDto struct {
	ID        string
	LastName  string
	FirstName string
	Email     string
	Posts     []string   // Change: Add a field for posts
	Idlimit   time.Time  // Change: Change the type to time.Time
}
//---e-

// --- 具象実装(リポジトリ) -s-
type PostgreSQLUserRepository struct {
	DB *sql.DB
}
//
func NewPostgreSQLUserRepository() (*PostgreSQLUserRepository, error) {
	db, err := sql.Open("postgres", "user=postgres dbname=postgres password=postgresgo sslmode=disable")
	if err != nil {
		return nil, err
	}

	// Create the users table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			last_name TEXT,
			first_name TEXT,
			email TEXT,
			posts TEXT,
			idlimit TEXT
		);
	`)
	if err != nil {
		return nil, err
	}

	return &PostgreSQLUserRepository{DB: db}, nil
}
//
func (r *PostgreSQLUserRepository) Save(user *userDomain.User) error {
	/*
	_, err := r.DB.Exec(
		"INSERT INTO users (id, last_name, first_name, email, posts, idlimit) VALUES ($1, $2, $3, $4, $5, $6)",
		user.ID(), user.LastName(), user.FirstName(), user.Email(), strings.Join(user.Posts(), ","), user.Idlimit().Format("2006-01-02"),
	)
	return err
	*/
    // トランザクションを開始-s-
    tx, err := r.DB.Begin()
    if err != nil {
        return err
    }
    defer func() {
        if p := recover(); p != nil {
            // パニックが発生した場合はロールバックして再パニックを発生させます
            tx.Rollback()
            panic(p)
        } else if err != nil {
            // エラーが発生した場合はロールバック
            tx.Rollback()
        } else {
            // 成功した場合はコミット
            err = tx.Commit()
        }
    }()
	//-e-
	
	_, err = tx.Exec(
		"INSERT INTO users (id, last_name, first_name, email, posts, idlimit) VALUES ($1, $2, $3, $4, $5, $6)",
		user.ID(), user.LastName(), user.FirstName(), user.Email(), strings.Join(user.Posts(), ","), user.Idlimit().Format("2006-01-02"),
	)
	if err != nil {
		return err
	}

	return nil		
}
//
func (r *PostgreSQLUserRepository) FindByFullName(lastName string, firstName string) (*userDomain.User, error) {
	var user FindUseCaseRepoDto
	row := r.DB.QueryRow("SELECT * FROM users WHERE last_name = $1 AND first_name = $2", lastName, firstName)
	var postsStr, idlimitStr string
	if err := row.Scan(&user.ID, &user.LastName, &user.FirstName, &user.Email, &postsStr, &idlimitStr); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, err // Handle other errors
	}

	// postsの文字列をスライスに分割
	user.Posts = strings.Split(postsStr, ",")

	idlimit, err := time.Parse("2006-01-02", idlimitStr)
	if err != nil {
		return nil, err
	}

	user.Idlimit = idlimit

	newUser, err := userDomain.Reconstruct(
		user.ID,
		user.LastName,
		user.FirstName,
		user.Email,
		user.Posts,
		user.Idlimit,
	)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
//
func (r *PostgreSQLUserRepository) FindFirstUser() (*userDomain.User, error) {
	rows, err := r.DB.Query("SELECT * FROM users LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var user FindUseCaseRepoDto
		var postsStr, idlimitStr string

		if err := rows.Scan(&user.ID, &user.LastName, &user.FirstName, &user.Email, &postsStr, &idlimitStr); err != nil {
			log.Printf("Error scanning row: %s", err)
			return nil, err
		}

		// postsの文字列をスライスに分割
		user.Posts = strings.Split(postsStr, ",")

		idlimit, err := time.Parse("2006-01-02", idlimitStr)
		if err != nil {
			log.Printf("Error parsing idlimit: %s", err)
			return nil, err
		}

		user.Idlimit = idlimit

		//型変換
		newUser, err := userDomain.Reconstruct(
			user.ID,
			user.LastName,
			user.FirstName,
			user.Email,
			user.Posts,
			user.Idlimit,
		)
		if err != nil {
			log.Printf("Error reconstructing user: %s", err)
			return nil, err
		}

		return newUser, nil
	}

	// No rows found
	return nil, sql.ErrNoRows
}
//
func (r *PostgreSQLUserRepository) FindAllUsers() ([]*userDomain.User, error) {
	rows, err := r.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*userDomain.User

	for rows.Next() {
		var user FindUseCaseRepoDto
		var postsStr, idlimitStr string
		err := rows.Scan(&user.ID, &user.LastName, &user.FirstName, &user.Email, &postsStr, &idlimitStr)
		if err != nil {
			return nil, err
		}

		// postsの文字列をスライスに分割
		user.Posts = strings.Split(postsStr, ",")

		idlimit, err := time.Parse("2006-01-02", idlimitStr)
		if err != nil {
			return nil, err
		}

		user.Idlimit = idlimit

		//型変換
		newUser, err := userDomain.Reconstruct(
			user.ID,
			user.LastName,
			user.FirstName,
			user.Email,
			user.Posts,
			user.Idlimit,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, newUser)
	}

	return users, nil
}
//---e-