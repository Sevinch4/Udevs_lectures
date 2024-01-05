package basket

import "main/packages/product"

type Basket struct {
	ProductList []product.Product
	Total       int
}
