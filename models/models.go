package models

type Token struct {
    Addr    string  `json:"addr"`
    Port    int     `json:"port"`
    Key     string  `json:"key"`
}
