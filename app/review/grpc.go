package review

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

// ReviewServiceInterface is implemented by ReviewService
type ReviewServiceInterface struct{
  app *common.App
}

// CreateReview implements ReviewService.CreateReview
func (s *ReviewServiceInterface) CreateReview(ctx context.Context, req *ReviewSchema) (*ReviewSchema, error) {
  testutils.Log(fmt.Sprint("ReviewService.CreateReview"))
  c := Review{
    Schema: req,
    Model: nil,
  }

  c.copySchema()

  if err := c.Model.createReview(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  c.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))

  return c.Schema, nil
}

// GetReview implements ReviewService.ReviewRequest
func (s *ReviewServiceInterface) GetReview(ctx context.Context, req *ReviewRequest) (*ReviewSchema, error) {
  testutils.Log(fmt.Sprint("ReviewService.GetReview"))
  c := Review{
    Model: &ReviewModel{ID: int(req.Id)},
    Schema: nil,
  }

  if err := c.Model.readReview(s.app.DB); err != nil {
    switch err {
    case sql.ErrNoRows:
      return nil, status.Error(codes.NotFound, "Review not found")
    default:
      return nil, status.Error(codes.Internal, err.Error())
    }
  }

  c.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))

  return c.Schema, nil
}

// UpdateReview implements ReviewService.UpdateReview
func (s *ReviewServiceInterface) UpdateReview(ctx context.Context, req *ReviewRequest) (*ReviewSchema, error) {
  testutils.Log(fmt.Sprint("ReviewService.UpdateReview"))
  c := Review{
    Schema: req.Review,
    Model: &ReviewModel{ID: int(req.Id)},
  }

  if err := c.Model.readReview(s.app.DB); err != nil {
    switch err {
    case sql.ErrNoRows:
      return nil, status.Error(codes.NotFound, "Review not found")
    default:
      return nil, status. Error(codes.Internal, err.Error())
    }
  }

  c.copySchema()

  if err := c.Model.updateReview(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  c.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))
  return c.Schema, nil
}

// DeleteReview implements ReviewService.DeleteReview
func (s *ReviewServiceInterface) DeleteReview(ctx context.Context, req *ReviewRequest) (*wrappers.StringValue, error) {
  testutils.Log(fmt.Sprint("ReviewService.DeleteReview"))

  c := Review{
    Model: &ReviewModel{ID: int(req.Id)},
    Schema: nil,
  }
  if err := c.Model.deleteReview(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  testutils.Log(fmt.Sprint("Response:\n{ value: \"success\" }"))

  m := &wrappers.StringValue{Value: "success"}
  return m, nil
}

// GetReviews implements ReviewService.GetReviews
func (s *ReviewServiceInterface) GetReviews(ctx context.Context, req *ReviewsRequest) (*ReviewsResponse, error) {
  testutils.Log(fmt.Sprint("ReviewService.GetReviews"))

  count, start := int(req.Count), int(req.Start)

  if count > 10 || count < 1 {
    count = 10
  }
  if start < 0 {
    start = 0
  }

  testutils.Log(fmt.Sprintf("{ count: %d, start: %d }", count, start))

  reviews, err := readReviews(s.app.DB, start, count)
  if err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  res := &ReviewsResponse{
    Reviews: []*ReviewSchema{},
  }

  for _, c := range reviews {
    tmp := &ReviewSchema{}
    copyModel(&c, tmp)
    res.Reviews = append(res.Reviews, tmp)
  }

  testutils.Log(fmt.Sprintf("Response:\n%#v", reviews))

  return res, nil
}
