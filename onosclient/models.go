package onosclient

type Intents struct {
	Intent []Intent `json:"intents"`
}

type Intent struct {
	AppID       string        `json:"appId"`
	ID          string        `json:"id,omitempty"`
	Key         string        `json:"key,omitempty"`
	State       string        `json:"state,omitempty"`
	Type        string        `json:"type"`
	Resources   []string      `json:"resources,omitempty"`
	Selector    *Selector     `json:"selector,omitempty"` //pointer so omitempty is honored
	Treatment   *Treatment    `json:"treatment,omitempty"`
	Priority    int           `json:"priority,omitempty"`
	Constraints []Constraints `json:"constraints,omitempty"`
	One         string        `json:"one"`
	Two         string        `json:"two"`
}
type Selector struct {
	Criteria []Criteria `json:"criteria,omitempty"`
}
type Criteria struct {
	EthType string `json:"ethtype,omitempty"`
	Mac     string `json:"mac,omitempty"`
	Port    int    `json:"port,omitempty"`
	Type    string `json:"type,omitempty"`
}

type Treatment struct {
	ClearDeferred bool           `json:"cleardeferred,omitempty"`
	Deferred      []Instructions `json:"deferred,omitempty"` //for deferred instructions
	Instructions  []Instructions `json:"instructions,omitempty"`
}

type Instructions struct {
	Port string `json:"port,omitempty"`
	Type string `json:"type,omitempty"`
}

type Constraints struct {
	Inclusive bool     `json:"inclusive,omitempty"`
	Types     []string `json:"types,omitempty"`
	Type      string   `json:"type,omitempty"`
}
