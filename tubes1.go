package main

import (
	"fmt"
	"math/rand"
	"time"
)

const Nmax = 80

type Passenger struct {
	Name      string
	Money     int
	BookingID int

	Reservation Reservation
	BookingTime time.Time
}

type Reservation struct {
	Train  string
	Seat   int
	Status string
}
type Pass [Nmax]Passenger

func inputData(per *Pass, n *int) {
	var ticket int
	for ticket != 4 {
		if *n > Nmax {
			*n = Nmax
		}
		fmt.Print("Input Name :")
		fmt.Scan(&per[*n].Name)
		fmt.Print("Input your Money :")
		fmt.Scan(&per[*n].Money)
		fmt.Print("Reservation Train (A/B) :")
		fmt.Scan(&per[*n].Reservation.Train)

		per[*n].Reservation.Seat = findAvailableSeat(per, per[*n].Reservation.Train)
		if per[*n].Reservation.Seat == -1 {
			fmt.Println("No available seats in train", per[*n].Reservation.Train)

		}

		per[*n].BookingID = rand.Intn(190)
		per[*n].BookingTime = time.Now()
		per[*n].Reservation.Status = "reserved"

		ticket++
		*n++

	}
}
func findAvailableSeat(per *Pass, train string) int {
	seats := 40
	for i := 1; i <= seats; i++ {
		seatOccupied := false
		isfound := false
		for j := 0; !isfound && j < Nmax; j++ {
			if per[j].Reservation.Train == train && per[j].Reservation.Seat == i {
				seatOccupied = true
				isfound = true
			}
		}
		if !seatOccupied {
			for j := 0; j < Nmax; j++ {
				if per[j].Reservation.Train == train && per[j].Reservation.Seat == 0 {
					per[j].Reservation.Seat = i
					return i
				}
			}
		}
	}
	return -1 // No available seats within the specified train
}

func Pay(A *Pass, n *int, bookID int) {
	for i := 0; i < *n; i++ {
		if A[i].BookingID == bookID {
			now := time.Since(A[i].BookingTime)
			if now <= 6*time.Minute {
				if A[i].Reservation.Train == "A" {
					A[i].Money -= 150000
					if A[i].Money < 0 {
						fmt.Print("Sorry you dont have enough money to train\n")
					} else {
						A[i].Reservation.Status = "paid"
						fmt.Println("Payment successful for Booking ID:", A[i].BookingID)
					}
				} else if A[i].Reservation.Train == "B" {
					A[i].Money -= 100000
					if A[i].Money < 0 {
						fmt.Print("Sorry you dont have enough money to train\n")
					} else {
						A[i].Reservation.Status = "paid"
						fmt.Println("Payment successful for Booking ID:", A[i].BookingID)
					}
				} else {
					fmt.Print("Invalid Train Reservation")
				}
			} else {
				fmt.Print("The Booking ID: ", A[i].BookingID, "Payment is not valid because its exceed the Time for the transaction\n")
				A[i] = A[i+1]
				*n--
				i--
				//for j := i; j < *n-1; j++ {
				//per[j] = per[j+1]
				//}
				//*n--

			}
		}
	}
}
func Book(A *Pass, n int) {
	var name string
	var BookID int
	var paid bool
	fmt.Print("Please enter the name and the Booking ID: ")
	fmt.Scan(&name, &BookID)

	for i := 0; i < n; i++ {
		if A[i].Reservation.Status == "reserved" {
			if BookID == A[i].BookingID && name == A[i].Name {
				Pay(A, &n, A[i].BookingID)
				paid = true
			}
		} else {
			fmt.Print("You've paid for the reservation\n")
			paid = true
		}

	}
	if !paid {
		fmt.Print("The data you input is not in the array\n")
	}
}
func Display(A Pass, n int) {
	for i := 0; i < n; i++ {
		if A[i].Reservation.Status == "paid" {
			fmt.Printf("Reserved seat %d in train %s for %s. Booking ID: %d\n", A[i].Reservation.Seat, A[i].Reservation.Train, A[i].Name, A[i].BookingID)
		} else {
			fmt.Printf("train %s for %s. Booking ID: %d\n", A[i].Reservation.Train, A[i].Name, A[i].BookingID)
		}
	}
}
func Search(A Pass, n int) bool {
	var nama string
	var bookId, left, right, mid, idx int
	var found bool
	SortingIDAsc(&A, n)
	left = 0
	right = n - 1
	idx = -1
	fmt.Print("Input name and Booking id: ")
	fmt.Scan(&nama, &bookId)
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if A[mid].Name == nama && A[mid].BookingID == bookId {
			fmt.Printf("Reserved seat %d in train %s for %s. Booking ID: %d\n", A[mid].Reservation.Seat, A[mid].Reservation.Train, A[mid].Name, A[mid].BookingID)
			idx = mid
			found = true
		} else if A[mid].BookingID < bookId {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if !found {
		fmt.Print("The id you're looking for is not in the array\n")
	}
	return false
}
func SortingIDAsc(A *Pass, n int) {
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if A[j].BookingID < A[idx].BookingID {
				A[j], A[idx] = A[idx], A[j]
			}
		}
	}
}
func SortingAsc(A *Pass, n int) {
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if A[j].Reservation.Seat < A[idx].Reservation.Seat {
				A[j], A[idx] = A[idx], A[j]
			}
		}
	}
}
func SortingDesc(A *Pass, n int) {
	pass := n - 1
	i := 1
	for i <= pass {
		tempp := A[i]
		j := i - 1
		for j >= 0 && A[j].Reservation.Seat < tempp.Reservation.Seat {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = tempp
		i++
	}
}
func menu() {

}

func main() {
	var per Pass
	var n, opsi int
	var option string
	for opsi != 6 {
		menu()
		fmt.Print("Enter Option")
		fmt.Scan(&opsi)
		if opsi == 1 {
			inputData(&per, &n)
		} else if opsi == 2 {
			Book(&per, n)
		} else if opsi == 3 {
			Display(per, n)
		} else if opsi == 4 {
			Search(per, n)
		} else if opsi == 5 {
			fmt.Print("Ascending / Descending")
			fmt.Scan(&option)
			if option == "Ascending" {
				SortingAsc(&per, n)
			} else {
				SortingDesc(&per, n)
			}
		}
	}

}
