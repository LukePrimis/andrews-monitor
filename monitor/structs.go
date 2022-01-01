package monitor

import "net/http"

type AndrewsTask struct {
	Client       *http.Client
	LoginSession string
	Tx           string
	DuoUrl       string
	PromptPath   string
	Sid          string
	Txid         string
	Auth         string
	App          string
	Saml         string
	Akey         string
	Ukey         string
	FormID       string
	Times        map[string]bool
}
