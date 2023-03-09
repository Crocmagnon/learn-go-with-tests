package main

import "testing"

func TestHello2(t *testing.T) {
	got := Hello("Gab")
	want := "Hello, Gab"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Empty", args: args{name: ""}, want: "Hello, World"},
		{name: "Gab", args: args{name: "Gab"}, want: "Hello, Gab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
