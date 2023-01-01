package test

import (
	"fmt"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/setup"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func Test_write_redis(t *testing.T) {
}

func Test_read_redis(t *testing.T) {
	redistClient := setup.NewRedisClient()
	ctx := context.Background()
	since := time.Now()
	result, err := redistClient.Get(ctx, "s8s978").Result()
	helpers.LogIfError(err)
	after := time.Since(since)
	fmt.Printf("done on %d miliseconds \n", after.Milliseconds())
	fmt.Println("productId", result)
}
