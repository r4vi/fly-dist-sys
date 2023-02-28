package main

import (
	"encoding/json"
	"fmt"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"log"
	"math/rand"
	"os"
	"time"
)

func getUnique() string {
	r := rand.Int31()
	pid := os.Getppid()
	ms := time.Now().UnixNano()
	return fmt.Sprintf("%v-%v-%v", ms, pid, r)
}

func main() {

	n := maelstrom.NewNode()
	n.Handle("generate", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Echo the original message back with the updated message type.
		return n.Reply(msg, map[string]any{"type": "generate_ok", "id": getUnique()})
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
