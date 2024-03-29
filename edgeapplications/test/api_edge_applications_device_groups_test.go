/*
Edge Application

Testing EdgeApplicationsDeviceGroupsApiService

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

func Test_edgeapplications_EdgeApplicationsDeviceGroupsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EdgeApplicationsDeviceGroupsApiService EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId int64
		var deviceGroupId int64

		httpRes, err := apiClient.EdgeApplicationsDeviceGroupsApi.EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdDelete(context.Background(), edgeApplicationId, deviceGroupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsDeviceGroupsApiService EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId int64
		var deviceGroupId int64

		resp, httpRes, err := apiClient.EdgeApplicationsDeviceGroupsApi.EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdGet(context.Background(), edgeApplicationId, deviceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsDeviceGroupsApiService EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId int64
		var deviceGroupId int64

		resp, httpRes, err := apiClient.EdgeApplicationsDeviceGroupsApi.EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPatch(context.Background(), edgeApplicationId, deviceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsDeviceGroupsApiService EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId int64
		var deviceGroupId int64

		resp, httpRes, err := apiClient.EdgeApplicationsDeviceGroupsApi.EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPut(context.Background(), edgeApplicationId, deviceGroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsDeviceGroupsApiService EdgeApplicationsEdgeApplicationIdDeviceGroupsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId int64

		resp, httpRes, err := apiClient.EdgeApplicationsDeviceGroupsApi.EdgeApplicationsEdgeApplicationIdDeviceGroupsGet(context.Background(), edgeApplicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsDeviceGroupsApiService EdgeApplicationsEdgeApplicationIdDeviceGroupsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeApplicationId int64

		resp, httpRes, err := apiClient.EdgeApplicationsDeviceGroupsApi.EdgeApplicationsEdgeApplicationIdDeviceGroupsPost(context.Background(), edgeApplicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
