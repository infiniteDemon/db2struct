package config

var SysConfig = &sysConfig{}

type sysConfig struct {
	Port       string `json:"Port"`
	DBUserName string `json:"DBUserName"`
	DBPassword string `json:"DBPassword"`
	DBIp       string `json:"DBIp"`
	DBPort     string `json:"DBPort"`
	DBName     string `json:"DBName"`
	Table      string `json:"table"`
	Path       string `json:"path"`
}

// DbTable 表类型
type DbTable struct {
	Name string `json:"name"`
}

// Column 数据字段类型
type Column struct {
	ColumnName    string `json:"column_name"`
	DataType      string `json:"data_type"`
	ColumnComment string `json:"column_comment"`
	ColumnKey     string `json:"column_key"`
	Extra         string `json:"extra"`
}

