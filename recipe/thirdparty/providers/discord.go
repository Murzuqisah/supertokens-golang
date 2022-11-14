/* Copyright (c) 2021, VRAI Labs and/or its affiliates. All rights reserved.
 *
 * This software is licensed under the Apache License, Version 2.0 (the
 * "License") as published by the Apache Software Foundation.
 *
 * You may not use this file except in compliance with the License. You may
 * obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package providers

import (
	"fmt"

	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

const discordID = "discord"

func Discord(input tpmodels.ProviderInput) tpmodels.TypeProvider {
	if input.ThirdPartyID == "" {
		input.ThirdPartyID = discordID
	}

	if input.Config.AuthorizationEndpoint == "" {
		input.Config.AuthorizationEndpoint = "https://discord.com/oauth2/authorize"
	}

	if input.Config.TokenEndpoint == "" {
		input.Config.TokenEndpoint = "https://discord.com/api/oauth2/token"
	}

	if input.Config.UserInfoEndpoint == "" {
		input.Config.UserInfoEndpoint = "https://discord.com/api/users/@me"
	}

	if input.Config.UserInfoMap.FromUserInfoAPI.UserId == "" {
		input.Config.UserInfoMap.FromUserInfoAPI.UserId = "id"
	}

	if input.Config.UserInfoMap.FromUserInfoAPI.Email == "" {
		input.Config.UserInfoMap.FromUserInfoAPI.Email = "email"
	}

	if input.Config.UserInfoMap.FromUserInfoAPI.EmailVerified == "" {
		input.Config.UserInfoMap.FromUserInfoAPI.EmailVerified = "verified"
	}

	oOverride := input.Override

	input.Override = func(provider *tpmodels.TypeProvider) *tpmodels.TypeProvider {
		oGetConfig := provider.GetConfigForClientType
		provider.GetConfigForClientType = func(clientType *string, input tpmodels.ProviderConfig, userContext supertokens.UserContext) (tpmodels.ProviderConfigForClientType, error) {
			config, err := oGetConfig(clientType, input, userContext)
			if err != nil {
				return tpmodels.ProviderConfigForClientType{}, err
			}

			if len(config.Scope) == 0 {
				config.Scope = []string{"identify", "email"}
			}

			return config, err
		}

		oGetUserInfo := provider.GetUserInfo
		provider.GetUserInfo = func(config tpmodels.ProviderConfigForClientType, oAuthTokens tpmodels.TypeOAuthTokens, userContext supertokens.UserContext) (tpmodels.TypeUserInfo, error) {
			result, err := oGetUserInfo(config, oAuthTokens, userContext)
			if err != nil {
				return result, err
			}

			if config.AdditionalConfig != nil && config.AdditionalConfig["requireEmail"] == true {
				if result.Email == nil {
					result.Email = &tpmodels.EmailStruct{
						ID:         fmt.Sprintf("%s@fakediscorduser.com", result.ThirdPartyUserId),
						IsVerified: true,
					}
				}
			}

			return result, nil
		}

		if oOverride != nil {
			provider = oOverride(provider)
		}
		return provider
	}

	return NewProvider(input)
}
