/*
Edge Functions Instances API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package edgefunctionsinstance_edgefirewall

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_edgefunctionsinstance_edgefirewall_DefaultApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultApiService EdgeFirewallEdgeFirewallIdFunctionsInstancesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService EdgeFirewallEdgeFirewallIdFunctionsInstancesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		httpRes, err := apiClient.DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidDelete(context.Background(), uuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}