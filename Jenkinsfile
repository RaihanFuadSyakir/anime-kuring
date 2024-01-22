pipeline {
    agent any
    environment{
        EMAIL_RECIPIENT = credentials('email-manager')
        DOCKER_ACCOUNT = credentials('docker-account')
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
                sh 'docker compose up -d --build --no-color --wait'
                sh 'docker compose ps'
            }
        }
        stage("test backend"){
            steps{
                sh 'curl http://127.0.0.1:5000/api/animes/1 | jq'
            }
        }
    }
    post{
        success {
            def version = 'v2'
            script {
                sh 'docker login --username $DOCKER_ACCOUNT_USR --password $DOCKER_ACCOUNT_PWD'
                sh 'docker images'
                sh 'docker compose ps'
                sh 'docker compose down --remove-orphans -v'
                sh 'docker compose ps'
                emailext subject: 'Build Successful',
                         from : '$EMAIL_RECIPIENT_USR'
                         body: 'The build was successful. stored at repository $version',
                         to: 'notsaya1@gmail.com'
            }
        }

        failure {
            script {
                sh 'docker logs docker-composetest-frontend-1'
                sh 'docker logs docker-composetest-backend-1'
                sh 'docker logs docker-composetest-mongodb-1'
                def buildError = currentBuild.rawBuild.logFile.text.readLines().join('\n')
                emailext subject: 'Build Failed',
                        mimeType: 'text/html',
                        from : '$EMAIL_RECIPIENT_USR',
                        body: """
                        The build has failed. Error details:
                        """,
                        to: 'notsaya1@gmail.com',
                        attachLog: true
            }
                sh 'docker compose down --remove-orphans -v'
                sh 'docker compose ps'
        }
    }
}