package strings_test

import (
	strings "github.com/synesissoftware/ANGoLS/strings"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"testing"
)

func Test_ParseKeyValuePairsList(t *testing.T) {

	tests := []struct {
		name                      string
		input                     string
		pairSeparator             string
		keyValueSeparator         string
		options                   strings.ParseKeyValuePairsListOption
		pairs                     []strings.KeyValuePair
		shouldFail                bool
		expectedErrStringFragment string
	}{
		{
			name:              "empty input",
			input:             "",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs:             []strings.KeyValuePair{},
		},
		{
			name:              "whitespace-only input",
			input:             "               \t          ",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs:             []strings.KeyValuePair{},
		},
		{
			name:              "1-pair",
			input:             "k1=v1",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs: []strings.KeyValuePair{
				{
					Key:   "k1",
					Value: "v1",
				},
			},
		},
		{
			name:              "1-pair with trailing pair-separator",
			input:             "k1=v1,",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs: []strings.KeyValuePair{
				{
					Key:   "k1",
					Value: "v1",
				},
			},
		},
		{
			name:              "1-pair with trailing pair-separator and leading and trailing whitespace",
			input:             "    k1=v1,   ",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs: []strings.KeyValuePair{
				{
					Key:   "k1",
					Value: "v1",
				},
			},
		},
		{
			name:              "2-pairs",
			input:             "k1=v1,k2=val2",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs: []strings.KeyValuePair{
				{
					Key:   "k1",
					Value: "v1",
				},
				{
					Key:   "k2",
					Value: "val2",
				},
			},
		},
		{
			name:              "2-pairs with leading and trailing pair-separators",
			input:             ",,,,,,k1=v1,k2=val2,,",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs: []strings.KeyValuePair{
				{
					Key:   "k1",
					Value: "v1",
				},
				{
					Key:   "k2",
					Value: "val2",
				},
			},
		},
		// anonymous value(s)
		{
			name:                      "1-pair with anonymous value, which fails",
			input:                     "=v1",
			pairSeparator:             ",",
			keyValueSeparator:         "=",
			shouldFail:                true,
			expectedErrStringFragment: "missing key",
		},
		{
			name:              "1-pair with anonymous value, which is ignored",
			input:             "=v1",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs:             []strings.KeyValuePair{},
			options:           strings.ParseKeyValuePairsListOption_IgnoreAnonymousValue,
		},
		{
			name:              "1-pair with anonymous value, which is permitted",
			input:             "=v1",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs: []strings.KeyValuePair{
				{
					Key:   "",
					Value: "v1",
				},
			},
			options: strings.ParseKeyValuePairsListOption_PermitAnonymousValue,
		},
		// valueless key(s)
		{
			name:                      "1-pair with valueless key, which fails",
			input:                     "k1",
			pairSeparator:             ",",
			keyValueSeparator:         "=",
			shouldFail:                true,
			expectedErrStringFragment: "missing value",
		},
		{
			name:              "1-pair with valueless key, which is ignored",
			input:             "k1",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs:             []strings.KeyValuePair{},
			options:           strings.ParseKeyValuePairsListOption_IgnoreValuelessKey,
		},
		{
			name:              "1-pair with valueless key, which is permitted",
			input:             "k1",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs: []strings.KeyValuePair{
				{
					Key:   "k1",
					Value: "",
				},
			},
			options: strings.ParseKeyValuePairsListOption_PermitValuelessKey,
		},
		{
			name:                      "1-pair with valueless key (and separator), which fails",
			input:                     "k1=",
			pairSeparator:             ",",
			keyValueSeparator:         "=",
			shouldFail:                true,
			expectedErrStringFragment: "missing value",
		},
		{
			name:              "1-pair with valueless key (and separator), which is ignored",
			input:             "k1=",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs:             []strings.KeyValuePair{},
			options:           strings.ParseKeyValuePairsListOption_IgnoreValuelessKey,
		},
		{
			name:              "1-pair with valueless key (and separator), which is permitted",
			input:             "k1=",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs: []strings.KeyValuePair{
				{
					Key:   "k1",
					Value: "",
				},
			},
			options: strings.ParseKeyValuePairsListOption_PermitValuelessKey,
		},
		// repeated key(s)
		{
			name:                      "2-pair with identical keys, which fails",
			input:                     "k1=v1,k2=v2,k1=v3",
			pairSeparator:             ",",
			keyValueSeparator:         "=",
			shouldFail:                true,
			expectedErrStringFragment: "repeated keys",
		},
		{
			name:              "2-pair with identical keys, which is permitted",
			input:             "k1=v1,k2=v2,k1=v3",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs: []strings.KeyValuePair{
				{
					Key:   "k1",
					Value: "v1",
				},
				{
					Key:   "k2",
					Value: "v2",
				},
				{
					Key:   "k1",
					Value: "v3",
				},
			},
			options: strings.ParseKeyValuePairsListOption_PermitRepeatedKeys,
		},
		{
			name:              "2-pair with identical keys, keeping the first",
			input:             "k1=v1,k2=v2,k1=v3",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs: []strings.KeyValuePair{
				{
					Key:   "k1",
					Value: "v1",
				},
				{
					Key:   "k2",
					Value: "v2",
				},
			},
			options: strings.ParseKeyValuePairsListOption_TakeFirstRepeatedKey,
		},
		{
			name:              "2-pair with identical keys, keeping the last",
			input:             "k1=v1,k2=v2,k1=v3",
			pairSeparator:     ",",
			keyValueSeparator: "=",
			pairs: []strings.KeyValuePair{
				{
					Key:   "k1",
					Value: "v3",
				},
				{
					Key:   "k2",
					Value: "v2",
				},
			},
			options: strings.ParseKeyValuePairsListOption_TakeLastRepeatedKey,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if pairs, err := strings.ParseKeyValuePairsList(
				tt.input,
				tt.pairSeparator,
				tt.keyValueSeparator,
				tt.options,
			); err != nil {

				require.True(t, tt.shouldFail)

				assert.Contains(t, tt.expectedErrStringFragment, err.Error())
			} else {

				require.False(t, tt.shouldFail)

				require.Equal(t, len(tt.pairs), len(pairs), "expected %d pair(s) but obtained %d", len(tt.pairs), len(pairs))

				for i := 0; i != len(tt.pairs); i++ {

					assert.Equal(t, tt.pairs[i], pairs[i])
				}
			}
		})
	}
}
