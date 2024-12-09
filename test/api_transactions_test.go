/*
Gr4vy API

Testing TransactionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/FuturHealth/gr4vy-go-sdk"
)

func Test_openapi_TransactionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TransactionsAPIService CaptureTransaction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.TransactionsAPI.CaptureTransaction(context.Background(), transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetRefund", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var transactionId string
		var refundId string

		resp, httpRes, err := apiClient.TransactionsAPI.GetRefund(context.Background(), transactionId, refundId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetSingleRefund", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var refundId string

		resp, httpRes, err := apiClient.TransactionsAPI.GetSingleRefund(context.Background(), refundId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetTransaction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.TransactionsAPI.GetTransaction(context.Background(), transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetTransactionSettlement", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var transactionId string
		var settlementId string

		resp, httpRes, err := apiClient.TransactionsAPI.GetTransactionSettlement(context.Background(), transactionId, settlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService GetTransactionSettlements", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.TransactionsAPI.GetTransactionSettlements(context.Background(), transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService ListTransactionRefunds", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.TransactionsAPI.ListTransactionRefunds(context.Background(), transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService ListTransactions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TransactionsAPI.ListTransactions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService NewRefund", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.TransactionsAPI.NewRefund(context.Background(), transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService NewTransaction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TransactionsAPI.NewTransaction(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService RefundAll", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.TransactionsAPI.RefundAll(context.Background(), transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService SyncTransaction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.TransactionsAPI.SyncTransaction(context.Background(), transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TransactionsAPIService VoidTransaction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.TransactionsAPI.VoidTransaction(context.Background(), transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
