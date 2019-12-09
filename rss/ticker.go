package rss

import (
	"context"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/capric98/t-rss/torrents"
)

type ticker struct {
	name         string
	client       *http.Client
	link, cookie string
	interval     time.Duration
	ctx          context.Context
}

func NewTicker(name string, link string, cookie string, interval time.Duration, wc *http.Client, ctx context.Context) (ch chan []torrents.Individ) {
	t := &ticker{
		name:     name,
		client:   wc,
		cookie:   cookie,
		link:     link,
		interval: interval,
		ctx:      ctx,
	}
	ch = make(chan []torrents.Individ)
	go t.tick(ch)
	return ch
}

func (t *ticker) tick(ch chan []torrents.Individ) {
	//tt := time.NewTicker(t.interval)
	//defer tt.Stop()

	//t.fetch(ch)
	for {
		select {
		case <-t.ctx.Done():
			close(ch)
			return
		default:
			go t.fetch(ch)
			time.Sleep(t.interval)
		}
	}
}

func (t *ticker) fetch(ch chan []torrents.Individ) {
	defer func() {
		e := recover()
		if e != nil {
			log.Println("rss ticker:", e)
		}
	}()
	startT := time.Now()

	// req, _ := http.NewRequest("GET", t.link, nil)
	// if t.cookie != "" {
	// 	req.Header.Add("Cookie", t.cookie)
	// }

	// resp, e := t.client.Do(req)
	// if e != nil {
	// 	return
	// }
	// defer resp.Body.Close()
	// rssFeed, _ := myfeed.Parse(resp.Body)

	// for k := range rssFeed.Items {
	// 	if rssFeed.Items[k].Enclosure.Url == "" {
	// 		rssFeed.Items[k].Enclosure.Url = rssFeed.Items[k].Link
	// 	}
	// 	if rssFeed.Items[k].GUID.Value == "" {
	// 		rssFeed.Items[k].GUID.Value = myfeed.NameRegularize(rssFeed.Items[k].Title)
	// 	}
	// 	rssFeed.Items[k].GUID.Value = myfeed.NameRegularize(rssFeed.Items[k].GUID.Value)
	// }

	log.Printf("%s fetched in %7.2fms.", t.name, time.Since(startT).Seconds()*1000.0)
	//ch <- rssFeed.Items
	runtime.GC()
}
