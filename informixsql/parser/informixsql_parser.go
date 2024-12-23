// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

// Code generated from InformixSQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // InformixSQLParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type InformixSQLParser struct {
	*antlr.BaseParser
}

var InformixSQLParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func informixsqlparserParserInit() {
	staticData := &InformixSQLParserParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "'.'", "','", "'('", "')'", "'ABORT'", "'ACTION'", "'ACCESS_METHOD'",
		"'ADD'", "'AFTER'", "'AGGREGATE'", "'ALL'", "'ALTER'", "'ANALYZE'",
		"'AND'", "'APPEND'", "'AS'", "'ASC'", "'ATTACH'", "'AUTOFREE'", "'AUTOINCREMENT'",
		"'BEFORE'", "'BEGIN'", "'BETWEEN'", "'BY'", "'CASCADE'", "'CASE'", "'CAST'",
		"'CHECK'", "'CLOSE'", "'COLLATE'", "'COLLATION'", "'COLUMN'", "'COMMIT'",
		"'COMPONENT'", "'CONFLICT'", "'CONSTRAINT'", "'CONTEXT'", "'CREATE'",
		"'CROSS'", "'CURRENT_DATE'", "'CURRENT_TIME'", "'CURRENT_TIMESTAMP'",
		"'DATABASE'", "'DATASKIP'", "'DEFAULT'", "'DEFERRABLE'", "'DEFERRED'",
		"'DEFERRED_PREPARE'", "'DELETE'", "'DEBUG'", "'DESC'", "'DETACH'", "'DISABLED'",
		"'DISTINCT'", "'DROP'", "'EACH'", "'ELSE'", "'ENABLED'", "'END'", "'ESCAPE'",
		"'EXCEPT'", "'EXCLUSIVE'", "'EXISTS'", "'EXPLAIN'", "'FAIL'", "'FOR'",
		"'FOREIGN'", "'FILE'", "'FROM'", "'FULL'", "'GLOB'", "'GROUP'", "'HAVING'",
		"'IF'", "'IGNORE'", "'IMMEDIATE'", "'IN'", "'INDEX'", "'INDEXED'", "'INITIALLY'",
		"'INNER'", "'INSERT'", "'INSTEAD'", "'INTERSECT'", "'INTO'", "'IS'",
		"'ISNULL'", "'JOIN'", "'KEY'", "'LABEL'", "'LEFT'", "'LIKE'", "'LIMIT'",
		"'MATCH'", "'NATURAL'", "'NO'", "'NOT'", "'NOTNULL'", "'NULL'", "'OF'",
		"'OFF'", "'OFFSET'", "'ON'", "'ONLINE'", "'OR'", "'ORDER'", "'OUTER'",
		"'POLICY'", "'PLAN'", "'PRAGMA'", "'PRIMARY'", "'QUERY'", "'RAISE'",
		"'RECURSIVE'", "'REFERENCES'", "'REGEXP'", "'REINDEX'", "'RELEASE'",
		"'RENAME'", "'REPLACE'", "'RESTRICT'", "'RIGHT'", "'ROLLBACK'", "'ROW'",
		"'ROLE'", "'ROWS'", "'SAVEPOINT'", "'SECURITY'", "'SELECT'", "'SET'",
		"'SEQUENCE'", "'SYNONYM'", "'TABLE'", "'TEMP'", "'TEMPORARY'", "'THEN'",
		"'TO'", "'TRANSACTION'", "'TRIGGER'", "'TRUSTED'", "'TYPE'", "'UNION'",
		"'UNIQUE'", "'UPDATE'", "'USER'", "'USING'", "'VACUUM'", "'VALUES'",
		"'VIEW'", "'VIRTUAL'", "'WHEN'", "'WHERE'", "'WITH'", "'WITHOUT'", "'WORK'",
		"'XADATASOURCE'", "'FIRST_VALUE'", "'OVER'", "'PARTITION'", "'RANGE'",
		"'PRECEDING'", "'UNBOUNDED'", "'CURRENT'", "'FOLLOWING'", "'CUME_DIST'",
		"'DENSE_RANK'", "'LAG'", "'LAST_VALUE'", "'LEAD'", "'NTH_VALUE'", "'NTILE'",
		"'PERCENT_RANK'", "'RANK'", "'ROW_NUMBER'", "'GENERATED'", "'ALWAYS'",
		"'STORED'", "'TRUE'", "'FALSE'", "'WINDOW'", "'NULLS'", "'FIRST'", "'LAST'",
		"'FILTER'", "'GROUPS'", "'EXCLUDE'",
	}
	staticData.SymbolicNames = []string{
		"", "SCOL", "DOT", "COMMA", "OPEN_PAR", "CLOSE_PAR", "ABORT", "ACTION",
		"ACCESS_METHOD", "ADD", "AFTER", "AGGREGATE", "ALL", "ALTER", "ANALYZE",
		"AND", "APPEND", "AS", "ASC", "ATTACH", "AUTOFREE", "AUTOINCREMENT",
		"BEFORE", "BEGIN", "BETWEEN", "BY", "CASCADE", "CASE", "CAST", "CHECK",
		"CLOSE", "COLLATE", "COLLATION", "COLUMN", "COMMIT", "COMPONENT", "CONFLICT",
		"CONSTRAINT", "CONTEXT", "CREATE", "CROSS", "CURRENT_DATE", "CURRENT_TIME",
		"CURRENT_TIMESTAMP", "DATABASE", "DATASKIP", "DEFAULT", "DEFERRABLE",
		"DEFERRED", "DEFERRED_PREPARE", "DELETE", "DEBUG", "DESC", "DETACH",
		"DISABLED", "DISTINCT", "DROP", "EACH", "ELSE", "ENABLED", "END", "ESCAPE",
		"EXCEPT", "EXCLUSIVE", "EXISTS", "EXPLAIN", "FAIL", "FOR", "FOREIGN",
		"FILE", "FROM", "FULL", "GLOB", "GROUP", "HAVING", "IF", "IGNORE", "IMMEDIATE",
		"IN", "INDEX", "INDEXED", "INITIALLY", "INNER", "INSERT", "INSTEAD",
		"INTERSECT", "INTO", "IS", "ISNULL", "JOIN", "KEY", "LABEL", "LEFT",
		"LIKE", "LIMIT", "MATCH", "NATURAL", "NO", "NOT", "NOTNULL", "NULL",
		"OF", "OFF", "OFFSET", "ON", "ONLINE", "OR", "ORDER", "OUTER", "POLICY",
		"PLAN", "PRAGMA", "PRIMARY", "QUERY", "RAISE", "RECURSIVE", "REFERENCES",
		"REGEXP", "REINDEX", "RELEASE", "RENAME", "REPLACE", "RESTRICT", "RIGHT",
		"ROLLBACK", "ROW", "ROLE", "ROWS", "SAVEPOINT", "SECURITY", "SELECT",
		"SET", "SEQUENCE", "SYNONYM", "TABLE", "TEMP", "TEMPORARY", "THEN",
		"TO", "TRANSACTION", "TRIGGER", "TRUSTED", "TYPE", "UNION", "UNIQUE",
		"UPDATE", "USER", "USING", "VACUUM", "VALUES", "VIEW", "VIRTUAL", "WHEN",
		"WHERE", "WITH", "WITHOUT", "WORK", "XADATASOURCE", "FIRST_VALUE", "OVER",
		"PARTITION", "RANGE", "PRECEDING", "UNBOUNDED", "CURRENT", "FOLLOWING",
		"CUME_DIST", "DENSE_RANK", "LAG", "LAST_VALUE", "LEAD", "NTH_VALUE",
		"NTILE", "PERCENT_RANK", "RANK", "ROW_NUMBER", "GENERATED", "ALWAYS",
		"STORED", "TRUE", "FALSE", "WINDOW", "NULLS", "FIRST", "LAST", "FILTER",
		"GROUPS", "EXCLUDE", "IDENTIFIER", "NUMERIC_LITERAL", "BIND_PARAMETER",
		"STRING_LITERAL", "CHAR_STRING", "SINGLE_LINE_COMMENT", "MULTILINE_COMMENT",
		"SPACES",
	}
	staticData.RuleNames = []string{
		"sqlScript", "unitStatement", "createRole", "dropRole", "dropSynonym",
		"dropTable", "dropTrigger", "dropTrustedContext", "dropType", "dropUser",
		"dropView", "dropXadatasource", "dropXadataTypeSource", "dropAccessMethod",
		"dropAggregate", "dropDatabase", "dropIndex", "closeStmt", "closeDatabaseStmt",
		"databaseStmt", "commitWorkStmt", "releaseSavepoint", "renameColumn",
		"renameConstraint", "renameDatabase", "renameIndex", "renameSecurity",
		"renameSequence", "renameTable", "renameTrustedContext", "renameUser",
		"rollbackWork", "savepointStmt", "setAutofree", "setCollation", "setDataskip",
		"setDebugFile", "setDeferredPrepareStatement", "quotedString", "anyName",
		"identifier", "keyword",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 195, 432, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 1,
		0, 5, 0, 86, 8, 0, 10, 0, 12, 0, 89, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 129, 8, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 138, 8, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 146, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 3, 4, 154, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 162, 8,
		5, 1, 5, 1, 5, 3, 5, 166, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 172, 8, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8,
		185, 8, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 3, 10, 198, 8, 10, 1, 10, 1, 10, 3, 10, 202, 8, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 3, 11, 208, 8, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 3, 12, 218, 8, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 3, 13, 227, 8, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 3, 14, 236, 8, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		3, 15, 244, 8, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 252,
		8, 16, 1, 16, 1, 16, 3, 16, 256, 8, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 19, 3, 19, 267, 8, 19, 1, 20, 1, 20, 3, 20,
		271, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 3, 26, 306, 8, 26, 1, 26, 3, 26, 309, 8, 26, 3,
		26, 311, 8, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 31, 1, 31, 3, 31, 344, 8, 31, 1, 31, 1, 31, 1, 31, 3, 31, 349, 8, 31,
		3, 31, 351, 8, 31, 1, 32, 1, 32, 1, 32, 3, 32, 356, 8, 32, 1, 33, 1, 33,
		1, 33, 3, 33, 361, 8, 33, 1, 33, 1, 33, 1, 33, 3, 33, 366, 8, 33, 3, 33,
		368, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 375, 8, 34, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 5, 35, 383, 8, 35, 10, 35, 12, 35, 386,
		9, 35, 3, 35, 388, 8, 35, 1, 35, 1, 35, 3, 35, 392, 8, 35, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 400, 8, 36, 1, 36, 1, 36, 3, 36, 404,
		8, 36, 1, 37, 1, 37, 1, 37, 3, 37, 409, 8, 37, 1, 38, 1, 38, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 420, 8, 39, 1, 40, 1, 40,
		1, 40, 5, 40, 425, 8, 40, 10, 40, 12, 40, 428, 9, 40, 1, 41, 1, 41, 1,
		41, 0, 0, 42, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
		68, 70, 72, 74, 76, 78, 80, 82, 0, 3, 2, 0, 26, 26, 122, 122, 2, 0, 54,
		54, 59, 59, 2, 0, 6, 125, 127, 187, 465, 0, 87, 1, 0, 0, 0, 2, 128, 1,
		0, 0, 0, 4, 132, 1, 0, 0, 0, 6, 141, 1, 0, 0, 0, 8, 149, 1, 0, 0, 0, 10,
		157, 1, 0, 0, 0, 12, 167, 1, 0, 0, 0, 14, 175, 1, 0, 0, 0, 16, 180, 1,
		0, 0, 0, 18, 189, 1, 0, 0, 0, 20, 193, 1, 0, 0, 0, 22, 203, 1, 0, 0, 0,
		24, 212, 1, 0, 0, 0, 26, 222, 1, 0, 0, 0, 28, 231, 1, 0, 0, 0, 30, 239,
		1, 0, 0, 0, 32, 247, 1, 0, 0, 0, 34, 257, 1, 0, 0, 0, 36, 260, 1, 0, 0,
		0, 38, 263, 1, 0, 0, 0, 40, 268, 1, 0, 0, 0, 42, 272, 1, 0, 0, 0, 44, 276,
		1, 0, 0, 0, 46, 282, 1, 0, 0, 0, 48, 288, 1, 0, 0, 0, 50, 294, 1, 0, 0,
		0, 52, 300, 1, 0, 0, 0, 54, 316, 1, 0, 0, 0, 56, 322, 1, 0, 0, 0, 58, 328,
		1, 0, 0, 0, 60, 335, 1, 0, 0, 0, 62, 341, 1, 0, 0, 0, 64, 352, 1, 0, 0,
		0, 66, 357, 1, 0, 0, 0, 68, 369, 1, 0, 0, 0, 70, 376, 1, 0, 0, 0, 72, 393,
		1, 0, 0, 0, 74, 405, 1, 0, 0, 0, 76, 410, 1, 0, 0, 0, 78, 419, 1, 0, 0,
		0, 80, 421, 1, 0, 0, 0, 82, 429, 1, 0, 0, 0, 84, 86, 3, 2, 1, 0, 85, 84,
		1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0,
		88, 90, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 91, 5, 0, 0, 1, 91, 1, 1, 0,
		0, 0, 92, 129, 3, 4, 2, 0, 93, 129, 3, 34, 17, 0, 94, 129, 3, 36, 18, 0,
		95, 129, 3, 40, 20, 0, 96, 129, 3, 26, 13, 0, 97, 129, 3, 28, 14, 0, 98,
		129, 3, 30, 15, 0, 99, 129, 3, 32, 16, 0, 100, 129, 3, 6, 3, 0, 101, 129,
		3, 8, 4, 0, 102, 129, 3, 10, 5, 0, 103, 129, 3, 12, 6, 0, 104, 129, 3,
		14, 7, 0, 105, 129, 3, 16, 8, 0, 106, 129, 3, 18, 9, 0, 107, 129, 3, 20,
		10, 0, 108, 129, 3, 22, 11, 0, 109, 129, 3, 24, 12, 0, 110, 129, 3, 38,
		19, 0, 111, 129, 3, 42, 21, 0, 112, 129, 3, 44, 22, 0, 113, 129, 3, 46,
		23, 0, 114, 129, 3, 48, 24, 0, 115, 129, 3, 50, 25, 0, 116, 129, 3, 52,
		26, 0, 117, 129, 3, 54, 27, 0, 118, 129, 3, 56, 28, 0, 119, 129, 3, 58,
		29, 0, 120, 129, 3, 60, 30, 0, 121, 129, 3, 62, 31, 0, 122, 129, 3, 64,
		32, 0, 123, 129, 3, 66, 33, 0, 124, 129, 3, 68, 34, 0, 125, 129, 3, 70,
		35, 0, 126, 129, 3, 72, 36, 0, 127, 129, 3, 74, 37, 0, 128, 92, 1, 0, 0,
		0, 128, 93, 1, 0, 0, 0, 128, 94, 1, 0, 0, 0, 128, 95, 1, 0, 0, 0, 128,
		96, 1, 0, 0, 0, 128, 97, 1, 0, 0, 0, 128, 98, 1, 0, 0, 0, 128, 99, 1, 0,
		0, 0, 128, 100, 1, 0, 0, 0, 128, 101, 1, 0, 0, 0, 128, 102, 1, 0, 0, 0,
		128, 103, 1, 0, 0, 0, 128, 104, 1, 0, 0, 0, 128, 105, 1, 0, 0, 0, 128,
		106, 1, 0, 0, 0, 128, 107, 1, 0, 0, 0, 128, 108, 1, 0, 0, 0, 128, 109,
		1, 0, 0, 0, 128, 110, 1, 0, 0, 0, 128, 111, 1, 0, 0, 0, 128, 112, 1, 0,
		0, 0, 128, 113, 1, 0, 0, 0, 128, 114, 1, 0, 0, 0, 128, 115, 1, 0, 0, 0,
		128, 116, 1, 0, 0, 0, 128, 117, 1, 0, 0, 0, 128, 118, 1, 0, 0, 0, 128,
		119, 1, 0, 0, 0, 128, 120, 1, 0, 0, 0, 128, 121, 1, 0, 0, 0, 128, 122,
		1, 0, 0, 0, 128, 123, 1, 0, 0, 0, 128, 124, 1, 0, 0, 0, 128, 125, 1, 0,
		0, 0, 128, 126, 1, 0, 0, 0, 128, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0,
		130, 131, 5, 1, 0, 0, 131, 3, 1, 0, 0, 0, 132, 133, 5, 39, 0, 0, 133, 137,
		5, 126, 0, 0, 134, 135, 5, 75, 0, 0, 135, 136, 5, 98, 0, 0, 136, 138, 5,
		64, 0, 0, 137, 134, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0,
		0, 139, 140, 3, 78, 39, 0, 140, 5, 1, 0, 0, 0, 141, 142, 5, 56, 0, 0, 142,
		145, 5, 126, 0, 0, 143, 144, 5, 75, 0, 0, 144, 146, 5, 64, 0, 0, 145, 143,
		1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 3, 78,
		39, 0, 148, 7, 1, 0, 0, 0, 149, 150, 5, 56, 0, 0, 150, 153, 5, 133, 0,
		0, 151, 152, 5, 75, 0, 0, 152, 154, 5, 64, 0, 0, 153, 151, 1, 0, 0, 0,
		153, 154, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 156, 3, 80, 40, 0, 156,
		9, 1, 0, 0, 0, 157, 158, 5, 56, 0, 0, 158, 161, 5, 134, 0, 0, 159, 160,
		5, 75, 0, 0, 160, 162, 5, 64, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1,
		0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 165, 3, 80, 40, 0, 164, 166, 7, 0,
		0, 0, 165, 164, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 11, 1, 0, 0, 0,
		167, 168, 5, 56, 0, 0, 168, 171, 5, 140, 0, 0, 169, 170, 5, 75, 0, 0, 170,
		172, 5, 64, 0, 0, 171, 169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 173,
		1, 0, 0, 0, 173, 174, 3, 80, 40, 0, 174, 13, 1, 0, 0, 0, 175, 176, 5, 56,
		0, 0, 176, 177, 5, 141, 0, 0, 177, 178, 5, 38, 0, 0, 178, 179, 3, 78, 39,
		0, 179, 15, 1, 0, 0, 0, 180, 181, 5, 56, 0, 0, 181, 184, 5, 142, 0, 0,
		182, 183, 5, 75, 0, 0, 183, 185, 5, 64, 0, 0, 184, 182, 1, 0, 0, 0, 184,
		185, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 187, 3, 80, 40, 0, 187, 188,
		5, 122, 0, 0, 188, 17, 1, 0, 0, 0, 189, 190, 5, 56, 0, 0, 190, 191, 5,
		146, 0, 0, 191, 192, 3, 78, 39, 0, 192, 19, 1, 0, 0, 0, 193, 194, 5, 56,
		0, 0, 194, 197, 5, 150, 0, 0, 195, 196, 5, 75, 0, 0, 196, 198, 5, 64, 0,
		0, 197, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199,
		201, 3, 80, 40, 0, 200, 202, 7, 0, 0, 0, 201, 200, 1, 0, 0, 0, 201, 202,
		1, 0, 0, 0, 202, 21, 1, 0, 0, 0, 203, 204, 5, 56, 0, 0, 204, 207, 5, 157,
		0, 0, 205, 206, 5, 75, 0, 0, 206, 208, 5, 64, 0, 0, 207, 205, 1, 0, 0,
		0, 207, 208, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 3, 80, 40, 0,
		210, 211, 5, 122, 0, 0, 211, 23, 1, 0, 0, 0, 212, 213, 5, 56, 0, 0, 213,
		214, 5, 157, 0, 0, 214, 217, 5, 142, 0, 0, 215, 216, 5, 75, 0, 0, 216,
		218, 5, 64, 0, 0, 217, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 219,
		1, 0, 0, 0, 219, 220, 3, 80, 40, 0, 220, 221, 5, 122, 0, 0, 221, 25, 1,
		0, 0, 0, 222, 223, 5, 56, 0, 0, 223, 226, 5, 8, 0, 0, 224, 225, 5, 75,
		0, 0, 225, 227, 5, 64, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0,
		227, 228, 1, 0, 0, 0, 228, 229, 3, 80, 40, 0, 229, 230, 5, 122, 0, 0, 230,
		27, 1, 0, 0, 0, 231, 232, 5, 56, 0, 0, 232, 235, 5, 11, 0, 0, 233, 234,
		5, 75, 0, 0, 234, 236, 5, 64, 0, 0, 235, 233, 1, 0, 0, 0, 235, 236, 1,
		0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 238, 3, 80, 40, 0, 238, 29, 1, 0, 0,
		0, 239, 240, 5, 56, 0, 0, 240, 243, 5, 44, 0, 0, 241, 242, 5, 75, 0, 0,
		242, 244, 5, 64, 0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244,
		245, 1, 0, 0, 0, 245, 246, 3, 80, 40, 0, 246, 31, 1, 0, 0, 0, 247, 248,
		5, 56, 0, 0, 248, 251, 5, 79, 0, 0, 249, 250, 5, 75, 0, 0, 250, 252, 5,
		64, 0, 0, 251, 249, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 253, 1, 0, 0,
		0, 253, 255, 3, 80, 40, 0, 254, 256, 5, 105, 0, 0, 255, 254, 1, 0, 0, 0,
		255, 256, 1, 0, 0, 0, 256, 33, 1, 0, 0, 0, 257, 258, 5, 30, 0, 0, 258,
		259, 3, 80, 40, 0, 259, 35, 1, 0, 0, 0, 260, 261, 5, 30, 0, 0, 261, 262,
		5, 44, 0, 0, 262, 37, 1, 0, 0, 0, 263, 264, 5, 44, 0, 0, 264, 266, 3, 78,
		39, 0, 265, 267, 5, 63, 0, 0, 266, 265, 1, 0, 0, 0, 266, 267, 1, 0, 0,
		0, 267, 39, 1, 0, 0, 0, 268, 270, 5, 34, 0, 0, 269, 271, 5, 156, 0, 0,
		270, 269, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 41, 1, 0, 0, 0, 272, 273,
		5, 119, 0, 0, 273, 274, 5, 128, 0, 0, 274, 275, 3, 80, 40, 0, 275, 43,
		1, 0, 0, 0, 276, 277, 5, 120, 0, 0, 277, 278, 5, 33, 0, 0, 278, 279, 3,
		80, 40, 0, 279, 280, 5, 138, 0, 0, 280, 281, 3, 80, 40, 0, 281, 45, 1,
		0, 0, 0, 282, 283, 5, 120, 0, 0, 283, 284, 5, 37, 0, 0, 284, 285, 3, 80,
		40, 0, 285, 286, 5, 138, 0, 0, 286, 287, 3, 80, 40, 0, 287, 47, 1, 0, 0,
		0, 288, 289, 5, 120, 0, 0, 289, 290, 5, 44, 0, 0, 290, 291, 3, 80, 40,
		0, 291, 292, 5, 138, 0, 0, 292, 293, 3, 80, 40, 0, 293, 49, 1, 0, 0, 0,
		294, 295, 5, 120, 0, 0, 295, 296, 5, 79, 0, 0, 296, 297, 3, 80, 40, 0,
		297, 298, 5, 138, 0, 0, 298, 299, 3, 80, 40, 0, 299, 51, 1, 0, 0, 0, 300,
		301, 5, 120, 0, 0, 301, 310, 5, 129, 0, 0, 302, 311, 5, 109, 0, 0, 303,
		308, 5, 91, 0, 0, 304, 306, 3, 80, 40, 0, 305, 304, 1, 0, 0, 0, 305, 306,
		1, 0, 0, 0, 306, 309, 1, 0, 0, 0, 307, 309, 5, 35, 0, 0, 308, 305, 1, 0,
		0, 0, 308, 307, 1, 0, 0, 0, 309, 311, 1, 0, 0, 0, 310, 302, 1, 0, 0, 0,
		310, 303, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 313, 3, 80, 40, 0, 313,
		314, 5, 138, 0, 0, 314, 315, 3, 80, 40, 0, 315, 53, 1, 0, 0, 0, 316, 317,
		5, 120, 0, 0, 317, 318, 5, 132, 0, 0, 318, 319, 3, 80, 40, 0, 319, 320,
		5, 138, 0, 0, 320, 321, 3, 80, 40, 0, 321, 55, 1, 0, 0, 0, 322, 323, 5,
		120, 0, 0, 323, 324, 5, 134, 0, 0, 324, 325, 3, 80, 40, 0, 325, 326, 5,
		138, 0, 0, 326, 327, 3, 80, 40, 0, 327, 57, 1, 0, 0, 0, 328, 329, 5, 120,
		0, 0, 329, 330, 5, 141, 0, 0, 330, 331, 5, 38, 0, 0, 331, 332, 3, 80, 40,
		0, 332, 333, 5, 138, 0, 0, 333, 334, 3, 80, 40, 0, 334, 59, 1, 0, 0, 0,
		335, 336, 5, 120, 0, 0, 336, 337, 5, 146, 0, 0, 337, 338, 3, 80, 40, 0,
		338, 339, 5, 138, 0, 0, 339, 340, 3, 80, 40, 0, 340, 61, 1, 0, 0, 0, 341,
		343, 5, 124, 0, 0, 342, 344, 5, 156, 0, 0, 343, 342, 1, 0, 0, 0, 343, 344,
		1, 0, 0, 0, 344, 350, 1, 0, 0, 0, 345, 346, 5, 138, 0, 0, 346, 348, 5,
		128, 0, 0, 347, 349, 3, 80, 40, 0, 348, 347, 1, 0, 0, 0, 348, 349, 1, 0,
		0, 0, 349, 351, 1, 0, 0, 0, 350, 345, 1, 0, 0, 0, 350, 351, 1, 0, 0, 0,
		351, 63, 1, 0, 0, 0, 352, 353, 5, 128, 0, 0, 353, 355, 3, 80, 40, 0, 354,
		356, 5, 144, 0, 0, 355, 354, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 65,
		1, 0, 0, 0, 357, 358, 5, 131, 0, 0, 358, 360, 5, 20, 0, 0, 359, 361, 7,
		1, 0, 0, 360, 359, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 367, 1, 0, 0,
		0, 362, 365, 5, 67, 0, 0, 363, 366, 3, 80, 40, 0, 364, 366, 3, 78, 39,
		0, 365, 363, 1, 0, 0, 0, 365, 364, 1, 0, 0, 0, 366, 368, 1, 0, 0, 0, 367,
		362, 1, 0, 0, 0, 367, 368, 1, 0, 0, 0, 368, 67, 1, 0, 0, 0, 369, 374, 5,
		131, 0, 0, 370, 371, 5, 32, 0, 0, 371, 375, 3, 76, 38, 0, 372, 373, 5,
		97, 0, 0, 373, 375, 5, 32, 0, 0, 374, 370, 1, 0, 0, 0, 374, 372, 1, 0,
		0, 0, 375, 69, 1, 0, 0, 0, 376, 377, 5, 131, 0, 0, 377, 391, 5, 45, 0,
		0, 378, 387, 5, 104, 0, 0, 379, 384, 3, 80, 40, 0, 380, 381, 5, 3, 0, 0,
		381, 383, 3, 80, 40, 0, 382, 380, 1, 0, 0, 0, 383, 386, 1, 0, 0, 0, 384,
		382, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 388, 1, 0, 0, 0, 386, 384,
		1, 0, 0, 0, 387, 379, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 392, 1, 0,
		0, 0, 389, 392, 5, 102, 0, 0, 390, 392, 5, 46, 0, 0, 391, 378, 1, 0, 0,
		0, 391, 389, 1, 0, 0, 0, 391, 390, 1, 0, 0, 0, 392, 71, 1, 0, 0, 0, 393,
		394, 5, 131, 0, 0, 394, 395, 5, 51, 0, 0, 395, 396, 5, 69, 0, 0, 396, 399,
		5, 138, 0, 0, 397, 400, 5, 192, 0, 0, 398, 400, 3, 80, 40, 0, 399, 397,
		1, 0, 0, 0, 399, 398, 1, 0, 0, 0, 400, 403, 1, 0, 0, 0, 401, 402, 5, 154,
		0, 0, 402, 404, 5, 16, 0, 0, 403, 401, 1, 0, 0, 0, 403, 404, 1, 0, 0, 0,
		404, 73, 1, 0, 0, 0, 405, 406, 5, 131, 0, 0, 406, 408, 5, 49, 0, 0, 407,
		409, 7, 1, 0, 0, 408, 407, 1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409, 75, 1,
		0, 0, 0, 410, 411, 5, 192, 0, 0, 411, 77, 1, 0, 0, 0, 412, 420, 5, 188,
		0, 0, 413, 420, 3, 82, 41, 0, 414, 420, 5, 191, 0, 0, 415, 416, 5, 4, 0,
		0, 416, 417, 3, 78, 39, 0, 417, 418, 5, 5, 0, 0, 418, 420, 1, 0, 0, 0,
		419, 412, 1, 0, 0, 0, 419, 413, 1, 0, 0, 0, 419, 414, 1, 0, 0, 0, 419,
		415, 1, 0, 0, 0, 420, 79, 1, 0, 0, 0, 421, 426, 3, 78, 39, 0, 422, 423,
		5, 2, 0, 0, 423, 425, 3, 78, 39, 0, 424, 422, 1, 0, 0, 0, 425, 428, 1,
		0, 0, 0, 426, 424, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 81, 1, 0, 0,
		0, 428, 426, 1, 0, 0, 0, 429, 430, 7, 2, 0, 0, 430, 83, 1, 0, 0, 0, 39,
		87, 128, 137, 145, 153, 161, 165, 171, 184, 197, 201, 207, 217, 226, 235,
		243, 251, 255, 266, 270, 305, 308, 310, 343, 348, 350, 355, 360, 365, 367,
		374, 384, 387, 391, 399, 403, 408, 419, 426,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// InformixSQLParserInit initializes any static state used to implement InformixSQLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewInformixSQLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func InformixSQLParserInit() {
	staticData := &InformixSQLParserParserStaticData
	staticData.once.Do(informixsqlparserParserInit)
}

// NewInformixSQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewInformixSQLParser(input antlr.TokenStream) *InformixSQLParser {
	InformixSQLParserInit()
	this := new(InformixSQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &InformixSQLParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "InformixSQLParser.g4"

	return this
}

// InformixSQLParser tokens.
const (
	InformixSQLParserEOF                 = antlr.TokenEOF
	InformixSQLParserSCOL                = 1
	InformixSQLParserDOT                 = 2
	InformixSQLParserCOMMA               = 3
	InformixSQLParserOPEN_PAR            = 4
	InformixSQLParserCLOSE_PAR           = 5
	InformixSQLParserABORT               = 6
	InformixSQLParserACTION              = 7
	InformixSQLParserACCESS_METHOD       = 8
	InformixSQLParserADD                 = 9
	InformixSQLParserAFTER               = 10
	InformixSQLParserAGGREGATE           = 11
	InformixSQLParserALL                 = 12
	InformixSQLParserALTER               = 13
	InformixSQLParserANALYZE             = 14
	InformixSQLParserAND                 = 15
	InformixSQLParserAPPEND              = 16
	InformixSQLParserAS                  = 17
	InformixSQLParserASC                 = 18
	InformixSQLParserATTACH              = 19
	InformixSQLParserAUTOFREE            = 20
	InformixSQLParserAUTOINCREMENT       = 21
	InformixSQLParserBEFORE              = 22
	InformixSQLParserBEGIN               = 23
	InformixSQLParserBETWEEN             = 24
	InformixSQLParserBY                  = 25
	InformixSQLParserCASCADE             = 26
	InformixSQLParserCASE                = 27
	InformixSQLParserCAST                = 28
	InformixSQLParserCHECK               = 29
	InformixSQLParserCLOSE               = 30
	InformixSQLParserCOLLATE             = 31
	InformixSQLParserCOLLATION           = 32
	InformixSQLParserCOLUMN              = 33
	InformixSQLParserCOMMIT              = 34
	InformixSQLParserCOMPONENT           = 35
	InformixSQLParserCONFLICT            = 36
	InformixSQLParserCONSTRAINT          = 37
	InformixSQLParserCONTEXT             = 38
	InformixSQLParserCREATE              = 39
	InformixSQLParserCROSS               = 40
	InformixSQLParserCURRENT_DATE        = 41
	InformixSQLParserCURRENT_TIME        = 42
	InformixSQLParserCURRENT_TIMESTAMP   = 43
	InformixSQLParserDATABASE            = 44
	InformixSQLParserDATASKIP            = 45
	InformixSQLParserDEFAULT             = 46
	InformixSQLParserDEFERRABLE          = 47
	InformixSQLParserDEFERRED            = 48
	InformixSQLParserDEFERRED_PREPARE    = 49
	InformixSQLParserDELETE              = 50
	InformixSQLParserDEBUG               = 51
	InformixSQLParserDESC                = 52
	InformixSQLParserDETACH              = 53
	InformixSQLParserDISABLED            = 54
	InformixSQLParserDISTINCT            = 55
	InformixSQLParserDROP                = 56
	InformixSQLParserEACH                = 57
	InformixSQLParserELSE                = 58
	InformixSQLParserENABLED             = 59
	InformixSQLParserEND                 = 60
	InformixSQLParserESCAPE              = 61
	InformixSQLParserEXCEPT              = 62
	InformixSQLParserEXCLUSIVE           = 63
	InformixSQLParserEXISTS              = 64
	InformixSQLParserEXPLAIN             = 65
	InformixSQLParserFAIL                = 66
	InformixSQLParserFOR                 = 67
	InformixSQLParserFOREIGN             = 68
	InformixSQLParserFILE                = 69
	InformixSQLParserFROM                = 70
	InformixSQLParserFULL                = 71
	InformixSQLParserGLOB                = 72
	InformixSQLParserGROUP               = 73
	InformixSQLParserHAVING              = 74
	InformixSQLParserIF                  = 75
	InformixSQLParserIGNORE              = 76
	InformixSQLParserIMMEDIATE           = 77
	InformixSQLParserIN                  = 78
	InformixSQLParserINDEX               = 79
	InformixSQLParserINDEXED             = 80
	InformixSQLParserINITIALLY           = 81
	InformixSQLParserINNER               = 82
	InformixSQLParserINSERT              = 83
	InformixSQLParserINSTEAD             = 84
	InformixSQLParserINTERSECT           = 85
	InformixSQLParserINTO                = 86
	InformixSQLParserIS                  = 87
	InformixSQLParserISNULL              = 88
	InformixSQLParserJOIN                = 89
	InformixSQLParserKEY                 = 90
	InformixSQLParserLABEL               = 91
	InformixSQLParserLEFT                = 92
	InformixSQLParserLIKE                = 93
	InformixSQLParserLIMIT               = 94
	InformixSQLParserMATCH               = 95
	InformixSQLParserNATURAL             = 96
	InformixSQLParserNO                  = 97
	InformixSQLParserNOT                 = 98
	InformixSQLParserNOTNULL             = 99
	InformixSQLParserNULL                = 100
	InformixSQLParserOF                  = 101
	InformixSQLParserOFF                 = 102
	InformixSQLParserOFFSET              = 103
	InformixSQLParserON                  = 104
	InformixSQLParserONLINE              = 105
	InformixSQLParserOR                  = 106
	InformixSQLParserORDER               = 107
	InformixSQLParserOUTER               = 108
	InformixSQLParserPOLICY              = 109
	InformixSQLParserPLAN                = 110
	InformixSQLParserPRAGMA              = 111
	InformixSQLParserPRIMARY             = 112
	InformixSQLParserQUERY               = 113
	InformixSQLParserRAISE               = 114
	InformixSQLParserRECURSIVE           = 115
	InformixSQLParserREFERENCES          = 116
	InformixSQLParserREGEXP              = 117
	InformixSQLParserREINDEX             = 118
	InformixSQLParserRELEASE             = 119
	InformixSQLParserRENAME              = 120
	InformixSQLParserREPLACE             = 121
	InformixSQLParserRESTRICT            = 122
	InformixSQLParserRIGHT               = 123
	InformixSQLParserROLLBACK            = 124
	InformixSQLParserROW                 = 125
	InformixSQLParserROLE                = 126
	InformixSQLParserROWS                = 127
	InformixSQLParserSAVEPOINT           = 128
	InformixSQLParserSECURITY            = 129
	InformixSQLParserSELECT              = 130
	InformixSQLParserSET                 = 131
	InformixSQLParserSEQUENCE            = 132
	InformixSQLParserSYNONYM             = 133
	InformixSQLParserTABLE               = 134
	InformixSQLParserTEMP                = 135
	InformixSQLParserTEMPORARY           = 136
	InformixSQLParserTHEN                = 137
	InformixSQLParserTO                  = 138
	InformixSQLParserTRANSACTION         = 139
	InformixSQLParserTRIGGER             = 140
	InformixSQLParserTRUSTED             = 141
	InformixSQLParserTYPE                = 142
	InformixSQLParserUNION               = 143
	InformixSQLParserUNIQUE              = 144
	InformixSQLParserUPDATE              = 145
	InformixSQLParserUSER                = 146
	InformixSQLParserUSING               = 147
	InformixSQLParserVACUUM              = 148
	InformixSQLParserVALUES              = 149
	InformixSQLParserVIEW                = 150
	InformixSQLParserVIRTUAL             = 151
	InformixSQLParserWHEN                = 152
	InformixSQLParserWHERE               = 153
	InformixSQLParserWITH                = 154
	InformixSQLParserWITHOUT             = 155
	InformixSQLParserWORK                = 156
	InformixSQLParserXADATASOURCE        = 157
	InformixSQLParserFIRST_VALUE         = 158
	InformixSQLParserOVER                = 159
	InformixSQLParserPARTITION           = 160
	InformixSQLParserRANGE               = 161
	InformixSQLParserPRECEDING           = 162
	InformixSQLParserUNBOUNDED           = 163
	InformixSQLParserCURRENT             = 164
	InformixSQLParserFOLLOWING           = 165
	InformixSQLParserCUME_DIST           = 166
	InformixSQLParserDENSE_RANK          = 167
	InformixSQLParserLAG                 = 168
	InformixSQLParserLAST_VALUE          = 169
	InformixSQLParserLEAD                = 170
	InformixSQLParserNTH_VALUE           = 171
	InformixSQLParserNTILE               = 172
	InformixSQLParserPERCENT_RANK        = 173
	InformixSQLParserRANK                = 174
	InformixSQLParserROW_NUMBER          = 175
	InformixSQLParserGENERATED           = 176
	InformixSQLParserALWAYS              = 177
	InformixSQLParserSTORED              = 178
	InformixSQLParserTRUE                = 179
	InformixSQLParserFALSE               = 180
	InformixSQLParserWINDOW              = 181
	InformixSQLParserNULLS               = 182
	InformixSQLParserFIRST               = 183
	InformixSQLParserLAST                = 184
	InformixSQLParserFILTER              = 185
	InformixSQLParserGROUPS              = 186
	InformixSQLParserEXCLUDE             = 187
	InformixSQLParserIDENTIFIER          = 188
	InformixSQLParserNUMERIC_LITERAL     = 189
	InformixSQLParserBIND_PARAMETER      = 190
	InformixSQLParserSTRING_LITERAL      = 191
	InformixSQLParserCHAR_STRING         = 192
	InformixSQLParserSINGLE_LINE_COMMENT = 193
	InformixSQLParserMULTILINE_COMMENT   = 194
	InformixSQLParserSPACES              = 195
)

// InformixSQLParser rules.
const (
	InformixSQLParserRULE_sqlScript                   = 0
	InformixSQLParserRULE_unitStatement               = 1
	InformixSQLParserRULE_createRole                  = 2
	InformixSQLParserRULE_dropRole                    = 3
	InformixSQLParserRULE_dropSynonym                 = 4
	InformixSQLParserRULE_dropTable                   = 5
	InformixSQLParserRULE_dropTrigger                 = 6
	InformixSQLParserRULE_dropTrustedContext          = 7
	InformixSQLParserRULE_dropType                    = 8
	InformixSQLParserRULE_dropUser                    = 9
	InformixSQLParserRULE_dropView                    = 10
	InformixSQLParserRULE_dropXadatasource            = 11
	InformixSQLParserRULE_dropXadataTypeSource        = 12
	InformixSQLParserRULE_dropAccessMethod            = 13
	InformixSQLParserRULE_dropAggregate               = 14
	InformixSQLParserRULE_dropDatabase                = 15
	InformixSQLParserRULE_dropIndex                   = 16
	InformixSQLParserRULE_closeStmt                   = 17
	InformixSQLParserRULE_closeDatabaseStmt           = 18
	InformixSQLParserRULE_databaseStmt                = 19
	InformixSQLParserRULE_commitWorkStmt              = 20
	InformixSQLParserRULE_releaseSavepoint            = 21
	InformixSQLParserRULE_renameColumn                = 22
	InformixSQLParserRULE_renameConstraint            = 23
	InformixSQLParserRULE_renameDatabase              = 24
	InformixSQLParserRULE_renameIndex                 = 25
	InformixSQLParserRULE_renameSecurity              = 26
	InformixSQLParserRULE_renameSequence              = 27
	InformixSQLParserRULE_renameTable                 = 28
	InformixSQLParserRULE_renameTrustedContext        = 29
	InformixSQLParserRULE_renameUser                  = 30
	InformixSQLParserRULE_rollbackWork                = 31
	InformixSQLParserRULE_savepointStmt               = 32
	InformixSQLParserRULE_setAutofree                 = 33
	InformixSQLParserRULE_setCollation                = 34
	InformixSQLParserRULE_setDataskip                 = 35
	InformixSQLParserRULE_setDebugFile                = 36
	InformixSQLParserRULE_setDeferredPrepareStatement = 37
	InformixSQLParserRULE_quotedString                = 38
	InformixSQLParserRULE_anyName                     = 39
	InformixSQLParserRULE_identifier                  = 40
	InformixSQLParserRULE_keyword                     = 41
)

// ISqlScriptContext is an interface to support dynamic dispatch.
type ISqlScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllUnitStatement() []IUnitStatementContext
	UnitStatement(i int) IUnitStatementContext

	// IsSqlScriptContext differentiates from other interfaces.
	IsSqlScriptContext()
}

type SqlScriptContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqlScriptContext() *SqlScriptContext {
	var p = new(SqlScriptContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_sqlScript
	return p
}

func InitEmptySqlScriptContext(p *SqlScriptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_sqlScript
}

func (*SqlScriptContext) IsSqlScriptContext() {}

func NewSqlScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlScriptContext {
	var p = new(SqlScriptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_sqlScript

	return p
}

func (s *SqlScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlScriptContext) EOF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEOF, 0)
}

func (s *SqlScriptContext) AllUnitStatement() []IUnitStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnitStatementContext); ok {
			len++
		}
	}

	tst := make([]IUnitStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnitStatementContext); ok {
			tst[i] = t.(IUnitStatementContext)
			i++
		}
	}

	return tst
}

func (s *SqlScriptContext) UnitStatement(i int) IUnitStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitStatementContext)
}

func (s *SqlScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SqlScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterSqlScript(s)
	}
}

func (s *SqlScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitSqlScript(s)
	}
}

