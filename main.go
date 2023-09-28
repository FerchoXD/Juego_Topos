package main

import (
	"myFirstGame/escenas"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	// Creamos una nueva aplicación de Fyne
	myApp := app.New()

	// Creamos una nueva ventana y le damos un título
	myWindow := myApp.NewWindow("Atrapa el topo!!!")

	// Centramos la ventana en la pantalla
	myWindow.CenterOnScreen()

	// Fijamos un tamaño para la ventana
	myWindow.SetFixedSize(true)

	// Cambiamos el tamaño de la ventana
	myWindow.Resize(fyne.NewSize(800, 600))

	// Creamos una nueva escena del juego
	juego := escenas.NuevaEscenaPrincipal(myWindow)

	// Cargamos la escena ejecuta todo
	juego.Cargar()

	// Mostramos la ventana y ejecutamos la aplicación
	myWindow.ShowAndRun()
}
