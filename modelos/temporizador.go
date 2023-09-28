package modelo

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// Estructura Timer para el temporizador
type Timer struct {
	tiempo uint
}

// NuevoTemporizadorModelo crea un nuevo modelo de temporizador con un tiempo inicial de 30 segundos
func NuevoTemporizadorModelo() *Timer {
	return &Timer{
		tiempo: 30,
	}
}

// IniciarTemporizador inicia el temporizador
func (t *Timer) IniciarTemporizador(topo *Topo, etiqueta *canvas.Text, botonFinal *widget.Button, etiquetaFinal *canvas.Text) {
	// Mientras el tiempo sea mayor a 0 y el topo tenga vida
	for t.tiempo > 0 && topo.vida {
		// Esperamos 1 segundo
		time.Sleep(time.Second * 1)
		// Reducimos el tiempo
		t.tiempo--
		// Creamos la cadena para la etiqueta del tiempo restante
		cadena := "Tiempo Restante: " + strconv.FormatUint(uint64(t.tiempo), 10) + " segundos"
		// Actualizamos la etiqueta
		etiqueta.Text = cadena
		etiqueta.Refresh()
	}
	// El topo ha perdido todo su tiempo
	topo.vida = false
	// Cambiamos la etiqueta para indicar que se ha acabado el tiempo
	cadena := "Se ha acabado!!!"
	etiqueta.Text = cadena
	etiqueta.Refresh()
	etiquetaFinal.Text = "Tiempo sobrante: " + strconv.FormatUint(uint64(t.tiempo), 10) + " segundos"
	etiquetaFinal.Refresh()
	// Mostramos el botón final y la etiqueta final
	botonFinal.Show()
	etiquetaFinal.Show()
	go parpadear(etiquetaFinal, botonFinal)
}
