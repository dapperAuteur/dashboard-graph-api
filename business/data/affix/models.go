package affix

// affix represents a part of a word or a part that may be added to a word
type Affix struct {
	ID        string   `json:"id"`
	Example   []string `json:"example"`
	Meaning   []string `json:"meaning"`
	Media     []string `json:"media"`
	Morpheme  string   `json:"morpheme"`
	Note      []string `json:"note"`
	Tongue    string   `json:"tongue"`
	AffixType []string `json:"affix_type"`
}

type addResult struct {
	AddAffix struct {
		Affix []struct {
			ID string `json:"id"`
		} `json:"affix"`
	} `json:"addAffix"`
}

func (addResult) document() string {
	return `{
		affix {
			id
		}
	}`
}
