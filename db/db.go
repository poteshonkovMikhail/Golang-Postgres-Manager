package db

import (
	"context"
	//"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

// var DB *sql.DB

func ConnectDB(c *gin.Context) {

	var dblogs models.NoSuchParameter

	if err := json.NewDecoder(c.Request.Body).Decode(&dblogs); err != nil {
		log.Fatalf("Failed to make login request: %v", err)
	}

	var err error
	var db *pgx.Conn
	var a int = 2

	connString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", dblogs.UserParameter, dblogs.PasswordParameter, dblogs.HostnameParameter, dblogs.DatabaseNameParameter)
	db, err = pgx.Connect(context.Background(), connString)

	if err != nil {
		//Отправляет отрицательный ответ клиенту
		authDatabase(c, dblogs.UserParameter, dblogs.PasswordParameter, dblogs.HostnameParameter, dblogs.DatabaseNameParameter)
		log.Printf("UNABLE TO CONNECT TO DATABASE: %v", err)
	} else {
		InterfaceEnableDB(c, dblogs.UserParameter, dblogs.PasswordParameter, dblogs.HostnameParameter, dblogs.DatabaseNameParameter)
		log.Printf("SUCSESFULLY CONNECTED TO DATABASE: %s", dblogs.DatabaseNameParameter)
		//handlers.GetItems(c, db)
	}

	if a == 1 {
		sql := `CREATE TABLE IF NOT EXISTS messages (
			id SERIAL PRIMARY KEY,
			content TEXT NOT NULL,
			processed BOOLEAN DEFAULT FALSE
		);`

		_, err = db.Exec(context.Background(), sql)
		if err != nil {
			log.Fatalf("unable to create table: %v", err)
		}
	}

}

func authDatabase(c *gin.Context, username string, password string, hostname string, dbName string) {
	var err error
	var db *pgx.Conn
	var a int = 2
	//var k int = 0
	var no_such_parameters models.NoSuchParameter

	connString := "postgres://postgres:workout+5@localhost:5432/user_databases"
	db, err = pgx.Connect(context.Background(), connString)

	if err != nil {
		//XXXXXX

	} else {
		err := db.QueryRow(context.Background(), "SELECT username FROM users WHERE username = $1", username).Scan(&username)
		if err != nil {
			//c.String(200, "No such user_")
			no_such_parameters.UserParameter = "No such user"
			log.Printf("SQL Command Error: %s", username)
		}
		err = db.QueryRow(context.Background(), "SELECT database_password FROM databases WHERE database_password = $1", password).Scan(&password)
		if err != nil {
			//c.String(200, "Incorrect password_")
			no_such_parameters.PasswordParameter = "Incorrect password"
			log.Printf("SQL Command Error: %v", err)
		}
		err = db.QueryRow(context.Background(), "SELECT database_hostname FROM databases WHERE database_hostname = $1", hostname).Scan(&hostname)
		if err != nil {
			//c.String(200, "No such hostname_")
			no_such_parameters.HostnameParameter = "No such hostname"
			log.Printf("SQL Command Error: %v", err)
		}
		err = db.QueryRow(context.Background(), "SELECT database_name FROM databases WHERE database_name = $1", dbName).Scan(&dbName)
		if err != nil {
			//c.String(200, "No such database name_")
			no_such_parameters.DatabaseNameParameter = "No such database name"
			log.Printf("SQL Command Error: %v", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"username_parameter":      no_such_parameters.UserParameter,
			"password_parameter":      no_such_parameters.PasswordParameter,
			"hostname_parameter":      no_such_parameters.HostnameParameter,
			"database_name_parameter": no_such_parameters.DatabaseNameParameter,
		})

		//handlers.GetItems(c, db)
	}
	if a == 1 {
		sql := `CREATE TABLE IF NOT EXISTS messages (
			id SERIAL PRIMARY KEY,
			content TEXT NOT NULL,
			processed BOOLEAN DEFAULT FALSE
		);`

		_, err = db.Exec(context.Background(), sql)
		if err != nil {
			log.Fatalf("unable to create table: %v", err)
		}
	}
}

func InterfaceEnableDB(c *gin.Context, username string, password string, hostname string, dbName string) {
	var db *pgx.Conn
	//var ownerName string
	var err error

	connString := "postgres://postgres:workout+5@localhost:5432/user_databases"
	db, err = pgx.Connect(context.Background(), connString)
	if err != nil {
		c.String(200, "Ошибка подключения к базе данных: ")

	}

	// Попытка обновить запись
	cmdTag, err := db.Exec(context.Background(), "UPDATE databases SET connected = TRUE WHERE database_name = $1", dbName)
	// Проверяем, обновилась ли запись
	rowsAffected := cmdTag.RowsAffected()
	if rowsAffected == 0 {
		//var insertedID int
		var insertedID0 int
		if checkIsUserOriginal(username) == true {
			_ = db.QueryRow(context.Background(), "INSERT INTO users(username) VALUES($1) RETURNING user_id", username).Scan(&insertedID0)

			_ = db.QueryRow(context.Background(), "INSERT INTO databases(user_id, database_name, database_password, database_hostname, connected) VALUES($1, $2, $3, $4, $5)", insertedID0, dbName, password, hostname, true).Scan(&insertedID0)

			c.JSON(http.StatusOK, gin.H{
				"username_parameter":      "V",
				"password_parameter":      "V",
				"hostname_parameter":      "V",
				"database_name_parameter": "V",
			})

		} else {

			_ = db.QueryRow(context.Background(), "SELECT user_id FROM users WHERE username = ($1)", username).Scan(&insertedID0)
			//log.Printf("%v", err)
			_ = db.QueryRow(context.Background(), "INSERT INTO databases(user_id, database_name, database_password, database_hostname, connected) VALUES($1, $2, $3, $4, $5)", insertedID0, dbName, password, hostname, true).Scan(&insertedID0)

			c.JSON(http.StatusOK, gin.H{
				"username_parameter":      "V",
				"password_parameter":      "V",
				"hostname_parameter":      "V",
				"database_name_parameter": "V",
			})

		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"username_parameter":      "This database already connected",
			"password_parameter":      "This database already connected",
			"hostname_parameter":      "This database already connected",
			"database_name_parameter": "This database already connected",
		})
	}

}

