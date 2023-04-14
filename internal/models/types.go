package models

type DbKey struct {
	DbConnection struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
	}
}

type Config struct {
	LocalHost struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
}

type Agents struct {
	Id     int    `gorm:"column:id"`
	Agents  string `gorm:"column:agents"`
}

type CardRule struct {
	Id       int    `gorm:"column:id"`
	Regexp   string `gorm:"column:regexp"`
	Type string `gorm:"column:type"`
	Agent    string `gorm: coolumn:agent`
}

func GetTableName() string {
	tableName := "agents"
	return tableName
}

func GetRuleTable() string {
	return "regexp"
}
