package runtime_test

import (
	"testing"

	"github.com/kyma-project/kyma-environment-broker/internal"
	"github.com/kyma-project/kyma-environment-broker/internal/runtime"
	"github.com/kyma-project/control-plane/components/provisioner/pkg/gqlschema"
	"github.com/stretchr/testify/assert"
)

func TestCustomDisablerExample(t *testing.T) {
	// given
	sut := runtime.NewCustomDisablerExample()

	givenComponents := internal.ComponentConfigurationInputList{
		{
			Component: runtime.CustomDisablerComponentName,
			Namespace: "kyma-system",
		},
	}
	expComponents := internal.ComponentConfigurationInputList{
		{
			Component: runtime.CustomDisablerComponentName,
			Namespace: "kyma-system",
			Configuration: []*gqlschema.ConfigEntryInput{
				{
					Key:   "component-x.enabled",
					Value: "false",
				},
				{
					Key:   "component-x.Output.conf.enabled",
					Value: "false",
				},
			},
		},
	}

	// when
	modifiedComponents := sut.Disable(givenComponents)

	// then
	assert.EqualValues(t, expComponents, modifiedComponents)
}
