package iam

import (
	"testing"
	"time"

	defsecTypes "github.com/aquasecurity/defsec/pkg/types"

	"github.com/aquasecurity/defsec/pkg/state"

	"github.com/aquasecurity/defsec/pkg/providers/aws/iam"
	"github.com/aquasecurity/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckAccessKeysRotated(t *testing.T) {
	tests := []struct {
		name     string
		input    iam.IAM
		expected bool
	}{
		{
			name: "Access key created a month ago",
			input: iam.IAM{
				Users: []iam.User{
					{
						Metadata:   defsecTypes.NewTestMetadata(),
						Name:       defsecTypes.String("user", defsecTypes.NewTestMetadata()),
						LastAccess: defsecTypes.TimeUnresolvable(defsecTypes.NewTestMetadata()),
						AccessKeys: []iam.AccessKey{
							{
								Metadata:     defsecTypes.NewTestMetadata(),
								AccessKeyId:  defsecTypes.String("AKIACKCEVSQ6C2EXAMPLE", defsecTypes.NewTestMetadata()),
								Active:       defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
								CreationDate: defsecTypes.Time(time.Now().Add(-time.Hour*24*30), defsecTypes.NewTestMetadata()),
								LastAccess:   defsecTypes.Time(time.Now(), defsecTypes.NewTestMetadata()),
							},
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "Access key created 4 months ago",
			input: iam.IAM{
				Users: []iam.User{
					{
						Metadata:   defsecTypes.NewTestMetadata(),
						Name:       defsecTypes.String("user", defsecTypes.NewTestMetadata()),
						LastAccess: defsecTypes.TimeUnresolvable(defsecTypes.NewTestMetadata()),
						AccessKeys: []iam.AccessKey{
							{
								Metadata:     defsecTypes.NewTestMetadata(),
								AccessKeyId:  defsecTypes.String("AKIACKCEVSQ6C2EXAMPLE", defsecTypes.NewTestMetadata()),
								Active:       defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
								CreationDate: defsecTypes.Time(time.Now().Add(-time.Hour*24*30*4), defsecTypes.NewTestMetadata()),
								LastAccess:   defsecTypes.Time(time.Now(), defsecTypes.NewTestMetadata()),
							},
						},
					},
				},
			},
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.AWS.IAM = test.input
			results := CheckAccessKeysRotated.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckAccessKeysRotated.LongID() {
					found = true
				}
			}
			if test.expected {
				assert.True(t, found, "Rule should have been found")
			} else {
				assert.False(t, found, "Rule should not have been found")
			}
		})
	}
}
