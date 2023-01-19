# WBA-BC-Project-02 Wemix-on
## 위믹스 플레이의 Dex 구조를 구현한 golang과 smart contract를 이용해 구현한 프로젝트

### 시작하기
- 파일 clone
```shell
git clone https://github.com/codestates/WBA-BC-Project-02.git
```

- go 의존성 패키지 설치

```shell
go mod init
go mod tidy
```  
  
- docker를 통한 redis와 mongoDB 실행
```shell
docker-compose up -d
```
  
- gin server 실행
```shell
go run "/Users/<프로젝트 루트 폴더까지의 경로>/was/main.go"
```

- signer Daemon 실행
```shell
go run "/Users/<프로젝트 루트 폴더까지의 경로>/signerdaemon/main.go"
```

- DB Daemon 실행
```shell
go run "/Users/<프로젝트 루트 폴더까지의 경로>/daemon/main.go"
```

### 지원 기능 및 url
- 회원가입
  POST - http://localhost:8080/app/v1/users/wallet
