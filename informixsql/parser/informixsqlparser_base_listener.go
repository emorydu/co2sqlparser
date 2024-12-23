// Copyright 2024 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by an Apache style
// license that can be found in the LICENSE file.

// Code generated from InformixSQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // InformixSQLParser

import "github.com/antlr4-go/antlr/v4"

// BaseInformixSQLParserListener is a complete listener for a parse tree produced by InformixSQLParser.
type BaseInformixSQLParserListener struct{}

var _ InformixSQLParserListener = &BaseInformixSQLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInformixSQLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInformixSQLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInformixSQLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInformixSQLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSqlScript is called when production sqlScript is entered.
func (s *BaseInformixSQLParserListener) EnterSqlScript(ctx *SqlScriptContext) {}

// ExitSqlScript is called when production sqlScript is exited.
func (s *BaseInformixSQLParserListener) ExitSqlScript(ctx *SqlScriptContext) {}

// EnterUnitStatement is called when production unitStatement is entered.
func (s *BaseInformixSQLParserListener) EnterUnitStatement(ctx *UnitStatementContext) {}

// ExitUnitStatement is called when production unitStatement is exited.
func (s *BaseInformixSQLParserListener) ExitUnitStatement(ctx *UnitStatementContext) {}

// EnterCreateRole is called when production createRole is entered.
func (s *BaseInformixSQLParserListener) EnterCreateRole(ctx *CreateRoleContext) {}

// ExitCreateRole is called when production createRole is exited.
func (s *BaseInformixSQLParserListener) ExitCreateRole(ctx *CreateRoleContext) {}

// EnterDropRole is called when production dropRole is entered.
func (s *BaseInformixSQLParserListener) EnterDropRole(ctx *DropRoleContext) {}

// ExitDropRole is called when production dropRole is exited.
func (s *BaseInformixSQLParserListener) ExitDropRole(ctx *DropRoleContext) {}

// EnterDropSynonym is called when production dropSynonym is entered.
func (s *BaseInformixSQLParserListener) EnterDropSynonym(ctx *DropSynonymContext) {}

// ExitDropSynonym is called when production dropSynonym is exited.
func (s *BaseInformixSQLParserListener) ExitDropSynonym(ctx *DropSynonymContext) {}

// EnterDropTable is called when production dropTable is entered.
func (s *BaseInformixSQLParserListener) EnterDropTable(ctx *DropTableContext) {}

// ExitDropTable is called when production dropTable is exited.
func (s *BaseInformixSQLParserListener) ExitDropTable(ctx *DropTableContext) {}

// EnterDropTrigger is called when production dropTrigger is entered.
func (s *BaseInformixSQLParserListener) EnterDropTrigger(ctx *DropTriggerContext) {}

// ExitDropTrigger is called when production dropTrigger is exited.
func (s *BaseInformixSQLParserListener) ExitDropTrigger(ctx *DropTriggerContext) {}

// EnterDropTrustedContext is called when production dropTrustedContext is entered.
func (s *BaseInformixSQLParserListener) EnterDropTrustedContext(ctx *DropTrustedContextContext) {}

// ExitDropTrustedContext is called when production dropTrustedContext is exited.
func (s *BaseInformixSQLParserListener) ExitDropTrustedContext(ctx *DropTrustedContextContext) {}

// EnterDropType is called when production dropType is entered.
func (s *BaseInformixSQLParserListener) EnterDropType(ctx *DropTypeContext) {}

// ExitDropType is called when production dropType is exited.
func (s *BaseInformixSQLParserListener) ExitDropType(ctx *DropTypeContext) {}

// EnterDropUser is called when production dropUser is entered.
func (s *BaseInformixSQLParserListener) EnterDropUser(ctx *DropUserContext) {}

// ExitDropUser is called when production dropUser is exited.
func (s *BaseInformixSQLParserListener) ExitDropUser(ctx *DropUserContext) {}

// EnterDropView is called when production dropView is entered.
func (s *BaseInformixSQLParserListener) EnterDropView(ctx *DropViewContext) {}

// ExitDropView is called when production dropView is exited.
func (s *BaseInformixSQLParserListener) ExitDropView(ctx *DropViewContext) {}

// EnterDropXadatasource is called when production dropXadatasource is entered.
func (s *BaseInformixSQLParserListener) EnterDropXadatasource(ctx *DropXadatasourceContext) {}

// ExitDropXadatasource is called when production dropXadatasource is exited.
func (s *BaseInformixSQLParserListener) ExitDropXadatasource(ctx *DropXadatasourceContext) {}

// EnterDropXadataTypeSource is called when production dropXadataTypeSource is entered.
func (s *BaseInformixSQLParserListener) EnterDropXadataTypeSource(ctx *DropXadataTypeSourceContext) {}

// ExitDropXadataTypeSource is called when production dropXadataTypeSource is exited.
func (s *BaseInformixSQLParserListener) ExitDropXadataTypeSource(ctx *DropXadataTypeSourceContext) {}