func InitDBlist(c *gin.Context) {
	var db *pgx.Conn
	var err error
	//dbnames := []string{}

	connString := "postgres://postgres:workout+5@localhost:5432/user_databases"
	db, err = pgx.Connect(context.Background(), connString)
	rows, err := db.Query(context.Background(), "SELECT database_name FROM databases WHERE connected = TRUE")
	if err != nil {
		c.String(200, "Error")
	} else {

		for rows.Next() {
			var dbname string
			if err := rows.Scan(&dbname); err != nil {
				log.Printf("Ошибка при сканировании строки:", err)
			} else {

				c.String(200, "%s|", dbname)

			}
		}

	}
}

func checkIsUserOriginal(username string) bool {
	var db *pgx.Conn
	var err error
	//var user_id int

	connString := "postgres://postgres:workout+5@localhost:5432/user_databases"
	db, err = pgx.Connect(context.Background(), connString)
	_, err = db.Query(context.Background(), "SELECT username FROM users WHERE username = $1", username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return true
		}
		return false
	}
	return false

}

func connectToDatabase(dbName string) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgresql://postgres:workout+5@localhost:5432/%s", dbName)
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return pool, nil
}

func getTableInfo(pool *pgxpool.Pool) ([]models.TableInfo, error) {
	var tables []models.TableInfo

	rows, err := pool.Query(context.Background(), "SELECT table_name FROM information_schema.tables WHERE table_schema='public'")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tables: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, fmt.Errorf("failed to scan table name: %v", err)
		}

		columns, err := getTableColumns(pool, tableName)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve columns for table %s: %v", tableName, err)
		}

		dataRows, err := getTableRows(pool, tableName)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve rows for table %s: %v", tableName, err)
		}

		tables = append(tables, models.TableInfo{
			TableName: tableName,
			Columns:   columns,
			Rows:      dataRows,
		})
	}

	return tables, nil
}

func getTableColumns(pool *pgxpool.Pool, tableName string) ([]string, error) {
	var columns []string

	rows, err := pool.Query(context.Background(), "SELECT column_name FROM information_schema.columns WHERE table_name=$1", tableName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve columns: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var columnName string
		if err := rows.Scan(&columnName); err != nil {
			return nil, fmt.Errorf("failed to scan column name: %v", err)
		}
		columns = append(columns, columnName)
	}

	return columns, nil
}

