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
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}

}

func GetRoomList() {

}

func AddRoom() {

}

func ReserveRoom() {

}

func CalculateRoomPrice() {

}

func GetRoomFromInput() {

}

func GenerateRooms() []Room {
	rooms := make([]Room, 20)

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

}
