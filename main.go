package main

const (
	VERSION       = "0.1.0"
	SCREEN_WIDTH  = 1024
	SCREEN_HEIGHT = 768
)

func main() {
	InitLog()
	Log.Info("Go Loi go! 👾")
	game, err := NewGame(SCREEN_WIDTH, SCREEN_HEIGHT)
	if err != nil {
		Log.Error(err)
		return
	}

	game.Run()
}
