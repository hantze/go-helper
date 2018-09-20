package db

// Condition ...
type Condition struct {
	Key           string
	Value         string
	Type          string
	Connector     string
	Operator      string
	Table         *string
	StartBrackets bool
	EndBrackets   bool
}