func (p *InformixSQLParser) SqlScript() (localctx ISqlScriptContext) {
	localctx = NewSqlScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, InformixSQLParserRULE_sqlScript)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72075754233397248) != 0) || ((int64((_la-119)) & ^0x3f) == 0 && ((int64(1)<<(_la-119))&4643) != 0) {
		{
			p.SetState(84)
			p.UnitStatement()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(InformixSQLParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnitStatementContext is an interface to support dynamic dispatch.
type IUnitStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SCOL() antlr.TerminalNode
	CreateRole() ICreateRoleContext
	CloseStmt() ICloseStmtContext
	CloseDatabaseStmt() ICloseDatabaseStmtContext
	CommitWorkStmt() ICommitWorkStmtContext
	DropAccessMethod() IDropAccessMethodContext
	DropAggregate() IDropAggregateContext
	DropDatabase() IDropDatabaseContext
	DropIndex() IDropIndexContext
	DropRole() IDropRoleContext
	DropSynonym() IDropSynonymContext
	DropTable() IDropTableContext
	DropTrigger() IDropTriggerContext
	DropTrustedContext() IDropTrustedContextContext
	DropType() IDropTypeContext
	DropUser() IDropUserContext
	DropView() IDropViewContext
	DropXadatasource() IDropXadatasourceContext
	DropXadataTypeSource() IDropXadataTypeSourceContext
	DatabaseStmt() IDatabaseStmtContext
	ReleaseSavepoint() IReleaseSavepointContext
	RenameColumn() IRenameColumnContext
	RenameConstraint() IRenameConstraintContext
	RenameDatabase() IRenameDatabaseContext
	RenameIndex() IRenameIndexContext
	RenameSecurity() IRenameSecurityContext
	RenameSequence() IRenameSequenceContext
	RenameTable() IRenameTableContext
	RenameTrustedContext() IRenameTrustedContextContext
	RenameUser() IRenameUserContext
	RollbackWork() IRollbackWorkContext
	SavepointStmt() ISavepointStmtContext
	SetAutofree() ISetAutofreeContext
	SetCollation() ISetCollationContext
	SetDataskip() ISetDataskipContext
	SetDebugFile() ISetDebugFileContext
	SetDeferredPrepareStatement() ISetDeferredPrepareStatementContext

	// IsUnitStatementContext differentiates from other interfaces.
	IsUnitStatementContext()
}

type UnitStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitStatementContext() *UnitStatementContext {
	var p = new(UnitStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_unitStatement
	return p
}

func InitEmptyUnitStatementContext(p *UnitStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_unitStatement
}

func (*UnitStatementContext) IsUnitStatementContext() {}

func NewUnitStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitStatementContext {
	var p = new(UnitStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_unitStatement

	return p
}

func (s *UnitStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitStatementContext) SCOL() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSCOL, 0)
}

func (s *UnitStatementContext) CreateRole() ICreateRoleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateRoleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateRoleContext)
}

func (s *UnitStatementContext) CloseStmt() ICloseStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICloseStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICloseStmtContext)
}

func (s *UnitStatementContext) CloseDatabaseStmt() ICloseDatabaseStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICloseDatabaseStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICloseDatabaseStmtContext)
}

func (s *UnitStatementContext) CommitWorkStmt() ICommitWorkStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommitWorkStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommitWorkStmtContext)
}

func (s *UnitStatementContext) DropAccessMethod() IDropAccessMethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropAccessMethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropAccessMethodContext)
}

func (s *UnitStatementContext) DropAggregate() IDropAggregateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropAggregateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropAggregateContext)
}

func (s *UnitStatementContext) DropDatabase() IDropDatabaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropDatabaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropDatabaseContext)
}

func (s *UnitStatementContext) DropIndex() IDropIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropIndexContext)
}

func (s *UnitStatementContext) DropRole() IDropRoleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropRoleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropRoleContext)
}

func (s *UnitStatementContext) DropSynonym() IDropSynonymContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropSynonymContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropSynonymContext)
}

func (s *UnitStatementContext) DropTable() IDropTableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropTableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropTableContext)
}

func (s *UnitStatementContext) DropTrigger() IDropTriggerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropTriggerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropTriggerContext)
}

func (s *UnitStatementContext) DropTrustedContext() IDropTrustedContextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropTrustedContextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropTrustedContextContext)
}

func (s *UnitStatementContext) DropType() IDropTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropTypeContext)
}

func (s *UnitStatementContext) DropUser() IDropUserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropUserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropUserContext)
}

func (s *UnitStatementContext) DropView() IDropViewContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropViewContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropViewContext)
}

func (s *UnitStatementContext) DropXadatasource() IDropXadatasourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropXadatasourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropXadatasourceContext)
}

func (s *UnitStatementContext) DropXadataTypeSource() IDropXadataTypeSourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropXadataTypeSourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropXadataTypeSourceContext)
}

func (s *UnitStatementContext) DatabaseStmt() IDatabaseStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatabaseStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatabaseStmtContext)
}

func (s *UnitStatementContext) ReleaseSavepoint() IReleaseSavepointContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReleaseSavepointContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReleaseSavepointContext)
}

func (s *UnitStatementContext) RenameColumn() IRenameColumnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRenameColumnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRenameColumnContext)
}

func (s *UnitStatementContext) RenameConstraint() IRenameConstraintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRenameConstraintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRenameConstraintContext)
}

func (s *UnitStatementContext) RenameDatabase() IRenameDatabaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRenameDatabaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRenameDatabaseContext)
}

func (s *UnitStatementContext) RenameIndex() IRenameIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRenameIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRenameIndexContext)
}

func (s *UnitStatementContext) RenameSecurity() IRenameSecurityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRenameSecurityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRenameSecurityContext)
}

func (s *UnitStatementContext) RenameSequence() IRenameSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRenameSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRenameSequenceContext)
}

func (s *UnitStatementContext) RenameTable() IRenameTableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRenameTableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRenameTableContext)
}

func (s *UnitStatementContext) RenameTrustedContext() IRenameTrustedContextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRenameTrustedContextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRenameTrustedContextContext)
}

func (s *UnitStatementContext) RenameUser() IRenameUserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRenameUserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRenameUserContext)
}

func (s *UnitStatementContext) RollbackWork() IRollbackWorkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRollbackWorkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRollbackWorkContext)
}

func (s *UnitStatementContext) SavepointStmt() ISavepointStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISavepointStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISavepointStmtContext)
}

func (s *UnitStatementContext) SetAutofree() ISetAutofreeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetAutofreeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetAutofreeContext)
}

func (s *UnitStatementContext) SetCollation() ISetCollationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetCollationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetCollationContext)
}

func (s *UnitStatementContext) SetDataskip() ISetDataskipContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetDataskipContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetDataskipContext)
}

func (s *UnitStatementContext) SetDebugFile() ISetDebugFileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetDebugFileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetDebugFileContext)
}

func (s *UnitStatementContext) SetDeferredPrepareStatement() ISetDeferredPrepareStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetDeferredPrepareStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetDeferredPrepareStatementContext)
}

func (s *UnitStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterUnitStatement(s)
	}
}

func (s *UnitStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitUnitStatement(s)
	}
}

func (p *InformixSQLParser) UnitStatement() (localctx IUnitStatementContext) {
	localctx = NewUnitStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InformixSQLParserRULE_unitStatement)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(92)
			p.CreateRole()
		}

	case 2:
		{
			p.SetState(93)
			p.CloseStmt()
		}

	case 3:
		{
			p.SetState(94)
			p.CloseDatabaseStmt()
		}

	case 4:
		{
			p.SetState(95)
			p.CommitWorkStmt()
		}

	case 5:
		{
			p.SetState(96)
			p.DropAccessMethod()
		}

	case 6:
		{
			p.SetState(97)
			p.DropAggregate()
		}

	case 7:
		{
			p.SetState(98)
			p.DropDatabase()
		}

	case 8:
		{
			p.SetState(99)
			p.DropIndex()
		}

	case 9:
		{
			p.SetState(100)
			p.DropRole()
		}

	case 10:
		{
			p.SetState(101)
			p.DropSynonym()
		}

	case 11:
		{
			p.SetState(102)
			p.DropTable()
		}

	case 12:
		{
			p.SetState(103)
			p.DropTrigger()
		}

	case 13:
		{
			p.SetState(104)
			p.DropTrustedContext()
		}

	case 14:
		{
			p.SetState(105)
			p.DropType()
		}

	case 15:
		{
			p.SetState(106)
			p.DropUser()
		}

	case 16:
		{
			p.SetState(107)
			p.DropView()
		}

	case 17:
		{
			p.SetState(108)
			p.DropXadatasource()
		}

	case 18:
		{
			p.SetState(109)
			p.DropXadataTypeSource()
		}

	case 19:
		{
			p.SetState(110)
			p.DatabaseStmt()
		}

	case 20:
		{
			p.SetState(111)
			p.ReleaseSavepoint()
		}

	case 21:
		{
			p.SetState(112)
			p.RenameColumn()
		}

	case 22:
		{
			p.SetState(113)
			p.RenameConstraint()
		}

	case 23:
		{
			p.SetState(114)
			p.RenameDatabase()
		}

	case 24:
		{
			p.SetState(115)
			p.RenameIndex()
		}

	case 25:
		{
			p.SetState(116)
			p.RenameSecurity()
		}

	case 26:
		{
			p.SetState(117)
			p.RenameSequence()
		}

	case 27:
		{
			p.SetState(118)
			p.RenameTable()
		}

	case 28:
		{
			p.SetState(119)
			p.RenameTrustedContext()
		}

	case 29:
		{
			p.SetState(120)
			p.RenameUser()
		}

	case 30:
		{
			p.SetState(121)
			p.RollbackWork()
		}

	case 31:
		{
			p.SetState(122)
			p.SavepointStmt()
		}

	case 32:
		{
			p.SetState(123)
			p.SetAutofree()
		}

	case 33:
		{
			p.SetState(124)
			p.SetCollation()
		}

	case 34:
		{
			p.SetState(125)
			p.SetDataskip()
		}

	case 35:
		{
			p.SetState(126)
			p.SetDebugFile()
		}

	case 36:
		{
			p.SetState(127)
			p.SetDeferredPrepareStatement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(130)
		p.Match(InformixSQLParserSCOL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICreateRoleContext is an interface to support dynamic dispatch.
type ICreateRoleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRoleName returns the roleName rule contexts.
	GetRoleName() IAnyNameContext

	// SetRoleName sets the roleName rule contexts.
	SetRoleName(IAnyNameContext)

	// Getter signatures
	CREATE() antlr.TerminalNode
	ROLE() antlr.TerminalNode
	AnyName() IAnyNameContext
	IF() antlr.TerminalNode
	NOT() antlr.TerminalNode
	EXISTS() antlr.TerminalNode

	// IsCreateRoleContext differentiates from other interfaces.
	IsCreateRoleContext()
}

type CreateRoleContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	roleName IAnyNameContext
}

func NewEmptyCreateRoleContext() *CreateRoleContext {
	var p = new(CreateRoleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_createRole
	return p
}

func InitEmptyCreateRoleContext(p *CreateRoleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_createRole
}

func (*CreateRoleContext) IsCreateRoleContext() {}

func NewCreateRoleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateRoleContext {
	var p = new(CreateRoleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_createRole

	return p
}

func (s *CreateRoleContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateRoleContext) GetRoleName() IAnyNameContext { return s.roleName }

func (s *CreateRoleContext) SetRoleName(v IAnyNameContext) { s.roleName = v }

func (s *CreateRoleContext) CREATE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCREATE, 0)
}

func (s *CreateRoleContext) ROLE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserROLE, 0)
}

func (s *CreateRoleContext) AnyName() IAnyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyNameContext)
}

func (s *CreateRoleContext) IF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIF, 0)
}

func (s *CreateRoleContext) NOT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserNOT, 0)
}

func (s *CreateRoleContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXISTS, 0)
}

func (s *CreateRoleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateRoleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateRoleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterCreateRole(s)
	}
}

func (s *CreateRoleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitCreateRole(s)
	}
}

func (p *InformixSQLParser) CreateRole() (localctx ICreateRoleContext) {
	localctx = NewCreateRoleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, InformixSQLParserRULE_createRole)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(InformixSQLParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(InformixSQLParserROLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(134)
			p.Match(InformixSQLParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Match(InformixSQLParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Match(InformixSQLParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(139)

		var _x = p.AnyName()

		localctx.(*CreateRoleContext).roleName = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropRoleContext is an interface to support dynamic dispatch.
type IDropRoleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRoleName returns the roleName rule contexts.
	GetRoleName() IAnyNameContext

	// SetRoleName sets the roleName rule contexts.
	SetRoleName(IAnyNameContext)

	// Getter signatures
	DROP() antlr.TerminalNode
	ROLE() antlr.TerminalNode
	AnyName() IAnyNameContext
	IF() antlr.TerminalNode
	EXISTS() antlr.TerminalNode

	// IsDropRoleContext differentiates from other interfaces.
	IsDropRoleContext()
}

type DropRoleContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	roleName IAnyNameContext
}

func NewEmptyDropRoleContext() *DropRoleContext {
	var p = new(DropRoleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropRole
	return p
}

func InitEmptyDropRoleContext(p *DropRoleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropRole
}

func (*DropRoleContext) IsDropRoleContext() {}

func NewDropRoleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropRoleContext {
	var p = new(DropRoleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_dropRole

	return p
}

func (s *DropRoleContext) GetParser() antlr.Parser { return s.parser }

func (s *DropRoleContext) GetRoleName() IAnyNameContext { return s.roleName }

func (s *DropRoleContext) SetRoleName(v IAnyNameContext) { s.roleName = v }

func (s *DropRoleContext) DROP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDROP, 0)
}

func (s *DropRoleContext) ROLE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserROLE, 0)
}

func (s *DropRoleContext) AnyName() IAnyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyNameContext)
}

func (s *DropRoleContext) IF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIF, 0)
}

func (s *DropRoleContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXISTS, 0)
}

func (s *DropRoleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropRoleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropRoleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterDropRole(s)
	}
}

func (s *DropRoleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitDropRole(s)
	}
}

func (p *InformixSQLParser) DropRole() (localctx IDropRoleContext) {
	localctx = NewDropRoleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, InformixSQLParserRULE_dropRole)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(InformixSQLParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Match(InformixSQLParserROLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(143)
			p.Match(InformixSQLParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.Match(InformixSQLParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(147)

		var _x = p.AnyName()

		localctx.(*DropRoleContext).roleName = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropSynonymContext is an interface to support dynamic dispatch.
type IDropSynonymContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSynonymName returns the synonymName rule contexts.
	GetSynonymName() IIdentifierContext

	// SetSynonymName sets the synonymName rule contexts.
	SetSynonymName(IIdentifierContext)

	// Getter signatures
	DROP() antlr.TerminalNode
	SYNONYM() antlr.TerminalNode
	Identifier() IIdentifierContext
	IF() antlr.TerminalNode
	EXISTS() antlr.TerminalNode

	// IsDropSynonymContext differentiates from other interfaces.
	IsDropSynonymContext()
}

type DropSynonymContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	synonymName IIdentifierContext
}

func NewEmptyDropSynonymContext() *DropSynonymContext {
	var p = new(DropSynonymContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropSynonym
	return p
}

func InitEmptyDropSynonymContext(p *DropSynonymContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropSynonym
}

func (*DropSynonymContext) IsDropSynonymContext() {}

func NewDropSynonymContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropSynonymContext {
	var p = new(DropSynonymContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_dropSynonym

	return p
}

func (s *DropSynonymContext) GetParser() antlr.Parser { return s.parser }

func (s *DropSynonymContext) GetSynonymName() IIdentifierContext { return s.synonymName }

func (s *DropSynonymContext) SetSynonymName(v IIdentifierContext) { s.synonymName = v }

func (s *DropSynonymContext) DROP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDROP, 0)
}

func (s *DropSynonymContext) SYNONYM() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSYNONYM, 0)
}

func (s *DropSynonymContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DropSynonymContext) IF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIF, 0)
}

func (s *DropSynonymContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXISTS, 0)
}

func (s *DropSynonymContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropSynonymContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropSynonymContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterDropSynonym(s)
	}
}

func (s *DropSynonymContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitDropSynonym(s)
	}
}

func (p *InformixSQLParser) DropSynonym() (localctx IDropSynonymContext) {
	localctx = NewDropSynonymContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, InformixSQLParserRULE_dropSynonym)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(InformixSQLParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(InformixSQLParserSYNONYM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(151)
			p.Match(InformixSQLParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Match(InformixSQLParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(155)

		var _x = p.Identifier()

		localctx.(*DropSynonymContext).synonymName = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropTableContext is an interface to support dynamic dispatch.
type IDropTableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTableName returns the tableName rule contexts.
	GetTableName() IIdentifierContext

	// SetTableName sets the tableName rule contexts.
	SetTableName(IIdentifierContext)

	// Getter signatures
	DROP() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	Identifier() IIdentifierContext
	IF() antlr.TerminalNode
	EXISTS() antlr.TerminalNode
	CASCADE() antlr.TerminalNode
	RESTRICT() antlr.TerminalNode

	// IsDropTableContext differentiates from other interfaces.
	IsDropTableContext()
}

type DropTableContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	tableName IIdentifierContext
}

func NewEmptyDropTableContext() *DropTableContext {
	var p = new(DropTableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropTable
	return p
}

func InitEmptyDropTableContext(p *DropTableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropTable
}

func (*DropTableContext) IsDropTableContext() {}

func NewDropTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropTableContext {
	var p = new(DropTableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_dropTable

	return p
}

func (s *DropTableContext) GetParser() antlr.Parser { return s.parser }

func (s *DropTableContext) GetTableName() IIdentifierContext { return s.tableName }

func (s *DropTableContext) SetTableName(v IIdentifierContext) { s.tableName = v }

func (s *DropTableContext) DROP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDROP, 0)
}

func (s *DropTableContext) TABLE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTABLE, 0)
}

func (s *DropTableContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DropTableContext) IF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIF, 0)
}

func (s *DropTableContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXISTS, 0)
}

func (s *DropTableContext) CASCADE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCASCADE, 0)
}

func (s *DropTableContext) RESTRICT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRESTRICT, 0)
}

func (s *DropTableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropTableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropTableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterDropTable(s)
	}
}

func (s *DropTableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitDropTable(s)
	}
}

func (p *InformixSQLParser) DropTable() (localctx IDropTableContext) {
	localctx = NewDropTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, InformixSQLParserRULE_dropTable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(InformixSQLParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(InformixSQLParserTABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(159)
			p.Match(InformixSQLParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Match(InformixSQLParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(163)

		var _x = p.Identifier()

		localctx.(*DropTableContext).tableName = _x
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InformixSQLParserCASCADE || _la == InformixSQLParserRESTRICT {
		{
			p.SetState(164)
			_la = p.GetTokenStream().LA(1)

			if !(_la == InformixSQLParserCASCADE || _la == InformixSQLParserRESTRICT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropTriggerContext is an interface to support dynamic dispatch.
type IDropTriggerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTriggerName returns the triggerName rule contexts.
	GetTriggerName() IIdentifierContext

	// SetTriggerName sets the triggerName rule contexts.
	SetTriggerName(IIdentifierContext)

	// Getter signatures
	DROP() antlr.TerminalNode
	TRIGGER() antlr.TerminalNode
	Identifier() IIdentifierContext
	IF() antlr.TerminalNode
	EXISTS() antlr.TerminalNode

	// IsDropTriggerContext differentiates from other interfaces.
	IsDropTriggerContext()
}

type DropTriggerContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	triggerName IIdentifierContext
}

func NewEmptyDropTriggerContext() *DropTriggerContext {
	var p = new(DropTriggerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropTrigger
	return p
}

func InitEmptyDropTriggerContext(p *DropTriggerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropTrigger
}

func (*DropTriggerContext) IsDropTriggerContext() {}

func NewDropTriggerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropTriggerContext {
	var p = new(DropTriggerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_dropTrigger

	return p
}

func (s *DropTriggerContext) GetParser() antlr.Parser { return s.parser }

func (s *DropTriggerContext) GetTriggerName() IIdentifierContext { return s.triggerName }

func (s *DropTriggerContext) SetTriggerName(v IIdentifierContext) { s.triggerName = v }

func (s *DropTriggerContext) DROP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDROP, 0)
}

func (s *DropTriggerContext) TRIGGER() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTRIGGER, 0)
}

func (s *DropTriggerContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DropTriggerContext) IF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIF, 0)
}

func (s *DropTriggerContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXISTS, 0)
}

func (s *DropTriggerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropTriggerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropTriggerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterDropTrigger(s)
	}
}

func (s *DropTriggerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitDropTrigger(s)
	}
}

func (p *InformixSQLParser) DropTrigger() (localctx IDropTriggerContext) {
	localctx = NewDropTriggerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, InformixSQLParserRULE_dropTrigger)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(InformixSQLParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(InformixSQLParserTRIGGER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(169)
			p.Match(InformixSQLParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Match(InformixSQLParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(173)

		var _x = p.Identifier()

		localctx.(*DropTriggerContext).triggerName = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropTrustedContextContext is an interface to support dynamic dispatch.
type IDropTrustedContextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetContextName returns the contextName rule contexts.
	GetContextName() IAnyNameContext

	// SetContextName sets the contextName rule contexts.
	SetContextName(IAnyNameContext)

	// Getter signatures
	DROP() antlr.TerminalNode
	TRUSTED() antlr.TerminalNode
	CONTEXT() antlr.TerminalNode
	AnyName() IAnyNameContext

	// IsDropTrustedContextContext differentiates from other interfaces.
	IsDropTrustedContextContext()
}

type DropTrustedContextContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	contextName IAnyNameContext
}

func NewEmptyDropTrustedContextContext() *DropTrustedContextContext {
	var p = new(DropTrustedContextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropTrustedContext
	return p
}

func InitEmptyDropTrustedContextContext(p *DropTrustedContextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropTrustedContext
}

func (*DropTrustedContextContext) IsDropTrustedContextContext() {}

func NewDropTrustedContextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropTrustedContextContext {
	var p = new(DropTrustedContextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_dropTrustedContext

	return p
}

func (s *DropTrustedContextContext) GetParser() antlr.Parser { return s.parser }

func (s *DropTrustedContextContext) GetContextName() IAnyNameContext { return s.contextName }

func (s *DropTrustedContextContext) SetContextName(v IAnyNameContext) { s.contextName = v }

func (s *DropTrustedContextContext) DROP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDROP, 0)
}

func (s *DropTrustedContextContext) TRUSTED() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTRUSTED, 0)
}

func (s *DropTrustedContextContext) CONTEXT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCONTEXT, 0)
}

func (s *DropTrustedContextContext) AnyName() IAnyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyNameContext)
}

func (s *DropTrustedContextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropTrustedContextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropTrustedContextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterDropTrustedContext(s)
	}
}

