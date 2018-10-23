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
  c := Product{
    Schema: req,
    Model: nil,
  }

  c.copySchema()

  if err := c.Model.createProduct(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  c.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))

  return c.Schema, nil
}

// GetProduct implements ProductService.ProductRequest
func (s *ProductServiceInterface) GetProduct(ctx context.Context, req *ProductRequest) (*ProductSchema, error) {
  testutils.Log(fmt.Sprint("ProductService.GetProduct"))
  c := Product{
    Model: &ProductModel{ID: int(req.Id)},
    Schema: nil,
  }

  if err := c.Model.readProduct(s.app.DB); err != nil {
    switch err {
    case sql.ErrNoRows:
      return nil, status.Error(codes.NotFound, "Product not found")
    default:
      return nil, status.Error(codes.Internal, err.Error())
    }
  }

  c.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))

  return c.Schema, nil
}

// UpdateProduct implements ProductService.UpdateProduct
func (s *ProductServiceInterface) UpdateProduct(ctx context.Context, req *ProductRequest) (*ProductSchema, error) {
  testutils.Log(fmt.Sprint("ProductService.UpdateProduct"))
  c := Product{
    Schema: req.Product,
    Model: &ProductModel{ID: int(req.Id)},
  }

  if err := c.Model.readProduct(s.app.DB); err != nil {
    switch err {
    case sql.ErrNoRows:
      return nil, status.Error(codes.NotFound, "Product not found")
    default:
      return nil, status. Error(codes.Internal, err.Error())
    }
  }

  c.copySchema()

  if err := c.Model.updateProduct(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  c.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))
  return c.Schema, nil
}

// DeleteProduct implements ProductService.DeleteProduct
func (s *ProductServiceInterface) DeleteProduct(ctx context.Context, req *ProductRequest) (*wrappers.StringValue, error) {
  testutils.Log(fmt.Sprint("ProductService.DeleteProduct"))

  c := Product{
    Model: &ProductModel{ID: int(req.Id)},
    Schema: nil,
  }
  if err := c.Model.deleteProduct(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  testutils.Log(fmt.Sprint("Response:\n{ value: \"success\" }"))

  m := &wrappers.StringValue{Value: "success"}
  return m, nil
}

// GetProducts implements ProductService.GetProducts
func (s *ProductServiceInterface) GetProducts(ctx context.Context, req *ProductsRequest) (*ProductsResponse, error) {
  testutils.Log(fmt.Sprint("ProductService.GetProducts"))

  count, start := int(req.Count), int(req.Start)

  if count > 10 || count < 1 {
    count = 10
  }
  if start < 0 {
    start = 0
  }

  testutils.Log(fmt.Sprintf("{ count: %d, start: %d }", count, start))

  products, err := readProducts(s.app.DB, start, count)
  if err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  res := &ProductsResponse{
    Products: []*ProductSchema{},
  }

  for _, c := range products {
    tmp := &ProductSchema{}
    copyModel(&c, tmp)
    res.Products = append(res.Products, tmp)
  }

  testutils.Log(fmt.Sprintf("Response:\n%#v", products))

  return res, nil
}
