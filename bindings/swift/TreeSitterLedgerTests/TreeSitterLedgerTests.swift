import XCTest
import SwiftTreeSitter
import TreeSitterLedger

final class TreeSitterLedgerTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_ledger())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Ledger grammar")
    }
}