// EnterDropAccessMethod is called when production dropAccessMethod is entered.
func (s *BaseInformixSQLParserListener) EnterDropAccessMethod(ctx *DropAccessMethodContext) {}

// ExitDropAccessMethod is called when production dropAccessMethod is exited.
func (s *BaseInformixSQLParserListener) ExitDropAccessMethod(ctx *DropAccessMethodContext) {}

// EnterDropAggregate is called when production dropAggregate is entered.
func (s *BaseInformixSQLParserListener) EnterDropAggregate(ctx *DropAggregateContext) {}

// ExitDropAggregate is called when production dropAggregate is exited.
func (s *BaseInformixSQLParserListener) ExitDropAggregate(ctx *DropAggregateContext) {}

// EnterDropDatabase is called when production dropDatabase is entered.
func (s *BaseInformixSQLParserListener) EnterDropDatabase(ctx *DropDatabaseContext) {}

// ExitDropDatabase is called when production dropDatabase is exited.
func (s *BaseInformixSQLParserListener) ExitDropDatabase(ctx *DropDatabaseContext) {}

// EnterDropIndex is called when production dropIndex is entered.
func (s *BaseInformixSQLParserListener) EnterDropIndex(ctx *DropIndexContext) {}

// ExitDropIndex is called when production dropIndex is exited.
func (s *BaseInformixSQLParserListener) ExitDropIndex(ctx *DropIndexContext) {}

// EnterCloseStmt is called when production closeStmt is entered.
func (s *BaseInformixSQLParserListener) EnterCloseStmt(ctx *CloseStmtContext) {}

// ExitCloseStmt is called when production closeStmt is exited.
func (s *BaseInformixSQLParserListener) ExitCloseStmt(ctx *CloseStmtContext) {}

// EnterCloseDatabaseStmt is called when production closeDatabaseStmt is entered.
func (s *BaseInformixSQLParserListener) EnterCloseDatabaseStmt(ctx *CloseDatabaseStmtContext) {}

// ExitCloseDatabaseStmt is called when production closeDatabaseStmt is exited.
func (s *BaseInformixSQLParserListener) ExitCloseDatabaseStmt(ctx *CloseDatabaseStmtContext) {}

// EnterDatabaseStmt is called when production databaseStmt is entered.
func (s *BaseInformixSQLParserListener) EnterDatabaseStmt(ctx *DatabaseStmtContext) {}

// ExitDatabaseStmt is called when production databaseStmt is exited.
func (s *BaseInformixSQLParserListener) ExitDatabaseStmt(ctx *DatabaseStmtContext) {}

// EnterCommitWorkStmt is called when production commitWorkStmt is entered.
func (s *BaseInformixSQLParserListener) EnterCommitWorkStmt(ctx *CommitWorkStmtContext) {}

// ExitCommitWorkStmt is called when production commitWorkStmt is exited.
func (s *BaseInformixSQLParserListener) ExitCommitWorkStmt(ctx *CommitWorkStmtContext) {}

// EnterReleaseSavepoint is called when production releaseSavepoint is entered.
func (s *BaseInformixSQLParserListener) EnterReleaseSavepoint(ctx *ReleaseSavepointContext) {}

// ExitReleaseSavepoint is called when production releaseSavepoint is exited.
func (s *BaseInformixSQLParserListener) ExitReleaseSavepoint(ctx *ReleaseSavepointContext) {}

// EnterRenameColumn is called when production renameColumn is entered.
func (s *BaseInformixSQLParserListener) EnterRenameColumn(ctx *RenameColumnContext) {}

// ExitRenameColumn is called when production renameColumn is exited.
func (s *BaseInformixSQLParserListener) ExitRenameColumn(ctx *RenameColumnContext) {}

// EnterRenameConstraint is called when production renameConstraint is entered.
func (s *BaseInformixSQLParserListener) EnterRenameConstraint(ctx *RenameConstraintContext) {}

// ExitRenameConstraint is called when production renameConstraint is exited.
func (s *BaseInformixSQLParserListener) ExitRenameConstraint(ctx *RenameConstraintContext) {}

// EnterRenameDatabase is called when production renameDatabase is entered.
func (s *BaseInformixSQLParserListener) EnterRenameDatabase(ctx *RenameDatabaseContext) {}

// ExitRenameDatabase is called when production renameDatabase is exited.
func (s *BaseInformixSQLParserListener) ExitRenameDatabase(ctx *RenameDatabaseContext) {}

// EnterRenameIndex is called when production renameIndex is entered.
func (s *BaseInformixSQLParserListener) EnterRenameIndex(ctx *RenameIndexContext) {}

// ExitRenameIndex is called when production renameIndex is exited.
func (s *BaseInformixSQLParserListener) ExitRenameIndex(ctx *RenameIndexContext) {}

// EnterRenameSecurity is called when production renameSecurity is entered.
func (s *BaseInformixSQLParserListener) EnterRenameSecurity(ctx *RenameSecurityContext) {}

