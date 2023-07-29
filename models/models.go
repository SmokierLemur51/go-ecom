package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {

	ID 				primitive.ObjectID
	First_Name		*string
	Last_Name		*string
	Password		*string
	Email			*string
	Phone			*string
	Token 			*string
	Refresh_Token   *string
	Created_At		time.Time
	Updated_At		time.Time
	User_ID			string
	UserCart		[]ProductUser	// a slice of type ProductUser
	Adress_Details	[]Adress
	Order_Status	[]Order
}

type Prodct struct {

	Product_ID  	primitive.ObjectID
	Product_Name	*string
	Price			*uint64		// uint cant be negative ... 
	Rating			*uint8
	Image			*string		// this will just be a stored url
}

type ProductUser struct {

	Product_ID		primitive.ObjectID
	Product_Name
	Price
	Rating
	Image
}

type Address struct {

	Address_ID	primitive.ObjectID
	House
	Street
	City
	Zip
}

type Order struct {

	Order_ID		primitive.ObjectID
	Order_Cart
	Ordered_At
	Price
	Discount
	Payment_Method
}

type Payment struct {

	Digital	bool
	COD		bool
}
