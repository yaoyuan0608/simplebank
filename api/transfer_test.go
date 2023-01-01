package api

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// 	mockdb "github.com/techschool/simplebank/db/mock"
// 	"github.com/techschool/simplebank/util"
// )

// func TestTransferAPI(t *testing.T) {
// 	amount := int64(10)

// 	user1, _ := randomUser(t)
// 	user2, _ := randomUser(t)
// 	user3, _ := randomUser(t)

// 	account1 := randomAccount(user1.Username)
// 	account2 := randomAccount(user2.Username)
// 	account3 := randomAccount(user3.Username)

// 	account1.Currency = util.USD
// 	account2.Currency = util.USD
// 	account3.Currency = util.EUR

// 	testCases := []struct {
// 		name          string
// 		body          gin.H
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: gin.H{
// 				"from_account_id": account1.ID,
// 				"to_account_id":   account2.ID,
// 				"amount":          amount,
// 				"currency":        util.USD,
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				// build stubs
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.ID)).Times(1).Return(account1, nil)
// 				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.ID)).Times(1).Return(account2, nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				// check response
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 			},
// 		},
// 	}
// 	for i := range testCases {
// 		tc := testCases[i]
// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			store := mockdb.NewMockStore(ctrl)
// 			tc.buildStubs(store)

// 			// start test server
// 			server := newTestServer(t, store)
// 			recorder := httptest.NewRecorder()

// 			url := "/transfers"
// 			request, err := http.NewRequest(http.MethodGet, url, nil)
// 			require.NoError(t, err)

// 			server.router.ServeHTTP(recorder, request)
// 			tc.checkResponse(t, recorder)
// 		})
// 	}
// }
