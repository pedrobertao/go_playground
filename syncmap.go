package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-faker/faker/v4"
)

var syncMap sync.Map
var simpleMap map[string]User

type User struct {
	Name  string `faker:"name"`
	UUID  string `faker:"uuid_digit"`
	Email string `faker:"email"`
}

func syncmap() {
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}
	syncMap = sync.Map{}
	simpleMap = make(map[string]User)
	iterations := 100

	for i := 0; i < iterations; i++ {
		u := User{}
		err := faker.FakeData(&u)
		if err != nil {
			fmt.Println(err)
		}

		syncMap.Store(u.UUID, u)
		simpleMap[u.UUID] = u
	}
	t1 := 0.0
	t2 := 0.0

	wg.Add(1)
	go func() {
		start := time.Now()
		defer wg.Done()
		syncMap.Range(func(key, value any) bool {
			if user, ok := value.(User); ok {
				mu.Lock()
				user.Name = user.Name + " test"
				syncMap.Swap(key, user)
				mu.Unlock()
			}
			return true
		})
		t1 = time.Since(start).Seconds()
		fmt.Println("Sync map done in: ", t1)

	}()

	// Time to get all from simple Map
	wg.Add(1)
	go func() {
		start := time.Now()
		defer wg.Done()
		for key, user := range simpleMap {
			mu.Lock()
			user.Name = user.Email + " test"
			simpleMap[key] = user
			mu.Unlock()
		}
		t2 = time.Since(start).Seconds()
		fmt.Println("Simple map done in: ", t2)
	}()

	fmt.Println("Waiting....")
	wg.Wait()
	fmt.Println("Result Dif(sync - simple): ", t1-t2)
	fmt.Println("Bye")
}
