/*
CrudSQL API

Testing DynamicAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_DynamicAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DynamicAPIService ModelFilterPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var model string

		resp, httpRes, err := apiClient.DynamicAPI.ModelFilterPost(context.Background(), model).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DynamicAPIService ModelGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var model string

		resp, httpRes, err := apiClient.DynamicAPI.ModelGet(context.Background(), model).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DynamicAPIService ModelIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var model string
		var id string

		resp, httpRes, err := apiClient.DynamicAPI.ModelIdDelete(context.Background(), model, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DynamicAPIService ModelIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var model string
		var id string

		resp, httpRes, err := apiClient.DynamicAPI.ModelIdGet(context.Background(), model, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DynamicAPIService ModelIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var model string
		var id string

		resp, httpRes, err := apiClient.DynamicAPI.ModelIdPut(context.Background(), model, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DynamicAPIService ModelPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var model string

		resp, httpRes, err := apiClient.DynamicAPI.ModelPost(context.Background(), model).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
