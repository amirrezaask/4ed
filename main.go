package main

import (
	"flag"
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.design/x/clipboard"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "cfg", path.Join(os.Getenv("HOME"), ".core"), "path to config file, defaults to: ~/.core.cfg")
	flag.Parse()

	// read config file
	cfg, err := readConfig(configPath)
	if err != nil {
		panic(err)
	}

	// basic setup
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowMaximized)
	rl.SetTraceLogLevel(rl.LogError)
	rl.InitWindow(1920, 1080, "4ed")

	if err := clipboard.Init(); err != nil {
		panic(err)
	}
	defer rl.CloseWindow()
	rl.SetTargetFPS(30)
	// create editor
	editor := Application{
		LineNumbers:  true,
		LineWrapping: true,
		Colors:       cfg.Colors,
	}

	err = loadFont(cfg.FontName, 20)
	if err != nil {
		panic(err)
	}

	filename := ""
	if len(flag.Args()) > 0 {
		filename = flag.Args()[0]
	}
	rl.SetTextLineSpacing(int(fontSize))

	// initialize first editor
	textEditorBuffer, err := NewEditorBuffer(EditorBufferOptions{
		Filename:       filename,
		LineNumbers:    true,
		TabSize:        4,
		MaxHeight:      int32(rl.GetRenderHeight()),
		MaxWidth:       int32(rl.GetRenderWidth()),
		Colors:         editor.Colors,
		CursorBlinking: false,
	})
	if err != nil {
		panic(err)
	}

	editor.Editors = append(editor.Editors, textEditorBuffer)

	for !rl.WindowShouldClose() {
		editor.HandleWindowResize()
		editor.HandleMouseEvents()
		editor.HandleKeyEvents()
		editor.Render()
	}

}
