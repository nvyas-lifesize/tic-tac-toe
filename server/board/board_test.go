package board

import (
	"reflect"
	"testing"
)

// Test the minimax algo.
func TestBoard_Minimax(t *testing.T) {
	type fields struct {
		Grid [3][3]int
	}
	type args struct {
		depth  int
		player bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   [3]int
	}{
		{name: "Test-Defence",
			fields: fields{Grid: [3][3]int{
				{0, 1, -1}, {-1, 1, -1}, {-1, -1, -1},
			}}, args: args{depth: 6, player: false}, want: [3]int{2, 1, 0}},
		{name: "Test-Win",
			fields: fields{Grid: [3][3]int{
				{0, 1, 1}, {1, 1, -1}, {0, 0, -1},
			}}, args: args{depth: 2, player: false}, want: [3]int{2, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := &Board{
				Grid: tt.fields.Grid,
			}
			if got := board.Minimax(tt.args.depth, tt.args.player); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Minimax() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test the IsGameover
func TestBoard_IsGameOver(t *testing.T) {
	type fields struct {
		Grid [3][3]int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
		want1  int
	}{
		{name: "Player win", fields: fields{Grid: [3][3]int{
			{1, -1, -1}, {0, 1, -1}, {-1, 0, 1},
		}}, want: true, want1: 1},
		{name: "Computer win", fields: fields{Grid: [3][3]int{
			{0, 1, 1}, {-1, 0, -1}, {-1, -1, 0},
		}}, want: true, want1: 0},
		{name: "Non result", fields: fields{Grid: [3][3]int{
			{0, -1, -1}, {-1, 1, -1}, {-1, -1, -1},
		}}, want: false, want1: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := &Board{
				Grid: tt.fields.Grid,
			}
			got, got1 := board.IsGameOver()
			if got != tt.want {
				t.Errorf("IsGameOver() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsGameOver() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

//Test to calculate Depth
func TestBoard_CalculateDepth(t *testing.T) {
	type fields struct {
		Grid [3][3]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "Test-Empty-grid", fields: fields{Grid: [3][3]int{
			{1, -1, -1}, {0, 1, -1}, {-1, 0, 1},
		}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := &Board{
				Grid: tt.fields.Grid,
			}
			if got := board.CalculateDepth(); got != tt.want {
				t.Errorf("CalculateDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
