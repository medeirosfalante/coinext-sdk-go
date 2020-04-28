package coinext

import (
	"encoding/json"
)

// Public struct
type Public struct {
	client *APIClient
	Crypto string
}

//Book retorno
type Book []BookItem

//BookQuery - objeto do post para buscar no book
type BookQuery struct {
	OMSId        int `json:"OMSId"`
	InstrumentID int `json:"InstrumentId"` // BTCBRL: 1  LTCBRL: 2 ETHBRL: 4 XRPBRL: 6
	Depth        int `json:"Depth"`
}

//BookItem - item do book
type BookItem []float64

//ID -  Identificador sequencial do "snapshot".
func (c BookItem) ID() int64 {
	return int64(c[0])
}

//Accounts - número de contas associadas
func (c BookItem) Accounts() int {
	return int(c[1])
}

//ActionTimeStamp -  "timestamp" da entrada (número de milisegundos desde 01/01/1970)
func (c BookItem) ActionTimeStamp() int64 {
	return int64(c[2])
}

//ActionType -  tipo do lance.
func (c BookItem) ActionType() int {
	return int(c[3])
}

//LastTradePrice -  valor da última negociação. Sim, é o mesmo número para todas os arrays retornados por este serviço.
func (c BookItem) LastTradePrice() float64 {
	return c[4]
}

//TotalOrders - número de ordens associadas.
func (c BookItem) TotalOrders() int {
	return int(c[5])
}

//Price - Preço do lance.
func (c BookItem) Price() float64 {
	return c[6]
}

//InstrumentID - Código do instrumento
func (c BookItem) InstrumentID() int64 {
	return int64(c[7])
}

//Qty - Quantidade disponível para compra ou venda neste lance.
func (c BookItem) Qty() float64 {
	return float64(c[8])
}

//Type - 0 = compra, 1 = venda
func (c BookItem) Type() int {
	return int(c[9])
}

//Public - Create a new instance struct
func (c *APIClient) Public() *Public {
	return &Public{client: c}
}

// OrderBook - OrderBook in exchange
func (p Public) OrderBook(query *BookQuery) (*Book, *Error, error) {
	var response *Book
	if query == nil {
		query = &BookQuery{
			InstrumentID: 1,
			OMSId:        1,
			Depth:        1,
		}
	}
	if query.InstrumentID == 0 {
		query.InstrumentID = 1
	}
	if query.OMSId == 0 {
		query.OMSId = 1
	}
	if query.Depth == 0 {
		query.Depth = 1
	}
	dataQuery, _ := json.Marshal(query)
	err, errAPI := p.client.Request("POST", "/GetL2Snapshot", dataQuery, nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
