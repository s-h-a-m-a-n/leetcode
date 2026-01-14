package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	keys := rooms[0]
	visitedRooms := map[int]struct{}{}
	visitedRooms[0] = struct{}{}
	for len(keys) > 0 {
		if _, ok := visitedRooms[keys[0]]; !ok {
			visitedRooms[keys[0]] = struct{}{}
			keys = append(keys, rooms[keys[0]]...)
		}
		keys = keys[1:]
	}

	return len(rooms) == len(visitedRooms)
}

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 0}, {3, 2, 1}, {2}, {0}}))

}
