package tree_sitter_ledger_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_ledger "github.com/tree-sitter/tree-sitter-ledger/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_ledger.Language())
	if language == nil {
		t.Errorf("Error loading Ledger grammar")
	}
}
