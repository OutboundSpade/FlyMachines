/*
Machines API

Testing SecretsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package machines

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/OutboundSpade/FlyMachines"
)

func Test_machines_SecretsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SecretsAPIService SecretCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecretsAPI.SecretCreate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecretsAPIService SecretDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appName string
		var secretLabel string

		httpRes, err := apiClient.SecretsAPI.SecretDelete(context.Background(), appName, secretLabel).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecretsAPIService SecretGenerate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecretsAPI.SecretGenerate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecretsAPIService SecretsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SecretsAPI.SecretsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
