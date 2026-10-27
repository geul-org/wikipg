//ff:type feature=dump type=model
//ff:what Term — a label, description or alias: language and text
package dump

// Term is a label, description or alias value: its language code and text.
type Term struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}