// ExitRenameSecurity is called when production renameSecurity is exited.
func (s *BaseInformixSQLParserListener) ExitRenameSecurity(ctx *RenameSecurityContext) {}

// EnterRenameSequence is called when production renameSequence is entered.
func (s *BaseInformixSQLParserListener) EnterRenameSequence(ctx *RenameSequenceContext) {}

// ExitRenameSequence is called when production renameSequence is exited.
func (s *BaseInformixSQLParserListener) ExitRenameSequence(ctx *RenameSequenceContext) {}

// EnterRenameTable is called when production renameTable is entered.
func (s *BaseInformixSQLParserListener) EnterRenameTable(ctx *RenameTableContext) {}

// ExitRenameTable is called when production renameTable is exited.
func (s *BaseInformixSQLParserListener) ExitRenameTable(ctx *RenameTableContext) {}

// EnterRenameTrustedContext is called when production renameTrustedContext is entered.
func (s *BaseInformixSQLParserListener) EnterRenameTrustedContext(ctx *RenameTrustedContextContext) {}

// ExitRenameTrustedContext is called when production renameTrustedContext is exited.
func (s *BaseInformixSQLParserListener) ExitRenameTrustedContext(ctx *RenameTrustedContextContext) {}

// EnterRenameUser is called when production renameUser is entered.
func (s *BaseInformixSQLParserListener) EnterRenameUser(ctx *RenameUserContext) {}

// ExitRenameUser is called when production renameUser is exited.
func (s *BaseInformixSQLParserListener) ExitRenameUser(ctx *RenameUserContext) {}

// EnterRollbackWork is called when production rollbackWork is entered.
func (s *BaseInformixSQLParserListener) EnterRollbackWork(ctx *RollbackWorkContext) {}

// ExitRollbackWork is called when production rollbackWork is exited.
func (s *BaseInformixSQLParserListener) ExitRollbackWork(ctx *RollbackWorkContext) {}

// EnterSavepointStmt is called when production savepointStmt is entered.
func (s *BaseInformixSQLParserListener) EnterSavepointStmt(ctx *SavepointStmtContext) {}

// ExitSavepointStmt is called when production savepointStmt is exited.
func (s *BaseInformixSQLParserListener) ExitSavepointStmt(ctx *SavepointStmtContext) {}

// EnterSetAutofree is called when production setAutofree is entered.
func (s *BaseInformixSQLParserListener) EnterSetAutofree(ctx *SetAutofreeContext) {}

// ExitSetAutofree is called when production setAutofree is exited.
func (s *BaseInformixSQLParserListener) ExitSetAutofree(ctx *SetAutofreeContext) {}

// EnterSetCollation is called when production setCollation is entered.
func (s *BaseInformixSQLParserListener) EnterSetCollation(ctx *SetCollationContext) {}

// ExitSetCollation is called when production setCollation is exited.
func (s *BaseInformixSQLParserListener) ExitSetCollation(ctx *SetCollationContext) {}

// EnterSetDataskip is called when production setDataskip is entered.
func (s *BaseInformixSQLParserListener) EnterSetDataskip(ctx *SetDataskipContext) {}

// ExitSetDataskip is called when production setDataskip is exited.
func (s *BaseInformixSQLParserListener) ExitSetDataskip(ctx *SetDataskipContext) {}

// EnterSetDebugFile is called when production setDebugFile is entered.
func (s *BaseInformixSQLParserListener) EnterSetDebugFile(ctx *SetDebugFileContext) {}

// ExitSetDebugFile is called when production setDebugFile is exited.
func (s *BaseInformixSQLParserListener) ExitSetDebugFile(ctx *SetDebugFileContext) {}

// EnterSetDeferredPrepareStatement is called when production setDeferredPrepareStatement is entered.
func (s *BaseInformixSQLParserListener) EnterSetDeferredPrepareStatement(ctx *SetDeferredPrepareStatementContext) {
}

// ExitSetDeferredPrepareStatement is called when production setDeferredPrepareStatement is exited.
func (s *BaseInformixSQLParserListener) ExitSetDeferredPrepareStatement(ctx *SetDeferredPrepareStatementContext) {
}

// EnterQuotedString is called when production quotedString is entered.
func (s *BaseInformixSQLParserListener) EnterQuotedString(ctx *QuotedStringContext) {}

// ExitQuotedString is called when production quotedString is exited.
func (s *BaseInformixSQLParserListener) ExitQuotedString(ctx *QuotedStringContext) {}

// EnterAnyName is called when production anyName is entered.
func (s *BaseInformixSQLParserListener) EnterAnyName(ctx *AnyNameContext) {}

// ExitAnyName is called when production anyName is exited.
func (s *BaseInformixSQLParserListener) ExitAnyName(ctx *AnyNameContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseInformixSQLParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseInformixSQLParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseInformixSQLParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseInformixSQLParserListener) ExitKeyword(ctx *KeywordContext) {}
