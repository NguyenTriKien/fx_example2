package main

import (
	"fmt"

	"go.uber.org/fx"
)

type Product struct {
	ID   string
	Name string
}

type ProductInterface interface {
	AddProduct(product Product) error
	GetProduct(id string) (Product, error)
} //Tao interface chua method

type ProductInstance struct {
	product map[string]Product
} // Product instance dung de quan li product va dung de truy xuat thong tin product

func NewProductInstance() ProductInstance {
	result := ProductInstance{
		product: make(map[string]Product),
	}
	return result
} // tao instance cua ProductInstance va trien khai ProductInstance la implement cua ProductInterface

func NewProductInstanceDecorate() ProductInstance {
	result := NewProductInstance()
	fmt.Println("Hello customer")
	return result

} //

// Implement AddProduct
func (p ProductInstance) AddProduct(product Product) error {
	p.product[string(product.ID)] = product
	return nil
}

// Implement GetProduct
func (p ProductInstance) GetProduct(id string) (Product, error) {
	product, ok := p.product[id]
	if !ok {
		return Product{}, fmt.Errorf("product not found")
	}
	return product, nil
}

var Module = fx.Options(
	/* fx.Provide(func() ProductInterface {
		return NewProductInstance()
	}), */
	fx.Provide(fx.Annotate(NewProductInstance, fx.As(new(ProductInterface)))),
	fx.Decorate(fx.Annotate(NewProductInstanceDecorate, fx.As(new(ProductInterface)))),
)

func main() {

	app := fx.New(
		Module,
		fx.Invoke(CreateProduct),
	)
	app.Run()

}

func CreateProduct(p ProductInterface) {

	product := Product{
		ID:   "1",
		Name: "Product",
	}

	p.AddProduct(product)
	fmt.Println("Product info:")
	fmt.Println(p.GetProduct("1"))

}
