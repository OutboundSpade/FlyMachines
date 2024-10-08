/*
Machines API

Testing AppsAPIService

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

func Test_machines_AppsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppsAPIService AppsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AppsAPI.AppsCreate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppsAPIService AppsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appName string

		httpRes, err := apiClient.AppsAPI.AppsDelete(context.Background(), appName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppsAPIService AppsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AppsAPI.AppsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppsAPIService AppsShow", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var appName string

		resp, httpRes, err := apiClient.AppsAPI.AppsShow(context.Background(), appName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
