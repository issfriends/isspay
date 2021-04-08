package httphandler

// import (
// 	"net/http"
// 	"testing"

// 	"github.com/issfriends/isspay/internal/app/model"
// 	"github.com/issfriends/isspay/internal/app/query"
// 	"github.com/issfriends/isspay/internal/test/testutil"
// 	"github.com/issfriends/isspay/pkg/factory"
// 	"github.com/stretchr/testify/suite"
// )

// func TestInventoryHandler(t *testing.T) {
// 	s := &inventorySuite{httpSuite: &httpSuite{}}
// 	suite.Run(t, s)
// }

// type inventorySuite struct {
// 	*httpSuite
// }

// func (s *inventorySuite) SetupTest() {
// 	s.TruncateTables("products")
// }

// func (s *inventorySuite) TestBatchCreateProducts() {
// 	tcs := []struct {
// 		name        string
// 		genProducts func() (valid, invalid []*model.Product)
// 	}{
// 		{
// 			"productsAllValid",
// 			func() (valid, invalid []*model.Product) {
// 				valid = factory.Product.Omit("UID").MustBuildN(10).([]*model.Product)
// 				return valid, []*model.Product{}
// 			},
// 		},
// 		{
// 			"partialValid",
// 			func() (valid, invalid []*model.Product) {
// 				valid = factory.Product.Omit("UID").MustBuildN(10).([]*model.Product)
// 				invalid = factory.Product.Omit("UID", "Price").MustBuildN(3).([]*model.Product)
// 				return valid, invalid
// 			},
// 		},
// 		{
// 			"duplicatedInvalid",
// 			func() (valid, invalid []*model.Product) {
// 				created := factory.Product.MustInsert().(*model.Product)
// 				valid = factory.Product.Omit("UID").MustBuildN(10).([]*model.Product)
// 				invalid = factory.Product.Name(created.Name).Omit("UID").MustBuildN(1).([]*model.Product)
// 				return valid, invalid
// 			},
// 		},
// 	}

// 	for _, tc := range tcs {
// 		s.SetupTest()
// 		s.Run(tc.name, func() {
// 			valid, invalid := tc.genProducts()
// 			products := append(valid, invalid...)
// 			productsReq := map[string][]*model.Product{
// 				"products": products,
// 			}
// 			req, resp := testutil.BuildRequest("POST", "/api/v1/products", productsReq)
// 			s.Serv.ServeHTTP(resp, req)
// 			result := resp.Result()
// 			names := []string{}
// 			for _, p := range valid {
// 				names = append(names, p.Name)
// 			}
// 			q := &query.ListProductsQuery{
// 				Names: names,
// 			}
// 			_, err := s.Database.Inventory().ListProducts(s.Ctx, q)

// 			if len(invalid) == 0 {
// 				s.Equal(http.StatusCreated, result.StatusCode)
// 			} else {
// 				s.NotEqual(http.StatusCreated, result.StatusCode)
// 				// s.Equal(resperr.ToErrResponse(goerr.ErrParitalUnprocessable).HTTPStatus, result.StatusCode)
// 			}

// 			s.Require().NoError(err)
// 			s.Len(q.Data, len(valid))
// 			s.AssertHelper.AssertProductsEq(q.Data, valid, true)
// 		})
// 	}
// }

// func (s *inventorySuite) TestUpdateProduct() {
// 	tcs := []struct {
// 		name     string
// 		updateTo func() (old, new *model.Product)
// 		hasErr   bool
// 	}{
// 		{
// 			"updatePriceAndQuantity",
// 			func() (old, new *model.Product) {
// 				old = factory.Product.MustInsert().(*model.Product)
// 				new = &model.Product{}
// 				new.Name = old.Name
// 				new.Price = old.Price.Add(old.Cost)
// 				new.Quantity = 12
// 				new.Cost = old.Cost
// 				new.ImageURL = old.ImageURL
// 				return old, new
// 			},
// 			false,
// 		},
// 	}

// 	for _, tc := range tcs {
// 		s.SetupTest()
// 		s.Run(tc.name, func() {
// 			old, updatedProduct := tc.updateTo()
// 			req, resp := testutil.BuildRequest("PUT", "/api/v1/products/"+old.UID, updatedProduct)
// 			s.Serv.ServeHTTP(resp, req)
// 			res := resp.Result()
// 			q := query.GetProductQuery{
// 				ID: old.ID,
// 			}

// 			err := s.Database.Inventory().GetProduct(s.Ctx, &q)
// 			s.Require().NoError(err)
// 			new := q.Data
// 			s.Require().NotNil(new)

// 			if tc.hasErr {
// 				s.NotEqual(res.StatusCode, http.StatusOK)
// 				s.Equal(new.Name, old.Name)
// 				s.Equal(new.ImageURL, old.ImageURL)
// 				s.True(new.Price.Round(2).Equal(old.Price.Round(2)))
// 				s.True(new.Cost.Round(2).Equal(old.Cost.Round(2)))
// 				s.Equal(new.Quantity, old.Quantity)
// 			} else {
// 				s.Equal(res.StatusCode, http.StatusOK)
// 				s.Equal(new.Name, updatedProduct.Name)
// 				s.Equal(new.ImageURL, updatedProduct.ImageURL)
// 				s.True(new.Price.Round(2).Equal(updatedProduct.Price.Round(2)))
// 				s.True(new.Cost.Round(2).Equal(updatedProduct.Cost.Round(2)))
// 				s.Equal(new.Quantity, updatedProduct.Quantity)
// 			}
// 		})
// 	}
// }

// func (s *inventorySuite) TestListProducts() {
// 	tcs := []struct {
// 		name     string
// 		getQuery func() (q string, expected []*model.Product)
// 	}{}

// 	for _, tc := range tcs {
// 		s.SetupTest()
// 		s.Run(tc.name, func() {
// 			q, expected := tc.getQuery()
// 			req, resp := testutil.BuildRequest("GET", "/api/v1/products?"+q, nil)
// 			actual := make(map[string][]*model.Product)
// 			s.Serv.ServeHTTP(resp, req)
// 			err := testutil.GetResponseData(resp, actual)
// 			s.Require().NoError(err)
// 			s.AssertHelper.AssertProductsEq(actual["data"], expected, true)
// 		})
// 	}
// }

// func (s *inventorySuite) TestDeleteProduct() {

// }
