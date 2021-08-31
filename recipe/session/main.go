package session

import (
	"net/http"

	"github.com/supertokens/supertokens-golang/recipe/session/api"
	"github.com/supertokens/supertokens-golang/recipe/session/models"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func SessionInit(config *models.TypeInput) supertokens.RecipeListFunction {
	return recipeInit(config)
}

func CreateNewSession(res http.ResponseWriter, userID string, jwtPayload interface{}, sessionData interface{}) (models.SessionContainer, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return models.SessionContainer{}, err
	}
	return instance.RecipeImpl.CreateNewSession(res, userID, jwtPayload, sessionData)
}

func GetSession(req *http.Request, res http.ResponseWriter, options *models.VerifySessionOptions) (*models.SessionContainer, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return nil, err
	}
	return instance.RecipeImpl.GetSession(req, res, options)
}

func GetSessionInformation(sessionHandle string) (models.SessionInformation, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return models.SessionInformation{}, err
	}
	return instance.RecipeImpl.GetSessionInformation(sessionHandle)
}

func RefreshSession(req *http.Request, res http.ResponseWriter) (models.SessionContainer, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return models.SessionContainer{}, err
	}
	return instance.RecipeImpl.RefreshSession(req, res)
}

func RevokeAllSessionsForUser(userID string) ([]string, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return nil, err
	}
	return instance.RecipeImpl.RevokeAllSessionsForUser(userID)
}

func GetAllSessionHandlesForUser(userID string) ([]string, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return nil, err
	}
	return instance.RecipeImpl.GetAllSessionHandlesForUser(userID)
}

func RevokeSession(sessionHandle string) (bool, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return false, err
	}
	return instance.RecipeImpl.RevokeSession(sessionHandle)
}

func RevokeMultipleSessions(sessionHandles []string) ([]string, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return nil, err
	}
	return instance.RecipeImpl.RevokeMultipleSessions(sessionHandles)
}

func UpdateSessionData(sessionHandle string, newSessionData interface{}) error {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return err
	}
	return instance.RecipeImpl.UpdateSessionData(sessionHandle, newSessionData)
}

func UpdateJWTPayload(sessionHandle string, newJWTPayload interface{}) error {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return err
	}
	return instance.RecipeImpl.UpdateJWTPayload(sessionHandle, newJWTPayload)
}

func VerifySession(options *models.VerifySessionOptions) func(w http.ResponseWriter, r *http.Request, otherHandler http.HandlerFunc) error {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		panic("can't fetch instance")
	}
	return api.VerifySession(*instance, options)
}
