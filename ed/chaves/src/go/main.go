package main

import (
	"bufio"
	
	"fmt"
	"os"
)

func (q *Queue[T]) Size() int {
	return q.items.Len()
}



func main() {
	queue := NewQueue[string]()

	for c := 'A'; c <= 'P'; c++ {
		queue.Enqueue(string(c))
	}

	scanner := bufio.NewScanner(os.Stdin)


	for queue.Size() > 1 {
		teamLeft := queue.Dequeue()
		teamRight := queue.Dequeue()

		if scanner.Scan() {
			line := scanner.Text()
			var goalsLeft, goalsRight int
			fmt.Sscanf(line, "%d %d", &goalsLeft, &goalsRight)

			
			if goalsLeft > goalsRight {
				
				queue.Enqueue(teamLeft)
			} else {
				
				queue.Enqueue(teamRight)
			}

		
		}
	}

	fmt.Println(queue.Dequeue())
}