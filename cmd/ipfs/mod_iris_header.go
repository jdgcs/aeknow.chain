package main

import (
	"html/template"
	"net/http"

	"github.com/aeternity/aepp-sdk-go/v7/account"

	"github.com/aeternity/aepp-sdk-go/v7/transactions"
	utils "github.com/aeternity/aepp-sdk-go/v7/utils"
)

type AccountInfo struct {
	// Name of authorization function for generalized account
	AuthFun string `json:"auth_fun,omitempty"`

	// Balance
	// Required: true
	Balance utils.BigInt `json:"balance"`

	// Id of authorization contract for generalized account
	ContractID string `json:"contract_id,omitempty"`

	// Public key
	// Required: true
	ID string `json:"id"`

	// kind
	// Enum: [basic generalized]
	Kind string `json:"kind,omitempty"`

	// Nonce
	// Required: true
	Nonce uint64 `json:"nonce"`

	// Payable
	Payable bool `json:"payable,omitempty"`
}

type HandleFnc func(http.ResponseWriter, *http.Request)

//HOMEPAGE
var aecommands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

type PageData struct {
	PageId      int
	PageContent template.HTML
	PageTitle   string
}

type PageReg struct {
	PageTitle string
	SubTitle  string
	Register  string
	Lang      langFile
}

type PageLogin struct {
	Options template.HTML
	Lang    langFile
}

type AeknowConfig struct {
	PublicNode string
	APINode    string
	IPFSNode   string
	LocalWeb   string
}

var myAccount account.Account
var globalAccount account.Account
var signAccount *account.Account
var NodeConfig AeknowConfig

//AENS
type TTLer func(offset uint64) (ttl uint64, err error)

type AENSNames struct {
	Aensname      string
	Expire_height int64
}

type NameSlice struct {
	Names []AENSNames
}

type AENSBidinfo struct {
	Aensname   string
	Lastbidder string
	Lastprice  string
}

type NameBidSlice struct {
	Names []AENSBidinfo
}

type AENSInfo struct {
	ID       string
	TTL      uint64
	Pointers []transactions.NamePointer
}

type PageAENS struct {
	PageId      int
	PageContent template.HTML
	PageTitle   string
	Account     string
	Balance     string
	Nonce       uint64
}
