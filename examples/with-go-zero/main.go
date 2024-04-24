package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/supertokens/supertokens-golang/recipe/dashboard"

	"github.com/supertokens/supertokens-golang/recipe/emailverification"
	"github.com/supertokens/supertokens-golang/recipe/emailverification/evmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
	"github.com/zeromicro/go-zero/rest"
)

func main() {
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: "https://try.supertokens.io",
		},
		AppInfo: supertokens.AppInfo{
			AppName:       "SuperTokens Demo App",
			APIDomain:     "http://localhost:3001",
			WebsiteDomain: "http://localhost:3000",
		},
		RecipeList: []supertokens.Recipe{
			emailverification.Init(evmodels.TypeInput{
				Mode: evmodels.ModeRequired,
			}),
			thirdparty.Init(&tpmodels.TypeInput{
				/*
				   We use different credentials for different platforms when required. For example the redirect URI for Github
				   is different for Web and mobile. In such a case we can provide multiple providers with different client Ids.
				   When the frontend makes a request and wants to use a specific clientId, it needs to send the clientId to use in the
				   request. In the absence of a clientId in the request the SDK uses the default provider, indicated by `isDefault: true`.
				   When adding multiple providers for the same type (Google, Github etc), make sure to set `isDefault: true`.
				*/
				SignInAndUpFeature: tpmodels.TypeInputSignInAndUp{
					Providers: []tpmodels.ProviderInput{
						// We have provided you with development keys which you can use for testsing.
						// IMPORTANT: Please replace them with your own OAuth keys for production use.
						{
							Config: tpmodels.ProviderConfig{
								ThirdPartyId: "google",
								Clients: []tpmodels.ProviderClientConfig{
									{
										ClientType:   "web",
										ClientID:     "1060725074195-kmeum4crr01uirfl2op9kd5acmi9jutn.apps.googleusercontent.com",
										ClientSecret: "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",
									},
									{
										// we use this for mobile apps
										ClientType:   "mobile",
										ClientID:     "1060725074195-c7mgk8p0h27c4428prfuo3lg7ould5o7.apps.googleusercontent.com",
										ClientSecret: "", // this is empty because we follow Authorization code grant flow via PKCE for mobile apps (Google doesn't issue a client secret for mobile apps).
									},
								},
							},
						},
						{
							Config: tpmodels.ProviderConfig{
								ThirdPartyId: "github",
								Clients: []tpmodels.ProviderClientConfig{
									{
										ClientType:   "web",
										ClientID:     "467101b197249757c71f",
										ClientSecret: "e97051221f4b6426e8fe8d51486396703012f5bd",
									},
									{
										// We use this for mobile apps
										ClientType:   "mobile",
										ClientID:     "8a9152860ce869b64c44",
										ClientSecret: "00e841f10f288363cd3786b1b1f538f05cfdbda2",
									},
								},
							},
						},
						/*
						   For Apple signin, iOS apps always use the bundle identifier as the client ID when communicating with Apple. Android, Web and other platforms
						   need to configure a Service ID on the Apple developer dashboard and use that as client ID.
						   In the example below 4398792-io.supertokens.example.service is the client ID for Web. Android etc and thus we mark it as default. For iOS
						   the frontend for the demo app sends the clientId in the request which is then used by the SDK.
						*/
						{
							Config: tpmodels.ProviderConfig{
								ThirdPartyId: "apple",
								Clients: []tpmodels.ProviderClientConfig{
									{
										// For Android and website apps
										ClientType: "web",
										ClientID:   "4398792-io.supertokens.example.service",
										AdditionalConfig: map[string]interface{}{
											"keyId":      "7M48Y4RYDL",
											"privateKey": "-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgu8gXs+XYkqXD6Ala9Sf/iJXzhbwcoG5dMh1OonpdJUmgCgYIKoZIzj0DAQehRANCAASfrvlFbFCYqn3I2zeknYXLwtH30JuOKestDbSfZYxZNMqhF/OzdZFTV0zc5u5s3eN+oCWbnvl0hM+9IW0UlkdA\n-----END PRIVATE KEY-----",
											"teamId":     "YWQCXGJRJL",
										},
									},
									{
										// For iOS Apps
										ClientType: "ios",
										ClientID:   "4398792-io.supertokens.example",
										AdditionalConfig: map[string]interface{}{
											"keyId":      "7M48Y4RYDL",
											"privateKey": "-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgu8gXs+XYkqXD6Ala9Sf/iJXzhbwcoG5dMh1OonpdJUmgCgYIKoZIzj0DAQehRANCAASfrvlFbFCYqn3I2zeknYXLwtH30JuOKestDbSfZYxZNMqhF/OzdZFTV0zc5u5s3eN+oCWbnvl0hM+9IW0UlkdA\n-----END PRIVATE KEY-----",
											"teamId":     "YWQCXGJRJL",
										},
									},
								},
							},
						},
						{
							Config: tpmodels.ProviderConfig{
								ThirdPartyId: "discord",
								Clients: []tpmodels.ProviderClientConfig{
									{
										ClientType:   "web",
										ClientID:     "4398792-907871294886928395",
										ClientSecret: "His4yXGEovVp5TZkZhEAt0ZXGh8uOVDm",
									},
									{
										// We use this for mobile apps
										ClientType:   "mobile",
										ClientID:     "4398792-907871294886928395",
										ClientSecret: "His4yXGEovVp5TZkZhEAt0ZXGh8uOVDm",
									},
								},
							},
						},
						{
							Config: tpmodels.ProviderConfig{
								ThirdPartyId: "google-workspaces",
								Clients: []tpmodels.ProviderClientConfig{
									{
										ClientType:   "web",
										ClientID:     "1060725074195-kmeum4crr01uirfl2op9kd5acmi9jutn.apps.googleusercontent.com",
										ClientSecret: "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",
									},
									{
										// We use this for mobile apps
										ClientType:   "mobile",
										ClientID:     "1060725074195-kmeum4crr01uirfl2op9kd5acmi9jutn.apps.googleusercontent.com",
										ClientSecret: "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",
									},
								},
							},
						},
					},
				},
			}),
			session.Init(nil),
			dashboard.Init(nil),
		},
	})
	if err != nil {
		log.Fatal("Something went wrong while starting up supertokens: ", err.Error())
	}
	server := rest.MustNewServer(
		rest.RestConf{
			Host: "0.0.0.0",
			Port: 3001,
		},
	)
	defer server.Stop()

	server.Use(corsMiddleware)
	server.Use(rest.ToMiddleware(supertokens.Middleware))
	addSupertokensRoutes("/auth", server) // go-zero doesn't execute middlewares if matching routes are not found
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/sessioninfo",
		Handler: func(rw http.ResponseWriter, r *http.Request) {
			session.VerifySession(nil, sessioninfo).ServeHTTP(rw, r)
		},
	})

	// go-zero does not honour middlewares for CORS requests unless we explicitly add a route for it.
	server.AddRoute(rest.Route{
		Method: http.MethodOptions,
		Path:   "/sessioninfo",
		Handler: func(rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(http.StatusOK)
		},
	})
	server.Start()
}

