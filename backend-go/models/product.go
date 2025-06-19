package models

type Product struct {
	ProductID int    `bson:"productId" json:"productId"`
	QRCodeURL string `bson:"qrCodeUrl" json:"qrCodeUrl"`
	ImageURL  string `bson:"imageURL" json:"imageURL"`
}