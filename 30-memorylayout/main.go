package main

import (
	"fmt"
	"unsafe"
)

// string size always 16 byte is golang
func main() {
	t1 := T1{ID: 11, name: "Anish", isMarried: false, address: "Bangalore"}
	fmt.Println("t1:", t1, "Size of t1:", unsafe.Sizeof(t1))

	t2 := T2{name: "Anish", address: "Bangalore", id: 1, isMarried: true, isVoted: false, ref: 23}
	fmt.Println("t2:", t2, "Size of t2:", unsafe.Sizeof(t2))

	arr1 := []string{"How", "are", "you", "doing"}
	//size := unsafe.Sizeof(arr1[0])

	arrptr := uintptr(unsafe.Pointer(&arr1[0]))
	//fmt.Println(*(*string)(unsafe.Pointer(arrptr)))

	for i := 0; i <= 3; i++ {
		fmt.Println(*(*string)(unsafe.Pointer(arrptr)))
		arrptr += 16 //string stores 16 byte

	}
}

// Ordered of variable define memeory layout work here
type T1 struct {
	ID        uint16 //8 byte   ---> 2 --6 left waste
	name      string //16 byte  --->16 --0 left
	isMarried bool   //8 byte   ---->8 --7 left waste
	// ---->> why this is 8 byte beacuse tha maximum is 16 byte in this struct so just after 16 max is 8 byte will assign to other varibale called padding bits
	address string //16 byte  ----->16 --0 left
} //---> total size is 48 byte

type T2 struct {
	name      string //16 ---->16 --0
	address   string //16 ---->16 --0
	id        uint16 //8 ----->2 --- 6 byte left
	isMarried bool   //user form above 6 byte ----->1 --- 5 byte letf wasted(unsed)
	ref       uint32 //------------------>4 ----> 0 byte left
	isVoted   bool   //--------------->1 --- >4 byte unsed

} // this total 48 bytes

type T3 struct {
	name      string // 16 bytes  | Offset: 0  → No padding needed
	address   string // 16 bytes  | Offset: 16 → No padding needed
	id        uint16 //  2 bytes  | Offset: 32 → 6 bytes padding needed (to align the next field)
	isMarried bool   //  1 byte   | Offset: 34 → Uses 1 byte, 5 bytes padding left
	isVoted   bool   //  1 byte   | Offset: 35 → Uses 1 byte, 4 bytes padding left
	ref       uint32 //  4 bytes  | Offset: 36 → Fits perfectly (no padding needed)
} // Total = 40 bytes ✅ (Efficient)

type T4 struct {
	name      string // 16 bytes  | Offset: 0  → No padding needed
	address   string // 16 bytes  | Offset: 16 → No padding needed
	id        uint16 //  2 bytes  | Offset: 32 → 6 bytes padding needed (to align the next field)
	isMarried bool   //  1 byte   | Offset: 34 → Uses 1 byte, 5 bytes padding left
	ref       uint32 //  4 bytes  | Offset: 36 → Fits perfectly (no padding needed)
	isVoted   bool   //  1 byte   | Offset: 40 → **New alignment rule forces an 8-byte boundary**
} // Total = 48 bytes ❌ (Less efficient)

/*
	Offset Calculation in Struct Alignment
At Offset 34:

isMarried bool (1 byte) occupies offset 34.

Offset 35 is still available.

At Offset 35:

isVoted bool (1 byte) could be placed here, but Golang prefers natural alignment for uint32.

At Offset 36:

ref uint32 (4 bytes) must be aligned to a 4-byte boundary.

Since offset 36 is already a multiple of 4, ref can directly go here.

At Offset 40:

isVoted bool (1 byte) could have been placed at 35, but Golang's alignment strategy pushes it to 40.

Because isVoted is alone at offset 40, and Golang rounds up to the next multiple of 8, it wastes extra 7 bytes.

Why Not Place isVoted at Offset 35?
Golang groups smaller types together when possible, but prefers natural alignment for larger types like uint32.
Since ref uint32 must be at offset 36, it leaves offset 35 unused instead of breaking alignment.

Fix: Better Ordering to Avoid Wasting Space
To make it more memory efficient, reorder the struct:

type T2 struct {
    name      string // 16 bytes  | Offset: 0
    address   string // 16 bytes  | Offset: 16
    id        uint16 //  2 bytes  | Offset: 32
    isMarried bool   //  1 byte   | Offset: 34
    isVoted   bool   //  1 byte   | Offset: 35
    ref       uint32 //  4 bytes  | Offset: 36 (naturally aligned)
} // Total = 40 bytes ✅ (No wasted space)
Now:

isVoted uses offset 35 (instead of getting pushed to 40).

ref remains properly aligned at 36.

No extra padding needed.

💡 Always group smaller fields together before adding larger fields to avoid wasted padding!
*/
