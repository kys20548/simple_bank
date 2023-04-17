package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/kys20548/simple_bank/db/mock"
	db "github.com/kys20548/simple_bank/db/sqlc"
	"github.com/kys20548/simple_bank/util"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func fixTestCreateAccount(t *testing.T) {
	user, _ := randomUser(t)
	account := randomAccount(user.Username)

	testCasts := []struct {
		name          string
		body          gin.H
		buildStubs    func(strore *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{{
		name: "OK",
		body: gin.H{
			"currency": account.Currency,
		},
		buildStubs: func(store *mockdb.MockStore) {
			arg := db.CreateAccountParams{
				Owner:    account.Owner,
				Currency: account.Currency,
				Balance:  0,
			}
			store.EXPECT().
				CreateAccount(gomock.Any(), gomock.Eq(arg)).
				Times(1).
				Return(account, nil)

		},
		checkResponse: func(recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusOK, recorder.Code)
			requireBodyMatchAccount(t, recorder.Body, account)
		},
	},
	}

	for i := range testCasts {
		tc := testCasts[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := "/accounts"
			request, err := http.NewRequest(http.MethodPost, url, nil)
			require.NoError(t, err)


			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}

}
func randomAccount(owner string) db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    owner,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}
