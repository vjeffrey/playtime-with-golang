install:
	glide install

play:
	go build . && ./game-time

clean:
	rm -rf game-time && rm -rf debug && rm -rf vendor
