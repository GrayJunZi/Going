package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	Id       int
	Comments []string
	Likes    int
	Friends  []int
}

type Response struct {
	data any
	err  error
}

func handleGetUserProfile(id int) (*UserProfile, error) {
	var (
		respch = make(chan Response, 3)
		wg     = &sync.WaitGroup{}
	)
	go getComments(id, respch, wg)
	go getLikes(id, respch, wg)
	go getFriends(id, respch, wg)
	wg.Add(3)
	wg.Wait()
	close(respch)

	userProfile := &UserProfile{}
	for resp := range respch {
		if resp.err != nil {
			return nil, resp.err
		}
		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		case []string:
			userProfile.Comments = msg
		}
	}

	return userProfile, nil
}

func getComments(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	comments := []string{
		"Hey, that was great",
		"Yeah Buddy",
		"Ow, I didnt know that",
	}
	respch <- Response{
		data: comments,
		err:  nil,
	}

	wg.Done()
}

func getLikes(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)

	respch <- Response{
		data: 14,
		err:  nil,
	}

	wg.Done()
}

func getFriends(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)

	friendIds := []int{12, 43, 65, 87}
	respch <- Response{
		data: friendIds,
		err:  nil,
	}

	wg.Done()
}

func main() {
	start := time.Now()
	userProfile, err := handleGetUserProfile(10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userProfile)
	fmt.Println("fetching te user profile took", time.Since(start))
}
