/*
Data Streaming - OpenAPI

Testing DataStreamingTemplatesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package data_streaming

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_data_streaming_DataStreamingTemplatesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DataStreamingTemplatesApiService GetDataStramingTemplateById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var templateId int32

		resp, httpRes, err := apiClient.DataStreamingTemplatesApi.GetDataStramingTemplateById(context.Background(), templateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataStreamingTemplatesApiService ListDataStreamingTemplates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DataStreamingTemplatesApi.ListDataStreamingTemplates(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}