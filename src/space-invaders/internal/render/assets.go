package render

import (
	"github.com/TobiAdeniyi/space-invaders/internal/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Assets struct {
	smallEnemyTexture  [2]rl.Texture2D
	mediumEnemyTexture [2]rl.Texture2D
	largeEnemyTexture  [2]rl.Texture2D
	// enemyDethTextures  [3]rl.Texture2D
}

func LoadAssets() Assets {
	assets := Assets{}

	// step 1: Load enemy textures
	// step 2: Resize enemy images to fit the game
	// step 3: Assign textures to the assets struct
	// step 4: Unload images after loading textures
	smallEnemyImage1 := rl.LoadImage("assets/invaders/small-enemy1.png")
	smallEnemyImage2 := rl.LoadImage("assets/invaders/small-enemy2.png")
	rl.ImageResize(smallEnemyImage2, game.SMALL_ENEMY_WIDTH, game.ENEMY_HEIGHT)
	rl.ImageResize(smallEnemyImage1, game.SMALL_ENEMY_WIDTH, game.ENEMY_HEIGHT)
	rl.ImageColorTint(smallEnemyImage1, rl.White)
	rl.ImageColorTint(smallEnemyImage2, rl.White)
	assets.smallEnemyTexture = [2]rl.Texture2D{
		rl.LoadTextureFromImage(smallEnemyImage1),
		rl.LoadTextureFromImage(smallEnemyImage2),
	}
	rl.UnloadImage(smallEnemyImage1)
	rl.UnloadImage(smallEnemyImage2)

	mediumEnemyImage1 := rl.LoadImage("assets/invaders/medium-enemy1.png")
	mediumEnemyImage2 := rl.LoadImage("assets/invaders/medium-enemy2.png")
	rl.ImageResize(mediumEnemyImage1, game.MEDIUM_ENEMY_WIDTH, game.ENEMY_HEIGHT)
	rl.ImageResize(mediumEnemyImage2, game.MEDIUM_ENEMY_WIDTH_2, game.ENEMY_HEIGHT)
	assets.mediumEnemyTexture = [2]rl.Texture2D{
		rl.LoadTextureFromImage(mediumEnemyImage1),
		rl.LoadTextureFromImage(mediumEnemyImage2),
	}
	rl.UnloadImage(mediumEnemyImage1)
	rl.UnloadImage(mediumEnemyImage2)

	largeEnemyImage1 := rl.LoadImage("assets/invaders/large-enemy1.png")
	largeEnemyImage2 := rl.LoadImage("assets/invaders/large-enemy2.png")
	rl.ImageResize(largeEnemyImage1, game.LARGE_ENEMY_WIDTH, game.ENEMY_HEIGHT)
	rl.ImageResize(largeEnemyImage2, game.LARGE_ENEMY_WIDTH, game.LARRGE_ENEMY_HEIGHT_2)
	assets.largeEnemyTexture = [2]rl.Texture2D{
		rl.LoadTextureFromImage(largeEnemyImage1),
		rl.LoadTextureFromImage(largeEnemyImage2),
	}
	rl.UnloadImage(largeEnemyImage1)
	rl.UnloadImage(largeEnemyImage2)

	// enemyDeathImage1 := rl.LoadImage("assets/invaders/enemy-death1.png")
	// enemyDeathImage2 := rl.LoadImage("assets/invaders/enemy-death2.png")
	// enemyDeathImage3 := rl.LoadImage("assets/invaders/enemy-death3.png")

	return assets
}
