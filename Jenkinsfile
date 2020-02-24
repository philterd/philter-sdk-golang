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
              sh "make build"
          }
        }
    }
}
