package main

import "fmt"

// Node mewakili sebuah elemen dalam list
type Node struct {
    Value int
    Next  *Node
}

// List mewakili sebuah list
type List struct {
    Head *Node
}

// Append menambahkan semua elemen dari list kedua ke akhir list pertama
func (l *List) Append(other *List) {
    // Jika list pertama kosong, setel head-nya ke head list kedua
    if l.Head == nil {
        l.Head = other.Head
        return
    }

    // Cari node terakhir dalam list pertama
    current := l.Head
    for current.Next != nil {
        current = current.Next
    }

    // Tambahkan list kedua ke akhir list pertama
    current.Next = other.Head
}

// Concat menggabungkan semua elemen dari beberapa list menjadi satu list
func Concat(lists []*List) *List {
    result := &List{}
    for _, list := range lists {
        result.Append(list)
    }
    return result
}

// Filter mengembalikan list baru yang hanya berisi elemen yang memenuhi kondisi
func (l *List) Filter(predicate func(int) bool) *List {
    result := &List{}
    current := l.Head
    for current != nil {
        if predicate(current.Value) {
            result.Append(&List{Head: &Node{Value: current.Value}})
        }
        current = current.Next
    }
    return result
}

// Map mengembalikan list baru yang berisi hasil penerapan fungsi pada setiap elemen
func (l *List) Map(function func(int) int) *List {
    result := &List{}
    current := l.Head
    for current != nil {
        result.Append(&List{Head: &Node{Value: function(current.Value)}})
        current = current.Next
    }
    return result
}

// Reverse mengembalikan list baru dengan elemen-elemen asli dalam urutan terbalik
func (l *List) Reverse() *List {
    result := &List{}
    current := l.Head
    for current != nil {
        result.Append(&List{Head: &Node{Value: current.Value}})
        current = current.Next
    }
    return result
}

// Print mencetak elemen-elemen list
func (l *List) Print() {
    current := l.Head
    for current != nil {
        fmt.Print(current.Value, " ")
        current = current.Next
    }
    fmt.Println()
}

func main() {
    // Contoh penggunaan
    list1 := &List{Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3}}}}
    list2 := &List{Head: &Node{Value: 4, Next: &Node{Value: 5, Next: &Node{Value: 6}}}}

    fmt.Println("Append:")
    list1.Append(list2)
    list1.Print()

    fmt.Println("Concat:")
    lists := []*List{{Head: &Node{Value: 1, Next: &Node{Value: 2}}}, {Head: &Node{Value: 3, Next: &Node{Value: 4}}}, {Head: &Node{Value: 5, Next: &Node{Value: 6}}}}
    result := Concat(lists)
    result.Print()

    fmt.Println("Filter:")
    list := &List{Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: 4, Next: &Node{Value: 5, Next: &Node{Value: 6}}}}}}}
    result = list.Filter(func(x int) bool { return x%2 == 0 })
    result.Print()

    fmt.Println("Map:")
    list = &List{Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3}}}}
    result = list.Map(func(x int) int { return x * 2 })
    result.Print()

    fmt.Println("Reverse:")
    list = &List{Head: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: 4}}}}}
    result = list.Reverse()
    result.Print()
}