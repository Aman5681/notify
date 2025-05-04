package handlers

import "github.com/Aman5681/notify/internal/payload"

func HandleGenerate(payload payload.Payload) (string, error) {
	return "Pretend I searched your code and found the answer! aman", nil
}
