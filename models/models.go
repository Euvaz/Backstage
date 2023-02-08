package models

type Token struct {
    Addr    string  `json:"addr"`
    Key     string  `json:"key"`
}

type TokenKey struct {
    Key string `json:"key"`
}
