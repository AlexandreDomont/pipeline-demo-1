package main

import (

               "fmt"
               "os"
               "github.com/joho/godotenv"
		"testing"
)

func main() {

               // Charger les variables d'environnement à partir du fichier .env

               err := godotenv.Load()

               if err != nil {

                              fmt.Println("Erreur lors du chargement des variables d'environnement")

               }

               // Accéder aux variables d'environnement chargées

               myVar := os.Getenv("MY_VAR")

               fmt.Println(myVar)

}


// Fonction de test unitaire

func TestMain(m *testing.M) {

               os.Setenv("MY_VAR", "valeur_test")

               retCode := m.Run()

               os.Exit(retCode)

}
