package session

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/supertokens/supertokens-golang/supertokens"
	"github.com/supertokens/supertokens-golang/test/unittesting"
)

func TestShouldNotChangeIfFetchValueReturnsNil(t *testing.T) {
	configValue := supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: "http://localhost:8080",
		},
		AppInfo: supertokens.AppInfo{
			AppName:       "SuperTokens",
			WebsiteDomain: "supertokens.io",
			APIDomain:     "api.supertokens.io",
		},
		RecipeList: []supertokens.Recipe{
			Init(nil),
		},
	}
	BeforeEach()
	unittesting.StartUpST("localhost", "8080")
	defer AfterEach()
	err := supertokens.Init(configValue)
	if err != nil {
		t.Error(err.Error())
	}

	res := fakeRes{}
	sessionContainer, err := CreateNewSession(res, "userId", map[string]interface{}{}, map[string]interface{}{})
	assert.NoError(t, err)

	err = sessionContainer.FetchAndSetClaim(NilClaim())
	assert.NoError(t, err)
	accessTokenPayload := sessionContainer.GetAccessTokenPayload()
	assert.Equal(t, 0, len(accessTokenPayload))
}

func TestShouldUpdateIfClaimFetchValueReturnsValue(t *testing.T) {
	configValue := supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: "http://localhost:8080",
		},
		AppInfo: supertokens.AppInfo{
			AppName:       "SuperTokens",
			WebsiteDomain: "supertokens.io",
			APIDomain:     "api.supertokens.io",
		},
		RecipeList: []supertokens.Recipe{
			Init(nil),
		},
	}
	BeforeEach()
	unittesting.StartUpST("localhost", "8080")
	defer AfterEach()
	err := supertokens.Init(configValue)
	if err != nil {
		t.Error(err.Error())
	}

	res := fakeRes{}
	sessionContainer, err := CreateNewSession(res, "userId", map[string]interface{}{}, map[string]interface{}{})
	assert.NoError(t, err)

	err = sessionContainer.FetchAndSetClaim(TrueClaim())
	assert.NoError(t, err)
	accessTokenPayload := sessionContainer.GetAccessTokenPayload()
	assert.Equal(t, 1, len(accessTokenPayload))
	assert.NotNil(t, accessTokenPayload["st-true"])
	assert.Equal(t, true, accessTokenPayload["st-true"].(map[string]interface{})["v"])
	assert.Greater(t, accessTokenPayload["st-true"].(map[string]interface{})["t"], float64(time.Now().UnixMilli()-1000))
}

func TestShouldUpdateUsingHandleIfClaimFetchValueReturnsValue(t *testing.T) {
	configValue := supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: "http://localhost:8080",
		},
		AppInfo: supertokens.AppInfo{
			AppName:       "SuperTokens",
			WebsiteDomain: "supertokens.io",
			APIDomain:     "api.supertokens.io",
		},
		RecipeList: []supertokens.Recipe{
			Init(nil),
		},
	}
	BeforeEach()
	unittesting.StartUpST("localhost", "8080")
	defer AfterEach()
	err := supertokens.Init(configValue)
	if err != nil {
		t.Error(err.Error())
	}

	res := fakeRes{}
	sessionContainer, err := CreateNewSession(res, "userId", map[string]interface{}{}, map[string]interface{}{})
	assert.NoError(t, err)

	ok, err := FetchAndSetClaim(sessionContainer.GetHandle(), TrueClaim())
	assert.NoError(t, err)
	assert.True(t, ok)

	sessInfo, err := GetSessionInformation(sessionContainer.GetHandle())
	assert.NoError(t, err)
	accessTokenPayload := sessInfo.AccessTokenPayload

	assert.Equal(t, 1, len(accessTokenPayload))
	assert.NotNil(t, accessTokenPayload["st-true"])
	assert.Equal(t, true, accessTokenPayload["st-true"].(map[string]interface{})["v"])
	assert.Greater(t, accessTokenPayload["st-true"].(map[string]interface{})["t"], float64(time.Now().UnixMilli()-1000))
}
