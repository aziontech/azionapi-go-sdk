/*
Intelligent DNS

Testing ZonesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package idns

import (
	"context"
	"testing"

	openapiclient "github.com/aziontech/azionapi-go-sdk/idns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_idns_ZonesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ZonesApiService DeleteZone", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var zoneId int32

		resp, httpRes, err := apiClient.ZonesApi.DeleteZone(context.Background(), zoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ZonesApiService GetZone", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var zoneId int32

		resp, httpRes, err := apiClient.ZonesApi.GetZone(context.Background(), zoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ZonesApiService GetZones", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ZonesApi.GetZones(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ZonesApiService PostZone", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ZonesApi.PostZone(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ZonesApiService PutZone", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var zoneId int32

		resp, httpRes, err := apiClient.ZonesApi.PutZone(context.Background(), zoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}