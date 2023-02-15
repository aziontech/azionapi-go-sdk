/*
Edge Application

Testing EdgeApplicationsMainSettingsApiService

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

func Test_edgeapplications_EdgeApplicationsMainSettingsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EdgeApplicationsMainSettingsApiService EdgeApplicationsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EdgeApplicationsMainSettingsApi.EdgeApplicationsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsMainSettingsApiService EdgeApplicationsIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.EdgeApplicationsMainSettingsApi.EdgeApplicationsIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsMainSettingsApiService EdgeApplicationsIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.EdgeApplicationsMainSettingsApi.EdgeApplicationsIdGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsMainSettingsApiService EdgeApplicationsIdPatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.EdgeApplicationsMainSettingsApi.EdgeApplicationsIdPatch(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsMainSettingsApiService EdgeApplicationsIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.EdgeApplicationsMainSettingsApi.EdgeApplicationsIdPut(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeApplicationsMainSettingsApiService EdgeApplicationsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EdgeApplicationsMainSettingsApi.EdgeApplicationsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
