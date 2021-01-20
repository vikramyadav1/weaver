package product

type Product struct {
	FirstName string `db:first_name`
	LastName  string `db:last_name`
	Email     string `db:email`

	GroupId int `db:group_id`
}

type ProductJson struct {
	FirstName string
	LastName  string
	Email     string

	References struct {
		Group int
	}
}

func (p *Product) ToJsonStruct() ProductJson {
	return ProductJson{
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,

		References: struct {
			Group int
		}{
			Group: p.GroupId,
		},
	}
}
