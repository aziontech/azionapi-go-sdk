/*
Intelligent DNS

Testing DNSSECApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package idns

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_idns_DNSSECApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DNSSECApiService GetZoneDnsSec", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var zoneId int32

		resp, httpRes, err := apiClient.DNSSECApi.GetZoneDnsSec(context.Background(), zoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DNSSECApiService PutZoneDnsSec", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var zoneId int32

		resp, httpRes, err := apiClient.DNSSECApi.PutZoneDnsSec(context.Background(), zoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
