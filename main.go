package main

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/exec"
	"time"
)

func UUID() string {
	milli := time.Now().UnixMilli()
	timeBuff := make([]byte, 8)
	binary.BigEndian.PutUint64(timeBuff, uint64(milli))
	randBuff := make([]byte, 8)
	_, _ = rand.Read(randBuff)
	buff := append(timeBuff, randBuff...)
	return base58.Encode(buff)
}

func main() {
	_, err := os.Stat("./temp")
	if os.IsNotExist(err) {
		err = os.Mkdir("./temp", 0777)
		if err != nil {
			log.Println(err)
			return
		}
	}

	app := fiber.New()

	api := app.Group("/fiber")

	api.Get("/stream_speech_trans", websocket.New(func(conn *websocket.Conn) {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				break
			}
			source := "./temp/" + UUID() + ".webm"
			err = os.WriteFile(source, message, 0777)
			if err != nil {
				log.Println(err)
				break
			}
			target := "./temp/" + UUID() + ".wav"
			log.Println(source)
			// ffmpeg -i input.webm -acodec pcm_s16le -ar 16000 -ac 1 output.wav
			cmd := exec.Command("ffmpeg", "-i", source, "-acodec", "pcm_s16le", "-ar", "16000", "-ac", "1", target)
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Println(string(output), err)
				break
			}
			_ = os.Remove(source)
		}
	}))

	log.Fatal(app.Listen(":8080"))
}
