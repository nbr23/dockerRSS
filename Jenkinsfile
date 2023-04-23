pipeline {
    agent any

    options {
        disableConcurrentBuilds()
        ansiColor('xterm')
    }

    stages {
        stage('Checkout'){
            steps {
                checkout scm
            }
        }
        stage('Tests'){
            steps {
                script {
                    env.REAL_PWD = getDockerPWD();
                    sh 'docker run -t -w /app -v $REAL_PWD:/app --rm golang:alpine go test ./dockerhub'
                }
            }
        }
        stage('Prep buildx') {
            steps {
                script {
                    env.BUILDX_BUILDER = getBuildxBuilder();
                }
            }
        }
        stage('Dockerhub login') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub', usernameVariable: 'DOCKERHUB_CREDENTIALS_USR', passwordVariable: 'DOCKERHUB_CREDENTIALS_PSW')]) {
                    sh 'docker login -u $DOCKERHUB_CREDENTIALS_USR -p "$DOCKERHUB_CREDENTIALS_PSW"'
                }
            }
        }
        stage('Build DockerRSS Image') {
            steps {
                sh """
                    docker buildx build --pull --builder \$BUILDX_BUILDER  --platform linux/amd64,linux/arm64 -t nbr23/dockerrss:`git rev-parse --short HEAD` -t nbr23/dockerrss:latest -f docker/Dockerfile --push .
                    """
            }
        }
        stage('Build DockerRSS-nginx Image') {
            steps {
                sh """
                    docker buildx build --pull --builder \$BUILDX_BUILDER  --platform linux/amd64,linux/arm64 -t nbr23/dockerrss:nginx-`git rev-parse --short HEAD` -t nbr23/dockerrss:nginx-latest -f docker/Dockerfile.nginx --push .
                    """
            }
        }
        stage('Sync github repo') {
            when { branch 'master' }
            steps {
                syncRemoteBranch('git@github.com:nbr23/dockerRSS.git', 'master')
            }
        }
    }
    post {
        always {
            sh 'docker buildx stop $BUILDX_BUILDER || true'
            sh 'docker buildx rm $BUILDX_BUILDER || true'
        }
    }
}
