pipeline {
    agent any
    environment {
        EMAIL_RECIPIENT = credentials('email-address')
    }
    stages {
        stage("verify tooling"){
            steps{
                sh '''
                    docker version
                    docker info
                    docker compose version
                    curl --version
                    jq --version           
                '''
            }
        }
        stage("deploy and deploy services within docker compose"){
            steps{
                sh 'docker compose up -d --no-color --wait'
                sh 'docker compose ps'
            }
        }
        stage("test backend"){
            steps{
                sh 'curl 127.0.0.1:5000/api/animes?propogate=10&page=1'
            }
        }
    }
    post{
        always{
            sh 'docker compose down --remove-orphans -v'
            sh 'docker compose ps'
        }
         post {
        success {
            script {
                emailext subject: 'Build Successful',
                          body: 'The build was successful. Congratulations!',
                          to: EMAIL_RECIPIENT
            }
        }

        failure {
            script {
                def buildError = currentBuild.rawBuild.logFile.text.readLines().join('\n')
                emailext subject: 'Build Failed',
                          body: "The build has failed. Error details:\n\n${buildError}",
                          to: EMAIL_RECIPIENT,
                          attachLog: true
            }
        }
    }
    }
}