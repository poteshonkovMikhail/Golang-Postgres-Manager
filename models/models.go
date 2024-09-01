package models

type DatabaseRequest struct {
	DatabaseName string `json:"database_name"`
}

type TableInfo struct {
	TableName string                   `json:"table_name"`
	Columns   []string                 `json:"columns"`
	Rows      []map[string]interface{} `json:"rows"`
}
type CreateDbLogs struct {
	Username string `json:"username_parameter" binding:"required"`
	Password string `json:"password_parameter" binding:"required"`
	Hostname string `json:"hostname_parameter" binding:"required"`
	DBName   string `json:"database_name_parameter" binding:"required"`
	//Port     string `json:"port_parameter" binding:"required"`
}

type Response struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
	Status   string `json:"status"`
	Error    string `json:"error,omitempty"`
}

type NoSuchParameter struct {
	UserParameter         string `json:"username_parameter"`
	PasswordParameter     string `json:"password_parameter"`
	HostnameParameter     string `json:"hostname_parameter"`
	DatabaseNameParameter string `json:"database_name_parameter"`
}

type SQLcommands struct {
	Username   string `json:"username_parameter" binding:"required"`
	Password   string `json:"password_parameter" binding:"required"`
	Hostname   string `json:"hostname_parameter" binding:"required"`
	DBName     string `json:"database_name_parameter" binding:"required"`
	Port       string `json:"port_parameter" binding:"required"`
	SqlCommand string `json:"sqlCommand_parameter" binding:"required"`
}

type DatabaseInfo struct {
	DatabaseID       int    `json:"database_id"`
	UserID           int    `json:"user_id"`
	DatabaseName     string `json:"database_name"`
	DatabasePassword string `json:"database_password"`
	DatabaseHostname string `json:"database_hostname"`
	Connected        bool   `json:"connected"`
	Username         string `json:"database_username"`
}
