pipeline {
    agent any
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
                sh 'curl http://localhost:5000/api/users | jq'
            }
        }
    }
    post{
        always{
            sh 'docker compose down --remove-orphans -v'
            sh 'docker compose ps'
        }
    }
}