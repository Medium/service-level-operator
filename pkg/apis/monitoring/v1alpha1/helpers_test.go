package v1alpha1_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	monitoringv1alpha1 "github.com/Medium/service-level-operator/pkg/apis/monitoring/v1alpha1"
)

func TestServiceLevelValidation(t *testing.T) {
	// Setup the different combinations of service level to validate.
	goodSL := &monitoringv1alpha1.ServiceLevel{
		ObjectMeta: metav1.ObjectMeta{
			Name: "fake-service0",
		},
		Spec: monitoringv1alpha1.ServiceLevelSpec{
			ServiceLevelObjectives: []monitoringv1alpha1.SLO{
				{
					Name:                         "fake_slo0",
					Description:                  "fake slo 0.",
					AvailabilityObjectivePercent: 99.99,
					ServiceLevelIndicator: monitoringv1alpha1.SLI{
						SLISource: monitoringv1alpha1.SLISource{
							Prometheus: &monitoringv1alpha1.PrometheusSLISource{
								Address:    "http://fake:9090",
								TotalQuery: `slo0_total`,
								ErrorQuery: `slo0_error`,
							},
						},
					},
					Output: monitoringv1alpha1.Output{
						Prometheus: &monitoringv1alpha1.PrometheusOutputSource{},
					},
				},
			},
		},
	}
	slWithoutSLO := goodSL.DeepCopy()
	slWithoutSLO.Spec.ServiceLevelObjectives = []monitoringv1alpha1.SLO{}
	slSLOWithoutName := goodSL.DeepCopy()
	slSLOWithoutName.Spec.ServiceLevelObjectives[0].Name = ""
	slSLOWithoutObjective := goodSL.DeepCopy()
	slSLOWithoutObjective.Spec.ServiceLevelObjectives[0].AvailabilityObjectivePercent = 0
	slSLOWithoutSLI := goodSL.DeepCopy()
	slSLOWithoutSLI.Spec.ServiceLevelObjectives[0].ServiceLevelIndicator.Prometheus = nil
	slSLOWithoutOutput := goodSL.DeepCopy()
	slSLOWithoutOutput.Spec.ServiceLevelObjectives[0].Output.Prometheus = nil

	tests := []struct {
		name         string
		serviceLevel *monitoringv1alpha1.ServiceLevel
		expErr       bool
	}{
		{
			name:         "A valid ServiceLevel should be valid.",
			serviceLevel: goodSL,
			expErr:       false,
		},
		{
			name:         "A ServiceLevel without SLOs houldn't be valid.",
			serviceLevel: slWithoutSLO,
			expErr:       true,
		},
		{
			name:         "A ServiceLevel with an SLO without name shouldn't be valid.",
			serviceLevel: slSLOWithoutName,
			expErr:       true,
		},
		{
			name:         "A ServiceLevel with an SLO without objective shouldn't be valid.",
			serviceLevel: slSLOWithoutObjective,
			expErr:       true,
		},
		{
			name:         "A ServiceLevel with an SLO without SLI shouldn't be valid.",
			serviceLevel: slSLOWithoutSLI,
			expErr:       true,
		},
		{
			name:         "A ServiceLevel with an SLO without output shouldn't be valid.",
			serviceLevel: slSLOWithoutOutput,
			expErr:       true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)

			err := test.serviceLevel.Validate()

			if test.expErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
		})
	}
}
