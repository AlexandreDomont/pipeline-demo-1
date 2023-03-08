pipeline {

    agent any

     stages {
	stage('Checkout') {

            steps {

                // Récupérer le code source depuis notre repo Git

                checkout scm

            }

        }
         

        stage('Build and Test') {

            steps {

                // Compiler le programme Go

                sh "go build -o mon_programme"

 

                // Lancer les tests unitaires

                sh "go test -v"

 

                // Lancer le programme avec les variables d'environnement chargées depuis .env

                sh "./mon_programme"

            }

        }

    }

}
