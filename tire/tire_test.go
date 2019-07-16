package tire_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/freman/roadlib/tire"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		input    string
		expected tire.Info
		err      error
	}{
		{
			"120/70R15",
			tire.Info{Type: 0x0, Width: 120, AspectRatio: 70, Construction: 0x1, Rim: 381, Diameter: 549, SideWall: 84, Circumference: 1724.7343668207964, Revolutions: 579.7994283857754},
			nil,
		}, {
			"140/70R14",
			tire.Info{Type: 0x0, Width: 140, AspectRatio: 70, Construction: 0x1, Rim: 355.59999999999997, Diameter: 551.5999999999999, SideWall: 98, Circumference: 1732.9025077201295, Revolutions: 577.066508672572},
			nil,
		}, {
			"1x0/70R14",
			tire.Info{},
			errors.New("expected / at 1"),
		}, {
			"140970R14",
			tire.Info{},
			errors.New("expected / at 6"),
		}, {
			"140/70H14",
			tire.Info{},
			errors.New("expected construction (R) at 6"),
		}, {
			"1.4.0/70R14",
			tire.Info{},
			errors.New("parsing width failed due to strconv.ParseFloat: parsing \"1.4.0\": invalid syntax at 0"),
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.input, func(t *testing.T) {
			t.Parallel()
			actual, err := tire.Parse(test.input)
			if test.err != nil && err == nil {
				t.Errorf("Error should have been %v", test.err)
			}
			if test.err == nil && err != nil {
				t.Errorf("Error should have been nil but was %v", err)
			}
			if test.err != nil && err != nil {
				require.Equal(t, test.err, err, "Unexpected error")
			}
			require.Equal(t, test.expected, actual, "Info didn't match expectation")
		})
	}
}

func TestRevolutions(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		input tire.Revolutions
		km    float64
		m     float64
		mile  float64
	}{
		{
			tire.Revolutions(611), 611, 0.611, 985.4777609738344,
		},
	}
	for _, test := range testCases {
		test := test
		t.Run(fmt.Sprintf("%f", test.input), func(t *testing.T) {
			t.Parallel()
			require.Equal(t, test.km, test.input.PerKilometre())
			require.Equal(t, test.m, test.input.PerMetre())
			require.Equal(t, test.mile, test.input.PerMile())
		})
	}
}
