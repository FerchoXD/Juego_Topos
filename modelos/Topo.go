package modelo

import (
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Estructura Topo para el modelo de topo
type Topo struct {
	boton   *widget.Button
	ventana fyne.Window
	vida    bool
}

// NuevoTopoModelo crea un nuevo modelo de topo
func NuevoTopoModelo(window fyne.Window, botonFinal *widget.Button, etiquetaFinal *canvas.Text, timer *Timer) *Topo {
	// Creamos un nuevo modelo de topo
	t := &Topo{
		boton:   widget.NewButton("", nil),
		ventana: window,
		vida:    true,
	}

	// Configuramos la función OnTapped del botón
	t.boton.OnTapped = func() {
		// El topo se ha pulsado
		t.vida = false
		// Cambiamos la etiqueta final para indicar el tiempo sobrante
		etiquetaFinal.Text = "Tiempo sobrante: " + strconv.FormatUint(uint64(timer.tiempo), 10) + " segundos"
		etiquetaFinal.Refresh()
		// Mostramos el botón final y la etiqueta final
		botonFinal.Show()
		etiquetaFinal.Show()
		// Iniciamos el parpadeo de la etiqueta y el botón final en un tercer hilo
		go parpadear(etiquetaFinal, botonFinal) // Tercer hilo (decorador)
		// Habilitamos el botón del topo
		t.boton.Enable()
	}

	// Redimensionamos el botón del topo
	t.boton.Resize(fyne.NewSize(50, 70))

	// Devolvemos un nuevo modelo de topo
	return t
}

// CrearContenedor crea un contenedor para el topo y la imagen del topo
func (t *Topo) CrearContenedor(image *canvas.Image) *fyne.Container {
	// Creamos un contenedor principal sin un layout
	contenedorPrincipal := container.NewWithoutLayout(t.boton)

	// Agregamos la imagen del topo al contenedor principal
	contenedorPrincipal.Add(image)

	// Redimensionamos y posicionamos el contenedor principal
	contenedorPrincipal.Resize(fyne.NewSize(50, 40))
	contenedorPrincipal.Move(fyne.NewPos(350, 260))

	// Devolvemos el contenedor principal
	return contenedorPrincipal
}

// Mover mueve el topo aleatoriamente
func (t *Topo) Mover(container *fyne.Container) {
	// Mientras el topo esté vivo
	for t.vida {
		// Creamos una nueva posición aleatoria
		newX := float32(rand.Intn(750))
		newY := float32(rand.Intn(550))

		// Movemos el contenedor del topo a la nueva posición
		container.Move(fyne.NewPos(newX, newY))

		// Esperamos 1 segundo
		time.Sleep(500 * time.Millisecond)
	}

	// Deshabilitamos el botón del topo
	t.boton.Disable()
}

// parpadear hace parpadear la etiqueta final y el botón final
func parpadear(etiquetaFinal *canvas.Text, botonFinal *widget.Button) {
	// Infinitamente
	for {
		// Mostramos la etiqueta final y el botón final
		time.Sleep(1 * time.Second)
		etiquetaFinal.Show()
		botonFinal.Show()

		// Esperamos 2 segundos y ocultamos la etiqueta final y el botón final
		time.Sleep(2 * time.Second)
		etiquetaFinal.Hide()
		botonFinal.Hide()
	}
}