func getTableRows(pool *pgxpool.Pool, tableName string) ([]map[string]interface{}, error) {
	rows, err := pool.Query(context.Background(), fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve rows: %v", err)
	}
	defer rows.Close()

	columnNames := rows.FieldDescriptions()
	var rowsData []map[string]interface{}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve row values: %v", err)
		}

		rowData := make(map[string]interface{})
		for i, columnName := range columnNames {
			rowData[string(columnName.Name)] = values[i]
		}
		rowsData = append(rowsData, rowData)
	}

	return rowsData, nil
}

func DatabaseHandler(c *gin.Context) {
	var request models.DatabaseRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	pool, err := connectToDatabase(request.DatabaseName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer pool.Close()

	tables, err := getTableInfo(pool)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve tables info: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"database_name": request.DatabaseName,
		"tables":        tables,
	})
}

// DatabaseExists проверяет, существует ли база данных с указанным именем
func DatabaseExists(conn *pgx.Conn, dbName string) (bool, error) {
	query := "SELECT 1 FROM pg_database WHERE datname=$1"
	var exists bool
	err := conn.QueryRow(context.Background(), query, dbName).Scan(&exists)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return exists, nil
}

func dbConnection(par models.CreateDbLogs) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", par.Username, par.Password, par.Hostname, par.DBName) // Получаем строку подключения из переменной окружения
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL не установлена")
	}

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %v", err)
	}

	return conn, nil
}

// CreateDatabase создает новую базу данных, если она не существует
func CreateDatabase(c *gin.Context) {
	var par models.CreateDbLogs
	if err := json.NewDecoder(c.Request.Body).Decode(&par); err != nil {
		log.Fatalf("Failed to make login request: %v", err)
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:5432/postgres", par.Username, par.Password, par.Hostname)
	db, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		c.String(200, "Err1 %v", err)
	}
	//defer db.Close(context.Background())

	exists, err := DatabaseExists(db, par.DBName)
	if err != nil {
		c.String(200, "Err2 %v", err)
	}

	if exists {
		fmt.Printf("База данных %s уже существует\n", par.DBName)
		c.String(200, "Exits ")
	}

	// Создаем SQL-запрос для создания базы данных
	query := fmt.Sprintf("CREATE DATABASE %s;", par.DBName)
	_, err = db.Exec(context.Background(), query)
	if err != nil {
		c.String(200, "Err3 ")
	}

	fmt.Printf("База данных %s успешно создана\n", par.DBName)
	c.String(200, "Kaif! ")
	InterfaceEnableDB(c, par.Username, par.Password, par.Hostname, par.DBName)
}

func ExecuteQuery(c *gin.Context) {
	var par models.SQLcommands
	if err := json.NewDecoder(c.Request.Body).Decode(&par); err != nil {
		log.Fatalf("Failed to make login request: %v", err)
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", par.Username, par.Password, par.Hostname, par.Port, par.DBName)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Failed to make login request: %v", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), par.SqlCommand)
	if err != nil {
		fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer rows.Close()

	var results []map[string]interface{}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			fmt.Errorf("ошибка получения значений строки: %v", err)
		}

		rowMap := make(map[string]interface{})
		for i, col := range rows.FieldDescriptions() {
			rowMap[string(col.Name)] = values[i]
		}

		results = append(results, rowMap)
	}

	if rows.Err() != nil {
		log.Printf("%v", rows.Err())
	}

	c.JSON(http.StatusOK, gin.H{
		"results": results,
	})

}

func InterfaceGetDbInformation(c *gin.Context) {
	var dbNameLine models.DatabaseRequest
	var pool *pgx.Conn
	var db models.DatabaseInfo
	if err := json.NewDecoder(c.Request.Body).Decode(&dbNameLine); err != nil {
		log.Fatalf("Failed to make login request: %v", err)
	}
	connString := "postgres://postgres:workout+5@localhost:5432/user_databases"
	pool, err := pgx.Connect(context.Background(), connString)
	err = pool.QueryRow(context.Background(),
		"SELECT database_id, user_id, database_name, database_password, database_hostname, connected FROM databases WHERE database_name = $1", dbNameLine.DatabaseName).Scan(&db.DatabaseID, &db.UserID, &db.DatabaseName, &db.DatabasePassword, &db.DatabaseHostname, &db.Connected)

	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Database not found"})
		} else {
			log.Printf("Query error: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	err = pool.QueryRow(context.Background(),
		"SELECT username FROM users WHERE user_id = $1", db.UserID).Scan(&db.Username)
	// Отправка результата в виде JSON
	c.JSON(http.StatusOK, db)
}
