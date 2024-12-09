/*
Gr4vy API

Testing GiftCardServiceDefinitionsAPIService

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

func Test_openapi_GiftCardServiceDefinitionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GiftCardServiceDefinitionsAPIService GetGiftCardServiceDefinition", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var giftCardServiceDefinitionId string

		resp, httpRes, err := apiClient.GiftCardServiceDefinitionsAPI.GetGiftCardServiceDefinition(context.Background(), giftCardServiceDefinitionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
