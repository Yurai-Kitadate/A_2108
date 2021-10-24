# dbテスト用DockerCompose
## 要約
- ./docker_composeでdocker-compose upをするとSQL serverとadminerが立ち上がります.
- 使い終わった後はdocker-compose downでコンテナとネットワークを消しましょう.
- 永続化はしてないので, データはdownすると消えます. 永続化したい人は[永続化](#Persistence)を読んでください.

## db設定, 初期データ投入について
### db設定
./docker_compose/sql_confに設定ファイルを投入するとMySQLdが読みます. 投入したファイルはreadonly属性にしてください. そうでない場合, MySQLdが読んでくれません.

### 初期投入データ
./docker_compose/sq1_init以下のSQLファイルが, 先頭の連番の順に実行されます. 
TODO: きちんと構造を定義する.

## db操作
### GUI
GUIでやりたい場合は, http://localhost:8080/ にAdminerが立ち上がっているので, ブラウザからアクセスしてください. 
- User: root
- Password: De3thM3rch

### CLI
CLIのmysqlでアクセスしたい場合, 3306をホストにフォワードしているので, localhostの3306にTCPでアクセスしてください.
- アクセス方法の例
``` sh
mysql -u root -p -h localhost -P 3306 --protocol=tcp
```

<div id="Persistence"></div>

## 永続化
- 基本はshell script叩いてください.
    - ./docker_compose/persistence.shに実行権限を与えて実行すると以下の処理を自動で行います. 

- 自分でやる方法
1. ./docker_compose以下にdbフォルダを作ってください. 
1. **dbフォルダとdocker-compose.ymlを.gitignoreに追加してください.** 理由は後述します.
1. docker-compose.ymlのdb::volumesに`- ./db:/var/lib/mysql`を追加してください.
1. compose upすると永続化されています.

dbフォルダは所有権とPermissionがDocker内部のものを引き継いでいるので, git addしてしまうと地獄になります.
docker-compose.ymlは永続化しないバージョンをGithubで管理しているので, 永続化後のyamlをpushされると困ります.