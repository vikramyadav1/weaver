package product

type Product struct {
	Id          int    `db:id`
	Name        string `db:name`
	Description string `db:name`
	Price       int    `db:price`

	GroupId int `db:group_id`
}
