pipeline {
  agent any
  triggers {
      pollSCM 'H/10 * * * *'
  }
  options {
      buildDiscarder(logRotator(numToKeepStr: '3'))
  }
  stages {
    stage ('Build') {
      steps {
        script {
          def root = tool name: 'go-1.13.8', type: 'go'
          withEnv(["GOPATH=${env.WORKSPACE}/go", "GOROOT=${root}", "GOBIN=${root}/bin", "PATH+GO=${root}/bin"]) {
            sh 'go env'
            sh 'make build'
          }
        }
      }
    }
  }
}
