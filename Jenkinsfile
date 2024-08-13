pipeline {
    agent any

    environment {
        GO_VERSION = "1.20"  // Specify the Go version you are using
    }

    stages {
        stage('Setup') {
            steps {
                script {
                    // Ensure the correct Go version is installed
                    if (!fileExists("/usr/local/go/bin/go")) {
                        echo 'Installing Go...'
                        sh 'wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz'
                        sh 'sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz'
                    }
                    env.PATH = "/usr/local/go/bin:$PATH"
                }
            }
        }

        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                script {
                    // Ensure go.mod is in place and dependencies are downloaded
                    sh 'go mod tidy'
                    sh 'go build -o myapp' // Replace 'myapp' with your preferred binary name
                }
            }
        }

        stage('Test') {
            steps {
                script {
                    sh 'go test ./... -v' // Run all tests with verbose output
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    echo 'Deploying the application...'
                    // Add deployment steps, such as copying the binary to a server or deploying to a cloud service
                    // Example:
                    // sh 'scp myapp user@server:/path/to/deploy/'
                }
            }
        }
    }

    post {
        always {
            cleanWs() // Clean workspace after build
        }
        success {
            echo 'Build succeeded!'
        }
        failure {
            echo 'Build failed!'
        }
    }
}
