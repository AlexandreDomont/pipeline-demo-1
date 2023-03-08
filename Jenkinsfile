
pipeline {

    agent any


	environment {
        	GO111MODULE = 'auto'
		py2Ana="0"	
		TAG_NAME="1"
    	}	
	tools {
		go 'go-1.20.1'
	}

     stages {
	stage('Checkout') {

            steps {

                // Récupérer le code source depuis notre repo Git

                checkout scm

            }

        }
         

        stage('Build and Test') {

            steps {


//		 sh "go install github.com/joho/godotenv/cmd/godotenv@latest"
		 sh "go get github.com/joho/godotenv"

                // Compiler le programme Go

                sh "go build -o mon_programme"

 

                // Lancer les tests unitaires

                sh "go test"

 

                // Lancer le programme avec les variables d'environnement chargées depuis .env

                sh "./mon_programme"

            }

        }

    }

}
