package transactions

type CreateTransactionRequest struct {
	CustomerId  string  `json:"customer_id"`
	Otr         float64 `json:"otr"`
	AdminFee    float64 `json:"admin_fee"`
	Installment float64 `json:"installment"`
	Interest    float64 `json:"interest"`
	AssetName   string  `json:"asset_name"`
	Tenor       int     `json:"tenor"`
}