func (s *DropTrustedContextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitDropTrustedContext(s)
	}
}

func (p *InformixSQLParser) DropTrustedContext() (localctx IDropTrustedContextContext) {
	localctx = NewDropTrustedContextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, InformixSQLParserRULE_dropTrustedContext)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(InformixSQLParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(InformixSQLParserTRUSTED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Match(InformixSQLParserCONTEXT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)

		var _x = p.AnyName()

		localctx.(*DropTrustedContextContext).contextName = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropTypeContext is an interface to support dynamic dispatch.
type IDropTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDataTypeName returns the dataTypeName rule contexts.
	GetDataTypeName() IIdentifierContext

	// SetDataTypeName sets the dataTypeName rule contexts.
	SetDataTypeName(IIdentifierContext)

	// Getter signatures
	DROP() antlr.TerminalNode
	TYPE() antlr.TerminalNode
	RESTRICT() antlr.TerminalNode
	Identifier() IIdentifierContext
	IF() antlr.TerminalNode
	EXISTS() antlr.TerminalNode

	// IsDropTypeContext differentiates from other interfaces.
	IsDropTypeContext()
}

type DropTypeContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	dataTypeName IIdentifierContext
}

func NewEmptyDropTypeContext() *DropTypeContext {
	var p = new(DropTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropType
	return p
}

func InitEmptyDropTypeContext(p *DropTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropType
}

func (*DropTypeContext) IsDropTypeContext() {}

func NewDropTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropTypeContext {
	var p = new(DropTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_dropType

	return p
}

func (s *DropTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DropTypeContext) GetDataTypeName() IIdentifierContext { return s.dataTypeName }

func (s *DropTypeContext) SetDataTypeName(v IIdentifierContext) { s.dataTypeName = v }

func (s *DropTypeContext) DROP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDROP, 0)
}

func (s *DropTypeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTYPE, 0)
}

func (s *DropTypeContext) RESTRICT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRESTRICT, 0)
}

func (s *DropTypeContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DropTypeContext) IF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIF, 0)
}

func (s *DropTypeContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXISTS, 0)
}

func (s *DropTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterDropType(s)
	}
}

func (s *DropTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitDropType(s)
	}
}

func (p *InformixSQLParser) DropType() (localctx IDropTypeContext) {
	localctx = NewDropTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, InformixSQLParserRULE_dropType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(InformixSQLParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Match(InformixSQLParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(182)
			p.Match(InformixSQLParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Match(InformixSQLParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(186)

		var _x = p.Identifier()

		localctx.(*DropTypeContext).dataTypeName = _x
	}
	{
		p.SetState(187)
		p.Match(InformixSQLParserRESTRICT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropUserContext is an interface to support dynamic dispatch.
type IDropUserContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetUserName returns the userName rule contexts.
	GetUserName() IAnyNameContext

	// SetUserName sets the userName rule contexts.
	SetUserName(IAnyNameContext)

	// Getter signatures
	DROP() antlr.TerminalNode
	USER() antlr.TerminalNode
	AnyName() IAnyNameContext

	// IsDropUserContext differentiates from other interfaces.
	IsDropUserContext()
}

type DropUserContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	userName IAnyNameContext
}

func NewEmptyDropUserContext() *DropUserContext {
	var p = new(DropUserContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropUser
	return p
}

func InitEmptyDropUserContext(p *DropUserContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropUser
}

func (*DropUserContext) IsDropUserContext() {}

func NewDropUserContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropUserContext {
	var p = new(DropUserContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_dropUser

	return p
}

func (s *DropUserContext) GetParser() antlr.Parser { return s.parser }

func (s *DropUserContext) GetUserName() IAnyNameContext { return s.userName }

func (s *DropUserContext) SetUserName(v IAnyNameContext) { s.userName = v }

func (s *DropUserContext) DROP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDROP, 0)
}

func (s *DropUserContext) USER() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserUSER, 0)
}

func (s *DropUserContext) AnyName() IAnyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyNameContext)
}

func (s *DropUserContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropUserContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropUserContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterDropUser(s)
	}
}

func (s *DropUserContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitDropUser(s)
	}
}

func (p *InformixSQLParser) DropUser() (localctx IDropUserContext) {
	localctx = NewDropUserContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, InformixSQLParserRULE_dropUser)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Match(InformixSQLParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
		p.Match(InformixSQLParserUSER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)

		var _x = p.AnyName()

		localctx.(*DropUserContext).userName = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropViewContext is an interface to support dynamic dispatch.
type IDropViewContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetViewName returns the viewName rule contexts.
	GetViewName() IIdentifierContext

	// SetViewName sets the viewName rule contexts.
	SetViewName(IIdentifierContext)

	// Getter signatures
	DROP() antlr.TerminalNode
	VIEW() antlr.TerminalNode
	Identifier() IIdentifierContext
	IF() antlr.TerminalNode
	EXISTS() antlr.TerminalNode
	CASCADE() antlr.TerminalNode
	RESTRICT() antlr.TerminalNode

	// IsDropViewContext differentiates from other interfaces.
	IsDropViewContext()
}

type DropViewContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	viewName IIdentifierContext
}

func NewEmptyDropViewContext() *DropViewContext {
	var p = new(DropViewContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropView
	return p
}

func InitEmptyDropViewContext(p *DropViewContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropView
}

func (*DropViewContext) IsDropViewContext() {}

func NewDropViewContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropViewContext {
	var p = new(DropViewContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_dropView

	return p
}

func (s *DropViewContext) GetParser() antlr.Parser { return s.parser }

func (s *DropViewContext) GetViewName() IIdentifierContext { return s.viewName }

func (s *DropViewContext) SetViewName(v IIdentifierContext) { s.viewName = v }

func (s *DropViewContext) DROP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDROP, 0)
}

func (s *DropViewContext) VIEW() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserVIEW, 0)
}

func (s *DropViewContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DropViewContext) IF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIF, 0)
}

func (s *DropViewContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXISTS, 0)
}

func (s *DropViewContext) CASCADE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCASCADE, 0)
}

func (s *DropViewContext) RESTRICT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRESTRICT, 0)
}

func (s *DropViewContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropViewContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropViewContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterDropView(s)
	}
}

func (s *DropViewContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitDropView(s)
	}
}

func (p *InformixSQLParser) DropView() (localctx IDropViewContext) {
	localctx = NewDropViewContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, InformixSQLParserRULE_dropView)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(InformixSQLParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.Match(InformixSQLParserVIEW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(195)
			p.Match(InformixSQLParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.Match(InformixSQLParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(199)

		var _x = p.Identifier()

		localctx.(*DropViewContext).viewName = _x
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InformixSQLParserCASCADE || _la == InformixSQLParserRESTRICT {
		{
			p.SetState(200)
			_la = p.GetTokenStream().LA(1)

			if !(_la == InformixSQLParserCASCADE || _la == InformixSQLParserRESTRICT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropXadatasourceContext is an interface to support dynamic dispatch.
type IDropXadatasourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetXaSourceName returns the xaSourceName rule contexts.
	GetXaSourceName() IIdentifierContext

	// SetXaSourceName sets the xaSourceName rule contexts.
	SetXaSourceName(IIdentifierContext)

	// Getter signatures
	DROP() antlr.TerminalNode
	XADATASOURCE() antlr.TerminalNode
	RESTRICT() antlr.TerminalNode
	Identifier() IIdentifierContext
	IF() antlr.TerminalNode
	EXISTS() antlr.TerminalNode

	// IsDropXadatasourceContext differentiates from other interfaces.
	IsDropXadatasourceContext()
}

type DropXadatasourceContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	xaSourceName IIdentifierContext
}

func NewEmptyDropXadatasourceContext() *DropXadatasourceContext {
	var p = new(DropXadatasourceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropXadatasource
	return p
}

func InitEmptyDropXadatasourceContext(p *DropXadatasourceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropXadatasource
}

func (*DropXadatasourceContext) IsDropXadatasourceContext() {}

func NewDropXadatasourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropXadatasourceContext {
	var p = new(DropXadatasourceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_dropXadatasource

	return p
}

func (s *DropXadatasourceContext) GetParser() antlr.Parser { return s.parser }

func (s *DropXadatasourceContext) GetXaSourceName() IIdentifierContext { return s.xaSourceName }

func (s *DropXadatasourceContext) SetXaSourceName(v IIdentifierContext) { s.xaSourceName = v }

func (s *DropXadatasourceContext) DROP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDROP, 0)
}

func (s *DropXadatasourceContext) XADATASOURCE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserXADATASOURCE, 0)
}

func (s *DropXadatasourceContext) RESTRICT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRESTRICT, 0)
}

func (s *DropXadatasourceContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DropXadatasourceContext) IF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIF, 0)
}

func (s *DropXadatasourceContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXISTS, 0)
}

func (s *DropXadatasourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropXadatasourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropXadatasourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterDropXadatasource(s)
	}
}

func (s *DropXadatasourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitDropXadatasource(s)
	}
}

func (p *InformixSQLParser) DropXadatasource() (localctx IDropXadatasourceContext) {
	localctx = NewDropXadatasourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, InformixSQLParserRULE_dropXadatasource)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(InformixSQLParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.Match(InformixSQLParserXADATASOURCE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(205)
			p.Match(InformixSQLParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.Match(InformixSQLParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(209)

		var _x = p.Identifier()

		localctx.(*DropXadatasourceContext).xaSourceName = _x
	}
	{
		p.SetState(210)
		p.Match(InformixSQLParserRESTRICT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropXadataTypeSourceContext is an interface to support dynamic dispatch.
type IDropXadataTypeSourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetXaSourceName returns the xaSourceName rule contexts.
	GetXaSourceName() IIdentifierContext

	// SetXaSourceName sets the xaSourceName rule contexts.
	SetXaSourceName(IIdentifierContext)

	// Getter signatures
	DROP() antlr.TerminalNode
	XADATASOURCE() antlr.TerminalNode
	TYPE() antlr.TerminalNode
	RESTRICT() antlr.TerminalNode
	Identifier() IIdentifierContext
	IF() antlr.TerminalNode
	EXISTS() antlr.TerminalNode

	// IsDropXadataTypeSourceContext differentiates from other interfaces.
	IsDropXadataTypeSourceContext()
}

type DropXadataTypeSourceContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	xaSourceName IIdentifierContext
}

func NewEmptyDropXadataTypeSourceContext() *DropXadataTypeSourceContext {
	var p = new(DropXadataTypeSourceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropXadataTypeSource
	return p
}

func InitEmptyDropXadataTypeSourceContext(p *DropXadataTypeSourceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropXadataTypeSource
}

func (*DropXadataTypeSourceContext) IsDropXadataTypeSourceContext() {}

func NewDropXadataTypeSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropXadataTypeSourceContext {
	var p = new(DropXadataTypeSourceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_dropXadataTypeSource

	return p
}

func (s *DropXadataTypeSourceContext) GetParser() antlr.Parser { return s.parser }

func (s *DropXadataTypeSourceContext) GetXaSourceName() IIdentifierContext { return s.xaSourceName }

func (s *DropXadataTypeSourceContext) SetXaSourceName(v IIdentifierContext) { s.xaSourceName = v }

func (s *DropXadataTypeSourceContext) DROP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDROP, 0)
}

func (s *DropXadataTypeSourceContext) XADATASOURCE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserXADATASOURCE, 0)
}

func (s *DropXadataTypeSourceContext) TYPE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTYPE, 0)
}

func (s *DropXadataTypeSourceContext) RESTRICT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRESTRICT, 0)
}

func (s *DropXadataTypeSourceContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DropXadataTypeSourceContext) IF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIF, 0)
}

func (s *DropXadataTypeSourceContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXISTS, 0)
}

func (s *DropXadataTypeSourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropXadataTypeSourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropXadataTypeSourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterDropXadataTypeSource(s)
	}
}

func (s *DropXadataTypeSourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitDropXadataTypeSource(s)
	}
}

func (p *InformixSQLParser) DropXadataTypeSource() (localctx IDropXadataTypeSourceContext) {
	localctx = NewDropXadataTypeSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, InformixSQLParserRULE_dropXadataTypeSource)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(InformixSQLParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Match(InformixSQLParserXADATASOURCE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(InformixSQLParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(215)
			p.Match(InformixSQLParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Match(InformixSQLParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(219)

		var _x = p.Identifier()

		localctx.(*DropXadataTypeSourceContext).xaSourceName = _x
	}
	{
		p.SetState(220)
		p.Match(InformixSQLParserRESTRICT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropAccessMethodContext is an interface to support dynamic dispatch.
type IDropAccessMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAccessMethodName returns the accessMethodName rule contexts.
	GetAccessMethodName() IIdentifierContext

	// SetAccessMethodName sets the accessMethodName rule contexts.
	SetAccessMethodName(IIdentifierContext)

	// Getter signatures
	DROP() antlr.TerminalNode
	ACCESS_METHOD() antlr.TerminalNode
	RESTRICT() antlr.TerminalNode
	Identifier() IIdentifierContext
	IF() antlr.TerminalNode
	EXISTS() antlr.TerminalNode

	// IsDropAccessMethodContext differentiates from other interfaces.
	IsDropAccessMethodContext()
}

type DropAccessMethodContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	accessMethodName IIdentifierContext
}

func NewEmptyDropAccessMethodContext() *DropAccessMethodContext {
	var p = new(DropAccessMethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropAccessMethod
	return p
}

func InitEmptyDropAccessMethodContext(p *DropAccessMethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropAccessMethod
}

func (*DropAccessMethodContext) IsDropAccessMethodContext() {}

func NewDropAccessMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropAccessMethodContext {
	var p = new(DropAccessMethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_dropAccessMethod

	return p
}

func (s *DropAccessMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *DropAccessMethodContext) GetAccessMethodName() IIdentifierContext { return s.accessMethodName }

func (s *DropAccessMethodContext) SetAccessMethodName(v IIdentifierContext) { s.accessMethodName = v }

func (s *DropAccessMethodContext) DROP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDROP, 0)
}

func (s *DropAccessMethodContext) ACCESS_METHOD() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserACCESS_METHOD, 0)
}

func (s *DropAccessMethodContext) RESTRICT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRESTRICT, 0)
}

func (s *DropAccessMethodContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DropAccessMethodContext) IF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIF, 0)
}

func (s *DropAccessMethodContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXISTS, 0)
}

func (s *DropAccessMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropAccessMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropAccessMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterDropAccessMethod(s)
	}
}

func (s *DropAccessMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitDropAccessMethod(s)
	}
}

func (p *InformixSQLParser) DropAccessMethod() (localctx IDropAccessMethodContext) {
	localctx = NewDropAccessMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, InformixSQLParserRULE_dropAccessMethod)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(InformixSQLParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.Match(InformixSQLParserACCESS_METHOD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(224)
			p.Match(InformixSQLParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.Match(InformixSQLParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(228)

		var _x = p.Identifier()

		localctx.(*DropAccessMethodContext).accessMethodName = _x
	}
	{
		p.SetState(229)
		p.Match(InformixSQLParserRESTRICT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropAggregateContext is an interface to support dynamic dispatch.
type IDropAggregateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAggregateName returns the aggregateName rule contexts.
	GetAggregateName() IIdentifierContext

	// SetAggregateName sets the aggregateName rule contexts.
	SetAggregateName(IIdentifierContext)

	// Getter signatures
	DROP() antlr.TerminalNode
	AGGREGATE() antlr.TerminalNode
	Identifier() IIdentifierContext
	IF() antlr.TerminalNode
	EXISTS() antlr.TerminalNode

	// IsDropAggregateContext differentiates from other interfaces.
	IsDropAggregateContext()
}

type DropAggregateContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	aggregateName IIdentifierContext
}

func NewEmptyDropAggregateContext() *DropAggregateContext {
	var p = new(DropAggregateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropAggregate
	return p
}

func InitEmptyDropAggregateContext(p *DropAggregateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropAggregate
}

func (*DropAggregateContext) IsDropAggregateContext() {}

func NewDropAggregateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropAggregateContext {
	var p = new(DropAggregateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_dropAggregate

	return p
}

func (s *DropAggregateContext) GetParser() antlr.Parser { return s.parser }

func (s *DropAggregateContext) GetAggregateName() IIdentifierContext { return s.aggregateName }

func (s *DropAggregateContext) SetAggregateName(v IIdentifierContext) { s.aggregateName = v }

func (s *DropAggregateContext) DROP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDROP, 0)
}

func (s *DropAggregateContext) AGGREGATE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserAGGREGATE, 0)
}

func (s *DropAggregateContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DropAggregateContext) IF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIF, 0)
}

func (s *DropAggregateContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXISTS, 0)
}

func (s *DropAggregateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropAggregateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropAggregateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterDropAggregate(s)
	}
}

func (s *DropAggregateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitDropAggregate(s)
	}
}

func (p *InformixSQLParser) DropAggregate() (localctx IDropAggregateContext) {
	localctx = NewDropAggregateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, InformixSQLParserRULE_dropAggregate)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Match(InformixSQLParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Match(InformixSQLParserAGGREGATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(233)
			p.Match(InformixSQLParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Match(InformixSQLParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(237)

		var _x = p.Identifier()

		localctx.(*DropAggregateContext).aggregateName = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropDatabaseContext is an interface to support dynamic dispatch.
type IDropDatabaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDatabaseName returns the databaseName rule contexts.
	GetDatabaseName() IIdentifierContext

	// SetDatabaseName sets the databaseName rule contexts.
	SetDatabaseName(IIdentifierContext)

	// Getter signatures
	DROP() antlr.TerminalNode
	DATABASE() antlr.TerminalNode
	Identifier() IIdentifierContext
	IF() antlr.TerminalNode
	EXISTS() antlr.TerminalNode

	// IsDropDatabaseContext differentiates from other interfaces.
	IsDropDatabaseContext()
}

type DropDatabaseContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	databaseName IIdentifierContext
}

func NewEmptyDropDatabaseContext() *DropDatabaseContext {
	var p = new(DropDatabaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropDatabase
	return p
}

func InitEmptyDropDatabaseContext(p *DropDatabaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropDatabase
}

func (*DropDatabaseContext) IsDropDatabaseContext() {}

func NewDropDatabaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropDatabaseContext {
	var p = new(DropDatabaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_dropDatabase

	return p
}

func (s *DropDatabaseContext) GetParser() antlr.Parser { return s.parser }

func (s *DropDatabaseContext) GetDatabaseName() IIdentifierContext { return s.databaseName }

func (s *DropDatabaseContext) SetDatabaseName(v IIdentifierContext) { s.databaseName = v }

func (s *DropDatabaseContext) DROP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDROP, 0)
}

func (s *DropDatabaseContext) DATABASE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDATABASE, 0)
}

func (s *DropDatabaseContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DropDatabaseContext) IF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIF, 0)
}

func (s *DropDatabaseContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXISTS, 0)
}

func (s *DropDatabaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropDatabaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropDatabaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterDropDatabase(s)
	}
}

func (s *DropDatabaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitDropDatabase(s)
	}
}

func (p *InformixSQLParser) DropDatabase() (localctx IDropDatabaseContext) {
	localctx = NewDropDatabaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, InformixSQLParserRULE_dropDatabase)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(InformixSQLParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		p.Match(InformixSQLParserDATABASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(241)
			p.Match(InformixSQLParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(InformixSQLParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(245)

		var _x = p.Identifier()

		localctx.(*DropDatabaseContext).databaseName = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropIndexContext is an interface to support dynamic dispatch.
type IDropIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIndexName returns the indexName rule contexts.
	GetIndexName() IIdentifierContext

	// SetIndexName sets the indexName rule contexts.
	SetIndexName(IIdentifierContext)

	// Getter signatures
	DROP() antlr.TerminalNode
	INDEX() antlr.TerminalNode
	Identifier() IIdentifierContext
	IF() antlr.TerminalNode
	EXISTS() antlr.TerminalNode
	ONLINE() antlr.TerminalNode

	// IsDropIndexContext differentiates from other interfaces.
	IsDropIndexContext()
}

type DropIndexContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	indexName IIdentifierContext
}

func NewEmptyDropIndexContext() *DropIndexContext {
	var p = new(DropIndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropIndex
	return p
}

func InitEmptyDropIndexContext(p *DropIndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_dropIndex
}

func (*DropIndexContext) IsDropIndexContext() {}

func NewDropIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropIndexContext {
	var p = new(DropIndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_dropIndex

	return p
}

func (s *DropIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *DropIndexContext) GetIndexName() IIdentifierContext { return s.indexName }

func (s *DropIndexContext) SetIndexName(v IIdentifierContext) { s.indexName = v }

func (s *DropIndexContext) DROP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDROP, 0)
}

func (s *DropIndexContext) INDEX() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserINDEX, 0)
}

func (s *DropIndexContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DropIndexContext) IF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIF, 0)
}

func (s *DropIndexContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXISTS, 0)
}

func (s *DropIndexContext) ONLINE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserONLINE, 0)
}

func (s *DropIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterDropIndex(s)
	}
}

func (s *DropIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitDropIndex(s)
	}
}

