/*
Gr4vy API

Testing GiftCardsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_GiftCardsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GiftCardsAPIService CheckGiftCardBalances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GiftCardsAPI.CheckGiftCardBalances(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GiftCardsAPIService DeleteGiftCard", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var giftCardId string

		httpRes, err := apiClient.GiftCardsAPI.DeleteGiftCard(context.Background(), giftCardId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GiftCardsAPIService GetGiftCard", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var giftCardId string

		resp, httpRes, err := apiClient.GiftCardsAPI.GetGiftCard(context.Background(), giftCardId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GiftCardsAPIService ListBuyerGiftCards", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GiftCardsAPI.ListBuyerGiftCards(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GiftCardsAPIService ListGiftCards", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GiftCardsAPI.ListGiftCards(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GiftCardsAPIService StoreGiftCard", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GiftCardsAPI.StoreGiftCard(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
