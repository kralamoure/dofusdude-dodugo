/*
dofusdude

Testing SetsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dodugo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/dofusdude/dodugo"
)

func Test_dodugo_SetsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SetsAPIService GetAllSetsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string
		var game string

		resp, httpRes, err := apiClient.SetsAPI.GetAllSetsList(context.Background(), language, game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SetsAPIService GetSetsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string
		var game string

		resp, httpRes, err := apiClient.SetsAPI.GetSetsList(context.Background(), language, game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SetsAPIService GetSetsSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string
		var game string

		resp, httpRes, err := apiClient.SetsAPI.GetSetsSearch(context.Background(), language, game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SetsAPIService GetSetsSingle", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string
		var ankamaId int32
		var game string

		resp, httpRes, err := apiClient.SetsAPI.GetSetsSingle(context.Background(), language, ankamaId, game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
