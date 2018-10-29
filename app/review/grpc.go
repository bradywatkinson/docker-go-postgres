package review

import (
  "database/sql"
  "context"
  "fmt"

  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
  wrappers "github.com/golang/protobuf/ptypes/wrappers"
  log "github.com/sirupsen/logrus"

  "app/common"
)

// ReviewServiceInterface is implemented by ReviewService
type ReviewServiceInterface struct{
  app *common.App
}

// CreateReview implements ReviewService.CreateReview
func (s *ReviewServiceInterface) CreateReview(ctx context.Context, req *ReviewSchema) (*ReviewSchema, error) {
  r := Review{
    Schema: req,
    Model: nil,
  }

  log.Debug(fmt.Sprintf("Review: %#v", r.Schema))

  r.copySchema()

  if err := r.Model.createReview(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  r.copyModel()

  return r.Schema, nil
}

// GetReview implements ReviewService.ReviewQuery
func (s *ReviewServiceInterface) GetReview(ctx context.Context, req *ReviewQuery) (*ReviewSchema, error) {
  r := Review{
    Model: &ReviewModel{ID: int(req.Id)},
    Schema: nil,
  }

  if err := r.Model.readReview(s.app.DB); err != nil {
    switch err {
    case sql.ErrNoRows:
      return nil, status.Error(codes.NotFound, "Review not found")
    default:
      return nil, status.Error(codes.Internal, err.Error())
    }
  }

  r.copyModel()

  return r.Schema, nil
}

// UpdateReview implements ReviewService.UpdateReview
func (s *ReviewServiceInterface) UpdateReview(ctx context.Context, req *ReviewQuery) (*ReviewSchema, error) {
  r := Review{
    Schema: req.Review,
    Model: &ReviewModel{ID: int(req.Id)},
  }

  if err := r.Model.readReview(s.app.DB); err != nil {
    switch err {
    case sql.ErrNoRows:
      return nil, status.Error(codes.NotFound, "Review not found")
    default:
      return nil, status. Error(codes.Internal, err.Error())
    }
  }

  r.copySchema()

  if err := r.Model.updateReview(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  r.copyModel()

  log.Debug(fmt.Sprintf("Response:\n%#v", r.Schema))
  return r.Schema, nil
}

// DeleteReview implements ReviewService.DeleteReview
func (s *ReviewServiceInterface) DeleteReview(ctx context.Context, req *ReviewQuery) (*wrappers.StringValue, error) {

  r := Review{
    Model: &ReviewModel{ID: int(req.Id)},
    Schema: nil,
  }
  if err := r.Model.deleteReview(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }


  m := &wrappers.StringValue{Value: "success"}
  return m, nil
}

// GetReviews implements ReviewService.GetReviews
func (s *ReviewServiceInterface) GetReviews(ctx context.Context, req *ReviewsQuery) (*ReviewsResponse, error) {

  count, start := int(req.Count), int(req.Start)

  log.Debug(fmt.Sprintf("{ count: %d, start: %d }", count, start))

  reviews, err := readReviews(s.app.DB, start, count)
  if err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  res := &ReviewsResponse{
    Reviews: []*ReviewSchema{},
  }

  for _, r := range reviews {
    tmp := &ReviewSchema{}
    copyModel(&r, tmp)
    res.Reviews = append(res.Reviews, tmp)
  }

  return res, nil
}
