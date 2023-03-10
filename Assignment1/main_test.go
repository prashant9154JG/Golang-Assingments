package main

import (
	"reflect"
	"sync"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"2s3rn4"},
		{"3h3s1l"},
		{"fid9nd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestVideo_incrementViewsCount(t *testing.T) {
	type fields struct {
		video_id string
		views    int
		mu       sync.Mutex
		wg       sync.WaitGroup
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"zero to one viewcount", fields{"d8s6ad", 0, sync.Mutex{}, sync.WaitGroup{}}},
		{"two to three viewcount", fields{"1kjb3d", 2, sync.Mutex{}, sync.WaitGroup{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prevViewCount := tt.fields.views

			v := &Video{
				video_id: tt.fields.video_id,
				views:    tt.fields.views,
				mu:       tt.fields.mu,
				wg:       tt.fields.wg,
			}
			v.wg.Add(1)
			v.incrementViewsCount()

			if prevViewCount != v.views-1 {
				t.Errorf("Expected incremented count but got same count")
			}
		})
	}
}

func TestVideo_showViewsCount(t *testing.T) {
	type fields struct {
		video_id string
		views    int
		mu       sync.Mutex
		wg       sync.WaitGroup
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"show 0 viewcount", fields{"d8s6ad", 0, sync.Mutex{}, sync.WaitGroup{}}},
		{"show 2 viewcount", fields{"1kjb3d", 2, sync.Mutex{}, sync.WaitGroup{}}},
		{"show 4 viewcount", fields{"8d1jns", 4, sync.Mutex{}, sync.WaitGroup{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Video{
				video_id: tt.fields.video_id,
				views:    tt.fields.views,
				mu:       tt.fields.mu,
				wg:       tt.fields.wg,
			}
			v.wg.Add(1)
			v.showViewsCount()
		})
	}
}

func Test_initVideo(t *testing.T) {
	type args struct {
		vid string
	}
	tests := []struct {
		name string
		args args
		want *Video
	}{
		{"base-test", args{"8ad8hv"}, &Video{"8ad8hv", 0, sync.Mutex{}, sync.WaitGroup{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initVideo(tt.args.vid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initVideo() = %v, want %v", got, tt.want)
			}
		})
	}
}
