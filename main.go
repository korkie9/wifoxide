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
	var ssids = []string{"helloo", "sdfsdf", "hwewe"}
	var showTextBox = true
	var profileTextBox = rl.Rectangle{X: float32(screenWidth)/2.0 - 350, Y: 180, Width: float32(screenWidth) - 125, Height: 50}
	var text string = ""
	rl.SetExitKey(rl.KeyNull)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		if showTextBox {
			key := rl.GetCharPressed()
			if (key >= 32) && (key <= 125) && len(text) < 35 {
				text = fmt.Sprintf("%s%s", text, string(key))
			}

			if rl.IsKeyPressed(rl.KeyBackspace) {
				if len(text) > 0 {
					text = text[:len(text)-1]
				}
			}
			if rl.IsKeyPressed(rl.KeyEscape) {
				showTextBox = false
			}

		}

		if (rl.IsKeyPressed(rl.KeyJ) || rl.IsKeyPressed(rl.KeyDown)) && !showTextBox {
			if index+1 < len(ssids) {
				index++
			}
		}
		if (rl.IsKeyPressed(rl.KeyK) || rl.IsKeyPressed(rl.KeyDown)) && !showTextBox {
			if index-1 > -1 {
				index--
			}
		}

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

		var buttonHeight float32 = 60
		var buttonSpacing float32 = 10
		var ycoord float32 = 20
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
		if showTextBox {
			rl.DrawRectangleRounded(profileTextBox, 0.3, 10, rl.LightGray)
			rl.DrawRectangleRoundedLines(profileTextBox, 0.3, 10, 2.0, rl.Blue)
			rl.DrawText(text, int32(profileTextBox.X)+5,
				int32(profileTextBox.Y)+8, 40, rl.Maroon)
		}

		rl.EndDrawing()
	}
}
