pipeline {
    agent any

    environment {
        GO_VERSION = "1.21"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Setup Go') {
            steps {
                sh '''
                sudo apt-get update
                sudo apt-get install -y wget
                wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz
                sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
                export PATH=$PATH:/usr/local/go/bin
                go version
                '''
            }
        }

        stage('Dependency Management') {
            steps {
                sh "go mod tidy"
            }
        }

        stage('Build') {
            steps {
                sh "go build ./..."
            }
        }

        stage('Run Tests') {
            steps {
                sh "go test ./... -coverprofile=coverage.out"
            }
        }

        stage('Static Analysis') {
            steps {
                sh '''
                go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
                golangci-lint run ./...
                go vet ./...
                '''
            }
        }

        stage('Security Scan') {
            steps {
                sh '''
                snyk test || true
                trivy fs ./ || true
                '''
            }
        }

        stage('Archive Reports') {
            steps {
                archiveArtifacts artifacts: 'coverage.out', allowEmptyArchive: true
            }
        }
    }

    post {
        always {
            echo 'Jenkins pipeline finished.'
        }
    }
}
