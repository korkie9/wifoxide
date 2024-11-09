package main

import (
	"fmt"
	"os/exec"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// TODO: Move this to models
type Button struct {
	x      int
	y      int
	width  int
	height int
}

func main() {
	var screenWidth int32 = 800
	var screenHeight int32 = 450
	rl.InitWindow(screenWidth, screenHeight, "Wifoxide")
	defer rl.CloseWindow()
	var index int = 0

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

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		// rl.DrawRectangle((screenWidth+400)/2, screenHeight-85, 90, 30, rl.LightGray)

		var buttonHeight float32 = 60
		var buttonSpacing float32 = 10
		var ycoord float32 = 20
		ssids := []string{"helloo", "sdfsdf", "hwewe"}
		for i := 0; i < len(ssids); i++ {
			buttonColor := rl.DarkGray
			buttonHoverColor := rl.LightGray
			button := rl.Rectangle{X: 50,
				Y: ycoord + (buttonHeight+buttonSpacing)*float32(i), Width: 600,
				Height: buttonHeight}

			if rl.CheckCollisionPointRec(rl.GetMousePosition(), button) {
				buttonColor = buttonHoverColor
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) && !(len(ssids) > 0) {
					// cmd := ssids[0].path + " -P " + profiles[index].name + " &"
					// openBrowserAndCloseProgram(db, &convertableTexture, cmd)
				}
			} else {
				buttonColor = rl.DarkGray
			}
			if index == i {
				buttonColor = rl.LightGray
			}
			rl.DrawRectangleRounded(button, 0.3, 10, buttonColor)
			rl.DrawText(ssids[i], int32(button.X)+30, int32(button.Y)+15, 20, rl.Black)
		}

		rl.EndDrawing()
	}
}
