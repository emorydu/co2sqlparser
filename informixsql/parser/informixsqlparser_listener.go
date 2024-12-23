// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

// Code generated from InformixSQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // InformixSQLParser

import "github.com/antlr4-go/antlr/v4"

// InformixSQLParserListener is a complete listener for a parse tree produced by InformixSQLParser.
type InformixSQLParserListener interface {
	antlr.ParseTreeListener

	// EnterSqlScript is called when entering the sqlScript production.
	EnterSqlScript(c *SqlScriptContext)

	// EnterUnitStatement is called when entering the unitStatement production.
	EnterUnitStatement(c *UnitStatementContext)

	// EnterCreateRole is called when entering the createRole production.
	EnterCreateRole(c *CreateRoleContext)

	// EnterDropRole is called when entering the dropRole production.
	EnterDropRole(c *DropRoleContext)

	// EnterDropSynonym is called when entering the dropSynonym production.
	EnterDropSynonym(c *DropSynonymContext)

	// EnterDropTable is called when entering the dropTable production.
	EnterDropTable(c *DropTableContext)

	// EnterDropTrigger is called when entering the dropTrigger production.
	EnterDropTrigger(c *DropTriggerContext)

	// EnterDropTrustedContext is called when entering the dropTrustedContext production.
	EnterDropTrustedContext(c *DropTrustedContextContext)

	// EnterDropType is called when entering the dropType production.
	EnterDropType(c *DropTypeContext)

	// EnterDropUser is called when entering the dropUser production.
	EnterDropUser(c *DropUserContext)

	// EnterDropView is called when entering the dropView production.
	EnterDropView(c *DropViewContext)

	// EnterDropXadatasource is called when entering the dropXadatasource production.
	EnterDropXadatasource(c *DropXadatasourceContext)

	// EnterDropXadataTypeSource is called when entering the dropXadataTypeSource production.
	EnterDropXadataTypeSource(c *DropXadataTypeSourceContext)

	// EnterDropAccessMethod is called when entering the dropAccessMethod production.
	EnterDropAccessMethod(c *DropAccessMethodContext)

	// EnterDropAggregate is called when entering the dropAggregate production.
	EnterDropAggregate(c *DropAggregateContext)

	// EnterDropDatabase is called when entering the dropDatabase production.
	EnterDropDatabase(c *DropDatabaseContext)

	// EnterDropIndex is called when entering the dropIndex production.
	EnterDropIndex(c *DropIndexContext)

	// EnterCloseStmt is called when entering the closeStmt production.
	EnterCloseStmt(c *CloseStmtContext)

	// EnterCloseDatabaseStmt is called when entering the closeDatabaseStmt production.
	EnterCloseDatabaseStmt(c *CloseDatabaseStmtContext)

	// EnterDatabaseStmt is called when entering the databaseStmt production.
	EnterDatabaseStmt(c *DatabaseStmtContext)

	// EnterCommitWorkStmt is called when entering the commitWorkStmt production.
	EnterCommitWorkStmt(c *CommitWorkStmtContext)

	// EnterReleaseSavepoint is called when entering the releaseSavepoint production.
	EnterReleaseSavepoint(c *ReleaseSavepointContext)

	// EnterRenameColumn is called when entering the renameColumn production.
	EnterRenameColumn(c *RenameColumnContext)

	// EnterRenameConstraint is called when entering the renameConstraint production.
	EnterRenameConstraint(c *RenameConstraintContext)

	// EnterRenameDatabase is called when entering the renameDatabase production.
	EnterRenameDatabase(c *RenameDatabaseContext)

	// EnterRenameIndex is called when entering the renameIndex production.
	EnterRenameIndex(c *RenameIndexContext)

	// EnterRenameSecurity is called when entering the renameSecurity production.
	EnterRenameSecurity(c *RenameSecurityContext)

	// EnterRenameSequence is called when entering the renameSequence production.
	EnterRenameSequence(c *RenameSequenceContext)

	// EnterRenameTable is called when entering the renameTable production.
	EnterRenameTable(c *RenameTableContext)

	// EnterRenameTrustedContext is called when entering the renameTrustedContext production.
	EnterRenameTrustedContext(c *RenameTrustedContextContext)

	// EnterRenameUser is called when entering the renameUser production.
	EnterRenameUser(c *RenameUserContext)

	// EnterRollbackWork is called when entering the rollbackWork production.
	EnterRollbackWork(c *RollbackWorkContext)

	// EnterSavepointStmt is called when entering the savepointStmt production.
	EnterSavepointStmt(c *SavepointStmtContext)

	// EnterSetAutofree is called when entering the setAutofree production.
	EnterSetAutofree(c *SetAutofreeContext)

	// EnterSetCollation is called when entering the setCollation production.
	EnterSetCollation(c *SetCollationContext)

	// EnterSetDataskip is called when entering the setDataskip production.
	EnterSetDataskip(c *SetDataskipContext)

	// EnterSetDebugFile is called when entering the setDebugFile production.
	EnterSetDebugFile(c *SetDebugFileContext)

	// EnterSetDeferredPrepareStatement is called when entering the setDeferredPrepareStatement production.
	EnterSetDeferredPrepareStatement(c *SetDeferredPrepareStatementContext)

	// EnterQuotedString is called when entering the quotedString production.
	EnterQuotedString(c *QuotedStringContext)

	// EnterAnyName is called when entering the anyName production.
	EnterAnyName(c *AnyNameContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// ExitSqlScript is called when exiting the sqlScript production.
	ExitSqlScript(c *SqlScriptContext)

	// ExitUnitStatement is called when exiting the unitStatement production.
	ExitUnitStatement(c *UnitStatementContext)

	// ExitCreateRole is called when exiting the createRole production.
	ExitCreateRole(c *CreateRoleContext)

	// ExitDropRole is called when exiting the dropRole production.
	ExitDropRole(c *DropRoleContext)

	// ExitDropSynonym is called when exiting the dropSynonym production.
	ExitDropSynonym(c *DropSynonymContext)

	// ExitDropTable is called when exiting the dropTable production.
	ExitDropTable(c *DropTableContext)

	// ExitDropTrigger is called when exiting the dropTrigger production.
	ExitDropTrigger(c *DropTriggerContext)

	// ExitDropTrustedContext is called when exiting the dropTrustedContext production.
	ExitDropTrustedContext(c *DropTrustedContextContext)

	// ExitDropType is called when exiting the dropType production.
	ExitDropType(c *DropTypeContext)

	// ExitDropUser is called when exiting the dropUser production.
	ExitDropUser(c *DropUserContext)

	// ExitDropView is called when exiting the dropView production.
	ExitDropView(c *DropViewContext)

	// ExitDropXadatasource is called when exiting the dropXadatasource production.
	ExitDropXadatasource(c *DropXadatasourceContext)

	// ExitDropXadataTypeSource is called when exiting the dropXadataTypeSource production.
	ExitDropXadataTypeSource(c *DropXadataTypeSourceContext)

	// ExitDropAccessMethod is called when exiting the dropAccessMethod production.
	ExitDropAccessMethod(c *DropAccessMethodContext)

	// ExitDropAggregate is called when exiting the dropAggregate production.
	ExitDropAggregate(c *DropAggregateContext)

	// ExitDropDatabase is called when exiting the dropDatabase production.
	ExitDropDatabase(c *DropDatabaseContext)

	// ExitDropIndex is called when exiting the dropIndex production.
	ExitDropIndex(c *DropIndexContext)

	// ExitCloseStmt is called when exiting the closeStmt production.
	ExitCloseStmt(c *CloseStmtContext)

	// ExitCloseDatabaseStmt is called when exiting the closeDatabaseStmt production.
	ExitCloseDatabaseStmt(c *CloseDatabaseStmtContext)

	// ExitDatabaseStmt is called when exiting the databaseStmt production.
	ExitDatabaseStmt(c *DatabaseStmtContext)

	// ExitCommitWorkStmt is called when exiting the commitWorkStmt production.
	ExitCommitWorkStmt(c *CommitWorkStmtContext)

	// ExitReleaseSavepoint is called when exiting the releaseSavepoint production.
	ExitReleaseSavepoint(c *ReleaseSavepointContext)

	// ExitRenameColumn is called when exiting the renameColumn production.
	ExitRenameColumn(c *RenameColumnContext)

	// ExitRenameConstraint is called when exiting the renameConstraint production.
	ExitRenameConstraint(c *RenameConstraintContext)

	// ExitRenameDatabase is called when exiting the renameDatabase production.
	ExitRenameDatabase(c *RenameDatabaseContext)

	// ExitRenameIndex is called when exiting the renameIndex production.
	ExitRenameIndex(c *RenameIndexContext)

	// ExitRenameSecurity is called when exiting the renameSecurity production.
	ExitRenameSecurity(c *RenameSecurityContext)

	// ExitRenameSequence is called when exiting the renameSequence production.
	ExitRenameSequence(c *RenameSequenceContext)

	// ExitRenameTable is called when exiting the renameTable production.
	ExitRenameTable(c *RenameTableContext)

	// ExitRenameTrustedContext is called when exiting the renameTrustedContext production.
	ExitRenameTrustedContext(c *RenameTrustedContextContext)

	// ExitRenameUser is called when exiting the renameUser production.
	ExitRenameUser(c *RenameUserContext)

	// ExitRollbackWork is called when exiting the rollbackWork production.
	ExitRollbackWork(c *RollbackWorkContext)

	// ExitSavepointStmt is called when exiting the savepointStmt production.
	ExitSavepointStmt(c *SavepointStmtContext)

	// ExitSetAutofree is called when exiting the setAutofree production.
	ExitSetAutofree(c *SetAutofreeContext)

	// ExitSetCollation is called when exiting the setCollation production.
	ExitSetCollation(c *SetCollationContext)

	// ExitSetDataskip is called when exiting the setDataskip production.
	ExitSetDataskip(c *SetDataskipContext)

	// ExitSetDebugFile is called when exiting the setDebugFile production.
	ExitSetDebugFile(c *SetDebugFileContext)

	// ExitSetDeferredPrepareStatement is called when exiting the setDeferredPrepareStatement production.
	ExitSetDeferredPrepareStatement(c *SetDeferredPrepareStatementContext)

	// ExitQuotedString is called when exiting the quotedString production.
	ExitQuotedString(c *QuotedStringContext)

	// ExitAnyName is called when exiting the anyName production.
	ExitAnyName(c *AnyNameContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)
}
