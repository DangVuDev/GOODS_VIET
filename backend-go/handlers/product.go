package handlers

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"product_trace/config"
	product "product_trace/contract" // Corrected import path from contract to product
	"product_trace/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// @title Product Trace API
// @version 1.0
// @description API for managing agricultural products on blockchain with QR code support
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:5000
// @BasePath /api
// @schemes http

var productTrace = product.NewProductTrace() // Updated to use product package

// AddProduct godoc
// @Summary Add a new agricultural product to the blockchain
// @Description Add a product with name, manufacturer, initial status, details, and image to the blockchain and store QR code in MongoDB
// @Tags Products
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Product Name"
// @Param manufacturer formData string true "Manufacturer"
// @Param initialStatus formData string true "Initial Status"
// @Param initialDetails formData string true "Initial Details"
// @Param image formData file true "Product Image"
// @Success 200 {object} object{productId=int64,qrCodeUrl=string,qrCodeData=string,imageUrl=string}
// @Failure 400 {object} object{error=string}
// @Failure 500 {object} object{error=string}
// @Router /product/add [post]
func AddProduct(c *gin.Context) {
	name := c.PostForm("name")
	manufacturer := c.PostForm("manufacturer")
	initialStatus := c.PostForm("initialStatus")
	initialDetails := c.PostForm("initialDetails")

	if name == "" || manufacturer == "" || initialStatus == "" || initialDetails == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	// Lấy file ảnh
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	// Lưu file tạm
	filePath := fmt.Sprintf("./tmp/%s", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded image"})
		return
	}

	// Upload ảnh lên ImageKit
	imageKitService := NewImageKitService("private_CJGv3VMsn3oKHHw94yHXKruc5WI=")
	uploadResp, err := imageKitService.UploadImage(filePath, file.Filename, "/product_images")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image to ImageKit"})
		return
	}

	// Xóa file tạm sau khi upload
	_ = os.Remove(filePath)

	// Thêm sản phẩm vào blockchain
	productID, err := productTrace.AddProduct(name, manufacturer, initialStatus, initialDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product to blockchain: " + err.Error()})
		return
	}

	// Tạo QR code
	qrCodeURL := "https://goods-viet-web.onrender.com/product/" + strconv.FormatInt(productID, 10)
	qr, err := qrcode.New(qrCodeURL, qrcode.Medium)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code: " + err.Error()})
		return
	}
	qrCodeBytes, err := qr.PNG(256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode QR code: " + err.Error()})
		return
	}

	// Lưu vào MongoDB
	collection := config.Client.Database("product_trace").Collection("products")
	product := models.Product{
		ProductID: int(productID),
		QRCodeURL: qrCodeURL,
		ImageURL:  uploadResp.URL,
	}
	_, err = collection.InsertOne(context.Background(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save product to MongoDB: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"productId":   productID,
		"qrCodeUrl":   qrCodeURL,
		"qrCodeData":  base64.StdEncoding.EncodeToString(qrCodeBytes),
		"imageUrl":    uploadResp.URL,
	})
}


// UpdateProductStatus godoc
// @Summary Update product status
// @Description Updates the status and details of a product on the blockchain using the provided product ID.
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID" Format(int64)
// @Param status body object{status=string,details=string} true "Status data"
// @Success 200 {object} object{message=string} "Status updated successfully"
// @Failure 400 {object} object{error=string} "Invalid product ID or missing required fields"
// @Failure 500 {object} object{error=string} "Failed to update status on blockchain"
// @Router /product/{id}/status [post]
func UpdateProductStatus(c *gin.Context) {
    productID, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    var req struct {
        Status  string `json:"status" binding:"required"`
        Details string `json:"details" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err = productTrace.UpdateStatus(productID, req.Status, req.Details)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status on blockchain: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully"})
}

// GetProduct godoc
// @Summary Get product details by ID
// @Description Retrieve product details and status history from the blockchain and MongoDB
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} object{productId=int64,name=string,manufacturer=string,statuses=array,qrCodeUrl=string}
// @Failure 400 {object} object{error=string}
// @Failure 404 {object} object{error=string}
// @Failure 500 {object} object{error=string}
// @Router /product/{id} [get]
func GetProduct(c *gin.Context) {
    productID, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    name, manufacturer, err := productTrace.GetProduct(productID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product from blockchain: " + err.Error()})
        return
    }

    statuses, err := productTrace.GetProductStatuses(productID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product statuses from blockchain: " + err.Error()})
        return
    }

    collection := config.Client.Database("product_trace").Collection("products")
    var product models.Product
    err = collection.FindOne(context.Background(), bson.M{"productId": int(productID)}).Decode(&product)
    if err == mongo.ErrNoDocuments {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found in MongoDB"})
        return
    } else if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query MongoDB: " + err.Error()})
        return
    }

    responseStatuses := make([]map[string]interface{}, len(statuses))
    for i, s := range statuses {
        responseStatuses[i] = map[string]interface{}{
            "status":    s.Status,
            "details":   s.Details,
            "timestamp": time.Unix(s.Timestamp, 0).Format(time.RFC3339),
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "productId":    productID,
        "name":         name,
        "manufacturer": manufacturer,
        "statuses":     responseStatuses,
        "qrCodeUrl":    product.QRCodeURL,
		"imageURL":     product.ImageURL,
    })
}

// GetProducts godoc
// @Summary Get all products
// @Description Retrieve all products with their latest status from the blockchain and MongoDB
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} object{productId=int64,name=string,manufacturer=string,statuses=array,qrCodeUrl=string}
// @Failure 500 {object} object{error=string}
// @Router /products [get]
func GetProducts(c *gin.Context) {
    collection := config.Client.Database("product_trace").Collection("products")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query MongoDB: " + err.Error()})
        return
    }
    defer cursor.Close(ctx)

    var products []map[string]interface{}
    for cursor.Next(ctx) {
        var product models.Product
        if err := cursor.Decode(&product); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode product from MongoDB: " + err.Error()})
            return
        }

        name, manufacturer, err := productTrace.GetProduct(int64(product.ProductID))
        if err != nil {
            continue // Skip products with errors
        }

        statuses, err := productTrace.GetProductStatuses(int64(product.ProductID))
        if err != nil {
            continue // Skip products with errors
        }

        responseStatuses := make([]map[string]interface{}, len(statuses))
        for i, s := range statuses {
            responseStatuses[i] = map[string]interface{}{
                "status":    s.Status,
                "details":   s.Details,
                "timestamp": time.Unix(s.Timestamp, 0).Format(time.RFC3339),
            }
        }

        products = append(products, map[string]interface{}{
            "productId":    product.ProductID,
            "name":         name,
            "manufacturer": manufacturer,
            "statuses":     responseStatuses,
            "qrCodeUrl":    product.QRCodeURL,
        })
    }

    if err := cursor.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor error: " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, products)
}