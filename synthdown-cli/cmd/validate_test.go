package cmd

import (
	"testing"

	"github.com/spf13/viper"
)

type returnedStatus struct {
	v int
}

func (r *returnedStatus) fail(i int) { r.v = i }

func TestValidate_Run(t *testing.T) {
	origFailer := failer
	defer func() {
		failer = origFailer
	}()

	for _, test := range []struct {
		fn           string
		expectStatus int
	}{
		{"testdata/complex.sdown", 0},
		{"validate.go", 1},
		{"testdata/nonsuch", 1},
	} {
		t.Run(test.fn, func(t *testing.T) {
			rs := returnedStatus{0}
			failer = rs.fail

			viper.Set("file", test.fn)
			validateCmd.Run(nil, []string{})

			if test.expectStatus != rs.v {
				t.Errorf("expected %d, received %d", test.expectStatus, rs.v)
			}
		})
	}

}
