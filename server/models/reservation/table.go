package reservation

import (
	"github.com/WaynerEP/restaurant-app/server/models/common"
)

type Table struct {
	common.ModelId
	TableNumber string `json:"tableNumber" gorm:"size:30;not null"`
	Description string `json:"description" gorm:"size:255"`
}

/*Piso:
Activo: El piso está abierto y disponible para uso.
Inactivo: El piso está cerrado o fuera de servicio por mantenimiento u otras razones.
En limpieza: El piso está siendo limpiado y no está disponible para uso.
Reservado: El piso ha sido reservado para un evento o función específica.
Ocupado: El piso está ocupado por clientes.
Cerrado: El piso está cerrado temporalmente por razones como feriados o cierres temporales.
Ambiente:
Activo: El ambiente está abierto y disponible para uso.
Inactivo: El ambiente está cerrado o fuera de servicio por mantenimiento u otras razones.
Reservado: El ambiente ha sido reservado para un evento o función específica.
Ocupado: El ambiente está ocupado por clientes.
Cerrado: El ambiente está cerrado temporalmente por razones como feriados o cierres temporales.
Mesa:
Disponible: La mesa está disponible para ser reservada o utilizada.
Reservada: La mesa ha sido reservada para un cliente en un momento específico.
Ocupada: La mesa está siendo utilizada por clientes.
Sucia: La mesa necesita ser limpiada antes de ser utilizada nuevamente.
Fuera de servicio: La mesa está fuera de servicio debido a daños u otros problemas.
Cerrada: La mesa está temporalmente fuera de servicio por razones como feriados o cierres temporales.*/
