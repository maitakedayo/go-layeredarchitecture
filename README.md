# go-layeredarchitecture

This Go program, named go-layeredarchitecture, provides functionality for managing user data in an SQLite database. It includes operations for saving users, finding users by full name, and retrieving all users. The project is structured to showcase a clean separation of concerns, with distinct layers for domain, application, and infrastructure.

## Overview

User Domain

The user domain encapsulates the core business logic related to users. It includes the User struct representing a user entity, and methods for creating and accessing user details.

User Application

The user application layer defines use cases for saving users, finding users, and retrieving all users. It utilizes the user service from the domain to interact with user data.

    SaveUserUseCase: Handles the use case of saving a new user.
    FindUserUseCase: Manages the use case of finding a user by full name.
    FindAllUsersUseCase: Handles the use case of retrieving all users.

User Infrastructure

The user infrastructure layer provides concrete implementations of interfaces defined in the domain. It includes specific implementations for SQLite, MySQL, and in-memory databases.

## Installation

Install the go-layeredarchitecture package using the following command:
```bash
$ go get github.com/maitakedayo/go-layeredarchitecture
```

## Usage

# SQLiteデータベースを使用する場合
go run ./app/cmd/sqlite/main.go

# MySQLデータベースを使用する場合
go run ./app/cmd/mysql/main.go

# インメモリデータベースを使用する場合
go run ./app/cmd/inmemory/main.go

## License

MIT

## Author

maitakedayo

## ライセンス

このプロジェクトは [MIT ライセンス](LICENSE) のもとで公開されています。