func (p *InformixSQLParser) DropIndex() (localctx IDropIndexContext) {
	localctx = NewDropIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, InformixSQLParserRULE_dropIndex)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(InformixSQLParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.Match(InformixSQLParserINDEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(249)
			p.Match(InformixSQLParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.Match(InformixSQLParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(253)

		var _x = p.Identifier()

		localctx.(*DropIndexContext).indexName = _x
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InformixSQLParserONLINE {
		{
			p.SetState(254)
			p.Match(InformixSQLParserONLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICloseStmtContext is an interface to support dynamic dispatch.
type ICloseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCursorId returns the cursorId rule contexts.
	GetCursorId() IIdentifierContext

	// SetCursorId sets the cursorId rule contexts.
	SetCursorId(IIdentifierContext)

	// Getter signatures
	CLOSE() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsCloseStmtContext differentiates from other interfaces.
	IsCloseStmtContext()
}

type CloseStmtContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	cursorId IIdentifierContext
}

func NewEmptyCloseStmtContext() *CloseStmtContext {
	var p = new(CloseStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_closeStmt
	return p
}

func InitEmptyCloseStmtContext(p *CloseStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_closeStmt
}

func (*CloseStmtContext) IsCloseStmtContext() {}

func NewCloseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CloseStmtContext {
	var p = new(CloseStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_closeStmt

	return p
}

func (s *CloseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CloseStmtContext) GetCursorId() IIdentifierContext { return s.cursorId }

func (s *CloseStmtContext) SetCursorId(v IIdentifierContext) { s.cursorId = v }

func (s *CloseStmtContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCLOSE, 0)
}

func (s *CloseStmtContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *CloseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CloseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CloseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterCloseStmt(s)
	}
}

func (s *CloseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitCloseStmt(s)
	}
}

func (p *InformixSQLParser) CloseStmt() (localctx ICloseStmtContext) {
	localctx = NewCloseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, InformixSQLParserRULE_closeStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(InformixSQLParserCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)

		var _x = p.Identifier()

		localctx.(*CloseStmtContext).cursorId = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICloseDatabaseStmtContext is an interface to support dynamic dispatch.
type ICloseDatabaseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CLOSE() antlr.TerminalNode
	DATABASE() antlr.TerminalNode

	// IsCloseDatabaseStmtContext differentiates from other interfaces.
	IsCloseDatabaseStmtContext()
}

type CloseDatabaseStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCloseDatabaseStmtContext() *CloseDatabaseStmtContext {
	var p = new(CloseDatabaseStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_closeDatabaseStmt
	return p
}

func InitEmptyCloseDatabaseStmtContext(p *CloseDatabaseStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_closeDatabaseStmt
}

func (*CloseDatabaseStmtContext) IsCloseDatabaseStmtContext() {}

func NewCloseDatabaseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CloseDatabaseStmtContext {
	var p = new(CloseDatabaseStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_closeDatabaseStmt

	return p
}

func (s *CloseDatabaseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CloseDatabaseStmtContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCLOSE, 0)
}

func (s *CloseDatabaseStmtContext) DATABASE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDATABASE, 0)
}

func (s *CloseDatabaseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CloseDatabaseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CloseDatabaseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterCloseDatabaseStmt(s)
	}
}

func (s *CloseDatabaseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitCloseDatabaseStmt(s)
	}
}

func (p *InformixSQLParser) CloseDatabaseStmt() (localctx ICloseDatabaseStmtContext) {
	localctx = NewCloseDatabaseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, InformixSQLParserRULE_closeDatabaseStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.Match(InformixSQLParserCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Match(InformixSQLParserDATABASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDatabaseStmtContext is an interface to support dynamic dispatch.
type IDatabaseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDatabaseName returns the databaseName rule contexts.
	GetDatabaseName() IAnyNameContext

	// SetDatabaseName sets the databaseName rule contexts.
	SetDatabaseName(IAnyNameContext)

	// Getter signatures
	DATABASE() antlr.TerminalNode
	AnyName() IAnyNameContext
	EXCLUSIVE() antlr.TerminalNode

	// IsDatabaseStmtContext differentiates from other interfaces.
	IsDatabaseStmtContext()
}

type DatabaseStmtContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	databaseName IAnyNameContext
}

func NewEmptyDatabaseStmtContext() *DatabaseStmtContext {
	var p = new(DatabaseStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_databaseStmt
	return p
}

func InitEmptyDatabaseStmtContext(p *DatabaseStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_databaseStmt
}

func (*DatabaseStmtContext) IsDatabaseStmtContext() {}

func NewDatabaseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatabaseStmtContext {
	var p = new(DatabaseStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_databaseStmt

	return p
}

func (s *DatabaseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DatabaseStmtContext) GetDatabaseName() IAnyNameContext { return s.databaseName }

func (s *DatabaseStmtContext) SetDatabaseName(v IAnyNameContext) { s.databaseName = v }

func (s *DatabaseStmtContext) DATABASE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDATABASE, 0)
}

func (s *DatabaseStmtContext) AnyName() IAnyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyNameContext)
}

func (s *DatabaseStmtContext) EXCLUSIVE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXCLUSIVE, 0)
}

func (s *DatabaseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatabaseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatabaseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterDatabaseStmt(s)
	}
}

func (s *DatabaseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitDatabaseStmt(s)
	}
}

func (p *InformixSQLParser) DatabaseStmt() (localctx IDatabaseStmtContext) {
	localctx = NewDatabaseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, InformixSQLParserRULE_databaseStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Match(InformixSQLParserDATABASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)

		var _x = p.AnyName()

		localctx.(*DatabaseStmtContext).databaseName = _x
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InformixSQLParserEXCLUSIVE {
		{
			p.SetState(265)
			p.Match(InformixSQLParserEXCLUSIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommitWorkStmtContext is an interface to support dynamic dispatch.
type ICommitWorkStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMIT() antlr.TerminalNode
	WORK() antlr.TerminalNode

	// IsCommitWorkStmtContext differentiates from other interfaces.
	IsCommitWorkStmtContext()
}

type CommitWorkStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommitWorkStmtContext() *CommitWorkStmtContext {
	var p = new(CommitWorkStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_commitWorkStmt
	return p
}

func InitEmptyCommitWorkStmtContext(p *CommitWorkStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_commitWorkStmt
}

func (*CommitWorkStmtContext) IsCommitWorkStmtContext() {}

func NewCommitWorkStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommitWorkStmtContext {
	var p = new(CommitWorkStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_commitWorkStmt

	return p
}

func (s *CommitWorkStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CommitWorkStmtContext) COMMIT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCOMMIT, 0)
}

func (s *CommitWorkStmtContext) WORK() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserWORK, 0)
}

func (s *CommitWorkStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommitWorkStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommitWorkStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterCommitWorkStmt(s)
	}
}

func (s *CommitWorkStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitCommitWorkStmt(s)
	}
}

func (p *InformixSQLParser) CommitWorkStmt() (localctx ICommitWorkStmtContext) {
	localctx = NewCommitWorkStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, InformixSQLParserRULE_commitWorkStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(InformixSQLParserCOMMIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InformixSQLParserWORK {
		{
			p.SetState(269)
			p.Match(InformixSQLParserWORK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReleaseSavepointContext is an interface to support dynamic dispatch.
type IReleaseSavepointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSavepointName returns the savepointName rule contexts.
	GetSavepointName() IIdentifierContext

	// SetSavepointName sets the savepointName rule contexts.
	SetSavepointName(IIdentifierContext)

	// Getter signatures
	RELEASE() antlr.TerminalNode
	SAVEPOINT() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsReleaseSavepointContext differentiates from other interfaces.
	IsReleaseSavepointContext()
}

type ReleaseSavepointContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	savepointName IIdentifierContext
}

func NewEmptyReleaseSavepointContext() *ReleaseSavepointContext {
	var p = new(ReleaseSavepointContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_releaseSavepoint
	return p
}

func InitEmptyReleaseSavepointContext(p *ReleaseSavepointContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_releaseSavepoint
}

func (*ReleaseSavepointContext) IsReleaseSavepointContext() {}

func NewReleaseSavepointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReleaseSavepointContext {
	var p = new(ReleaseSavepointContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_releaseSavepoint

	return p
}

func (s *ReleaseSavepointContext) GetParser() antlr.Parser { return s.parser }

func (s *ReleaseSavepointContext) GetSavepointName() IIdentifierContext { return s.savepointName }

func (s *ReleaseSavepointContext) SetSavepointName(v IIdentifierContext) { s.savepointName = v }

func (s *ReleaseSavepointContext) RELEASE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRELEASE, 0)
}

func (s *ReleaseSavepointContext) SAVEPOINT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSAVEPOINT, 0)
}

func (s *ReleaseSavepointContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ReleaseSavepointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReleaseSavepointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReleaseSavepointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterReleaseSavepoint(s)
	}
}

func (s *ReleaseSavepointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitReleaseSavepoint(s)
	}
}

func (p *InformixSQLParser) ReleaseSavepoint() (localctx IReleaseSavepointContext) {
	localctx = NewReleaseSavepointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, InformixSQLParserRULE_releaseSavepoint)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(InformixSQLParserRELEASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		p.Match(InformixSQLParserSAVEPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)

		var _x = p.Identifier()

		localctx.(*ReleaseSavepointContext).savepointName = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRenameColumnContext is an interface to support dynamic dispatch.
type IRenameColumnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOldColumn returns the oldColumn rule contexts.
	GetOldColumn() IIdentifierContext

	// GetNewColumn returns the newColumn rule contexts.
	GetNewColumn() IIdentifierContext

	// SetOldColumn sets the oldColumn rule contexts.
	SetOldColumn(IIdentifierContext)

	// SetNewColumn sets the newColumn rule contexts.
	SetNewColumn(IIdentifierContext)

	// Getter signatures
	RENAME() antlr.TerminalNode
	COLUMN() antlr.TerminalNode
	TO() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext

	// IsRenameColumnContext differentiates from other interfaces.
	IsRenameColumnContext()
}

type RenameColumnContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	oldColumn IIdentifierContext
	newColumn IIdentifierContext
}

func NewEmptyRenameColumnContext() *RenameColumnContext {
	var p = new(RenameColumnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameColumn
	return p
}

func InitEmptyRenameColumnContext(p *RenameColumnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameColumn
}

func (*RenameColumnContext) IsRenameColumnContext() {}

func NewRenameColumnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RenameColumnContext {
	var p = new(RenameColumnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_renameColumn

	return p
}

func (s *RenameColumnContext) GetParser() antlr.Parser { return s.parser }

func (s *RenameColumnContext) GetOldColumn() IIdentifierContext { return s.oldColumn }

func (s *RenameColumnContext) GetNewColumn() IIdentifierContext { return s.newColumn }

func (s *RenameColumnContext) SetOldColumn(v IIdentifierContext) { s.oldColumn = v }

func (s *RenameColumnContext) SetNewColumn(v IIdentifierContext) { s.newColumn = v }

func (s *RenameColumnContext) RENAME() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRENAME, 0)
}

func (s *RenameColumnContext) COLUMN() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCOLUMN, 0)
}

func (s *RenameColumnContext) TO() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTO, 0)
}

func (s *RenameColumnContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *RenameColumnContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *RenameColumnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RenameColumnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RenameColumnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterRenameColumn(s)
	}
}

func (s *RenameColumnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitRenameColumn(s)
	}
}

func (p *InformixSQLParser) RenameColumn() (localctx IRenameColumnContext) {
	localctx = NewRenameColumnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, InformixSQLParserRULE_renameColumn)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(InformixSQLParserRENAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)
		p.Match(InformixSQLParserCOLUMN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)

		var _x = p.Identifier()

		localctx.(*RenameColumnContext).oldColumn = _x
	}
	{
		p.SetState(279)
		p.Match(InformixSQLParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)

		var _x = p.Identifier()

		localctx.(*RenameColumnContext).newColumn = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRenameConstraintContext is an interface to support dynamic dispatch.
type IRenameConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOldConstraint returns the oldConstraint rule contexts.
	GetOldConstraint() IIdentifierContext

	// GetNewConstraint returns the newConstraint rule contexts.
	GetNewConstraint() IIdentifierContext

	// SetOldConstraint sets the oldConstraint rule contexts.
	SetOldConstraint(IIdentifierContext)

	// SetNewConstraint sets the newConstraint rule contexts.
	SetNewConstraint(IIdentifierContext)

	// Getter signatures
	RENAME() antlr.TerminalNode
	CONSTRAINT() antlr.TerminalNode
	TO() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext

	// IsRenameConstraintContext differentiates from other interfaces.
	IsRenameConstraintContext()
}

type RenameConstraintContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	oldConstraint IIdentifierContext
	newConstraint IIdentifierContext
}

func NewEmptyRenameConstraintContext() *RenameConstraintContext {
	var p = new(RenameConstraintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameConstraint
	return p
}

func InitEmptyRenameConstraintContext(p *RenameConstraintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameConstraint
}

func (*RenameConstraintContext) IsRenameConstraintContext() {}

func NewRenameConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RenameConstraintContext {
	var p = new(RenameConstraintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_renameConstraint

	return p
}

func (s *RenameConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *RenameConstraintContext) GetOldConstraint() IIdentifierContext { return s.oldConstraint }

func (s *RenameConstraintContext) GetNewConstraint() IIdentifierContext { return s.newConstraint }

func (s *RenameConstraintContext) SetOldConstraint(v IIdentifierContext) { s.oldConstraint = v }

func (s *RenameConstraintContext) SetNewConstraint(v IIdentifierContext) { s.newConstraint = v }

func (s *RenameConstraintContext) RENAME() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRENAME, 0)
}

func (s *RenameConstraintContext) CONSTRAINT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCONSTRAINT, 0)
}

func (s *RenameConstraintContext) TO() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTO, 0)
}

func (s *RenameConstraintContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *RenameConstraintContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *RenameConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RenameConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RenameConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterRenameConstraint(s)
	}
}

func (s *RenameConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitRenameConstraint(s)
	}
}

func (p *InformixSQLParser) RenameConstraint() (localctx IRenameConstraintContext) {
	localctx = NewRenameConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, InformixSQLParserRULE_renameConstraint)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(InformixSQLParserRENAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(283)
		p.Match(InformixSQLParserCONSTRAINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)

		var _x = p.Identifier()

		localctx.(*RenameConstraintContext).oldConstraint = _x
	}
	{
		p.SetState(285)
		p.Match(InformixSQLParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)

		var _x = p.Identifier()

		localctx.(*RenameConstraintContext).newConstraint = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRenameDatabaseContext is an interface to support dynamic dispatch.
type IRenameDatabaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOldDatabase returns the oldDatabase rule contexts.
	GetOldDatabase() IIdentifierContext

	// GetNewDatabase returns the newDatabase rule contexts.
	GetNewDatabase() IIdentifierContext

	// SetOldDatabase sets the oldDatabase rule contexts.
	SetOldDatabase(IIdentifierContext)

	// SetNewDatabase sets the newDatabase rule contexts.
	SetNewDatabase(IIdentifierContext)

	// Getter signatures
	RENAME() antlr.TerminalNode
	DATABASE() antlr.TerminalNode
	TO() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext

	// IsRenameDatabaseContext differentiates from other interfaces.
	IsRenameDatabaseContext()
}

type RenameDatabaseContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	oldDatabase IIdentifierContext
	newDatabase IIdentifierContext
}

func NewEmptyRenameDatabaseContext() *RenameDatabaseContext {
	var p = new(RenameDatabaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameDatabase
	return p
}

func InitEmptyRenameDatabaseContext(p *RenameDatabaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameDatabase
}

func (*RenameDatabaseContext) IsRenameDatabaseContext() {}

func NewRenameDatabaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RenameDatabaseContext {
	var p = new(RenameDatabaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_renameDatabase

	return p
}

func (s *RenameDatabaseContext) GetParser() antlr.Parser { return s.parser }

func (s *RenameDatabaseContext) GetOldDatabase() IIdentifierContext { return s.oldDatabase }

func (s *RenameDatabaseContext) GetNewDatabase() IIdentifierContext { return s.newDatabase }

func (s *RenameDatabaseContext) SetOldDatabase(v IIdentifierContext) { s.oldDatabase = v }

func (s *RenameDatabaseContext) SetNewDatabase(v IIdentifierContext) { s.newDatabase = v }

func (s *RenameDatabaseContext) RENAME() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRENAME, 0)
}

func (s *RenameDatabaseContext) DATABASE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDATABASE, 0)
}

func (s *RenameDatabaseContext) TO() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTO, 0)
}

func (s *RenameDatabaseContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *RenameDatabaseContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *RenameDatabaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RenameDatabaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RenameDatabaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterRenameDatabase(s)
	}
}

func (s *RenameDatabaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitRenameDatabase(s)
	}
}

func (p *InformixSQLParser) RenameDatabase() (localctx IRenameDatabaseContext) {
	localctx = NewRenameDatabaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, InformixSQLParserRULE_renameDatabase)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.Match(InformixSQLParserRENAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(289)
		p.Match(InformixSQLParserDATABASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)

		var _x = p.Identifier()

		localctx.(*RenameDatabaseContext).oldDatabase = _x
	}
	{
		p.SetState(291)
		p.Match(InformixSQLParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)

		var _x = p.Identifier()

		localctx.(*RenameDatabaseContext).newDatabase = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRenameIndexContext is an interface to support dynamic dispatch.
type IRenameIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOldIndex returns the oldIndex rule contexts.
	GetOldIndex() IIdentifierContext

	// GetNewIndex returns the newIndex rule contexts.
	GetNewIndex() IIdentifierContext

	// SetOldIndex sets the oldIndex rule contexts.
	SetOldIndex(IIdentifierContext)

	// SetNewIndex sets the newIndex rule contexts.
	SetNewIndex(IIdentifierContext)

	// Getter signatures
	RENAME() antlr.TerminalNode
	INDEX() antlr.TerminalNode
	TO() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext

	// IsRenameIndexContext differentiates from other interfaces.
	IsRenameIndexContext()
}

type RenameIndexContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	oldIndex IIdentifierContext
	newIndex IIdentifierContext
}

func NewEmptyRenameIndexContext() *RenameIndexContext {
	var p = new(RenameIndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameIndex
	return p
}

func InitEmptyRenameIndexContext(p *RenameIndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameIndex
}

func (*RenameIndexContext) IsRenameIndexContext() {}

func NewRenameIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RenameIndexContext {
	var p = new(RenameIndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_renameIndex

	return p
}

func (s *RenameIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *RenameIndexContext) GetOldIndex() IIdentifierContext { return s.oldIndex }

func (s *RenameIndexContext) GetNewIndex() IIdentifierContext { return s.newIndex }

func (s *RenameIndexContext) SetOldIndex(v IIdentifierContext) { s.oldIndex = v }

func (s *RenameIndexContext) SetNewIndex(v IIdentifierContext) { s.newIndex = v }

func (s *RenameIndexContext) RENAME() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRENAME, 0)
}

func (s *RenameIndexContext) INDEX() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserINDEX, 0)
}

func (s *RenameIndexContext) TO() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTO, 0)
}

func (s *RenameIndexContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *RenameIndexContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *RenameIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RenameIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RenameIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterRenameIndex(s)
	}
}

func (s *RenameIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitRenameIndex(s)
	}
}

func (p *InformixSQLParser) RenameIndex() (localctx IRenameIndexContext) {
	localctx = NewRenameIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, InformixSQLParserRULE_renameIndex)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.Match(InformixSQLParserRENAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)
		p.Match(InformixSQLParserINDEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)

		var _x = p.Identifier()

		localctx.(*RenameIndexContext).oldIndex = _x
	}
	{
		p.SetState(297)
		p.Match(InformixSQLParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)

		var _x = p.Identifier()

		localctx.(*RenameIndexContext).newIndex = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRenameSecurityContext is an interface to support dynamic dispatch.
type IRenameSecurityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPolicy returns the policy rule contexts.
	GetPolicy() IIdentifierContext

	// GetOldSecurity returns the oldSecurity rule contexts.
	GetOldSecurity() IIdentifierContext

	// GetNewSecurity returns the newSecurity rule contexts.
	GetNewSecurity() IIdentifierContext

	// SetPolicy sets the policy rule contexts.
	SetPolicy(IIdentifierContext)

	// SetOldSecurity sets the oldSecurity rule contexts.
	SetOldSecurity(IIdentifierContext)

	// SetNewSecurity sets the newSecurity rule contexts.
	SetNewSecurity(IIdentifierContext)

	// Getter signatures
	RENAME() antlr.TerminalNode
	SECURITY() antlr.TerminalNode
	TO() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	POLICY() antlr.TerminalNode
	LABEL() antlr.TerminalNode
	COMPONENT() antlr.TerminalNode

	// IsRenameSecurityContext differentiates from other interfaces.
	IsRenameSecurityContext()
}

type RenameSecurityContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	policy      IIdentifierContext
	oldSecurity IIdentifierContext
	newSecurity IIdentifierContext
}

func NewEmptyRenameSecurityContext() *RenameSecurityContext {
	var p = new(RenameSecurityContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameSecurity
	return p
}

func InitEmptyRenameSecurityContext(p *RenameSecurityContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameSecurity
}

func (*RenameSecurityContext) IsRenameSecurityContext() {}

func NewRenameSecurityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RenameSecurityContext {
	var p = new(RenameSecurityContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_renameSecurity

	return p
}

func (s *RenameSecurityContext) GetParser() antlr.Parser { return s.parser }

func (s *RenameSecurityContext) GetPolicy() IIdentifierContext { return s.policy }

func (s *RenameSecurityContext) GetOldSecurity() IIdentifierContext { return s.oldSecurity }

func (s *RenameSecurityContext) GetNewSecurity() IIdentifierContext { return s.newSecurity }

func (s *RenameSecurityContext) SetPolicy(v IIdentifierContext) { s.policy = v }

func (s *RenameSecurityContext) SetOldSecurity(v IIdentifierContext) { s.oldSecurity = v }

func (s *RenameSecurityContext) SetNewSecurity(v IIdentifierContext) { s.newSecurity = v }

func (s *RenameSecurityContext) RENAME() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRENAME, 0)
}

func (s *RenameSecurityContext) SECURITY() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSECURITY, 0)
}

func (s *RenameSecurityContext) TO() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTO, 0)
}

func (s *RenameSecurityContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *RenameSecurityContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *RenameSecurityContext) POLICY() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserPOLICY, 0)
}

func (s *RenameSecurityContext) LABEL() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserLABEL, 0)
}

func (s *RenameSecurityContext) COMPONENT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCOMPONENT, 0)
}

func (s *RenameSecurityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RenameSecurityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RenameSecurityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterRenameSecurity(s)
	}
}

func (s *RenameSecurityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitRenameSecurity(s)
	}
}

