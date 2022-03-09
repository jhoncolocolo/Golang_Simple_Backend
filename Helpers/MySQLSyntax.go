package Helpers

type MysqlSyntax struct {
    Table string
    Id string
    Columns  []string
    Timesas [2]string
}

func Structure(table string, id string,columns []string, timesass [2]string) *MysqlSyntax {
    // set only specific field value with field key
    return &MysqlSyntax{
        Table: table,
        Id: id,
        Columns: columns,
        Timesas: timesass,
    }
}