package nodequery

import "log"

var databasesNodes = map[string]*QueryResult{}
var queryNodes = map[string]map[string]*QueryResult{}

func AttachQueryResult(query string, dbKey string, queryResult *QueryResult) {
	attachQueryToDatabasesNodes(dbKey, queryResult)
	attachQueryToQueryNodes(query, dbKey, queryResult)
}

func attachQueryToDatabasesNodes(dbKey string, queryResult *QueryResult) {
	databasesNodes[dbKey] = queryResult
}

func attachQueryToQueryNodes(query string, dbKey string, queryResult *QueryResult) {
	if queryNodes[query] == nil {
		queryNodes[query] = map[string]*QueryResult{}
	}

	if queryNodes[query][dbKey] == nil {
		queryNodes[query][dbKey] = queryResult
		return
	}

	currHead := queryNodes[query][dbKey]
	queryResult.Next = currHead
	queryNodes[query][dbKey] = queryResult
}

func GetMaxLenghtOfQueryNodes(query string) int {
    log.Println("GetMaxLenghtOfQueryNodes")
	if queryNodes[query] == nil {
        log.Println("queryNodes[query] is nil")
		return 0
	}

	dbQueries := queryNodes[query]
	max := 0
	for key, dbNode := range dbQueries {
		dbNodeLength := nodeLength(dbNode)
        log.Println("key", key, dbNodeLength)
		if dbNodeLength > max {
			max = dbNodeLength
		}
	}

	return max
}

func GetQueryNode(query string, dbKey string, indexToFind int) *QueryResult {
	if queryNodes[query] == nil || queryNodes[query][dbKey] == nil {
		return nil
	}

    initialNode := queryNodes[query][dbKey] 
    if initialNode == nil {
        return nil
    }

    for i := 0; i <= indexToFind; i++ {
        if i == indexToFind {
            return initialNode

        }

        initialNode = initialNode.Next
    }

    return initialNode
}

func nodeLength(node *QueryResult) int {
	if node != nil && node.Next != nil {
		node = node.Next
		return 1 + nodeLength(node)
	}

	return 0
}

func GetLastQueryForDatabase(dbKey string) *QueryResult {
	if databasesNodes[dbKey] == nil {
		return nil
	}

	return databasesNodes[dbKey]
}
