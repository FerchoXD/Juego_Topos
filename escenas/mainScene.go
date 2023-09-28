package escenas

import (
	"image/color"
	modelo "myFirstGame/modelos"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type EscenaPrincipal struct {
	window fyne.Window
}

// Creamos una nueva escena principal del juego
func NuevaEscenaPrincipal(window fyne.Window) *EscenaPrincipal {
	// Devolvemos una nueva instancia de EscenaPrincipal con el campo window definido
	return &EscenaPrincipal{window: window}
}

// Cargamos la escena principal cargar no puede ser llamado mas que por un objeto tipo EscenaPrincipal
func (e EscenaPrincipal) Cargar() {
	// Creamos las imágenes para los topos (3 imágenes)
	img1 := canvas.NewImageFromURI(storage.NewFileURI("./assets/marmota.png"))
	img2 := canvas.NewImageFromURI(storage.NewFileURI("./assets/marmota2.png"))
	img3 := canvas.NewImageFromURI(storage.NewFileURI("./assets/marmota3.png"))

	// Redimensionamos las imágenes
	img1.Resize(fyne.NewSize(50, 70))
	img2.Resize(fyne.NewSize(50, 70))
	img3.Resize(fyne.NewSize(50, 70))

	// Creamos una etiqueta para el temporizador
	etiquetaTiempo := canvas.NewText("Temporizador Restante: 30 Segundos Restantes", color.Black)

	// Creamos un mensaje final para cuando se acabe el tiempo
	mensajeFinal := canvas.NewText("Tu tiempo se acabo", color.Black)

	// Redimensionamos y posicionamos el mensaje final
	mensajeFinal.Resize(fyne.NewSize(120, 130))
	mensajeFinal.Move(fyne.NewPos(330, 150))

	// Creamos un botón para finalizar el juego
	botonFinal := widget.NewButton("Finalizar", func() {
		os.Exit(0)
	})

	// Redimensionamos y posicionamos el botón
	botonFinal.Resize(fyne.NewSize(150, 30))
	botonFinal.Move(fyne.NewPos(320, 250))

	// Escondemos el botón y el mensaje final por defecto
	botonFinal.Hide()
	mensajeFinal.Hide()

	// Creamos un modelo para el temporizador
	temporizadorModelo := modelo.NuevoTemporizadorModelo()

	// Creamos un modelo para el topo
	topoModelo := modelo.NuevoTopoModelo(e.window, botonFinal, mensajeFinal, temporizadorModelo)

	// Iniciamos el temporizador en un hilo lógico
	go temporizadorModelo.IniciarTemporizador(topoModelo, etiquetaTiempo, botonFinal, mensajeFinal)

	// Creamos contenedores para los topos
	topo1 := topoModelo.CrearContenedor(img1)
	topo2 := topoModelo.CrearContenedor(img2)
	topo3 := topoModelo.CrearContenedor(img3)

	// Creamos un fondo para la escena principal
	fondo := canvas.NewImageFromURI(storage.NewFileURI("./assets/fondo_fer.jpg"))
	fondo.Resize(fyne.NewSize(800, 600))

	// Movemos los topos en hilos lógicos separados
	go topoModelo.Mover(topo1)
	go topoModelo.Mover(topo2)
	go topoModelo.Mover(topo3)

	// Agregamos los componentes a la ventana (fondo, topos, etiqueta, mensaje final, botón final)
	e.window.SetContent(container.NewWithoutLayout(fondo, topo1, topo2, topo3, etiquetaTiempo, mensajeFinal, botonFinal))
}
