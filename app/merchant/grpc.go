// merchant/grpm.go

package merchant

import (
  "context"

  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
  wrappers "github.com/golang/protobuf/ptypes/wrappers"
  "github.com/jinzhu/gorm"

  "app/common"
)

// MerchantServiceInterface is implemented by MerchantService
type MerchantServiceInterface struct{
  app *common.App
}

// CreateMerchant implements MerchantService.CreateMerchant
func (s *MerchantServiceInterface) CreateMerchant(ctx context.Context, req *MerchantSchema) (*MerchantSchema, error) {
  m := Merchant{
    Schema: req,
    Model: nil,
  }

  m.copySchema()

  if err := m.Model.createMerchant(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  m.copyModel()

  return m.Schema, nil
}

// GetMerchant implements MerchantService.MerchantQuery
func (s *MerchantServiceInterface) GetMerchant(ctx context.Context, req *MerchantQuery) (*MerchantSchema, error) {
  m := Merchant{
    Model: &MerchantModel{ID: int(req.Id)},
    Schema: nil,
  }

  if err := m.Model.readMerchant(s.app.DB); err != nil {
    switch err {
    case gorm.ErrRecordNotFound:
      return nil, status.Error(codes.NotFound, "Merchant not found")
    default:
      return nil, status.Error(codes.Internal, err.Error())
    }
  }

  m.copyModel()

  return m.Schema, nil
}

// UpdateMerchant implements MerchantService.UpdateMerchant
func (s *MerchantServiceInterface) UpdateMerchant(ctx context.Context, req *MerchantQuery) (*MerchantSchema, error) {
  m := Merchant{
    Schema: req.Merchant,
    Model: &MerchantModel{ID: int(req.Id)},
  }

  if err := m.Model.readMerchant(s.app.DB); err != nil {
    switch err {
    case gorm.ErrRecordNotFound:
      return nil, status.Error(codes.NotFound, "Merchant not found")
    default:
      return nil, status. Error(codes.Internal, err.Error())
    }
  }

  m.copySchema()

  if err := m.Model.updateMerchant(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  m.copyModel()

  return m.Schema, nil
}

// DeleteMerchant implements MerchantService.DeleteMerchant
func (s *MerchantServiceInterface) DeleteMerchant(ctx context.Context, req *MerchantQuery) (*wrappers.StringValue, error) {

  m := Merchant{
    Model: &MerchantModel{ID: int(req.Id)},
    Schema: nil,
  }
  if err := m.Model.deleteMerchant(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }


  res := &wrappers.StringValue{Value: "success"}
  return res, nil
}

// GetMerchants implements MerchantService.GetMerchants
func (s *MerchantServiceInterface) GetMerchants(ctx context.Context, req *MerchantsQuery) (*MerchantsResponse, error) {

  count, start := int(req.Count), int(req.Start)

  merchants, err := readMerchants(s.app.DB, start, count)
  if err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  res := &MerchantsResponse{
    Merchants: []*MerchantSchema{},
  }

  for _, m := range merchants {
    tmp := &MerchantSchema{}
    copyModel(&m, tmp)
    res.Merchants = append(res.Merchants, tmp)
  }

  return res, nil
}
