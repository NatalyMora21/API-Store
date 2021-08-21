package schema

type Buyer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Transaction struct {
	Id       string   `json:"id"`
	Buyer_id string   `json:"buyer_id"`
	Ip       string   `json:"ip"`
	Device   string   `json:"device"`
	Products []string `json:"products"`
}

//Info para crear los nods en Dgraph
//La info de los productos sería el id que se trae al crearlos en la Bd
type infoBuyer struct {
	buyer        Buyer
	transactions []Transaction
}
