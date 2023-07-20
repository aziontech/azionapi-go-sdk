/*
Personal Tokens - OpenAPI

Testing PersonalTokenApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package personal_tokens

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_personal_tokens_PersonalTokenApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PersonalTokenApiService CreatePersonalToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PersonalTokenApi.CreatePersonalToken(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonalTokenApiService DeletePersonalToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personalTokenId string

		httpRes, err := apiClient.PersonalTokenApi.DeletePersonalToken(context.Background(), personalTokenId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonalTokenApiService GetPersonalToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personalTokenId string

		resp, httpRes, err := apiClient.PersonalTokenApi.GetPersonalToken(context.Background(), personalTokenId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonalTokenApiService ListPersonalToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PersonalTokenApi.ListPersonalToken(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}