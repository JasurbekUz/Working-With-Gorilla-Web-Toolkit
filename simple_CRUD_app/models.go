package main

type Category struct {
	CategoryId uint16	`json:"categoryId"`
	Name string			`json:"name"`
}

type GetClassified struct {
	ClassifiedId uint16	`json:"cassifiedId"`
	Title string		`json:"title"`
	Price float64		`json:"price"`
	CreatedAt string	`json:"createdAt"`
	Category Category 	`json:category`
} 

type PostClassifed struct {

	Title string		`json:"title"` 	
	Price float64		`json:"price"`
	CategoryId uint16	`json:"categoryId"`
} 

type PutClassified struct {
	ClassifiedId uint16	`json:"cassifiedId"`
	Title string		`json:"title"`
	Price float64		`json:"price"`
	CategoryId uint16	`json:"categoryId"`
}

var SELECT_ALL = `
	select
		cl.classified_id,
		cl.title,
		cl.price,
		cl.created_at,
		ct.category_id,
		ct.name
	from classifieds cl
	join categories as ct using (category_id)
`

var SELECT_ONCE = `
	select
		cl.classified_id,
		cl.title,
		cl.price,
		cl.created_at,
		ct.category_id,
		ct.name
	from classifieds cl
	join categories as ct using (category_id)
	where cl.classified_id = $1
`

var INSERT = `
	insert into 
		classifieds (title, price, category_id) 
	values 
		($1, $2, $3)
	returning 
		*

`

var UPDATE = `
	update classifieds
		set classified_id = coalesce($1, classified_id),
    	title = coalesce($2, title),
    	price = coalesce($3, price),
    	category_id = coalesce($4, category_id)
	WHERE 
		classified_id = $1
	returning
		*
`

var DELETE = `
	delete from 
		classifieds
	where
		classified_id = $1
	returning 
		*
`

