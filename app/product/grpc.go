// product/grpp.go

package product

import (
  "database/sql"
  "context"
  "fmt"

  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
  wrappers "github.com/golang/protobuf/ptypes/wrappers"

  "app/common"
  "app/test_utils"
)

// ProductServiceInterface is implemented by ProductService
type ProductServiceInterface struct{
  app *common.App
}

// CreateProduct implements ProductService.CreateProduct
func (s *ProductServiceInterface) CreateProduct(ctx context.Context, req *ProductSchema) (*ProductSchema, error) {
  testutils.Log(fmt.Sprint("ProductService.CreateProduct"))
  p := Product{
    Schema: req,
    Model: nil,
  }

  testutils.Log(fmt.Sprintf("Product: %#v", p.Schema))

  p.copySchema()

  if err := p.Model.createProduct(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  p.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", p.Schema))

  return p.Schema, nil
}

// GetProduct implements ProductService.ProductQuery
func (s *ProductServiceInterface) GetProduct(ctx context.Context, req *ProductQuery) (*ProductSchema, error) {
  testutils.Log(fmt.Sprint("ProductService.GetProduct"))
  p := Product{
    Model: &ProductModel{ID: int(req.Id)},
    Schema: nil,
  }

  if err := p.Model.readProduct(s.app.DB); err != nil {
    switch err {
    case sql.ErrNoRows:
      return nil, status.Error(codes.NotFound, "Product not found")
    default:
      return nil, status.Error(codes.Internal, err.Error())
    }
  }

  p.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", p.Schema))

  return p.Schema, nil
}

// UpdateProduct implements ProductService.UpdateProduct
func (s *ProductServiceInterface) UpdateProduct(ctx context.Context, req *ProductQuery) (*ProductSchema, error) {
  testutils.Log(fmt.Sprint("ProductService.UpdateProduct"))
  p := Product{
    Schema: req.Product,
    Model: &ProductModel{ID: int(req.Id)},
  }

  if err := p.Model.readProduct(s.app.DB); err != nil {
    switch err {
    case sql.ErrNoRows:
      return nil, status.Error(codes.NotFound, "Product not found")
    default:
      return nil, status. Error(codes.Internal, err.Error())
    }
  }

  p.copySchema()

  if err := p.Model.updateProduct(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  p.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", p.Schema))
  return p.Schema, nil
}

// DeleteProduct implements ProductService.DeleteProduct
func (s *ProductServiceInterface) DeleteProduct(ctx context.Context, req *ProductQuery) (*wrappers.StringValue, error) {
  testutils.Log(fmt.Sprint("ProductService.DeleteProduct"))

  p := Product{
    Model: &ProductModel{ID: int(req.Id)},
    Schema: nil,
  }
  if err := p.Model.deleteProduct(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  testutils.Log(fmt.Sprint("Response:\n{ value: \"success\" }"))

  m := &wrappers.StringValue{Value: "success"}
  return m, nil
}

// GetProducts implements ProductService.GetProducts
func (s *ProductServiceInterface) GetProducts(ctx context.Context, req *ProductsQuery) (*ProductsResponse, error) {
  testutils.Log(fmt.Sprint("ProductService.GetProducts"))

  count, start := int(req.Count), int(req.Start)

  testutils.Log(fmt.Sprintf("{ count: %d, start: %d }", count, start))

  products, err := readProducts(s.app.DB, start, count)
  if err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  res := &ProductsResponse{
    Products: []*ProductSchema{},
  }

  for _, p := range products {
    tmp := &ProductSchema{}
    copyModel(&p, tmp)
    res.Products = append(res.Products, tmp)
  }

  testutils.Log(fmt.Sprintf("Response:\n%#v", products))

  return res, nil
}