func (p *InformixSQLParser) RenameSecurity() (localctx IRenameSecurityContext) {
	localctx = NewRenameSecurityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, InformixSQLParserRULE_renameSecurity)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(InformixSQLParserRENAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.Match(InformixSQLParserSECURITY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InformixSQLParserPOLICY:
		{
			p.SetState(302)
			p.Match(InformixSQLParserPOLICY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case InformixSQLParserLABEL:
		{
			p.SetState(303)
			p.Match(InformixSQLParserLABEL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
		case 1:
			p.SetState(305)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(304)

					var _x = p.Identifier()

					localctx.(*RenameSecurityContext).policy = _x
				}

			} else if p.HasError() { // JIM
				goto errorExit
			}

		case 2:
			{
				p.SetState(307)
				p.Match(InformixSQLParserCOMPONENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(312)

		var _x = p.Identifier()

		localctx.(*RenameSecurityContext).oldSecurity = _x
	}
	{
		p.SetState(313)
		p.Match(InformixSQLParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(314)

		var _x = p.Identifier()

		localctx.(*RenameSecurityContext).newSecurity = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRenameSequenceContext is an interface to support dynamic dispatch.
type IRenameSequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOldSequence returns the oldSequence rule contexts.
	GetOldSequence() IIdentifierContext

	// GetNewSequence returns the newSequence rule contexts.
	GetNewSequence() IIdentifierContext

	// SetOldSequence sets the oldSequence rule contexts.
	SetOldSequence(IIdentifierContext)

	// SetNewSequence sets the newSequence rule contexts.
	SetNewSequence(IIdentifierContext)

	// Getter signatures
	RENAME() antlr.TerminalNode
	SEQUENCE() antlr.TerminalNode
	TO() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext

	// IsRenameSequenceContext differentiates from other interfaces.
	IsRenameSequenceContext()
}

type RenameSequenceContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	oldSequence IIdentifierContext
	newSequence IIdentifierContext
}

func NewEmptyRenameSequenceContext() *RenameSequenceContext {
	var p = new(RenameSequenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameSequence
	return p
}

func InitEmptyRenameSequenceContext(p *RenameSequenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameSequence
}

func (*RenameSequenceContext) IsRenameSequenceContext() {}

func NewRenameSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RenameSequenceContext {
	var p = new(RenameSequenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_renameSequence

	return p
}

func (s *RenameSequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *RenameSequenceContext) GetOldSequence() IIdentifierContext { return s.oldSequence }

func (s *RenameSequenceContext) GetNewSequence() IIdentifierContext { return s.newSequence }

func (s *RenameSequenceContext) SetOldSequence(v IIdentifierContext) { s.oldSequence = v }

func (s *RenameSequenceContext) SetNewSequence(v IIdentifierContext) { s.newSequence = v }

func (s *RenameSequenceContext) RENAME() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRENAME, 0)
}

func (s *RenameSequenceContext) SEQUENCE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSEQUENCE, 0)
}

func (s *RenameSequenceContext) TO() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTO, 0)
}

func (s *RenameSequenceContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *RenameSequenceContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *RenameSequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RenameSequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RenameSequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterRenameSequence(s)
	}
}

func (s *RenameSequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitRenameSequence(s)
	}
}

func (p *InformixSQLParser) RenameSequence() (localctx IRenameSequenceContext) {
	localctx = NewRenameSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, InformixSQLParserRULE_renameSequence)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.Match(InformixSQLParserRENAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(317)
		p.Match(InformixSQLParserSEQUENCE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(318)

		var _x = p.Identifier()

		localctx.(*RenameSequenceContext).oldSequence = _x
	}
	{
		p.SetState(319)
		p.Match(InformixSQLParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(320)

		var _x = p.Identifier()

		localctx.(*RenameSequenceContext).newSequence = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRenameTableContext is an interface to support dynamic dispatch.
type IRenameTableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOldTableName returns the oldTableName rule contexts.
	GetOldTableName() IIdentifierContext

	// GetNewTableName returns the newTableName rule contexts.
	GetNewTableName() IIdentifierContext

	// SetOldTableName sets the oldTableName rule contexts.
	SetOldTableName(IIdentifierContext)

	// SetNewTableName sets the newTableName rule contexts.
	SetNewTableName(IIdentifierContext)

	// Getter signatures
	RENAME() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	TO() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext

	// IsRenameTableContext differentiates from other interfaces.
	IsRenameTableContext()
}

type RenameTableContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	oldTableName IIdentifierContext
	newTableName IIdentifierContext
}

func NewEmptyRenameTableContext() *RenameTableContext {
	var p = new(RenameTableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameTable
	return p
}

func InitEmptyRenameTableContext(p *RenameTableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameTable
}

func (*RenameTableContext) IsRenameTableContext() {}

func NewRenameTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RenameTableContext {
	var p = new(RenameTableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_renameTable

	return p
}

func (s *RenameTableContext) GetParser() antlr.Parser { return s.parser }

func (s *RenameTableContext) GetOldTableName() IIdentifierContext { return s.oldTableName }

func (s *RenameTableContext) GetNewTableName() IIdentifierContext { return s.newTableName }

func (s *RenameTableContext) SetOldTableName(v IIdentifierContext) { s.oldTableName = v }

func (s *RenameTableContext) SetNewTableName(v IIdentifierContext) { s.newTableName = v }

func (s *RenameTableContext) RENAME() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRENAME, 0)
}

func (s *RenameTableContext) TABLE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTABLE, 0)
}

func (s *RenameTableContext) TO() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTO, 0)
}

func (s *RenameTableContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *RenameTableContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *RenameTableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RenameTableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RenameTableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterRenameTable(s)
	}
}

func (s *RenameTableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitRenameTable(s)
	}
}

func (p *InformixSQLParser) RenameTable() (localctx IRenameTableContext) {
	localctx = NewRenameTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, InformixSQLParserRULE_renameTable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.Match(InformixSQLParserRENAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.Match(InformixSQLParserTABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(324)

		var _x = p.Identifier()

		localctx.(*RenameTableContext).oldTableName = _x
	}
	{
		p.SetState(325)
		p.Match(InformixSQLParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(326)

		var _x = p.Identifier()

		localctx.(*RenameTableContext).newTableName = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRenameTrustedContextContext is an interface to support dynamic dispatch.
type IRenameTrustedContextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOldTrustedContextName returns the oldTrustedContextName rule contexts.
	GetOldTrustedContextName() IIdentifierContext

	// GetNewTrustedContextName returns the newTrustedContextName rule contexts.
	GetNewTrustedContextName() IIdentifierContext

	// SetOldTrustedContextName sets the oldTrustedContextName rule contexts.
	SetOldTrustedContextName(IIdentifierContext)

	// SetNewTrustedContextName sets the newTrustedContextName rule contexts.
	SetNewTrustedContextName(IIdentifierContext)

	// Getter signatures
	RENAME() antlr.TerminalNode
	TRUSTED() antlr.TerminalNode
	CONTEXT() antlr.TerminalNode
	TO() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext

	// IsRenameTrustedContextContext differentiates from other interfaces.
	IsRenameTrustedContextContext()
}

type RenameTrustedContextContext struct {
	antlr.BaseParserRuleContext
	parser                antlr.Parser
	oldTrustedContextName IIdentifierContext
	newTrustedContextName IIdentifierContext
}

func NewEmptyRenameTrustedContextContext() *RenameTrustedContextContext {
	var p = new(RenameTrustedContextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameTrustedContext
	return p
}

func InitEmptyRenameTrustedContextContext(p *RenameTrustedContextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameTrustedContext
}

func (*RenameTrustedContextContext) IsRenameTrustedContextContext() {}

func NewRenameTrustedContextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RenameTrustedContextContext {
	var p = new(RenameTrustedContextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_renameTrustedContext

	return p
}

func (s *RenameTrustedContextContext) GetParser() antlr.Parser { return s.parser }

func (s *RenameTrustedContextContext) GetOldTrustedContextName() IIdentifierContext {
	return s.oldTrustedContextName
}

func (s *RenameTrustedContextContext) GetNewTrustedContextName() IIdentifierContext {
	return s.newTrustedContextName
}

func (s *RenameTrustedContextContext) SetOldTrustedContextName(v IIdentifierContext) {
	s.oldTrustedContextName = v
}

func (s *RenameTrustedContextContext) SetNewTrustedContextName(v IIdentifierContext) {
	s.newTrustedContextName = v
}

func (s *RenameTrustedContextContext) RENAME() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRENAME, 0)
}

func (s *RenameTrustedContextContext) TRUSTED() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTRUSTED, 0)
}

func (s *RenameTrustedContextContext) CONTEXT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCONTEXT, 0)
}

func (s *RenameTrustedContextContext) TO() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTO, 0)
}

func (s *RenameTrustedContextContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *RenameTrustedContextContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *RenameTrustedContextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RenameTrustedContextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RenameTrustedContextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterRenameTrustedContext(s)
	}
}

func (s *RenameTrustedContextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitRenameTrustedContext(s)
	}
}

func (p *InformixSQLParser) RenameTrustedContext() (localctx IRenameTrustedContextContext) {
	localctx = NewRenameTrustedContextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, InformixSQLParserRULE_renameTrustedContext)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(InformixSQLParserRENAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(329)
		p.Match(InformixSQLParserTRUSTED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(330)
		p.Match(InformixSQLParserCONTEXT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)

		var _x = p.Identifier()

		localctx.(*RenameTrustedContextContext).oldTrustedContextName = _x
	}
	{
		p.SetState(332)
		p.Match(InformixSQLParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(333)

		var _x = p.Identifier()

		localctx.(*RenameTrustedContextContext).newTrustedContextName = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRenameUserContext is an interface to support dynamic dispatch.
type IRenameUserContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOldUserName returns the oldUserName rule contexts.
	GetOldUserName() IIdentifierContext

	// GetNewUserName returns the newUserName rule contexts.
	GetNewUserName() IIdentifierContext

	// SetOldUserName sets the oldUserName rule contexts.
	SetOldUserName(IIdentifierContext)

	// SetNewUserName sets the newUserName rule contexts.
	SetNewUserName(IIdentifierContext)

	// Getter signatures
	RENAME() antlr.TerminalNode
	USER() antlr.TerminalNode
	TO() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext

	// IsRenameUserContext differentiates from other interfaces.
	IsRenameUserContext()
}

type RenameUserContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	oldUserName IIdentifierContext
	newUserName IIdentifierContext
}

func NewEmptyRenameUserContext() *RenameUserContext {
	var p = new(RenameUserContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameUser
	return p
}

func InitEmptyRenameUserContext(p *RenameUserContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_renameUser
}

func (*RenameUserContext) IsRenameUserContext() {}

func NewRenameUserContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RenameUserContext {
	var p = new(RenameUserContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_renameUser

	return p
}

func (s *RenameUserContext) GetParser() antlr.Parser { return s.parser }

func (s *RenameUserContext) GetOldUserName() IIdentifierContext { return s.oldUserName }

func (s *RenameUserContext) GetNewUserName() IIdentifierContext { return s.newUserName }

func (s *RenameUserContext) SetOldUserName(v IIdentifierContext) { s.oldUserName = v }

func (s *RenameUserContext) SetNewUserName(v IIdentifierContext) { s.newUserName = v }

func (s *RenameUserContext) RENAME() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRENAME, 0)
}

func (s *RenameUserContext) USER() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserUSER, 0)
}

func (s *RenameUserContext) TO() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTO, 0)
}

func (s *RenameUserContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *RenameUserContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *RenameUserContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RenameUserContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RenameUserContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterRenameUser(s)
	}
}

func (s *RenameUserContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitRenameUser(s)
	}
}

func (p *InformixSQLParser) RenameUser() (localctx IRenameUserContext) {
	localctx = NewRenameUserContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, InformixSQLParserRULE_renameUser)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Match(InformixSQLParserRENAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(336)
		p.Match(InformixSQLParserUSER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(337)

		var _x = p.Identifier()

		localctx.(*RenameUserContext).oldUserName = _x
	}
	{
		p.SetState(338)
		p.Match(InformixSQLParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(339)

		var _x = p.Identifier()

		localctx.(*RenameUserContext).newUserName = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRollbackWorkContext is an interface to support dynamic dispatch.
type IRollbackWorkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSavepoint returns the savepoint rule contexts.
	GetSavepoint() IIdentifierContext

	// SetSavepoint sets the savepoint rule contexts.
	SetSavepoint(IIdentifierContext)

	// Getter signatures
	ROLLBACK() antlr.TerminalNode
	WORK() antlr.TerminalNode
	TO() antlr.TerminalNode
	SAVEPOINT() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsRollbackWorkContext differentiates from other interfaces.
	IsRollbackWorkContext()
}

type RollbackWorkContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	savepoint IIdentifierContext
}

func NewEmptyRollbackWorkContext() *RollbackWorkContext {
	var p = new(RollbackWorkContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_rollbackWork
	return p
}

func InitEmptyRollbackWorkContext(p *RollbackWorkContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_rollbackWork
}

func (*RollbackWorkContext) IsRollbackWorkContext() {}

func NewRollbackWorkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RollbackWorkContext {
	var p = new(RollbackWorkContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_rollbackWork

	return p
}

func (s *RollbackWorkContext) GetParser() antlr.Parser { return s.parser }

func (s *RollbackWorkContext) GetSavepoint() IIdentifierContext { return s.savepoint }

func (s *RollbackWorkContext) SetSavepoint(v IIdentifierContext) { s.savepoint = v }

func (s *RollbackWorkContext) ROLLBACK() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserROLLBACK, 0)
}

func (s *RollbackWorkContext) WORK() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserWORK, 0)
}

func (s *RollbackWorkContext) TO() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTO, 0)
}

func (s *RollbackWorkContext) SAVEPOINT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSAVEPOINT, 0)
}

func (s *RollbackWorkContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *RollbackWorkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RollbackWorkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RollbackWorkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterRollbackWork(s)
	}
}

func (s *RollbackWorkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitRollbackWork(s)
	}
}

func (p *InformixSQLParser) RollbackWork() (localctx IRollbackWorkContext) {
	localctx = NewRollbackWorkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, InformixSQLParserRULE_rollbackWork)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.Match(InformixSQLParserROLLBACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InformixSQLParserWORK {
		{
			p.SetState(342)
			p.Match(InformixSQLParserWORK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InformixSQLParserTO {
		{
			p.SetState(345)
			p.Match(InformixSQLParserTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.Match(InformixSQLParserSAVEPOINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-48) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-4611686018427387905) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-6917529027641081857) != 0) {
			{
				p.SetState(347)

				var _x = p.Identifier()

				localctx.(*RollbackWorkContext).savepoint = _x
			}

		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISavepointStmtContext is an interface to support dynamic dispatch.
type ISavepointStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSavepoint returns the savepoint rule contexts.
	GetSavepoint() IIdentifierContext

	// SetSavepoint sets the savepoint rule contexts.
	SetSavepoint(IIdentifierContext)

	// Getter signatures
	SAVEPOINT() antlr.TerminalNode
	Identifier() IIdentifierContext
	UNIQUE() antlr.TerminalNode

	// IsSavepointStmtContext differentiates from other interfaces.
	IsSavepointStmtContext()
}

type SavepointStmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	savepoint IIdentifierContext
}

func NewEmptySavepointStmtContext() *SavepointStmtContext {
	var p = new(SavepointStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_savepointStmt
	return p
}

func InitEmptySavepointStmtContext(p *SavepointStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_savepointStmt
}

func (*SavepointStmtContext) IsSavepointStmtContext() {}

func NewSavepointStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SavepointStmtContext {
	var p = new(SavepointStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_savepointStmt

	return p
}

func (s *SavepointStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SavepointStmtContext) GetSavepoint() IIdentifierContext { return s.savepoint }

func (s *SavepointStmtContext) SetSavepoint(v IIdentifierContext) { s.savepoint = v }

func (s *SavepointStmtContext) SAVEPOINT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSAVEPOINT, 0)
}

func (s *SavepointStmtContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SavepointStmtContext) UNIQUE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserUNIQUE, 0)
}

func (s *SavepointStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SavepointStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SavepointStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterSavepointStmt(s)
	}
}

func (s *SavepointStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitSavepointStmt(s)
	}
}

func (p *InformixSQLParser) SavepointStmt() (localctx ISavepointStmtContext) {
	localctx = NewSavepointStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, InformixSQLParserRULE_savepointStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.Match(InformixSQLParserSAVEPOINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(353)

		var _x = p.Identifier()

		localctx.(*SavepointStmtContext).savepoint = _x
	}
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InformixSQLParserUNIQUE {
		{
			p.SetState(354)
			p.Match(InformixSQLParserUNIQUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISetAutofreeContext is an interface to support dynamic dispatch.
type ISetAutofreeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCursorId returns the cursorId rule contexts.
	GetCursorId() IIdentifierContext

	// GetCursorIdVar returns the cursorIdVar rule contexts.
	GetCursorIdVar() IAnyNameContext

	// SetCursorId sets the cursorId rule contexts.
	SetCursorId(IIdentifierContext)

	// SetCursorIdVar sets the cursorIdVar rule contexts.
	SetCursorIdVar(IAnyNameContext)

	// Getter signatures
	SET() antlr.TerminalNode
	AUTOFREE() antlr.TerminalNode
	FOR() antlr.TerminalNode
	ENABLED() antlr.TerminalNode
	DISABLED() antlr.TerminalNode
	Identifier() IIdentifierContext
	AnyName() IAnyNameContext

	// IsSetAutofreeContext differentiates from other interfaces.
	IsSetAutofreeContext()
}

type SetAutofreeContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	cursorId    IIdentifierContext
	cursorIdVar IAnyNameContext
}

func NewEmptySetAutofreeContext() *SetAutofreeContext {
	var p = new(SetAutofreeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_setAutofree
	return p
}

func InitEmptySetAutofreeContext(p *SetAutofreeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_setAutofree
}

func (*SetAutofreeContext) IsSetAutofreeContext() {}

func NewSetAutofreeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetAutofreeContext {
	var p = new(SetAutofreeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_setAutofree

	return p
}

func (s *SetAutofreeContext) GetParser() antlr.Parser { return s.parser }

func (s *SetAutofreeContext) GetCursorId() IIdentifierContext { return s.cursorId }

func (s *SetAutofreeContext) GetCursorIdVar() IAnyNameContext { return s.cursorIdVar }

func (s *SetAutofreeContext) SetCursorId(v IIdentifierContext) { s.cursorId = v }

func (s *SetAutofreeContext) SetCursorIdVar(v IAnyNameContext) { s.cursorIdVar = v }

func (s *SetAutofreeContext) SET() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSET, 0)
}

func (s *SetAutofreeContext) AUTOFREE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserAUTOFREE, 0)
}

func (s *SetAutofreeContext) FOR() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserFOR, 0)
}

func (s *SetAutofreeContext) ENABLED() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserENABLED, 0)
}

func (s *SetAutofreeContext) DISABLED() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDISABLED, 0)
}

func (s *SetAutofreeContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SetAutofreeContext) AnyName() IAnyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyNameContext)
}

func (s *SetAutofreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetAutofreeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetAutofreeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterSetAutofree(s)
	}
}

func (s *SetAutofreeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitSetAutofree(s)
	}
}

func (p *InformixSQLParser) SetAutofree() (localctx ISetAutofreeContext) {
	localctx = NewSetAutofreeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, InformixSQLParserRULE_setAutofree)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(InformixSQLParserSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(358)
		p.Match(InformixSQLParserAUTOFREE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InformixSQLParserDISABLED || _la == InformixSQLParserENABLED {
		{
			p.SetState(359)
			_la = p.GetTokenStream().LA(1)

			if !(_la == InformixSQLParserDISABLED || _la == InformixSQLParserENABLED) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InformixSQLParserFOR {
		{
			p.SetState(362)
			p.Match(InformixSQLParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(363)

				var _x = p.Identifier()

				localctx.(*SetAutofreeContext).cursorId = _x
			}

		case 2:
			{
				p.SetState(364)

				var _x = p.AnyName()

				localctx.(*SetAutofreeContext).cursorIdVar = _x
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISetCollationContext is an interface to support dynamic dispatch.
type ISetCollationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLocale returns the locale rule contexts.
	GetLocale() IQuotedStringContext

	// SetLocale sets the locale rule contexts.
	SetLocale(IQuotedStringContext)

	// Getter signatures
	SET() antlr.TerminalNode
	COLLATION() antlr.TerminalNode
	NO() antlr.TerminalNode
	QuotedString() IQuotedStringContext

	// IsSetCollationContext differentiates from other interfaces.
	IsSetCollationContext()
}

type SetCollationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	locale IQuotedStringContext
}

func NewEmptySetCollationContext() *SetCollationContext {
	var p = new(SetCollationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_setCollation
	return p
}

func InitEmptySetCollationContext(p *SetCollationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_setCollation
}

func (*SetCollationContext) IsSetCollationContext() {}

func NewSetCollationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetCollationContext {
	var p = new(SetCollationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_setCollation

	return p
}

func (s *SetCollationContext) GetParser() antlr.Parser { return s.parser }

func (s *SetCollationContext) GetLocale() IQuotedStringContext { return s.locale }

func (s *SetCollationContext) SetLocale(v IQuotedStringContext) { s.locale = v }

func (s *SetCollationContext) SET() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSET, 0)
}

func (s *SetCollationContext) COLLATION() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCOLLATION, 0)
}

func (s *SetCollationContext) NO() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserNO, 0)
}

func (s *SetCollationContext) QuotedString() IQuotedStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuotedStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuotedStringContext)
}

func (s *SetCollationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetCollationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetCollationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterSetCollation(s)
	}
}

func (s *SetCollationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitSetCollation(s)
	}
}

