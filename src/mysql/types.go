package mysql

type State struct {
	Stmt string
	Params []interface{}
}