package quiethn

import (
	"strconv"
	"testing"
	"time"
)

// TestGetTopStories tests the GetStories function of the client struct
func TestGetTopStories(t *testing.T) {
	for _, v := range testTable.getTopStories {
		m, err := strconv.Atoi(v)
		client := NewClient(m, 1*time.Hour)
		if err != nil {
			t.Fatal(err)
		}

		stories := client.GetTopStories(m)
		if len(stories) != m {
			t.Fatalf("exepected %d but got %d. v: %s", m, len(stories), v)
		}

		for _, story := range stories {
			if story.Type != "story" || story.Title == "" {
				t.Fatalf("got invalid data %s. v:%s", story.String(), v)
			}
		}
	}
}

// TestRefreshCache tests the refreshCache function of the client struct
func TestRefreshCache(t *testing.T) {
	for _, v := range testTable.refreshCache {
		second, err := strconv.Atoi(v)
		if err != nil {
			t.Fatal(err)
		}
		totalStories := 3
		refreshTime := time.Duration(second) * time.Second
		client := NewClient(totalStories, refreshTime)
		time.Sleep(refreshTime + 5*time.Second)

		if len(client.tmpStories) != totalStories {
			t.Fatalf("expected %d in len but got %d. v: %s", totalStories, len(client.tmpStories), v)
		}
		t.Logf("%d\n", len(client.tmpStories))
		for _, story := range client.tmpStories {
			t.Logf("%s\n", story.String())
			if story.Type != "story" && story.Title == "" {
				t.Fatalf("got invalid data %s. v:%s", story.String(), v)
			}
		}
		client.timer.Stop()
	}
}
