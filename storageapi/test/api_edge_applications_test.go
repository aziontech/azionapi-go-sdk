/*
Edge Application Statics API

Testing EdgeApplicationsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package storageapi

import (
	"context"
	"testing"

	openapiclient "github.com/aziontech/azionapi-go-sdk/storageapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_storageapi_EdgeApplicationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EdgeApplicationsApiService CreateVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var edgeApplicationId string

		resp, httpRes, err := apiClient.EdgeApplicationsApi.CreateVersion(context.Background(), edgeApplicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsApiService EdgeApplicationsEdgeApplicationIdStaticsVersionIdFilesPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var edgeApplicationId string
		var versionId string

		resp, httpRes, err := apiClient.EdgeApplicationsApi.EdgeApplicationsEdgeApplicationIdStaticsVersionIdFilesPost(context.Background(), edgeApplicationId, versionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
