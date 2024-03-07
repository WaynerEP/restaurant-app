package initialize

import (
	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/WaynerEP/restaurant-app/server/utils"
	"github.com/songzhibin97/gkit/cache/local_cache"
)

// OtherInit realiza la inicialización de configuraciones adicionales.
// Parsea la duración de expiración del token JWT y la duración de buffer.
// Crea una nueva instancia de caché usando la duración de expiración como valor predeterminado.
// Si hay errores al analizar, la función entra en pánico.
func OtherInit() {
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}
