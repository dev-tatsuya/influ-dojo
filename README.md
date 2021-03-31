# influ-dojo(Close)

Twitterで競い合いながらインフルエンサーを目指すことができるランキングサイトです。

運営Twitterアカウント: https://twitter.com/infludojo \
ランキングサイトURL: https://influ-dojo.work

## 概要
モダンな技術を習得することを目的とした、本番稼働中の個人開発のサービスになります。
現在、400名以上のユーザ様にご利用いただいております。

本サービスは、SNSのフォロワー情報を取得、保存し、ランキング付けをする機能が主な機能となります。具体的には外部API(Twitter API)を利用してSNSのフォロワー情報を取得し、それをRDS(MySQL)に格納し、日/週/月単位のバッチ処理でランキング付けをするロジックを回し、その結果をElastiCache(Redis)に格納しています。ランキング結果をKVSに格納することで、リクエスト毎のRDBへのアクセスを避けるようにしました。
またインフラに関して、様々なAWSサービスを活用(※GitHub参照)していますが、主にはECS on Fargateを導入し、本番環境でのコンテナ基盤の構築、運用をよしなに行ってくれるインフラを整備することで、アプリケーションの開発に集中することができました。
またCI/CDに関してはGitHub Actionsを導入し、GitHubにpushされたら自動テスト、マスターブランチにマージされたらECR経由でECSへ自動デプロイされる機構を整備しました。
最後に、インフラのプロビジョニングツールとしてTerraformを導入し、AWS環境のバージョニングに努めました。コード一行で一括で構築、削除が可能となり、個人開発においては金額面の考慮もあるため非常に役立ちました。

## 技術スタック
- フロントエンド: Flutter Web
  - β版のFlutter Webを選定した理由は以下が挙げられます。
    - 個人開発でFlutter製アプリを作った経験があったこと
    - いずれモバイル版アプリの制作を見越した判断（Web, iOS, Androidの一元コード化）
    - コンテナ上でどのように動かすことができるのかという興味
- バックエンド: Go言語
- インフラ: AWS（ECS on Fargate, ECR等）
- CI/CD: GitHub Actions
  - 選定理由は以下が挙げられます。
    - GitHubとの親和性
    - CI/CD以外にもWebhook利用も可能なため、CircleCI等を包含すると考えたため
    - publicリポジトリでの利用は無料
- IaC: Terraform 

## インフラ構成図
### production（stagingから移行中）
<img width="600" alt="スクリーンショット 2020-09-08 22 29 07" src="https://user-images.githubusercontent.com/44150538/92482626-c8ce6200-f222-11ea-9ce3-df0fe3e51107.png">

### staging
<img width="932" alt="future-archi" src="https://user-images.githubusercontent.com/44150538/93420109-42341780-f8e9-11ea-957e-3e3b4d7e8d5f.png">
