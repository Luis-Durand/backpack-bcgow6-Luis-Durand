package main

/* Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura de usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones.
La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:
Cambiar nombre: me permite cambiar el nombre y apellido.
Cambiar edad: me permite cambiar la edad.
Cambiar correo: me permite cambiar el correo.
Cambiar contraseña: me permite cambiar la contraseña.
*/

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func (u *Usuario) cambiarNombre(nombre, apellido string) {

	u.Nombre = nombre
	u.Apellido = apellido

}

func (u *Usuario) cambiarEdad(edad int) {

	u.Edad = edad

}

func (u *Usuario) cambiarCorreo(correo string) {
	u.Correo = correo
}

func (u *Usuario) cambiarContraseña(contra string) {

	u.Contraseña = contra
}

func main() {
	/*
		 	juan := Usuario{"pepe", "martin", 23, "dadssa", "sdasdasd"}

			juan.cambiarNombre("sancho", "panza")
			juan.cambiarEdad(10)
			juan.cambiarCorreo("juanchi@gmail.com")
			juan.cambiarContraseña("Alalal")
			fmt.Println(juan)
	*/
}