func (p *InformixSQLParser) SetCollation() (localctx ISetCollationContext) {
	localctx = NewSetCollationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, InformixSQLParserRULE_setCollation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.Match(InformixSQLParserSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InformixSQLParserCOLLATION:
		{
			p.SetState(370)
			p.Match(InformixSQLParserCOLLATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)

			var _x = p.QuotedString()

			localctx.(*SetCollationContext).locale = _x
		}

	case InformixSQLParserNO:
		{
			p.SetState(372)
			p.Match(InformixSQLParserNO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
			p.Match(InformixSQLParserCOLLATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISetDataskipContext is an interface to support dynamic dispatch.
type ISetDataskipContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SET() antlr.TerminalNode
	DATASKIP() antlr.TerminalNode
	ON() antlr.TerminalNode
	OFF() antlr.TerminalNode
	DEFAULT() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsSetDataskipContext differentiates from other interfaces.
	IsSetDataskipContext()
}

type SetDataskipContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetDataskipContext() *SetDataskipContext {
	var p = new(SetDataskipContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_setDataskip
	return p
}

func InitEmptySetDataskipContext(p *SetDataskipContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_setDataskip
}

func (*SetDataskipContext) IsSetDataskipContext() {}

func NewSetDataskipContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetDataskipContext {
	var p = new(SetDataskipContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_setDataskip

	return p
}

func (s *SetDataskipContext) GetParser() antlr.Parser { return s.parser }

func (s *SetDataskipContext) SET() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSET, 0)
}

func (s *SetDataskipContext) DATASKIP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDATASKIP, 0)
}

func (s *SetDataskipContext) ON() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserON, 0)
}

func (s *SetDataskipContext) OFF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserOFF, 0)
}

func (s *SetDataskipContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDEFAULT, 0)
}

func (s *SetDataskipContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *SetDataskipContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SetDataskipContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(InformixSQLParserCOMMA)
}

func (s *SetDataskipContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCOMMA, i)
}

func (s *SetDataskipContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetDataskipContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetDataskipContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterSetDataskip(s)
	}
}

func (s *SetDataskipContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitSetDataskip(s)
	}
}

func (p *InformixSQLParser) SetDataskip() (localctx ISetDataskipContext) {
	localctx = NewSetDataskipContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, InformixSQLParserRULE_setDataskip)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(InformixSQLParserSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)
		p.Match(InformixSQLParserDATASKIP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InformixSQLParserON:
		{
			p.SetState(378)
			p.Match(InformixSQLParserON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(387)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-48) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-4611686018427387905) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&-6917529027641081857) != 0) {
			{
				p.SetState(379)
				p.Identifier()
			}
			p.SetState(384)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == InformixSQLParserCOMMA {
				{
					p.SetState(380)
					p.Match(InformixSQLParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(381)
					p.Identifier()
				}

				p.SetState(386)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}

	case InformixSQLParserOFF:
		{
			p.SetState(389)
			p.Match(InformixSQLParserOFF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case InformixSQLParserDEFAULT:
		{
			p.SetState(390)
			p.Match(InformixSQLParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISetDebugFileContext is an interface to support dynamic dispatch.
type ISetDebugFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SET() antlr.TerminalNode
	DEBUG() antlr.TerminalNode
	FILE() antlr.TerminalNode
	TO() antlr.TerminalNode
	CHAR_STRING() antlr.TerminalNode
	Identifier() IIdentifierContext
	WITH() antlr.TerminalNode
	APPEND() antlr.TerminalNode

	// IsSetDebugFileContext differentiates from other interfaces.
	IsSetDebugFileContext()
}

type SetDebugFileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetDebugFileContext() *SetDebugFileContext {
	var p = new(SetDebugFileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_setDebugFile
	return p
}

func InitEmptySetDebugFileContext(p *SetDebugFileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_setDebugFile
}

func (*SetDebugFileContext) IsSetDebugFileContext() {}

func NewSetDebugFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetDebugFileContext {
	var p = new(SetDebugFileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_setDebugFile

	return p
}

func (s *SetDebugFileContext) GetParser() antlr.Parser { return s.parser }

func (s *SetDebugFileContext) SET() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSET, 0)
}

func (s *SetDebugFileContext) DEBUG() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDEBUG, 0)
}

func (s *SetDebugFileContext) FILE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserFILE, 0)
}

func (s *SetDebugFileContext) TO() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTO, 0)
}

func (s *SetDebugFileContext) CHAR_STRING() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCHAR_STRING, 0)
}

func (s *SetDebugFileContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SetDebugFileContext) WITH() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserWITH, 0)
}

func (s *SetDebugFileContext) APPEND() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserAPPEND, 0)
}

func (s *SetDebugFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetDebugFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetDebugFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterSetDebugFile(s)
	}
}

func (s *SetDebugFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitSetDebugFile(s)
	}
}

func (p *InformixSQLParser) SetDebugFile() (localctx ISetDebugFileContext) {
	localctx = NewSetDebugFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, InformixSQLParserRULE_setDebugFile)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)
		p.Match(InformixSQLParserSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(394)
		p.Match(InformixSQLParserDEBUG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(395)
		p.Match(InformixSQLParserFILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(396)
		p.Match(InformixSQLParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InformixSQLParserCHAR_STRING:
		{
			p.SetState(397)
			p.Match(InformixSQLParserCHAR_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case InformixSQLParserOPEN_PAR, InformixSQLParserABORT, InformixSQLParserACTION, InformixSQLParserACCESS_METHOD, InformixSQLParserADD, InformixSQLParserAFTER, InformixSQLParserAGGREGATE, InformixSQLParserALL, InformixSQLParserALTER, InformixSQLParserANALYZE, InformixSQLParserAND, InformixSQLParserAPPEND, InformixSQLParserAS, InformixSQLParserASC, InformixSQLParserATTACH, InformixSQLParserAUTOFREE, InformixSQLParserAUTOINCREMENT, InformixSQLParserBEFORE, InformixSQLParserBEGIN, InformixSQLParserBETWEEN, InformixSQLParserBY, InformixSQLParserCASCADE, InformixSQLParserCASE, InformixSQLParserCAST, InformixSQLParserCHECK, InformixSQLParserCLOSE, InformixSQLParserCOLLATE, InformixSQLParserCOLLATION, InformixSQLParserCOLUMN, InformixSQLParserCOMMIT, InformixSQLParserCOMPONENT, InformixSQLParserCONFLICT, InformixSQLParserCONSTRAINT, InformixSQLParserCONTEXT, InformixSQLParserCREATE, InformixSQLParserCROSS, InformixSQLParserCURRENT_DATE, InformixSQLParserCURRENT_TIME, InformixSQLParserCURRENT_TIMESTAMP, InformixSQLParserDATABASE, InformixSQLParserDATASKIP, InformixSQLParserDEFAULT, InformixSQLParserDEFERRABLE, InformixSQLParserDEFERRED, InformixSQLParserDEFERRED_PREPARE, InformixSQLParserDELETE, InformixSQLParserDEBUG, InformixSQLParserDESC, InformixSQLParserDETACH, InformixSQLParserDISABLED, InformixSQLParserDISTINCT, InformixSQLParserDROP, InformixSQLParserEACH, InformixSQLParserELSE, InformixSQLParserENABLED, InformixSQLParserEND, InformixSQLParserESCAPE, InformixSQLParserEXCEPT, InformixSQLParserEXCLUSIVE, InformixSQLParserEXISTS, InformixSQLParserEXPLAIN, InformixSQLParserFAIL, InformixSQLParserFOR, InformixSQLParserFOREIGN, InformixSQLParserFILE, InformixSQLParserFROM, InformixSQLParserFULL, InformixSQLParserGLOB, InformixSQLParserGROUP, InformixSQLParserHAVING, InformixSQLParserIF, InformixSQLParserIGNORE, InformixSQLParserIMMEDIATE, InformixSQLParserIN, InformixSQLParserINDEX, InformixSQLParserINDEXED, InformixSQLParserINITIALLY, InformixSQLParserINNER, InformixSQLParserINSERT, InformixSQLParserINSTEAD, InformixSQLParserINTERSECT, InformixSQLParserINTO, InformixSQLParserIS, InformixSQLParserISNULL, InformixSQLParserJOIN, InformixSQLParserKEY, InformixSQLParserLABEL, InformixSQLParserLEFT, InformixSQLParserLIKE, InformixSQLParserLIMIT, InformixSQLParserMATCH, InformixSQLParserNATURAL, InformixSQLParserNO, InformixSQLParserNOT, InformixSQLParserNOTNULL, InformixSQLParserNULL, InformixSQLParserOF, InformixSQLParserOFF, InformixSQLParserOFFSET, InformixSQLParserON, InformixSQLParserONLINE, InformixSQLParserOR, InformixSQLParserORDER, InformixSQLParserOUTER, InformixSQLParserPOLICY, InformixSQLParserPLAN, InformixSQLParserPRAGMA, InformixSQLParserPRIMARY, InformixSQLParserQUERY, InformixSQLParserRAISE, InformixSQLParserRECURSIVE, InformixSQLParserREFERENCES, InformixSQLParserREGEXP, InformixSQLParserREINDEX, InformixSQLParserRELEASE, InformixSQLParserRENAME, InformixSQLParserREPLACE, InformixSQLParserRESTRICT, InformixSQLParserRIGHT, InformixSQLParserROLLBACK, InformixSQLParserROW, InformixSQLParserROWS, InformixSQLParserSAVEPOINT, InformixSQLParserSECURITY, InformixSQLParserSELECT, InformixSQLParserSET, InformixSQLParserSEQUENCE, InformixSQLParserSYNONYM, InformixSQLParserTABLE, InformixSQLParserTEMP, InformixSQLParserTEMPORARY, InformixSQLParserTHEN, InformixSQLParserTO, InformixSQLParserTRANSACTION, InformixSQLParserTRIGGER, InformixSQLParserTRUSTED, InformixSQLParserTYPE, InformixSQLParserUNION, InformixSQLParserUNIQUE, InformixSQLParserUPDATE, InformixSQLParserUSER, InformixSQLParserUSING, InformixSQLParserVACUUM, InformixSQLParserVALUES, InformixSQLParserVIEW, InformixSQLParserVIRTUAL, InformixSQLParserWHEN, InformixSQLParserWHERE, InformixSQLParserWITH, InformixSQLParserWITHOUT, InformixSQLParserWORK, InformixSQLParserXADATASOURCE, InformixSQLParserFIRST_VALUE, InformixSQLParserOVER, InformixSQLParserPARTITION, InformixSQLParserRANGE, InformixSQLParserPRECEDING, InformixSQLParserUNBOUNDED, InformixSQLParserCURRENT, InformixSQLParserFOLLOWING, InformixSQLParserCUME_DIST, InformixSQLParserDENSE_RANK, InformixSQLParserLAG, InformixSQLParserLAST_VALUE, InformixSQLParserLEAD, InformixSQLParserNTH_VALUE, InformixSQLParserNTILE, InformixSQLParserPERCENT_RANK, InformixSQLParserRANK, InformixSQLParserROW_NUMBER, InformixSQLParserGENERATED, InformixSQLParserALWAYS, InformixSQLParserSTORED, InformixSQLParserTRUE, InformixSQLParserFALSE, InformixSQLParserWINDOW, InformixSQLParserNULLS, InformixSQLParserFIRST, InformixSQLParserLAST, InformixSQLParserFILTER, InformixSQLParserGROUPS, InformixSQLParserEXCLUDE, InformixSQLParserIDENTIFIER, InformixSQLParserSTRING_LITERAL:
		{
			p.SetState(398)
			p.Identifier()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InformixSQLParserWITH {
		{
			p.SetState(401)
			p.Match(InformixSQLParserWITH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)
			p.Match(InformixSQLParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISetDeferredPrepareStatementContext is an interface to support dynamic dispatch.
type ISetDeferredPrepareStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SET() antlr.TerminalNode
	DEFERRED_PREPARE() antlr.TerminalNode
	ENABLED() antlr.TerminalNode
	DISABLED() antlr.TerminalNode

	// IsSetDeferredPrepareStatementContext differentiates from other interfaces.
	IsSetDeferredPrepareStatementContext()
}

type SetDeferredPrepareStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetDeferredPrepareStatementContext() *SetDeferredPrepareStatementContext {
	var p = new(SetDeferredPrepareStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_setDeferredPrepareStatement
	return p
}

func InitEmptySetDeferredPrepareStatementContext(p *SetDeferredPrepareStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_setDeferredPrepareStatement
}

func (*SetDeferredPrepareStatementContext) IsSetDeferredPrepareStatementContext() {}

func NewSetDeferredPrepareStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetDeferredPrepareStatementContext {
	var p = new(SetDeferredPrepareStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_setDeferredPrepareStatement

	return p
}

func (s *SetDeferredPrepareStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SetDeferredPrepareStatementContext) SET() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSET, 0)
}

func (s *SetDeferredPrepareStatementContext) DEFERRED_PREPARE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDEFERRED_PREPARE, 0)
}

func (s *SetDeferredPrepareStatementContext) ENABLED() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserENABLED, 0)
}

func (s *SetDeferredPrepareStatementContext) DISABLED() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDISABLED, 0)
}

func (s *SetDeferredPrepareStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetDeferredPrepareStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetDeferredPrepareStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterSetDeferredPrepareStatement(s)
	}
}

func (s *SetDeferredPrepareStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitSetDeferredPrepareStatement(s)
	}
}

func (p *InformixSQLParser) SetDeferredPrepareStatement() (localctx ISetDeferredPrepareStatementContext) {
	localctx = NewSetDeferredPrepareStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, InformixSQLParserRULE_setDeferredPrepareStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(405)
		p.Match(InformixSQLParserSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(406)
		p.Match(InformixSQLParserDEFERRED_PREPARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == InformixSQLParserDISABLED || _la == InformixSQLParserENABLED {
		{
			p.SetState(407)
			_la = p.GetTokenStream().LA(1)

			if !(_la == InformixSQLParserDISABLED || _la == InformixSQLParserENABLED) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQuotedStringContext is an interface to support dynamic dispatch.
type IQuotedStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CHAR_STRING() antlr.TerminalNode

	// IsQuotedStringContext differentiates from other interfaces.
	IsQuotedStringContext()
}

type QuotedStringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedStringContext() *QuotedStringContext {
	var p = new(QuotedStringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_quotedString
	return p
}

func InitEmptyQuotedStringContext(p *QuotedStringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_quotedString
}

func (*QuotedStringContext) IsQuotedStringContext() {}

func NewQuotedStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedStringContext {
	var p = new(QuotedStringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_quotedString

	return p
}

func (s *QuotedStringContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedStringContext) CHAR_STRING() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCHAR_STRING, 0)
}

func (s *QuotedStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterQuotedString(s)
	}
}

func (s *QuotedStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitQuotedString(s)
	}
}

func (p *InformixSQLParser) QuotedString() (localctx IQuotedStringContext) {
	localctx = NewQuotedStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, InformixSQLParserRULE_quotedString)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(410)
		p.Match(InformixSQLParserCHAR_STRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnyNameContext is an interface to support dynamic dispatch.
type IAnyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Keyword() IKeywordContext
	STRING_LITERAL() antlr.TerminalNode
	OPEN_PAR() antlr.TerminalNode
	AnyName() IAnyNameContext
	CLOSE_PAR() antlr.TerminalNode

	// IsAnyNameContext differentiates from other interfaces.
	IsAnyNameContext()
}

type AnyNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnyNameContext() *AnyNameContext {
	var p = new(AnyNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_anyName
	return p
}

func InitEmptyAnyNameContext(p *AnyNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_anyName
}

func (*AnyNameContext) IsAnyNameContext() {}

func NewAnyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnyNameContext {
	var p = new(AnyNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_anyName

	return p
}

func (s *AnyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *AnyNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIDENTIFIER, 0)
}

func (s *AnyNameContext) Keyword() IKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *AnyNameContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSTRING_LITERAL, 0)
}

func (s *AnyNameContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserOPEN_PAR, 0)
}

func (s *AnyNameContext) AnyName() IAnyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyNameContext)
}

func (s *AnyNameContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCLOSE_PAR, 0)
}

func (s *AnyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnyNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterAnyName(s)
	}
}

func (s *AnyNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitAnyName(s)
	}
}

func (p *InformixSQLParser) AnyName() (localctx IAnyNameContext) {
	localctx = NewAnyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, InformixSQLParserRULE_anyName)
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case InformixSQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(412)
			p.Match(InformixSQLParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case InformixSQLParserABORT, InformixSQLParserACTION, InformixSQLParserACCESS_METHOD, InformixSQLParserADD, InformixSQLParserAFTER, InformixSQLParserAGGREGATE, InformixSQLParserALL, InformixSQLParserALTER, InformixSQLParserANALYZE, InformixSQLParserAND, InformixSQLParserAPPEND, InformixSQLParserAS, InformixSQLParserASC, InformixSQLParserATTACH, InformixSQLParserAUTOFREE, InformixSQLParserAUTOINCREMENT, InformixSQLParserBEFORE, InformixSQLParserBEGIN, InformixSQLParserBETWEEN, InformixSQLParserBY, InformixSQLParserCASCADE, InformixSQLParserCASE, InformixSQLParserCAST, InformixSQLParserCHECK, InformixSQLParserCLOSE, InformixSQLParserCOLLATE, InformixSQLParserCOLLATION, InformixSQLParserCOLUMN, InformixSQLParserCOMMIT, InformixSQLParserCOMPONENT, InformixSQLParserCONFLICT, InformixSQLParserCONSTRAINT, InformixSQLParserCONTEXT, InformixSQLParserCREATE, InformixSQLParserCROSS, InformixSQLParserCURRENT_DATE, InformixSQLParserCURRENT_TIME, InformixSQLParserCURRENT_TIMESTAMP, InformixSQLParserDATABASE, InformixSQLParserDATASKIP, InformixSQLParserDEFAULT, InformixSQLParserDEFERRABLE, InformixSQLParserDEFERRED, InformixSQLParserDEFERRED_PREPARE, InformixSQLParserDELETE, InformixSQLParserDEBUG, InformixSQLParserDESC, InformixSQLParserDETACH, InformixSQLParserDISABLED, InformixSQLParserDISTINCT, InformixSQLParserDROP, InformixSQLParserEACH, InformixSQLParserELSE, InformixSQLParserENABLED, InformixSQLParserEND, InformixSQLParserESCAPE, InformixSQLParserEXCEPT, InformixSQLParserEXCLUSIVE, InformixSQLParserEXISTS, InformixSQLParserEXPLAIN, InformixSQLParserFAIL, InformixSQLParserFOR, InformixSQLParserFOREIGN, InformixSQLParserFILE, InformixSQLParserFROM, InformixSQLParserFULL, InformixSQLParserGLOB, InformixSQLParserGROUP, InformixSQLParserHAVING, InformixSQLParserIF, InformixSQLParserIGNORE, InformixSQLParserIMMEDIATE, InformixSQLParserIN, InformixSQLParserINDEX, InformixSQLParserINDEXED, InformixSQLParserINITIALLY, InformixSQLParserINNER, InformixSQLParserINSERT, InformixSQLParserINSTEAD, InformixSQLParserINTERSECT, InformixSQLParserINTO, InformixSQLParserIS, InformixSQLParserISNULL, InformixSQLParserJOIN, InformixSQLParserKEY, InformixSQLParserLABEL, InformixSQLParserLEFT, InformixSQLParserLIKE, InformixSQLParserLIMIT, InformixSQLParserMATCH, InformixSQLParserNATURAL, InformixSQLParserNO, InformixSQLParserNOT, InformixSQLParserNOTNULL, InformixSQLParserNULL, InformixSQLParserOF, InformixSQLParserOFF, InformixSQLParserOFFSET, InformixSQLParserON, InformixSQLParserONLINE, InformixSQLParserOR, InformixSQLParserORDER, InformixSQLParserOUTER, InformixSQLParserPOLICY, InformixSQLParserPLAN, InformixSQLParserPRAGMA, InformixSQLParserPRIMARY, InformixSQLParserQUERY, InformixSQLParserRAISE, InformixSQLParserRECURSIVE, InformixSQLParserREFERENCES, InformixSQLParserREGEXP, InformixSQLParserREINDEX, InformixSQLParserRELEASE, InformixSQLParserRENAME, InformixSQLParserREPLACE, InformixSQLParserRESTRICT, InformixSQLParserRIGHT, InformixSQLParserROLLBACK, InformixSQLParserROW, InformixSQLParserROWS, InformixSQLParserSAVEPOINT, InformixSQLParserSECURITY, InformixSQLParserSELECT, InformixSQLParserSET, InformixSQLParserSEQUENCE, InformixSQLParserSYNONYM, InformixSQLParserTABLE, InformixSQLParserTEMP, InformixSQLParserTEMPORARY, InformixSQLParserTHEN, InformixSQLParserTO, InformixSQLParserTRANSACTION, InformixSQLParserTRIGGER, InformixSQLParserTRUSTED, InformixSQLParserTYPE, InformixSQLParserUNION, InformixSQLParserUNIQUE, InformixSQLParserUPDATE, InformixSQLParserUSER, InformixSQLParserUSING, InformixSQLParserVACUUM, InformixSQLParserVALUES, InformixSQLParserVIEW, InformixSQLParserVIRTUAL, InformixSQLParserWHEN, InformixSQLParserWHERE, InformixSQLParserWITH, InformixSQLParserWITHOUT, InformixSQLParserWORK, InformixSQLParserXADATASOURCE, InformixSQLParserFIRST_VALUE, InformixSQLParserOVER, InformixSQLParserPARTITION, InformixSQLParserRANGE, InformixSQLParserPRECEDING, InformixSQLParserUNBOUNDED, InformixSQLParserCURRENT, InformixSQLParserFOLLOWING, InformixSQLParserCUME_DIST, InformixSQLParserDENSE_RANK, InformixSQLParserLAG, InformixSQLParserLAST_VALUE, InformixSQLParserLEAD, InformixSQLParserNTH_VALUE, InformixSQLParserNTILE, InformixSQLParserPERCENT_RANK, InformixSQLParserRANK, InformixSQLParserROW_NUMBER, InformixSQLParserGENERATED, InformixSQLParserALWAYS, InformixSQLParserSTORED, InformixSQLParserTRUE, InformixSQLParserFALSE, InformixSQLParserWINDOW, InformixSQLParserNULLS, InformixSQLParserFIRST, InformixSQLParserLAST, InformixSQLParserFILTER, InformixSQLParserGROUPS, InformixSQLParserEXCLUDE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(413)
			p.Keyword()
		}

	case InformixSQLParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(414)
			p.Match(InformixSQLParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case InformixSQLParserOPEN_PAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(415)
			p.Match(InformixSQLParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.AnyName()
		}
		{
			p.SetState(417)
			p.Match(InformixSQLParserCLOSE_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAnyName() []IAnyNameContext
	AnyName(i int) IAnyNameContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) AllAnyName() []IAnyNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnyNameContext); ok {
			len++
		}
	}

	tst := make([]IAnyNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnyNameContext); ok {
			tst[i] = t.(IAnyNameContext)
			i++
		}
	}

	return tst
}

