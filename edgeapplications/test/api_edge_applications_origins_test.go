/*
Edge Application

Testing EdgeApplicationsOriginsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package edgeapplications

import (
	"context"
	"testing"

	openapiclient "github.com/aziontech/azionapi-go-sdk/edgeapplications"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_edgeapplications_EdgeApplicationsOriginsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EdgeApplicationsOriginsApiService EdgeApplicationsEdgeApplicationIdOriginsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var edgeApplicationId int64

		resp, httpRes, err := apiClient.EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsGet(context.Background(), edgeApplicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsOriginsApiService EdgeApplicationsEdgeApplicationIdOriginsOriginKeyDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var edgeApplicationId int64
		var originKey string

		httpRes, err := apiClient.EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyDelete(context.Background(), edgeApplicationId, originKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsOriginsApiService EdgeApplicationsEdgeApplicationIdOriginsOriginKeyGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var edgeApplicationId int64
		var originKey string

		resp, httpRes, err := apiClient.EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyGet(context.Background(), edgeApplicationId, originKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsOriginsApiService EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var edgeApplicationId int64
		var originKey string

		resp, httpRes, err := apiClient.EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatch(context.Background(), edgeApplicationId, originKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsOriginsApiService EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var edgeApplicationId int64
		var originKey string

		resp, httpRes, err := apiClient.EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPut(context.Background(), edgeApplicationId, originKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsOriginsApiService EdgeApplicationsEdgeApplicationIdOriginsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var edgeApplicationId int64

		resp, httpRes, err := apiClient.EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsPost(context.Background(), edgeApplicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}