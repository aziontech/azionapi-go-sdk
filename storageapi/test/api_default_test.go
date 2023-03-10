/*
Storage API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package storageapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_storageapi_DefaultApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultApiService DeleteVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var versionId string

		httpRes, err := apiClient.DefaultApi.DeleteVersion(context.Background(), versionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService StorageVersionIdPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var versionId string

		resp, httpRes, err := apiClient.DefaultApi.StorageVersionIdPost(context.Background(), versionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}