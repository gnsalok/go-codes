package test

import (
	"github.com/gnsalok/master-golang/golang-all/Testing"
	"testing"
)

func TestWhatTime(t *testing.T) {

	tests := []struct {
		name     string
		input    int
		wantErr  bool
		wantTime string
	}{
		{
			name:    "Failure - invalid time not in range, below 0",
			input:   -1,
			wantErr: true,
		},
		{
			name:    "Failure - invalid time not in range, above 86399",
			input:   86400,
			wantErr: true,
		},
		{
			name:     "Success - got time for 0 as 0:0:0 ",
			input:    0,
			wantErr:  false,
			wantTime: "0:0:0",
		},
		{
			name:     "Success - got time for 3661 as 1:1:1 ",
			input:    3661,
			wantErr:  false,
			wantTime: "1:1:1",
		},
		{
			name:     "Success - got time for 5436 as 1:30:36 ",
			input:    5436,
			wantErr:  false,
			wantTime: "1:30:36",
		},
		{
			name:     "Success - got time for 0 as 0:0:0  ",
			input:    86399,
			wantErr:  false,
			wantTime: "23:59:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time, err := Testing.WhatTime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("log it")
			}
			if err == nil && time != tt.wantTime {
				t.Errorf("log it again")
			}
		})
	}
}
