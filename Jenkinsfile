pipeline {
    agent any
    environment{
        EMAIL_RECIPIENT = credentials('email-manager')
        DOCKER_ACCOUNT = credentials('docker-account')
        REPO_VERSION = "v2"
        CONTAINER_NAME = "docker-composetest"
        PIC_EMAIL = "notsaya1@gmail.com"
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
            script {
                sh('docker login --username $DOCKER_ACCOUNT_USR --password $DOCKER_ACCOUNT_PSW')
                sh 'docker images'
                sh 'docker compose ps'
                sh 'docker commit $CONTAINER_NAME-backend-1 anime-kr-backend:$REPO_VERSION'
                sh 'docker commit $CONTAINER_NAME-frontend-1 anime-kr-frontend:$REPO_VERSION'
                sh 'docker commit $CONTAINER_NAME-mongodb-1 anime-kr-db:$REPO_VERSION'
                sh 'docker tag anime-kr-backend:$REPO_VERSION docker.io/caltfasy/anime-kr-backend:$REPO_VERSION'
                sh 'docker tag anime-kr-frontend:$REPO_VERSION docker.io/caltfasy/anime-kr-frontend:$REPO_VERSION'
                sh 'docker tag anime-kr-db:$REPO_VERSION docker.io/caltfasy/anime-kr-db:$REPO_VERSION'
                sh 'docker push docker.io/caltfasy/anime-kr-backend:$REPO_VERSION'
                sh 'docker push docker.io/caltfasy/anime-kr-frontend:$REPO_VERSION'
                sh 'docker push docker.io/caltfasy/anime-kr-db:$REPO_VERSION'
                sh 'docker logout'
                sh 'docker compose down --remove-orphans -v'
                sh 'docker compose ps'
                emailext subject: 'Build Successful',
                         from : '$EMAIL_RECIPIENT_USR',
                         body: """
                         The build was successful. stored at repository ${REPO_VERSION}.
                         you can check at
                         https://hub.docker.com/repository/docker/caltfasy/anime-kr-db/general
                         https://hub.docker.com/repository/docker/caltfasy/anime-kr-frontend/general
                         https://hub.docker.com/repository/docker/caltfasy/anime-kr-backend/general
                         """,
                         to: '$PIC_EMAIL'
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
                        to: '$PIC_EMAIL',
                        attachLog: true
            }
                sh 'docker compose down --remove-orphans -v'
                sh 'docker compose ps'
        }
    }
}