package main

import (
	"fmt"
	"os/exec"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var screenWidth int32 = 800
	var screenHeight int32 = 450
	rl.InitWindow(screenWidth, screenHeight, "Wifoxide")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeyEnter) {

			app := "echo"

			arg0 := "-e"
			arg1 := "Hello world"
			arg2 := "\n\tfrom"
			arg3 := "golang"

			cmd := exec.Command(app, arg0, arg1, arg2, arg3)
			stdout, err := cmd.Output()

			fmt.Println(string(stdout))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		// Print the output

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawRectangle((screenWidth+400)/2, screenHeight-85, 90, 30, rl.LightGray)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
