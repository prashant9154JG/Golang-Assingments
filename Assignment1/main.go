package main

import (
	"fmt"
	"sync"
	"time"
)

// Video is a structure of video
type Video struct {
	video_id string
	views    int
	mu       sync.Mutex
	wg       sync.WaitGroup
}

// initVideo initialise the video using given video id and initial view count to 0
func initVideo(vid string) *Video {
	return &Video{
		video_id: vid,
		views:    0,
		mu:       sync.Mutex{},
		wg:       sync.WaitGroup{},
	}
}

// showViewsCount prints the view counts of video on terminal
func (v *Video) showViewsCount() {
	v.mu.Lock()
	defer v.wg.Done()
	fmt.Println("Video with id:", v.video_id, " has ", v.views, " views")
	v.mu.Unlock()
}

// incrementViewsCount increments views count of video by 1
func (v *Video) incrementViewsCount() {
	v.mu.Lock()
	defer v.wg.Done()
	time.Sleep(time.Second)
	v.views += 1
	fmt.Println("Incremented views for video id ", v.video_id)
	v.mu.Unlock()
}

func main() {

	// initialising slice of videos
	videos := []Video{
		*initVideo("d5jn13"),
		*initVideo("g1j40a"),
		*initVideo("abc123"),
	}

	for _, video := range videos {
		for i := 0; i < 5; i++ {
			video.wg.Add(1)
			go video.showViewsCount()
			video.wg.Add(1)
			go video.incrementViewsCount()

			video.wg.Wait()
		}
	}
}
