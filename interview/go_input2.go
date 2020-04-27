package main
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        s := scanner.Text()
		if len(s) == 0 {
			break
		}
        fmt.Println(s)
    }
    if err := scanner.Err(); err != nil {
        os.Exit(1)
    }
}
