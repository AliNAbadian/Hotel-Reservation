package main

import "fmt"

type Room struct {
	ID       int
	Type     string
	Status   bool
	BedCount int
	Price    int
}

var Rooms []Room = GenerateRooms()

func main() {

	input := ""

	for input != "exit" {
		fmt.Println("Choose an option:")
		fmt.Println("1. Room list")
		fmt.Println("2. Add Room")
		fmt.Println("3. Reserve Room ")
		fmt.Println("4. Exit")
		fmt.Scanln(&input)
		switch input {
		case "1":
			GetRoomList()
		case "2":
			AddRoom()
		case "3":
			ReserveRoom()
		case "4":
			fmt.Println("Exiting...")
			break
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}

}

func GetRoomList() {
	for _, room := range Rooms {
		fmt.Printf("%+v\n", room)
	}

}
func GetRoomFromInput() Room {
	var room Room = Room{Status: false}
	fmt.Println("Enter room line by line:")
	fmt.Scanln(&room.ID)
	fmt.Scanln(&room.Type)
	fmt.Scanln(&room.BedCount)
	fmt.Scanln(&room.Price)

	return room

}

func AddRoom() {
	room := GetRoomFromInput()

	Rooms = append(Rooms, room)
}

func ReserveRoom() {

}

func CalculateRoomPrice() {

}

// GenerateRooms creates a list of Room instances with predefined data.
//
// The function returns a slice of Room instances. Each Room instance has an ID, Type, Status, BedCount, and Price.
// The Status field indicates whether the room is available (true) or not (false).
//
// The function initializes a slice of Room instances and appends 10 predefined Room instances to it.
// The predefined Room instances have the following properties:
//
// Room ID | Type  | Status | BedCount | Price
// --------|-------|--------|----------|------
// 1       | Single| true   | 1        | 100
// 2       | Double| true   | 2        | 200
// 3       | Suite | true   | 4        | 300
// 4       | Single| false  | 1        | 100
// 5       | Double| false  | 2        | 200
// 6       | Suite | false  | 4        | 300
// 7       | Single| true   | 1        | 100
// 8       | Double| true   | 2        | 200
// 9       | Suite | true   | 4        | 300
// 10      | Single| false  | 1        | 100
//
// The function does not take any parameters.
func GenerateRooms() []Room {
	rooms := []Room{}

	rooms = append(rooms, Room{ID: 1, Type: "Single", Status: true, BedCount: 1, Price: 100})
	rooms = append(rooms, Room{ID: 2, Type: "Double", Status: true, BedCount: 2, Price: 200})
	rooms = append(rooms, Room{ID: 3, Type: "Suite", Status: true, BedCount: 4, Price: 300})
	rooms = append(rooms, Room{ID: 4, Type: "Single", Status: false, BedCount: 1, Price: 100})
	rooms = append(rooms, Room{ID: 5, Type: "Double", Status: false, BedCount: 2, Price: 200})
	rooms = append(rooms, Room{ID: 6, Type: "Suite", Status: false, BedCount: 4, Price: 300})
	rooms = append(rooms, Room{ID: 7, Type: "Single", Status: true, BedCount: 1, Price: 100})
	rooms = append(rooms, Room{ID: 8, Type: "Double", Status: true, BedCount: 2, Price: 200})
	rooms = append(rooms, Room{ID: 9, Type: "Suite", Status: true, BedCount: 4, Price: 300})
	rooms = append(rooms, Room{ID: 10, Type: "Single", Status: false, BedCount: 1, Price: 100})

	return rooms
}
