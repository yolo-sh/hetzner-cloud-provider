package userconfig_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/yolo-sh/hetzner-cloud-provider/config"
	"github.com/yolo-sh/hetzner-cloud-provider/mocks"
	"github.com/yolo-sh/hetzner-cloud-provider/userconfig"
)

func TestFilesResolving(t *testing.T) {
	unknownError := errors.New("UnknownError")

	testCases := []struct {
		test                  string
		configInFile          *userconfig.Config
		regionOpts            string
		contextOpts           string
		errorReturnedByLoader error
		expectedError         error
		expectedConfig        *userconfig.Config
	}{
		{
			test:           "valid",
			configInFile:   userconfig.NewConfig("a", ""),
			regionOpts:     "b",
			expectedConfig: userconfig.NewConfig("a", "b"),
			expectedError:  nil,
		},

		{
			test:           "missing region opt",
			configInFile:   userconfig.NewConfig("a", ""),
			expectedError:  userconfig.ErrMissingRegion,
			expectedConfig: nil,
		},

		{
			test:           "missing API token and region",
			configInFile:   userconfig.NewConfig("", ""),
			expectedError:  userconfig.ErrMissingConfig,
			expectedConfig: nil,
		},

		{
			test:                  "missing context",
			contextOpts:           "context",
			regionOpts:            "region",
			errorReturnedByLoader: config.ErrContextNotFound{},
			expectedError:         config.ErrContextNotFound{},
			expectedConfig:        nil,
		},

		{
			test:                  "unknown error",
			errorReturnedByLoader: unknownError,
			expectedError:         unknownError,
			expectedConfig:        nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			contextToLoad := tc.contextOpts

			configAsReturnedByProfileLoader := &userconfig.Config{}
			if tc.configInFile != nil {
				configAsReturnedByProfileLoader.Credentials.APIToken = tc.configInFile.Credentials.APIToken
				configAsReturnedByProfileLoader.Region = tc.configInFile.Region
			}

			contextLoaderMock := mocks.NewUserConfigContextLoader(mockCtrl)
			contextLoaderMock.
				EXPECT().
				Load(contextToLoad, userconfig.DefaultConfigFilePath).
				Return(configAsReturnedByProfileLoader, tc.errorReturnedByLoader).
				AnyTimes()

			resolver := userconfig.NewFilesResolver(
				contextLoaderMock,
				userconfig.FilesResolverOpts{
					Region:  tc.regionOpts,
					Context: tc.contextOpts,
				},
			)

			resolvedConfig, err := resolver.Resolve()

			if tc.expectedError == nil && err != nil {
				t.Fatalf("expected no error, got '%+v'", err)
			}

			if tc.expectedError != nil && !errors.Is(err, tc.expectedError) && !errors.As(err, &tc.expectedError) {
				t.Fatalf("expected error to equal '%+v', got '%+v'", tc.expectedError, err)
			}

			if tc.expectedConfig != nil && *resolvedConfig != *tc.expectedConfig {
				t.Fatalf("expected config to equal '%+v', got '%+v'", *tc.expectedConfig, *resolvedConfig)
			}

			if tc.expectedConfig == nil && resolvedConfig != nil {
				t.Fatalf("expected no config, got '%+v'", *resolvedConfig)
			}
		})
	}
}
