package json

// GetTransactionDetail SevenBank's Out Entity
type GetTransactionDetail struct {
	MessageType  string `json:"message_type"`
	TranFlg      string `json:"tran_flg"`
	Amount       string `json:"amount"`
	BTime        string `json:"b_time"`
	BadiID       string `json:"badi_id"`
	ErrorCode    string `json:"error_code"`
	ErrorDetail  string `json:"error_detail"`
	FeeFlg       string `json:"fee_flg"`
	CustFee      string `json:"cust_fee"`
	PrintFlg2    string `json:"print_flg_2"`
	MarkFlg      string `json:"mark_flg"`
	DispFlg      string `json:"disp_flg"`
	BalanceFlg   string `json:"balance_flg"`
	PrintMes1    string `json:"print_mes1"`
	PrintMes2    string `json:"print_mes2"`
	PrintMes3    string `json:"print_mes3"`
	PrintMes4    string `json:"print_mes4"`
	PrintMes5    string `json:"print_mes5"`
	PrintMes6    string `json:"print_mes6"`
	DispMes1     string `json:"disp_mes1"`
	DispMes2     string `json:"disp_mes2"`
	DispMes3     string `json:"disp_mes3"`
	DispMes4     string `json:"disp_mes4"`
	DispBalance1 string `json:"disp_balance1"`
	DispBalance2 string `json:"disp_balance2"`
	SecInf       string `json:"sec_inf"`
	CipKeyInf    string `json:"cip_key_inf"`
}

// AcceptRepayment SevenBank's Out Entity
type AcceptRepayment struct {
	MessageType  string `json:"message_type"`
	TranFlg      string `json:"tran_flg"`
	Amount       string `json:"amount"`
	BTime        string `json:"b_time"`
	BadiID       string `json:"badi_id"`
	ErrorCode    string `json:"error_code"`
	ErrorDetail  string `json:"error_detail"`
	Balance1     string `json:"balance_1"`
	FeeFlg       string `json:"fee_flg"`
	CustFee      string `json:"cust_fee"`
	PrintFlg2    string `json:"print_flg_2"`
	MarkFlg      string `json:"mark_flg"`
	DispFlg      string `json:"disp_flg"`
	BalanceFlg   string `json:"balance_flg"`
	PrintMes1    string `json:"print_mes1"`
	PrintMes2    string `json:"print_mes2"`
	PrintMes3    string `json:"print_mes3"`
	PrintMes4    string `json:"print_mes4"`
	PrintMes5    string `json:"print_mes5"`
	PrintMes6    string `json:"print_mes6"`
	DispMes1     string `json:"disp_mes1"`
	DispMes2     string `json:"disp_mes2"`
	DispMes3     string `json:"disp_mes3"`
	DispMes4     string `json:"disp_mes4"`
	DispBalance1 string `json:"disp_balance1"`
	DispBalance2 string `json:"disp_balance2"`
	SecInf       string `json:"sec_inf"`
	CipKeyInf    string `json:"cip_key_inf"`
}
