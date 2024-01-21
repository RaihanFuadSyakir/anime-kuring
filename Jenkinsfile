pipeline {
    agent any

    stages {
        stage("verify tooling") {
            steps {
                node {
                    environment {
                        EMAIL_RECIPIENT = credentials('email-address')
                    }
                    sh '''
                        docker version
                        docker info
                        docker compose version
                        curl --version
                        jq --version           
                    '''
                }
            }
        }

        stage("deploy and deploy services within docker compose") {
            steps {
                node {
                    sh 'docker compose up -d --no-color --wait'
                    sh 'docker compose ps'
                }
            }
        }

        stage("test backend") {
            steps {
                node {
                    sh 'curl http://127.0.0.1:5000/api/animes/1 | jq'
                }
            }
        }
    }

    post {
        always {
            node {
                sh 'docker compose down --remove-orphans -v'
                sh 'docker compose ps'
            }
        }

        success {
            node {
                script {
                    emailext subject: 'Build Successful',
                              body: 'The build was successful. Congratulations!',
                              to: EMAIL_RECIPIENT
                }
            }
        }

        failure {
            node {
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
