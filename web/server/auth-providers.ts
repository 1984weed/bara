import { Provider, ProviderProfile } from "./model"
import passportTwitter from "passport-twitter"
import passportGithub from "passport-github"


export type ProviderKeys = {
    consumerKey: string
    consumerSecret: string
}

export type ProviderConfig = {
	twitter?: ProviderKeys
	github?: ProviderKeys
}

export const createProviders = ({ twitter, github }: ProviderConfig): Provider[] => {
    const providers = []
    if (twitter) {
        providers.push({
            name: "twitter",
            strategy: passportTwitter.Strategy,
            strategyOptions: {...twitter, includeEmail: true},
            getProfile: function(profile: any): ProviderProfile {
                return {
                    id: profile.id,
                    name: profile.displayName,
                    email: profile.emails && profile.emails[0].value ? profile.emails[0].value : "",
                    displayName: profile.displayName,
                    imageUrl: profile.photos[0].value,
                }
            },
        })
	}
	
	if(github) {
        providers.push({
            name: "github",
            strategy: passportGithub.Strategy,
            strategyOptions: {
				clientID: github.consumerKey,
				clientSecret: github.consumerSecret
			},
            getProfile: function(profile: any): ProviderProfile {
                return {
                    id: profile.id,
                    name: profile.displayName,
                    email: profile.emails && profile.emails[0].value ? profile.emails[0].value : "",
                    displayName: profile.displayName,
                    imageUrl: profile.photos[0].value,
                }
            },
        })
	}

    return providers
}
