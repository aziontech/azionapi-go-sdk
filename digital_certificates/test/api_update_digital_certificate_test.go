/*
Digital Certificates API

Testing UpdateDigitalCertificateApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package digital_certificates

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_digital_certificates_UpdateDigitalCertificateApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UpdateDigitalCertificateApiService UpdateDigitalCertificate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var digitalCertificateId int32

		resp, httpRes, err := apiClient.UpdateDigitalCertificateApi.UpdateDigitalCertificate(context.Background(), digitalCertificateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}