func (s *IdentifierContext) AnyName(i int) IAnyNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyNameContext)
}

func (s *IdentifierContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(InformixSQLParserDOT)
}

func (s *IdentifierContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDOT, i)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *InformixSQLParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, InformixSQLParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.AnyName()
	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == InformixSQLParserDOT {
		{
			p.SetState(422)
			p.Match(InformixSQLParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
			p.AnyName()
		}

		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ABORT() antlr.TerminalNode
	ACTION() antlr.TerminalNode
	ACCESS_METHOD() antlr.TerminalNode
	ADD() antlr.TerminalNode
	AFTER() antlr.TerminalNode
	AGGREGATE() antlr.TerminalNode
	ALL() antlr.TerminalNode
	ALTER() antlr.TerminalNode
	ANALYZE() antlr.TerminalNode
	AND() antlr.TerminalNode
	APPEND() antlr.TerminalNode
	AS() antlr.TerminalNode
	ASC() antlr.TerminalNode
	ATTACH() antlr.TerminalNode
	AUTOFREE() antlr.TerminalNode
	AUTOINCREMENT() antlr.TerminalNode
	BEFORE() antlr.TerminalNode
	BEGIN() antlr.TerminalNode
	BETWEEN() antlr.TerminalNode
	BY() antlr.TerminalNode
	CASCADE() antlr.TerminalNode
	CASE() antlr.TerminalNode
	CAST() antlr.TerminalNode
	CHECK() antlr.TerminalNode
	CLOSE() antlr.TerminalNode
	COLLATE() antlr.TerminalNode
	COLLATION() antlr.TerminalNode
	COLUMN() antlr.TerminalNode
	COMPONENT() antlr.TerminalNode
	COMMIT() antlr.TerminalNode
	CONFLICT() antlr.TerminalNode
	CONSTRAINT() antlr.TerminalNode
	CONTEXT() antlr.TerminalNode
	CREATE() antlr.TerminalNode
	CROSS() antlr.TerminalNode
	CURRENT_DATE() antlr.TerminalNode
	CURRENT_TIME() antlr.TerminalNode
	CURRENT_TIMESTAMP() antlr.TerminalNode
	DATABASE() antlr.TerminalNode
	DATASKIP() antlr.TerminalNode
	DEBUG() antlr.TerminalNode
	DEFAULT() antlr.TerminalNode
	DEFERRABLE() antlr.TerminalNode
	DEFERRED() antlr.TerminalNode
	DEFERRED_PREPARE() antlr.TerminalNode
	DELETE() antlr.TerminalNode
	DESC() antlr.TerminalNode
	DETACH() antlr.TerminalNode
	DISABLED() antlr.TerminalNode
	DISTINCT() antlr.TerminalNode
	DROP() antlr.TerminalNode
	EACH() antlr.TerminalNode
	ELSE() antlr.TerminalNode
	ENABLED() antlr.TerminalNode
	END() antlr.TerminalNode
	ESCAPE() antlr.TerminalNode
	EXCEPT() antlr.TerminalNode
	EXCLUSIVE() antlr.TerminalNode
	EXISTS() antlr.TerminalNode
	EXPLAIN() antlr.TerminalNode
	FAIL() antlr.TerminalNode
	FOR() antlr.TerminalNode
	FOREIGN() antlr.TerminalNode
	FROM() antlr.TerminalNode
	FILE() antlr.TerminalNode
	FULL() antlr.TerminalNode
	GLOB() antlr.TerminalNode
	GROUP() antlr.TerminalNode
	HAVING() antlr.TerminalNode
	IF() antlr.TerminalNode
	IGNORE() antlr.TerminalNode
	IMMEDIATE() antlr.TerminalNode
	IN() antlr.TerminalNode
	INDEX() antlr.TerminalNode
	INDEXED() antlr.TerminalNode
	INITIALLY() antlr.TerminalNode
	INNER() antlr.TerminalNode
	INSERT() antlr.TerminalNode
	INSTEAD() antlr.TerminalNode
	INTERSECT() antlr.TerminalNode
	INTO() antlr.TerminalNode
	IS() antlr.TerminalNode
	ISNULL() antlr.TerminalNode
	JOIN() antlr.TerminalNode
	KEY() antlr.TerminalNode
	LABEL() antlr.TerminalNode
	LEFT() antlr.TerminalNode
	LIKE() antlr.TerminalNode
	LIMIT() antlr.TerminalNode
	MATCH() antlr.TerminalNode
	NATURAL() antlr.TerminalNode
	NO() antlr.TerminalNode
	NOT() antlr.TerminalNode
	NOTNULL() antlr.TerminalNode
	NULL() antlr.TerminalNode
	OF() antlr.TerminalNode
	OFF() antlr.TerminalNode
	OFFSET() antlr.TerminalNode
	ON() antlr.TerminalNode
	ONLINE() antlr.TerminalNode
	OR() antlr.TerminalNode
	ORDER() antlr.TerminalNode
	OUTER() antlr.TerminalNode
	POLICY() antlr.TerminalNode
	PLAN() antlr.TerminalNode
	PRAGMA() antlr.TerminalNode
	PRIMARY() antlr.TerminalNode
	QUERY() antlr.TerminalNode
	RAISE() antlr.TerminalNode
	RECURSIVE() antlr.TerminalNode
	REFERENCES() antlr.TerminalNode
	REGEXP() antlr.TerminalNode
	REINDEX() antlr.TerminalNode
	RELEASE() antlr.TerminalNode
	RENAME() antlr.TerminalNode
	REPLACE() antlr.TerminalNode
	RESTRICT() antlr.TerminalNode
	RIGHT() antlr.TerminalNode
	ROLLBACK() antlr.TerminalNode
	ROW() antlr.TerminalNode
	ROWS() antlr.TerminalNode
	SAVEPOINT() antlr.TerminalNode
	SECURITY() antlr.TerminalNode
	SELECT() antlr.TerminalNode
	SET() antlr.TerminalNode
	SEQUENCE() antlr.TerminalNode
	SYNONYM() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	TEMP() antlr.TerminalNode
	TEMPORARY() antlr.TerminalNode
	THEN() antlr.TerminalNode
	TO() antlr.TerminalNode
	TRANSACTION() antlr.TerminalNode
	TRIGGER() antlr.TerminalNode
	TRUSTED() antlr.TerminalNode
	TYPE() antlr.TerminalNode
	UNION() antlr.TerminalNode
	UNIQUE() antlr.TerminalNode
	UPDATE() antlr.TerminalNode
	USER() antlr.TerminalNode
	USING() antlr.TerminalNode
	VACUUM() antlr.TerminalNode
	VALUES() antlr.TerminalNode
	VIEW() antlr.TerminalNode
	VIRTUAL() antlr.TerminalNode
	WHEN() antlr.TerminalNode
	WHERE() antlr.TerminalNode
	WITH() antlr.TerminalNode
	WITHOUT() antlr.TerminalNode
	WORK() antlr.TerminalNode
	XADATASOURCE() antlr.TerminalNode
	FIRST_VALUE() antlr.TerminalNode
	OVER() antlr.TerminalNode
	PARTITION() antlr.TerminalNode
	RANGE() antlr.TerminalNode
	PRECEDING() antlr.TerminalNode
	UNBOUNDED() antlr.TerminalNode
	CURRENT() antlr.TerminalNode
	FOLLOWING() antlr.TerminalNode
	CUME_DIST() antlr.TerminalNode
	DENSE_RANK() antlr.TerminalNode
	LAG() antlr.TerminalNode
	LAST_VALUE() antlr.TerminalNode
	LEAD() antlr.TerminalNode
	NTH_VALUE() antlr.TerminalNode
	NTILE() antlr.TerminalNode
	PERCENT_RANK() antlr.TerminalNode
	RANK() antlr.TerminalNode
	ROW_NUMBER() antlr.TerminalNode
	GENERATED() antlr.TerminalNode
	ALWAYS() antlr.TerminalNode
	STORED() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	WINDOW() antlr.TerminalNode
	NULLS() antlr.TerminalNode
	FIRST() antlr.TerminalNode
	LAST() antlr.TerminalNode
	FILTER() antlr.TerminalNode
	GROUPS() antlr.TerminalNode
	EXCLUDE() antlr.TerminalNode

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_keyword
	return p
}

func InitEmptyKeywordContext(p *KeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = InformixSQLParserRULE_keyword
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = InformixSQLParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) ABORT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserABORT, 0)
}

func (s *KeywordContext) ACTION() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserACTION, 0)
}

func (s *KeywordContext) ACCESS_METHOD() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserACCESS_METHOD, 0)
}

func (s *KeywordContext) ADD() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserADD, 0)
}

func (s *KeywordContext) AFTER() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserAFTER, 0)
}

func (s *KeywordContext) AGGREGATE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserAGGREGATE, 0)
}

func (s *KeywordContext) ALL() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserALL, 0)
}

func (s *KeywordContext) ALTER() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserALTER, 0)
}

func (s *KeywordContext) ANALYZE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserANALYZE, 0)
}

func (s *KeywordContext) AND() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserAND, 0)
}

func (s *KeywordContext) APPEND() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserAPPEND, 0)
}

func (s *KeywordContext) AS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserAS, 0)
}

func (s *KeywordContext) ASC() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserASC, 0)
}

func (s *KeywordContext) ATTACH() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserATTACH, 0)
}

func (s *KeywordContext) AUTOFREE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserAUTOFREE, 0)
}

func (s *KeywordContext) AUTOINCREMENT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserAUTOINCREMENT, 0)
}

func (s *KeywordContext) BEFORE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserBEFORE, 0)
}

func (s *KeywordContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserBEGIN, 0)
}

func (s *KeywordContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserBETWEEN, 0)
}

func (s *KeywordContext) BY() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserBY, 0)
}

func (s *KeywordContext) CASCADE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCASCADE, 0)
}

func (s *KeywordContext) CASE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCASE, 0)
}

func (s *KeywordContext) CAST() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCAST, 0)
}

func (s *KeywordContext) CHECK() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCHECK, 0)
}

func (s *KeywordContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCLOSE, 0)
}

func (s *KeywordContext) COLLATE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCOLLATE, 0)
}

func (s *KeywordContext) COLLATION() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCOLLATION, 0)
}

func (s *KeywordContext) COLUMN() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCOLUMN, 0)
}

func (s *KeywordContext) COMPONENT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCOMPONENT, 0)
}

func (s *KeywordContext) COMMIT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCOMMIT, 0)
}

func (s *KeywordContext) CONFLICT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCONFLICT, 0)
}

func (s *KeywordContext) CONSTRAINT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCONSTRAINT, 0)
}

func (s *KeywordContext) CONTEXT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCONTEXT, 0)
}

func (s *KeywordContext) CREATE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCREATE, 0)
}

func (s *KeywordContext) CROSS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCROSS, 0)
}

func (s *KeywordContext) CURRENT_DATE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCURRENT_DATE, 0)
}

func (s *KeywordContext) CURRENT_TIME() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCURRENT_TIME, 0)
}

func (s *KeywordContext) CURRENT_TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCURRENT_TIMESTAMP, 0)
}

func (s *KeywordContext) DATABASE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDATABASE, 0)
}

func (s *KeywordContext) DATASKIP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDATASKIP, 0)
}

func (s *KeywordContext) DEBUG() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDEBUG, 0)
}

func (s *KeywordContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDEFAULT, 0)
}

func (s *KeywordContext) DEFERRABLE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDEFERRABLE, 0)
}

func (s *KeywordContext) DEFERRED() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDEFERRED, 0)
}

func (s *KeywordContext) DEFERRED_PREPARE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDEFERRED_PREPARE, 0)
}

func (s *KeywordContext) DELETE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDELETE, 0)
}

func (s *KeywordContext) DESC() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDESC, 0)
}

func (s *KeywordContext) DETACH() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDETACH, 0)
}

func (s *KeywordContext) DISABLED() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDISABLED, 0)
}

func (s *KeywordContext) DISTINCT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDISTINCT, 0)
}

func (s *KeywordContext) DROP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDROP, 0)
}

func (s *KeywordContext) EACH() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEACH, 0)
}

func (s *KeywordContext) ELSE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserELSE, 0)
}

func (s *KeywordContext) ENABLED() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserENABLED, 0)
}

func (s *KeywordContext) END() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEND, 0)
}

func (s *KeywordContext) ESCAPE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserESCAPE, 0)
}

func (s *KeywordContext) EXCEPT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXCEPT, 0)
}

func (s *KeywordContext) EXCLUSIVE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXCLUSIVE, 0)
}

func (s *KeywordContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXISTS, 0)
}

func (s *KeywordContext) EXPLAIN() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXPLAIN, 0)
}

func (s *KeywordContext) FAIL() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserFAIL, 0)
}

func (s *KeywordContext) FOR() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserFOR, 0)
}

func (s *KeywordContext) FOREIGN() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserFOREIGN, 0)
}

func (s *KeywordContext) FROM() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserFROM, 0)
}

func (s *KeywordContext) FILE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserFILE, 0)
}

func (s *KeywordContext) FULL() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserFULL, 0)
}

func (s *KeywordContext) GLOB() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserGLOB, 0)
}

func (s *KeywordContext) GROUP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserGROUP, 0)
}

func (s *KeywordContext) HAVING() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserHAVING, 0)
}

func (s *KeywordContext) IF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIF, 0)
}

func (s *KeywordContext) IGNORE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIGNORE, 0)
}

func (s *KeywordContext) IMMEDIATE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIMMEDIATE, 0)
}

func (s *KeywordContext) IN() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIN, 0)
}

func (s *KeywordContext) INDEX() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserINDEX, 0)
}

func (s *KeywordContext) INDEXED() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserINDEXED, 0)
}

func (s *KeywordContext) INITIALLY() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserINITIALLY, 0)
}

func (s *KeywordContext) INNER() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserINNER, 0)
}

func (s *KeywordContext) INSERT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserINSERT, 0)
}

func (s *KeywordContext) INSTEAD() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserINSTEAD, 0)
}

func (s *KeywordContext) INTERSECT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserINTERSECT, 0)
}

func (s *KeywordContext) INTO() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserINTO, 0)
}

func (s *KeywordContext) IS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserIS, 0)
}

func (s *KeywordContext) ISNULL() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserISNULL, 0)
}

func (s *KeywordContext) JOIN() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserJOIN, 0)
}

func (s *KeywordContext) KEY() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserKEY, 0)
}

func (s *KeywordContext) LABEL() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserLABEL, 0)
}

func (s *KeywordContext) LEFT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserLEFT, 0)
}

func (s *KeywordContext) LIKE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserLIKE, 0)
}

func (s *KeywordContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserLIMIT, 0)
}

func (s *KeywordContext) MATCH() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserMATCH, 0)
}

func (s *KeywordContext) NATURAL() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserNATURAL, 0)
}

func (s *KeywordContext) NO() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserNO, 0)
}

func (s *KeywordContext) NOT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserNOT, 0)
}

func (s *KeywordContext) NOTNULL() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserNOTNULL, 0)
}

func (s *KeywordContext) NULL() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserNULL, 0)
}

func (s *KeywordContext) OF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserOF, 0)
}

func (s *KeywordContext) OFF() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserOFF, 0)
}

func (s *KeywordContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserOFFSET, 0)
}

func (s *KeywordContext) ON() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserON, 0)
}

func (s *KeywordContext) ONLINE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserONLINE, 0)
}

func (s *KeywordContext) OR() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserOR, 0)
}

func (s *KeywordContext) ORDER() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserORDER, 0)
}

func (s *KeywordContext) OUTER() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserOUTER, 0)
}

func (s *KeywordContext) POLICY() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserPOLICY, 0)
}

func (s *KeywordContext) PLAN() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserPLAN, 0)
}

func (s *KeywordContext) PRAGMA() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserPRAGMA, 0)
}

func (s *KeywordContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserPRIMARY, 0)
}

func (s *KeywordContext) QUERY() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserQUERY, 0)
}

func (s *KeywordContext) RAISE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRAISE, 0)
}

func (s *KeywordContext) RECURSIVE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRECURSIVE, 0)
}

func (s *KeywordContext) REFERENCES() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserREFERENCES, 0)
}

func (s *KeywordContext) REGEXP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserREGEXP, 0)
}

func (s *KeywordContext) REINDEX() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserREINDEX, 0)
}

func (s *KeywordContext) RELEASE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRELEASE, 0)
}

func (s *KeywordContext) RENAME() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRENAME, 0)
}

func (s *KeywordContext) REPLACE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserREPLACE, 0)
}

func (s *KeywordContext) RESTRICT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRESTRICT, 0)
}

func (s *KeywordContext) RIGHT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRIGHT, 0)
}

func (s *KeywordContext) ROLLBACK() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserROLLBACK, 0)
}

func (s *KeywordContext) ROW() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserROW, 0)
}

func (s *KeywordContext) ROWS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserROWS, 0)
}

func (s *KeywordContext) SAVEPOINT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSAVEPOINT, 0)
}

func (s *KeywordContext) SECURITY() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSECURITY, 0)
}

func (s *KeywordContext) SELECT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSELECT, 0)
}

func (s *KeywordContext) SET() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSET, 0)
}

func (s *KeywordContext) SEQUENCE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSEQUENCE, 0)
}

func (s *KeywordContext) SYNONYM() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSYNONYM, 0)
}

func (s *KeywordContext) TABLE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTABLE, 0)
}

func (s *KeywordContext) TEMP() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTEMP, 0)
}

func (s *KeywordContext) TEMPORARY() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTEMPORARY, 0)
}

func (s *KeywordContext) THEN() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTHEN, 0)
}

func (s *KeywordContext) TO() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTO, 0)
}

func (s *KeywordContext) TRANSACTION() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTRANSACTION, 0)
}

func (s *KeywordContext) TRIGGER() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTRIGGER, 0)
}

func (s *KeywordContext) TRUSTED() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTRUSTED, 0)
}

func (s *KeywordContext) TYPE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTYPE, 0)
}

func (s *KeywordContext) UNION() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserUNION, 0)
}

func (s *KeywordContext) UNIQUE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserUNIQUE, 0)
}

func (s *KeywordContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserUPDATE, 0)
}

func (s *KeywordContext) USER() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserUSER, 0)
}

func (s *KeywordContext) USING() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserUSING, 0)
}

func (s *KeywordContext) VACUUM() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserVACUUM, 0)
}

func (s *KeywordContext) VALUES() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserVALUES, 0)
}

func (s *KeywordContext) VIEW() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserVIEW, 0)
}

func (s *KeywordContext) VIRTUAL() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserVIRTUAL, 0)
}

func (s *KeywordContext) WHEN() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserWHEN, 0)
}

func (s *KeywordContext) WHERE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserWHERE, 0)
}

func (s *KeywordContext) WITH() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserWITH, 0)
}

func (s *KeywordContext) WITHOUT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserWITHOUT, 0)
}

func (s *KeywordContext) WORK() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserWORK, 0)
}

func (s *KeywordContext) XADATASOURCE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserXADATASOURCE, 0)
}

func (s *KeywordContext) FIRST_VALUE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserFIRST_VALUE, 0)
}

func (s *KeywordContext) OVER() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserOVER, 0)
}

func (s *KeywordContext) PARTITION() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserPARTITION, 0)
}

func (s *KeywordContext) RANGE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRANGE, 0)
}

func (s *KeywordContext) PRECEDING() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserPRECEDING, 0)
}

func (s *KeywordContext) UNBOUNDED() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserUNBOUNDED, 0)
}

func (s *KeywordContext) CURRENT() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCURRENT, 0)
}

func (s *KeywordContext) FOLLOWING() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserFOLLOWING, 0)
}

func (s *KeywordContext) CUME_DIST() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserCUME_DIST, 0)
}

func (s *KeywordContext) DENSE_RANK() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserDENSE_RANK, 0)
}

func (s *KeywordContext) LAG() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserLAG, 0)
}

func (s *KeywordContext) LAST_VALUE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserLAST_VALUE, 0)
}

func (s *KeywordContext) LEAD() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserLEAD, 0)
}

func (s *KeywordContext) NTH_VALUE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserNTH_VALUE, 0)
}

func (s *KeywordContext) NTILE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserNTILE, 0)
}

func (s *KeywordContext) PERCENT_RANK() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserPERCENT_RANK, 0)
}

func (s *KeywordContext) RANK() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserRANK, 0)
}

func (s *KeywordContext) ROW_NUMBER() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserROW_NUMBER, 0)
}

func (s *KeywordContext) GENERATED() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserGENERATED, 0)
}

func (s *KeywordContext) ALWAYS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserALWAYS, 0)
}

func (s *KeywordContext) STORED() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserSTORED, 0)
}

func (s *KeywordContext) TRUE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserTRUE, 0)
}

func (s *KeywordContext) FALSE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserFALSE, 0)
}

func (s *KeywordContext) WINDOW() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserWINDOW, 0)
}

func (s *KeywordContext) NULLS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserNULLS, 0)
}

func (s *KeywordContext) FIRST() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserFIRST, 0)
}

func (s *KeywordContext) LAST() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserLAST, 0)
}

func (s *KeywordContext) FILTER() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserFILTER, 0)
}

func (s *KeywordContext) GROUPS() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserGROUPS, 0)
}

func (s *KeywordContext) EXCLUDE() antlr.TerminalNode {
	return s.GetToken(InformixSQLParserEXCLUDE, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InformixSQLParserListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *InformixSQLParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, InformixSQLParserRULE_keyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(429)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-64) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&-4611686018427387905) != 0) || ((int64((_la-128)) & ^0x3f) == 0 && ((int64(1)<<(_la-128))&1152921504606846975) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
