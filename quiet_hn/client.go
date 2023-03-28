package quiethn

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"
)

type story struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	orderId int
}

// String returns the string representaion of this object
func (s *story) String() string {
	return fmt.Sprintf("id: %d, url: %s, type: %s, title: %s", s.Id, s.Url, s.Type, s.Title)
}

type client struct {
	baseAPI     string
	stories     []*story
	tmpStories  []*story
	timer       *time.Timer
	refreshTime time.Duration
}

// init prepares the necessary setups for this object
func (c *client) init(initialTotalStories int) {
	c.stories = c.getNewTopStories(initialTotalStories, 0, []*story{})
	c.timer = time.AfterFunc(c.refreshTime, c.refreshCache)
}

// refreshCache re-fetches data from the APIs to collect the most updated ones
func (c *client) refreshCache() {
	c.tmpStories = c.getNewTopStories(len(c.stories), 0, []*story{})
	log.Printf("Added new data to the tmpStories. stories: %d, tmpStories: %d\n", len(c.stories), len(c.tmpStories))

	c.timer.Reset(c.refreshTime)
}

// GetTopStories returns the top-m stories. It will return data from the cache, except if the cache contains data less than m
func (c *client) GetTopStories(m int) []*story {
	if len(c.tmpStories) > 0 {
		if len(c.tmpStories) >= len(c.stories) {
			c.stories = c.tmpStories
		}
		c.tmpStories = nil
	}

	if len(c.stories) < m {
		c.stories = c.getNewTopStories(m, 0, c.stories)
	}
	return c.stories[:m]
}

// getTopStories returns the top-m sorted stories
func (c *client) getNewTopStories(m, p int, stories []*story) []*story {
	stories = append(stories, c.getStories(m, p, stories)...)
	sort.Slice(stories, func(i, j int) bool {
		return stories[i].orderId < stories[j].orderId
	})
	return stories[:m]
}

// getStories returns m stories from the API
func (c *client) getStories(m, p int, stories []*story) []*story {
	var wg sync.WaitGroup
	dataToGet := int(float64(m) * 1.5)
	ch := make(chan story, dataToGet)

	for i := p; i < p+dataToGet; i++ {
		wg.Add(1)
		go c.getStory(i, &wg, ch)
	}

	wg.Wait()
	close(ch)
	for chVal := range ch {
		if chVal.Type == "story" && chVal.Title != "" {
			story := chVal
			stories = append(stories, &story)
		}
	}

	if len(stories) < m {
		return c.getStories(m, p+dataToGet, stories)
	}
	return stories
}

// getStories gets data from the API and stores it in the channel that will then be processed
func (c *client) getStory(i int, wg *sync.WaitGroup, ch chan story) {
	story := story{}
	defer wg.Done()

	resp, err := http.Get(fmt.Sprintf("%s/v0/item/%d.json", c.baseAPI, i+1))
	if err != nil {
		ch <- story
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&story)
	if err != nil {
		ch <- story
		return
	}

	story.orderId = i
	ch <- story
}

// NewClient creates a client that can retrieve top-m stories from the hacker-news API
func NewClient(initialTotalStories int, refreshTime time.Duration) *client {
	if refreshTime < time.Second {
		refreshTime = time.Second
	}

	result := &client{
		baseAPI:     "https://hacker-news.firebaseio.com",
		refreshTime: refreshTime,
	}
	result.init(initialTotalStories)

	return result
}
