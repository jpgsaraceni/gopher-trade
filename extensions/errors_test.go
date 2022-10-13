package extensions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Extensions_ErrStack(t *testing.T) {
	t.Parallel()

	type args struct {
		err error
		op  string
	}

	tableTests := []struct {
		name string
		args
		want string
	}{
		{
			name: "should return error stack with a single error",
			args: args{
				op:  "SomeFunc",
				err: fmt.Errorf("bad stuff"),
			},
			want: "SomeFunc(): bad stuff",
		},
		{
			name: "should return error stack with three nested errors with operations and errors",
			args: args{
				op:  "Level3Operation",
				err: fmt.Errorf("level 3 err %w", childErrors(t, 2, true)),
			},
			want: "Level3Operation(): level 3 err Level2Operation(): level 2 err: Level1Operation(): original error",
		},
		{
			name: "should return error stack with three nested errors with only operations",
			args: args{
				op:  "Level3Operation",
				err: childErrors(t, 2, false),
			},
			want: "Level3Operation()-> Level2Operation()-> Level1Operation(): original error",
		},
	}
	for _, tt := range tableTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotErr := ErrStack(tt.args.op, tt.args.err)
			assert.ErrorIs(t, gotErr, tt.err)
			assert.Equal(t, tt.want, gotErr.Error())
		})
	}
}

func childErrors(t *testing.T, n int, showErrs bool) error { //nolint:revive
	t.Helper()

	err := ErrStack("Level1Operation", fmt.Errorf("original error"))

	if showErrs {
		for i := 2; i <= n; i++ {
			err = ErrStack(fmt.Sprintf("Level%dOperation", i), fmt.Errorf("level %d err: %w", i, err))
		}
	} else {
		for i := 2; i <= n; i++ {
			err = ErrStack(fmt.Sprintf("Level%dOperation", i), err)
		}
	}

	return err
}
