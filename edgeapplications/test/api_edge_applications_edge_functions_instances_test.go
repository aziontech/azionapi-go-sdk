/*
Edge Application

Testing EdgeApplicationsEdgeFunctionsInstancesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package edgeapplications

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_edgeapplications_EdgeApplicationsEdgeFunctionsInstancesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EdgeApplicationsEdgeFunctionsInstancesApiService EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId string
		var functionsInstancesId string

		httpRes, err := apiClient.EdgeApplicationsEdgeFunctionsInstancesApi.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDelete(context.Background(), edgeApplicationId, functionsInstancesId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsEdgeFunctionsInstancesApiService EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId int64
		var functionsInstancesId int64

		resp, httpRes, err := apiClient.EdgeApplicationsEdgeFunctionsInstancesApi.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGet(context.Background(), edgeApplicationId, functionsInstancesId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsEdgeFunctionsInstancesApiService EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId string
		var functionsInstancesId string

		resp, httpRes, err := apiClient.EdgeApplicationsEdgeFunctionsInstancesApi.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatch(context.Background(), edgeApplicationId, functionsInstancesId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsEdgeFunctionsInstancesApiService EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId string
		var functionsInstancesId string

		resp, httpRes, err := apiClient.EdgeApplicationsEdgeFunctionsInstancesApi.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPut(context.Background(), edgeApplicationId, functionsInstancesId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsEdgeFunctionsInstancesApiService EdgeApplicationsEdgeApplicationIdFunctionsInstancesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId int64

		resp, httpRes, err := apiClient.EdgeApplicationsEdgeFunctionsInstancesApi.EdgeApplicationsEdgeApplicationIdFunctionsInstancesGet(context.Background(), edgeApplicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsEdgeFunctionsInstancesApiService EdgeApplicationsEdgeApplicationIdFunctionsInstancesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId int64

		resp, httpRes, err := apiClient.EdgeApplicationsEdgeFunctionsInstancesApi.EdgeApplicationsEdgeApplicationIdFunctionsInstancesPost(context.Background(), edgeApplicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
