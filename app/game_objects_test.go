package app

import "testing"

func TestPath_Follow(t *testing.T) {
	tests := []struct {
		name       string
		points     []PathPoint
		length     float64
		x, y       float64
		distance   float64
		expX, expY float64
	}{
		{
			name: "casual line",
			points: []PathPoint{
				{0, 0},
				{0, 1},
			},
			length:   1,
			distance: 1,
			expX:     0,
			expY:     1,
		},
		{
			name: "line with more distance",
			points: []PathPoint{
				{0, 0},
				{0, 1},
			},
			length:   1,
			distance: 2,
			expX:     0,
			expY:     1,
		},
		{
			name: "line with less distance",
			points: []PathPoint{
				{0, 0},
				{0, 2},
			},
			length:   2,
			distance: 1,
			expX:     0,
			expY:     1,
		},
		{
			name: "casual turn",
			points: []PathPoint{
				{0, 0},
				{0, 1},
				{1, 1},
			},
			length:   2,
			distance: 2,
			expX:     1,
			expY:     1,
		},
		{
			name: "several turns, less distance",
			points: []PathPoint{
				{0, 0},
				{0, 1},
				{1, 1},
				{1, 2},
				{2, 2},
			},
			length:   4,
			distance: 3,
			expX:     1,
			expY:     2,
		},
		{
			name: "several turns, less distance one more time",
			points: []PathPoint{
				{0, 0},
				{0, 1},
				{1, 1},
				{1, 2},
				{2, 2},
			},
			length:   4,
			distance: 3.5,
			expX:     1.5,
			expY:     2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var path = Path{
				Points: tt.points,
				Length: tt.length,
			}
			actX, actY := path.Follow(tt.distance)
			if actX != tt.expX || actY != tt.expY {
				t.Errorf("incorrect: (%f, %f), expected (%f, %f)\n", actX, actY, tt.expX, tt.expY)
			}
		})
	}
}
