package line

type lineConfig struct {
	chanToken  string
	chanSecret string
}

func NewLineConfig(ctk, cs string) *lineConfig {
	return &lineConfig{
		chanToken:  ctk,
		chanSecret: cs,
	}
}
