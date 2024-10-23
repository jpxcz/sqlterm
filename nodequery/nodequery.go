package nodequery

var databasesNodes = map[string]*QueryResult{}
// var queryNodes = map[string]*QueryResult{}

func AttachQueryResult(query string, dbKey string, queryResult *QueryResult) {
	databasesNodes[dbKey] = queryResult
}

func GetLastQueriesResults() map[string]*QueryResult {
	return databasesNodes
}

func GetLastQueryForDatabase(dbKey string) *QueryResult {
    if databasesNodes[dbKey] == nil {
        return nil
    }

	return databasesNodes[dbKey]
}
