package main

// APIResult das
type APIResult struct {
	Ret []struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Domain   int    `json:"domain"`
		Currency string `json:"currency"`
		Cash     struct {
			ID          int     `json:"id"`
			UserID      int     `json:"user_id"`
			Balance     float64 `json:"balance"`
			PreSub      int     `json:"pre_sub"`
			PreAdd      int     `json:"pre_add"`
			Currency    string  `json:"currency"`
			LastEntryAt string  `json:"last_entry_at"`
		} `json:"cash"`
	} `json:"ret"`
	Result  string `json:"result"`
	Profile struct {
		ExecutionTime int    `json:"execution_time"`
		ServerName    string `json:"server_name"`
	} `json:"profile"`
}

type conn struct {
	IP   string
	Host string
	Port string
}
