package main

import "testing"

func TestHello1(t *testing.T) {
	type args struct {
		name     string
		language string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Empty name & language", args: args{name: "", language: ""}, want: "Hello, World"},
		{name: "Empty name & Spanish", args: args{name: "", language: "Spanish"}, want: "Hola, World"},
		{name: "Elodie & Spanish", args: args{name: "Elodie", language: "Spanish"}, want: "Hola, Elodie"},
		{name: "Gab & French", args: args{name: "Gab", language: "French"}, want: "Bonjour, Gab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name, tt.args.language); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
