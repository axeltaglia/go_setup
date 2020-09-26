package main

type Validatable interface {
	Validate() bool
}

type DatabaseConfig struct {
	Host *string `json:"host"`
	Port *string `json:"port"`
}

type EnvVars struct {
	Environment    *string         `json:"environment"`
	DatabaseConfig *DatabaseConfig `json:"databaseConfig"`
}

func (o *EnvVars) Validate() bool {
	if o.DatabaseConfig == nil {
		return false
	}
	return true
}
