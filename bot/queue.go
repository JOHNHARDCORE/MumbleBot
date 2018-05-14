package bot

import (
	"bytes"
	"fmt"
	"math"
	"sync"

	"github.com/MumbleBot/abstract"
	"layeh.com/gumble/gumble"
	"layeh.com/gumble/gumbleffmpeg"
)

// Queue holds videos and manages playback
type Queue struct {
	Queue       []abstract.Video
	maxCapacity int
	maxLength   int
	mut         sync.RWMutex
	voteMut     sync.RWMutex
	enabled     bool
	skips       map[string]int
}

func (q *Queue) Enabled() bool {
	return q.enabled
}

func newQueue() *Queue {
	return &Queue{Queue: make([]abstract.Video, 0), maxCapacity: 10, skips: make(map[string]int, 0)}
}

func (q *Queue) Clear() {
	q.mut.Lock()
	defer q.mut.Unlock()
	if Bot.Stream != nil {
		Bot.Stream.Stop()
	}
	q.enabled = false
	for i := 0; i < len(q.Queue); i++ {
		q.Queue[i] = nil
	}
	q.Queue = make([]abstract.Video, 0)
}
func (q *Queue) Length() int {
	q.mut.RLock()
	defer q.mut.RUnlock()
	return len(q.Queue)
}

func (q *Queue) Add(v abstract.Video) error {
	q.mut.Lock()
	defer q.mut.Unlock()
	fmt.Println("locking")

	if len(q.Queue) > q.maxCapacity {
		return fmt.Errorf("Too many videos in queue")
	}
	fmt.Println("adding video")
	/*
		if v.Length() >= q.maxLength {
			/* TODO: display a message saying that track is too long

			return
		}
	*/
	q.Queue = append(q.Queue, v)
	return nil
}

func (q *Queue) Pop() abstract.Video {
	q.mut.Lock()
	defer q.mut.Unlock()
	if len(q.Queue) == 0 {
		return nil
	}
	vid := q.Queue[0]
	q.Remove(1)
	fmt.Println("poppin 'deo")
	return vid
}
func (q *Queue) Remove(ind int) {
	// This is because users provide index + 1
	ind--
	if ind < 0 || ind >= len(q.Queue) {
		return
	}
	copy(q.Queue[ind:], q.Queue[ind+1:])
	q.Queue[len(q.Queue)-1] = nil
	q.Queue = q.Queue[:len(q.Queue)-1]
}

func (q *Queue) GetQueue() string {
	return q.contents(0, q.Length())
}

func (q *Queue) Peek() string {
	return q.contents(0, 1)
}

func (q *Queue) contents(start, end int) string {
	q.mut.RLock()
	defer q.mut.RUnlock()
	var buf bytes.Buffer
	if start < 0 || end > q.Length() {
		return "Error: Invalid index\n"
	}
	buf.WriteString(fmt.Sprintf("Tracks in Queue: %d\n", q.Length()))
	for i := start; i < end; i++ {
		buf.WriteString(fmt.Sprintf("%d. %-10s Posted By: %s<br>", i+1, q.Queue[i].Title(),
			q.Queue[i].Poster().Name))
	}
	return buf.String()
}

func (q *Queue) voteSkip(user *gumble.User) error {
	q.voteMut.Lock()
	defer q.voteMut.Unlock()
	if q.skips == nil {
		q.skips = make(map[string]int, 0)
	}
	if user.Channel != Bot.Client.Self.Channel {
		return fmt.Errorf("You can't skip from another channel")
	}
	if _, ok := q.skips[user.Name]; ok {
		return fmt.Errorf("You already voted")
	} else {
		q.skips[user.Name] = 1
		// add to queue and check to see if shits skipped
		// TODO: figure out some edge cases
		users := len(user.Channel.Users)
		votes := sum(q.skips)
		if users == 0 {
			return fmt.Errorf("idk how this was hit but somehow there's 0 people in the channel. can't skip")
		}
		pct := float64(votes) / float64(users)
		fmt.Println("Skip %:", pct)
		if pct >= .3 {
			if Bot.Stream != nil && Bot.Stream.State() == gumbleffmpeg.StatePlaying {
				Bot.Stream.Stop()
			}
			q.skips = make(map[string]int, 0)
			return fmt.Errorf("Video Skipped")
		}

		return fmt.Errorf("%s voted to skip. Current Votes: %d. Needed: %d", user.Name, votes, 1+int(math.Ceil(pct*float64(users))))
	}
}

func sum(votes map[string]int) int {
	total := 0
	for _, v := range votes {
		v++ // this is useless
		total++
	}

	return total
}
func (q *Queue) Skip(user *gumble.User) error {
	q.mut.Lock()
	defer q.mut.Unlock()
	if Bot.Stream == nil {
		return fmt.Errorf("No videos are playing")
	}

	return q.voteSkip(user)
}
func (q *Queue) Stop() {
	/* TODO: readd mutexs, clean shit up */
	q.mut.Lock()
	if Bot.Stream != nil {
		fmt.Println("Stopping...")
		q.enabled = false
		if Bot.Stream.State() != 0 {
			Bot.Stream.Stop()
		}
		Bot.Stream = nil
		fmt.Println("stopped")
	}
	q.mut.Unlock()
}
func (q *Queue) Start() error {
	q.enabled = true
	if Bot.Stream != nil && Bot.Stream.State() == gumbleffmpeg.StatePlaying {
		return nil
	}
	/* TODO: figure out how to get autoplay working */
	fmt.Println("Starting to process queue")
	for vid := q.Pop(); vid != nil && q.enabled; vid = vid {
		fmt.Println("Clearing voteskip")
		q.voteMut.Lock()
		q.skips = make(map[string]int, 0)
		q.voteMut.Unlock()
		err := q.Play(vid)
		if err != nil {
			Bot.Client.Self.Channel.Send(fmt.Sprintf("Couldn't play %s", vid.Title), false)
		}
		vid = q.Pop()
	}
	q.enabled = false

	return nil

}

func (q *Queue) Play(v abstract.Video) error {
	err := v.Download()
	if err != nil {
		return err
	}
	fmt.Println("Downloaded the video\n")
	err = v.Play()
	if v.Path() == "temp.mp3" {
		v.Remove()
	}
	return err
}
