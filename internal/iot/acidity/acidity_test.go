package acidity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidUnit(t *testing.T) {
	testcases := []struct {
		input   float32
		want    Unit
		wantErr error
	}{
		{
			input:   7,
			want:    Unit{value: 7},
			wantErr: nil,
		},
		{
			input:   -1,
			want:    Unit{},
			wantErr: ErrInvalidValue,
		},
		{
			input:   14.01,
			want:    Unit{},
			wantErr: ErrInvalidValue,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			got, err := NewUnit(tc.input)
			if assert.ErrorIs(t, err, tc.wantErr, fmt.Sprintf("Want: %v Got: %v", tc.wantErr, err)) {
				assert.Equal(t, tc.want, got, fmt.Sprintf("Want: %v Got: %v", tc.want, got))
			}
		})
	}
}
