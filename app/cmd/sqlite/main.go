package main

import (
	"fmt"
	"log"
	"time"

	userDomain "github.com/maitakedayo/go-layeredarchitecture/app/domain/user"
	userAppli "github.com/maitakedayo/go-layeredarchitecture/app/application/user"
	userRepository "github.com/maitakedayo/go-layeredarchitecture/app/infrastructure/sqlite/repository"
	_ "github.com/mattn/go-sqlite3" // SQLite用サードパーティードライバ
)


func main() {
	// 練習用に毎回初期化しているので注意
	// Initialize repository with the SQLite database
	sqliteUserRepository, err := userRepository.NewSQLiteUserRepository()
	if err != nil {
		log.Fatal("Error initializing SQLite repository:", err)
	}
	defer sqliteUserRepository.DB.Close()

	// Initialize service with the repository
	userService := userDomain.NewUserService(sqliteUserRepository)

	// Initialize use cases with the service
	saveUserUseCase := userAppli.NewSaveUserUseCase(userService)
	findUserUseCase := userAppli.NewFindUserUseCase(userService)
	findAllUsersUseCase := userAppli.NewFindAllUsersUseCase(userService)

	// Save user
	saveDto := userAppli.SaveUseCaseDto{
		LastName:  "Smith",
		FirstName: "John",
		Email:     "john.smith@example.com",
		Posts:     []string{"Post1", "Post2"},
		Idlimit:   time.Now().AddDate(50, 0, 0),
	}
	err = saveUserUseCase.Run(saveDto)
	if err != nil {
		log.Fatal("Error saving user:", err)
	}

	// Save another user
	saveDto2 := userAppli.SaveUseCaseDto{
		LastName:  "Doe",
		FirstName: "Jane",
		Email:     "jane.doe@example.com",
		Posts:     []string{"Post1", "Post2"},
		Idlimit:   time.Now().AddDate(50, 0, 0),
	}
	err = saveUserUseCase.Run(saveDto2)
	if err != nil {
		log.Fatal("Error saving user:", err)
	}

	// Find user by full name
	findDto := userAppli.FindUseCaseDto{
		LastName:  "Smith",
		FirstName: "John",
	}
	foundUserDto, err := findUserUseCase.Run(findDto)
	if err != nil {
		log.Fatal("Error finding user:", err)
	}
	fmt.Println("===Found User:===")
	fmt.Printf("ID: %s\nLastName: %s\nFirstName: %s\nEmail: %s\nPosts: %v\nIdlimit: %s\n",
		foundUserDto.ID, foundUserDto.LastName, foundUserDto.FirstName, foundUserDto.Email, foundUserDto.Posts, foundUserDto.Idlimit)

	// Fetch and display all users
	allUsersDto, err := findAllUsersUseCase.Run()
	if err != nil {
		log.Fatal("Error fetching all users:", err)
	}
	fmt.Println("\n===All Users:===")
	for _, userDto := range allUsersDto.Users {
		fmt.Printf("ID: %s\nLastName: %s\nFirstName: %s\nEmail: %s\nPosts: %v\nIdlimit: %s\n",
			userDto.ID, userDto.LastName, userDto.FirstName, userDto.Email, userDto.Posts, userDto.Idlimit)
	}
}