package solarzero

import (
	"context"
	"time"

	cognitosrp "github.com/alexrudd/cognito-srp/v4"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero/jsontypes"
)

var (
	NewAWSInterface = NewAWSInterfaceImpl
)

type AWSInterface interface {
	Authenticate() bool
	GetUser() bool
	UserAttributes() map[string]string
}

type awsInterfaceImpl struct {
	cognitoSvc   *cip.Client
	accessToken  string
	refreshToken string
	idToken      string

	config         *jsontypes.Configuration
	userAttributes map[string]string
}

func NewAWSInterfaceImpl(config *jsontypes.Configuration) AWSInterface {
	s := &awsInterfaceImpl{
		config:         config,
		userAttributes: make(map[string]string),
	}
	return s
}

func (impl *awsInterfaceImpl) UserAttributes() map[string]string {
	return impl.userAttributes
}

// Authenticate implements AWSInterface
func (impl *awsInterfaceImpl) Authenticate() bool {
	Logger.Info().Msg("INFO: Starting Cognito Authentication")
	csrp, _ := cognitosrp.NewCognitoSRP(
		impl.config.SolarZero.Username,
		impl.config.SolarZero.Password,
		impl.config.SolarZero.UserPoolID,
		impl.config.SolarZero.ClientID, nil)

	// configure cognito identity provider
	cfg, _ := awsconfig.LoadDefaultConfig(context.Background(),
		awsconfig.WithRegion(impl.config.SolarZero.API.Region),
		awsconfig.WithCredentialsProvider(aws.AnonymousCredentials{}),
	)

	svc := cip.NewFromConfig(cfg)
	resp, err := svc.InitiateAuth(context.Background(), &cip.InitiateAuthInput{
		AuthFlow:       types.AuthFlowTypeUserSrpAuth,
		ClientId:       aws.String(csrp.GetClientId()),
		AuthParameters: csrp.GetAuthParams(),
	})
	if err != nil {
		Logger.Error().Msgf("Cognito Authentication (InitiateAuth): %s ", err.Error())
		return false
	}

	// respond to password verifier challenge
	if resp.ChallengeName == types.ChallengeNameTypePasswordVerifier {
		challengeResponses, _ := csrp.PasswordVerifierChallenge(resp.ChallengeParameters, time.Now())

		resp, err := svc.RespondToAuthChallenge(context.Background(), &cip.RespondToAuthChallengeInput{
			ChallengeName:      types.ChallengeNameTypePasswordVerifier,
			ChallengeResponses: challengeResponses,
			ClientId:           aws.String(csrp.GetClientId()),
		})
		if err != nil {
			Logger.Error().Msgf("Cognito Authentication (RespondToAuthChallenge): %s", err.Error())
			return false
		}

		impl.accessToken = *resp.AuthenticationResult.AccessToken
		impl.idToken = *resp.AuthenticationResult.IdToken
		impl.refreshToken = *resp.AuthenticationResult.RefreshToken
		impl.cognitoSvc = svc
		Logger.Debug().Msgf("Access Token %s", *resp.AuthenticationResult.AccessToken)
		Logger.Debug().Msgf("ID Token %s", *resp.AuthenticationResult.IdToken)
		Logger.Debug().Msgf("Refresh Token %s", *resp.AuthenticationResult.RefreshToken)
		Logger.Info().Msg("Cognito Authentication Success")
		return true
	} else {
		Logger.Error().Msgf("Cognito Authentication (Unhandled Challenge): %s", resp.ChallengeName)
		return false
	}
}

func (impl *awsInterfaceImpl) GetUser() bool {
	Logger.Info().Msg("Cognito GetUser")

	getUserOutput, err := impl.cognitoSvc.GetUser(context.Background(), &cip.GetUserInput{
		AccessToken: &impl.accessToken,
	})

	if err != nil {
		Logger.Error().Msgf("Cognito GetUser (GetUser): %s", err.Error())
		return false
	}
	for _, attr := range getUserOutput.UserAttributes {
		impl.userAttributes[*attr.Name] = *attr.Value
	}

	Logger.Info().Msg("Cognito GetUser Success")

	return true
}
