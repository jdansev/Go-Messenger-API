package main

import (
	"fmt"
	"strconv"
	"time"
)

// TEST helpers
func addTestHubs() {

	// Manual creation of hubs
	createTestUsers()

	h := createHub("hub1", "private")
	h.addTestUsersToHub()
	h.addTestMessagesToHub()

	addHub(h)

	// user creation of hubs
	h2 := p1.createHub("p1s-hub", "public")
	go h2.MessageHandler()
	// p1.leaveHub(h2)

}

var p1 *User
var p2 *User
var p3 *User

func createTestUsers() {
	p1 = createUser("testuser1", "secret-key")
	p2 = createUser("testuser2", "secret-key")
	p3 = createUser("testuser3", "secret-key")

	// add friends
	p1.sendFriendRequestTo(p3)
	p3.acceptFriendRequest(p1)

	p2.sendFriendRequestTo(p1)
	p1.declineFriendRequest(p2)

	p1.sendFriendRequestTo(p2)

	p3.sendFriendRequestTo(p1)
	p1.sendFriendRequestTo(p2)

	p2.acceptFriendRequest(p3)

}

func (h *Hub) addTestUsersToHub() {
	h.addUserToHub(p1)
	h.addUserToHub(p2)
	h.addUserToHub(p3)
}

func (h *Hub) addTestMessagesToHub() {
	h.Messages = append(h.Messages, &Message{"1", "hey there", "one", []*UserTag{}})
	h.Messages = append(h.Messages, &Message{"2", "whats up", "two", []*UserTag{}})
	h.Messages = append(h.Messages, &Message{"3", "how's it going", "three", []*UserTag{}})
}

var count = 1

func registerTestUserLoop() {
	for {
		createUser("new-user-"+strconv.Itoa(count), "my-secret-key")
		fmt.Printf("created new user, %d\n", count)
		count++
		time.Sleep(2 * time.Second)
	}
}