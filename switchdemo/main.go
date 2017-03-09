// Twelve days of Christmas:
// https://en.wikipedia.org/wiki/The_Twelve_Days_of_Christmas_(song)
package main

import "fmt"

func main() {
    fmt.Println("On the 1 of Christmas my true love sent to me:")
    fmt.Println("A Pattridge in a Peer Tree")
    fmt.Println("")

    for day :=2; day <=12; day++ {
        fmt.Println("On the",day,"of Christmas my true love sent to me:")
        switch day {
        case 12:
            fmt.Println("Twelve Drummers Drumming,")
            fallthrough
        case 11:
            fmt.Println("Eleven Pipers Piping,")
            fallthrough
        case 10:
            fmt.Println("Ten Lords Leaping,")
            fallthrough
        case 9:
            fmt.Println("Nine Ladies Dancing,")
            fallthrough
        case 8:
            fmt.Println("Eight Maids Milking,")
            fallthrough
        case 7:
            fmt.Println("Seven Swans Swimming,")
            fallthrough
        case 6:
            fmt.Println("Six Geese Laying,")
            fallthrough
        case 5:
            fmt.Println("Five Gold Rings,")
            fallthrough
        case 4:
            fmt.Println("Four Calling Birds,")
            fallthrough
        case 3:
            fmt.Println("Three French Hens,")
            fallthrough
        case 2:
            fmt.Println("Two Turtle Doves,")
            fallthrough
        default:
            fmt.Println("and a and a Partridge in a Pear Tree.")
        }
        fmt.Println("")
    }
}
