package utils

import "fmt"

func hi() {
	fmt.Println("Hi")
}

// import (
// 	"os"
// 	"sync"

// 	//Autoload the env
// 	_ "github.com/joho/godotenv/autoload"
// )

// var (

// 	/*Mapa que trae clave-valor*/
// 	evironment = map[string]string{}
// 	/*utilizada para sincronizar el acceso a los recursos compartidos por goroutines en un entorno concurrente.*/
// 	evironmentMutex = sync.RWMutex{}
// )

// func AccessENV(key string) string {
// 	//Obtine el valor del map, si existe regresa el valor

// 	/*Esto se hace para asegurar que no haya escrituras concurrentes en el mapa */
// 	evironmentMutex.RLock()

// 	/* Aquí se obtiene el valor asociado con la clave key del mapa environment y se asigna a la variable val.*/
// 	val := evironment[key]

// 	/*Después de leer el valor del mapa, se libera el bloqueo de lectura (RUnlock) del evironmentMutex.*/
// 	evironmentMutex.RUnlock()

// 	if evironment[key] != "" {
// 		return val
// 	}

// 	//Si el valor no exixte, lo obtine de ENV
// 	val = os.Getenv(key)

// 	if val == "" || len(val) <= 0 {
// 		return ""
// 	}

// 	evironmentMutex.Lock()
// 	evironment[key] = val
// 	evironmentMutex.Unlock()

// 	return val
// }
