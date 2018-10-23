package merchant

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

// MerchantServiceInterface is implemented by MerchantService
type MerchantServiceInterface struct{
  app *common.App
}

// CreateMerchant implements MerchantService.CreateMerchant
func (s *MerchantServiceInterface) CreateMerchant(ctx context.Context, req *MerchantSchema) (*MerchantSchema, error) {
  testutils.Log(fmt.Sprint("MerchantService.CreateMerchant"))
  c := Merchant{
    Schema: req,
    Model: nil,
  }

  c.copySchema()

  if err := c.Model.createMerchant(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  c.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))

  return c.Schema, nil
}

// GetMerchant implements MerchantService.MerchantRequest
func (s *MerchantServiceInterface) GetMerchant(ctx context.Context, req *MerchantRequest) (*MerchantSchema, error) {
  testutils.Log(fmt.Sprint("MerchantService.GetMerchant"))
  c := Merchant{
    Model: &MerchantModel{ID: int(req.Id)},
    Schema: nil,
  }

  if err := c.Model.readMerchant(s.app.DB); err != nil {
    switch err {
    case sql.ErrNoRows:
      return nil, status.Error(codes.NotFound, "Merchant not found")
    default:
      return nil, status.Error(codes.Internal, err.Error())
    }
  }

  c.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))

  return c.Schema, nil
}

// UpdateMerchant implements MerchantService.UpdateMerchant
func (s *MerchantServiceInterface) UpdateMerchant(ctx context.Context, req *MerchantRequest) (*MerchantSchema, error) {
  testutils.Log(fmt.Sprint("MerchantService.UpdateMerchant"))
  c := Merchant{
    Schema: req.Merchant,
    Model: &MerchantModel{ID: int(req.Id)},
  }

  if err := c.Model.readMerchant(s.app.DB); err != nil {
    switch err {
    case sql.ErrNoRows:
      return nil, status.Error(codes.NotFound, "Merchant not found")
    default:
      return nil, status. Error(codes.Internal, err.Error())
    }
  }

  c.copySchema()

  if err := c.Model.updateMerchant(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  c.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))
  return c.Schema, nil
}

// DeleteMerchant implements MerchantService.DeleteMerchant
func (s *MerchantServiceInterface) DeleteMerchant(ctx context.Context, req *MerchantRequest) (*wrappers.StringValue, error) {
  testutils.Log(fmt.Sprint("MerchantService.DeleteMerchant"))

  c := Merchant{
    Model: &MerchantModel{ID: int(req.Id)},
    Schema: nil,
  }
  if err := c.Model.deleteMerchant(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  testutils.Log(fmt.Sprint("Response:\n{ value: \"success\" }"))

  m := &wrappers.StringValue{Value: "success"}
  return m, nil
}

// GetMerchants implements MerchantService.GetMerchants
func (s *MerchantServiceInterface) GetMerchants(ctx context.Context, req *MerchantsRequest) (*MerchantsResponse, error) {
  testutils.Log(fmt.Sprint("MerchantService.GetMerchants"))

  count, start := int(req.Count), int(req.Start)

  if count > 10 || count < 1 {
    count = 10
  }
  if start < 0 {
    start = 0
  }

  testutils.Log(fmt.Sprintf("{ count: %d, start: %d }", count, start))

  merchants, err := readMerchants(s.app.DB, start, count)
  if err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  res := &MerchantsResponse{
    Merchants: []*MerchantSchema{},
  }

  for _, c := range merchants {
    tmp := &MerchantSchema{}
    copyModel(&c, tmp)
    res.Merchants = append(res.Merchants, tmp)
  }

  testutils.Log(fmt.Sprintf("Response:\n%#v", merchants))

  return res, nil
}
