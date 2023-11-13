package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/ramdanariadi/grocery-product-service/main/controller"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/model"
	"github.com/ramdanariadi/grocery-product-service/main/setup"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	env := os.Getenv("ENVIRONMENT")
	if "" == env {
		env = "development"
	}
	err := godotenv.Load(".env." + env)
	utils.LogIfError(err)
	err = godotenv.Load()
	utils.LogIfError(err)
	connection, err := setup.NewDbConnection()
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: connection}))
	utils.PanicIfError(err)
	err = db.AutoMigrate(&model.Category{}, &model.Product{}, &model.Wishlist{}, &model.Cart{}, &model.Transaction{}, &model.TransactionDetail{}, &model.Shop{})
	utils.LogIfError(err)

	client := setup.NewRedisClient()

	router := gin.Default()
	router.Use(gin.CustomRecovery(exception.Handler))

	shopGroup := router.Group("api/v1/shop")
	{
		shopController := controller.NewShopController(db)
		shopGroup.POST("", shopController.AddShop)
		shopGroup.PUT("", shopController.EditShop)
		shopGroup.GET("", shopController.GetShop)
		shopGroup.DELETE("", shopController.DeleteShop)
	}

	categoryRoute := router.Group("api/v1/category")
	{
		categoryController := controller.NewCategoryController(db)
		categoryRoute.POST("", categoryController.Save)
		categoryRoute.GET("/:id", categoryController.FindById)
		categoryRoute.GET("", categoryController.FindAll)
		categoryRoute.PUT("/:id", categoryController.Update)
		categoryRoute.DELETE("/:id", categoryController.Delete)
	}

	productRoute := router.Group("api/v1/product")
	{
		productController := controller.NewProductController(db, client)
		productRoute.POST("", productController.Save)
		productRoute.GET("/:id", productController.FindById)
		productRoute.GET("", productController.FindAll)
		productRoute.PUT("/:id", productController.Update)
		productRoute.DELETE("/:id", productController.Delete)
		productRoute.PUT("/top/:id", productController.SetTopProduct)
		productRoute.PUT("/recommendation/:id", productController.SetRecommendationProduct)
	}

	cartRoute := router.Group("api/v1/cart")
	{
		cartController := controller.NewController(db)
		cartRoute.POST("/:productId/:total", cartController.Store)
		cartRoute.DELETE("/:id", cartController.Destroy)
		cartRoute.GET("", cartController.Find)
	}

	wishlistRoute := router.Group("api/v1/wishlist")
	{
		wishlistController := controller.NewWishlistController(db)
		wishlistRoute.POST("/:productId", wishlistController.Store)
		wishlistRoute.DELETE("/:productId", wishlistController.Destroy)
		wishlistRoute.GET("", wishlistController.Find)
		wishlistRoute.GET("/:productId", wishlistController.FindByProductId)
	}

	transactionGroup := router.Group("api/v1/transaction")
	{
		transactionController := controller.NewTransactionController(db)
		transactionGroup.POST("", transactionController.Save)
		transactionGroup.GET("", transactionController.Find)
	}

	err = router.Run(":10000")
	utils.LogIfError(err)
}
