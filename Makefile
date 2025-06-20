.PHONY: start-emulator
start-emulator:
	gcloud beta emulators pubsub start --project=local

.PHONY: run
run:
	go run main.go
	