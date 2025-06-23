pipeline {
  agent any
  stages {
    stage('Initialize'){
      steps {
        script {
          def dockerHome = tool 'docker'
          env.PATH = "${dockerHome}/bin:${env.PATH}"
        }
      }
    }
    stage('Clone Repository') {
      steps {
        script {
          checkout scm
        }
      }
    }

    stage('Build and push') {
      steps {
        script {
          // use your own docker registry address, or any docker registry platform
          docker.withRegistry('http://10.4.125.116:5000/go-cms') {
            def app = docker.build("go-cms")
            app.push("${env.BUILD_NUMBER}")
            app.push("latest")
          }
        }
      }
    }

    stage('Cleanup') {
      steps {
        script {
          sh 'docker image prune -f'
          sh 'curl -L https://github.com/regclient/regclient/releases/latest/download/regctl-linux-amd64 >regctl'
          sh 'chmod 755 regctl'
          sh './regctl registry set --tls disabled 10.4.125.116:5000'
          sh "./regctl tag rm 10.4.125.116:5000/sekolahku-smanssa-api:\$(echo ${env.BUILD_NUMBER} | awk '{print \$1 - 3}')"
        }
      }
    }
    
    stage('Deploy') {
      steps {
        script {
          withKubeCredentials([
              [credentialsId: '886bc7db-2d66-49a7-9c67-292a787b3c26']
          ]) {
              sh 'curl -LO https://dl.k8s.io/release/v1.25.0/bin/linux/amd64/kubectl'
              sh 'chmod u+x ./kubectl'
              sh "./kubectl set image deployment/go-cms go-cms=10.4.125.116:5000/go-cms:${env.BUILD_NUMBER} -n default"
          }
        }
      }
    }
  }
}