//ff:func feature=dump type=decoder control=sequence
//ff:what sampleLine — a dump line of a small item with labels, aliases, a sitelink and claims with qualifiers, references and novalue
package dump

// sampleLine is one dump line (with its trailing comma) of a small item Q42:
// two labels, a description, two English aliases, a sitelink, a P69 claim
// with two qualifiers (in qualifiers-order P580 then P582) and a reference,
// a quantity claim with unit and bounds, and a novalue P40 claim.
func sampleLine() []byte {
	return []byte(`{"type":"item","id":"Q42","lastrevid":2401234567,"modified":"2026-10-18T09:00:00Z",` +
		`"labels":{"en":{"language":"en","value":"Douglas Adams"},"ko":{"language":"ko","value":"더글러스 애덤스"}},` +
		`"descriptions":{"en":{"language":"en","value":"English writer"}},` +
		`"aliases":{"en":[{"language":"en","value":"DNA"},{"language":"en","value":"Douglas Noël Adams"}]},` +
		`"claims":{"P69":[{"mainsnak":{"snaktype":"value","property":"P69","hash":"m1","datavalue":{"value":{"entity-type":"item","numeric-id":691283,"id":"Q691283"},"type":"wikibase-entityid"},"datatype":"wikibase-item"},` +
		`"type":"statement","id":"Q42$A","rank":"normal",` +
		`"qualifiers":{"P582":[{"snaktype":"value","property":"P582","hash":"q2","datavalue":{"value":{"time":"+1974-00-00T00:00:00Z","precision":9},"type":"time"},"datatype":"time"}],` +
		`"P580":[{"snaktype":"value","property":"P580","hash":"q1","datavalue":{"value":{"time":"+1971-00-00T00:00:00Z","precision":9},"type":"time"},"datatype":"time"}]},` +
		`"qualifiers-order":["P580","P582"],` +
		`"references":[{"hash":"r1","snaks":{"P248":[{"snaktype":"value","property":"P248","datavalue":{"value":{"entity-type":"item","id":"Q5375741"},"type":"wikibase-entityid"},"datatype":"wikibase-item"}]},"snaks-order":["P248"]}]}],` +
		`"P2048":[{"mainsnak":{"snaktype":"value","property":"P2048","hash":"m2","datavalue":{"value":{"amount":"+1.96","unit":"http://www.wikidata.org/entity/Q11573","upperBound":"+1.97","lowerBound":"+1.95"},"type":"quantity"},"datatype":"quantity"},"type":"statement","id":"Q42$B","rank":"preferred"}],` +
		`"P40":[{"mainsnak":{"snaktype":"novalue","property":"P40","hash":"m3","datatype":"wikibase-item"},"type":"statement","id":"Q42$C","rank":"normal"}]},` +
		`"sitelinks":{"enwiki":{"site":"enwiki","title":"Douglas Adams","badges":["Q17437796"]}}},` + "\n")
}
