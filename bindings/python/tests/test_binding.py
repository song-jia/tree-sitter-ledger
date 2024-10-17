from unittest import TestCase

import tree_sitter, tree_sitter_ledger


class TestLanguage(TestCase):
    def test_can_load_grammar(self):
        try:
            tree_sitter.Language(tree_sitter_ledger.language())
        except Exception:
            self.fail("Error loading Ledger grammar")
