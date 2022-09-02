package session

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
	"github.com/supertokens/supertokens-golang/test/unittesting"
)

func TestGetClaimValueReturnsRightValue(t *testing.T) {
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
			Init(&sessmodels.TypeInput{
				Override: &sessmodels.OverrideStruct{
					Functions: func(originalImplementation sessmodels.RecipeInterface) sessmodels.RecipeInterface {
						oCreateNewSession := *originalImplementation.CreateNewSession
						nCreateNewSession := func(res http.ResponseWriter, userID string, accessTokenPayload map[string]interface{}, sessionData map[string]interface{}, userContext supertokens.UserContext) (sessmodels.SessionContainer, error) {
							accessTokenPayloadUpdate, err := TrueClaim().Build(userID, userContext)
							if err != nil {
								return sessmodels.SessionContainer{}, err
							}
							for k, v := range accessTokenPayloadUpdate {
								accessTokenPayload[k] = v
							}
							return oCreateNewSession(res, userID, accessTokenPayload, sessionData, userContext)
						}
						*originalImplementation.CreateNewSession = nCreateNewSession
						return originalImplementation
					},
				},
			}),
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

	assert.True(t, sessionContainer.GetClaimValue(TrueClaim()).(bool))
}

func TestGetClaimValueFromHandleReturnsRightValue(t *testing.T) {
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
			Init(&sessmodels.TypeInput{
				Override: &sessmodels.OverrideStruct{
					Functions: func(originalImplementation sessmodels.RecipeInterface) sessmodels.RecipeInterface {
						oCreateNewSession := *originalImplementation.CreateNewSession
						nCreateNewSession := func(res http.ResponseWriter, userID string, accessTokenPayload map[string]interface{}, sessionData map[string]interface{}, userContext supertokens.UserContext) (sessmodels.SessionContainer, error) {
							accessTokenPayloadUpdate, err := TrueClaim().Build(userID, userContext)
							if err != nil {
								return sessmodels.SessionContainer{}, err
							}
							for k, v := range accessTokenPayloadUpdate {
								accessTokenPayload[k] = v
							}
							return oCreateNewSession(res, userID, accessTokenPayload, sessionData, userContext)
						}
						*originalImplementation.CreateNewSession = nCreateNewSession
						return originalImplementation
					},
				},
			}),
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

	getRes, err := GetClaimValue(sessionContainer.GetHandle(), TrueClaim())
	assert.NoError(t, err)
	assert.NotNil(t, getRes.OK)
	assert.True(t, getRes.OK.Value.(bool))
}

func TestGetClaimValueForNonExistantSessionhandle(t *testing.T) {
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
			Init(&sessmodels.TypeInput{
				Override: &sessmodels.OverrideStruct{
					Functions: func(originalImplementation sessmodels.RecipeInterface) sessmodels.RecipeInterface {
						oCreateNewSession := *originalImplementation.CreateNewSession
						nCreateNewSession := func(res http.ResponseWriter, userID string, accessTokenPayload map[string]interface{}, sessionData map[string]interface{}, userContext supertokens.UserContext) (sessmodels.SessionContainer, error) {
							accessTokenPayloadUpdate, err := TrueClaim().Build(userID, userContext)
							if err != nil {
								return sessmodels.SessionContainer{}, err
							}
							for k, v := range accessTokenPayloadUpdate {
								accessTokenPayload[k] = v
							}
							return oCreateNewSession(res, userID, accessTokenPayload, sessionData, userContext)
						}
						*originalImplementation.CreateNewSession = nCreateNewSession
						return originalImplementation
					},
				},
			}),
		},
	}
	BeforeEach()
	unittesting.StartUpST("localhost", "8080")
	defer AfterEach()
	err := supertokens.Init(configValue)
	if err != nil {
		t.Error(err.Error())
	}

	getRes, err := GetClaimValue("invalidSessionHandle", TrueClaim())
	assert.NoError(t, err)
	assert.Nil(t, getRes.OK)
	assert.NotNil(t, getRes.SessionDoesNotExistError)
}
