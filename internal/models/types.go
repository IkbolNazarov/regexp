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
	Id    int    `gorm:"column:id"`
	Agent string `gorm:"column:agent"`
	Sort  int    `gorm:"column:sort"`
}

type Regexp struct {
	Id      int    `gorm:"column:id"`
	Regexp  string `gorm:"column:regexp"`
	Type    string `gorm:"column:type"`
	AgentID int    `gorm:"column:agent_id"`
	Agent   Agents `gorm:"foreignKey:AgentID"`
}

func GetRuleTable() string {
	return "regexp"
}

func GetTableName() string {
	return "agents"
}
