pipeline {
    agent any

    environment {
        // SonarQube server name as configured in Jenkins
        SONARQUBE_SERVER = 'SonarQube'
        // SonarQube project key and name
        SONAR_PROJECT_KEY = 'golang-bug-analysis'
        SONAR_PROJECT_NAME = 'GoLang Bug Analysis'
        // Go binary path (if required)
        PATH = "${env.PATH}"
    }

    tools {
        go 'Go-1.21' // Name of Go installation in Jenkins
    }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/SyedRehanAli25/golang-bug-analysis.git'
            }
        }

        stage('Build') {
            steps {
                sh 'go mod tidy'
                sh 'go build ./...'
            }
        }

        stage('Run Unit Tests') {
            steps {
                sh 'go test -v ./... -coverprofile=coverage.out'
            }
        }

        stage('Static Code Analysis (Optional)') {
            steps {
                sh 'golangci-lint run ./... || true'
                // '|| true' prevents pipeline failure if linter finds issues, optional
            }
        }

        stage('SonarQube Analysis') {
            environment {
                SONAR_SCANNER_HOME = tool name: 'SonarScanner', type: 'hudson.plugins.sonar.SonarRunnerInstallation'
            }
            steps {
                withSonarQubeEnv(SONARQUBE_SERVER) {
                    sh """
                    sonar-scanner \
                      -Dsonar.projectKey=${SONAR_PROJECT_KEY} \
                      -Dsonar.projectName="${SONAR_PROJECT_NAME}" \
                      -Dsonar.sources=. \
                      -Dsonar.go.coverage.reportPaths=coverage.out \
                      -Dsonar.host.url=${env.SONAR_HOST_URL} \
                      -Dsonar.login=${env.SONAR_AUTH_TOKEN}
                    """
                }
            }
        }

        stage('Quality Gate') {
            steps {
                timeout(time: 5, unit: 'MINUTES') {
                    waitForQualityGate abortPipeline: true
                }
            }
        }
    }

    post {
        success {
            echo "Pipeline completed successfully. Code passed SonarQube Quality Gate."
        }
        failure {
            echo "Pipeline failed. Check SonarQube dashboard for issues."
        }
    }
}
