//ff:func feature=dump type=decoder control=sequence
//ff:what TestDecodeEntityLexeme — a lexeme line gives its lemmas, language, category, forms with features and senses with glosses and claims
package dump

import "testing"

func TestDecodeEntityLexeme(t *testing.T) {
	line := `{"type":"lexeme","id":"L4","lemmas":{"en":{"language":"en","value":"windsurf"}},"lexicalCategory":"Q24905","language":"Q1860","claims":{},` +
		`"forms":[{"id":"L4-F1","representations":{"en":{"language":"en","value":"windsurfing"}},"grammaticalFeatures":["Q10345583"],"claims":{}}],` +
		`"senses":[{"id":"L4-S1","glosses":{"fr":{"language":"fr","value":"faire de la planche à voile"}},"claims":{"P5137":[{"mainsnak":{"snaktype":"value","property":"P5137","datatype":"wikibase-item","datavalue":{"value":{"entity-type":"item","id":"Q168"},"type":"wikibase-entityid"}},"type":"statement","id":"L4-S1$x","rank":"normal"}]}}],` +
		`"lastrevid":7,"modified":"2026-10-20T00:00:00Z"}`
	e, err := DecodeEntity([]byte(line))
	if err != nil || e.Type != "lexeme" || e.Lemmas["en"].Value != "windsurf" || e.Language != "Q1860" || e.LexicalCategory != "Q24905" {
		t.Fatalf("head %+v %v", e, err)
	}
	if len(e.Forms) != 1 || e.Forms[0].Representations["en"].Value != "windsurfing" || e.Forms[0].GrammaticalFeatures[0] != "Q10345583" {
		t.Fatalf("forms %+v", e.Forms)
	}
	if len(e.Senses) != 1 || e.Senses[0].Glosses["fr"].Value == "" || len(e.Senses[0].Claims["P5137"]) != 1 {
		t.Fatalf("senses %+v", e.Senses)
	}
}
