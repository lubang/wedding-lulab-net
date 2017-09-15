# Server 설정

## sqlite 설정

wedding-lulab-net.db

### Table Definition

예약 관련 정보를 저장하기 위한 DB 테이블 구조입니다.

  CREATE TABLE `reservations` (
    `uid`	INTEGER PRIMARY KEY AUTOINCREMENT,
    `name`	VARCHAR(256) NOT NULL,
    `mobile`	VARCHAR(16) NOT NULL UNIQUE,
    `created`	INTEGER NOT NULL
  );


## Build 설정

cgo 기반으로 sqlite가 빌드되기 때문에 gcc이 설치되어 있어야 합니다. 다음의 명령어를 통해 배포 버전을 생성 가능합니다.

  $ go build --tags "libsqlite3 darwin"

위 명령을 위해서 sqlite3가 설치되어 있어야 합니다. `$ brew install sqlite3` 명령을 이용하여 설치합니다.