func sessioninfo(w http.ResponseWriter, r *http.Request) {
	sessionContainer := session.GetSessionFromRequestContext(r.Context())

	if sessionContainer == nil {
		w.WriteHeader(500)
		w.Write([]byte("no session found"))
		return
	}
	sessionData, err := sessionContainer.GetSessionDataInDatabase()
	if err != nil {
		err = supertokens.ErrorHandler(err, r, w)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}
		return
	}
	w.WriteHeader(200)
	w.Header().Add("content-type", "application/json")
	bytes, err := json.Marshal(map[string]interface{}{
		"sessionHandle":      sessionContainer.GetHandle(),
		"userId":             sessionContainer.GetUserID(),
		"accessTokenPayload": sessionContainer.GetAccessTokenPayload(),
		"sessionData":        sessionData,
	})
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("error in converting to json"))
	} else {
		w.Write(bytes)
	}
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, r *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		response.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			response.Header().Set("Access-Control-Allow-Headers", strings.Join(append([]string{"Content-Type"}, supertokens.GetAllCORSHeaders()...), ","))
			response.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,HEAD,OPTIONS")
		} else {
			next.ServeHTTP(response, r)
		}
	}
}

func addSupertokensRoutes(prefix string, server *rest.Server) {
	methods := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions}
	routes := []string{
		prefix,
		prefix + "/:a",
		prefix + "/:a/:b",
		prefix + "/:a/:b/:c",
		prefix + "/:a/:b/:c/:d",
		prefix + "/:a/:b/:c/:d/:e",
		prefix + "/:a/:b/:c/:d/:e/:f",
	}

	for _, method := range methods {
		for _, route := range routes {
			server.AddRoute(rest.Route{
				Method: method,
				Path:   route,
				Handler: func(w http.ResponseWriter, r *http.Request) {
					http.NotFound(w, r) // Return not found if not handled by middleware
				},
			})
		}
	}
}
