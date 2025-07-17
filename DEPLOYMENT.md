# FortuneSpinner デプロイメントガイド

## 本番環境へのデプロイ手順

### 1. 環境変数の設定

#### 方法1: クラウドサービスを使用（推奨）

**AWS Parameter Store**
```bash
# 1. パラメータの初期設定
./deploy/setup-aws-parameters.sh production

# 2. データベースパスワードを設定
aws ssm put-parameter \
  --name "/fortunespinner/production/db_password" \
  --value "your-secure-password-here" \
  --type "SecureString" \
  --overwrite

# 3. デプロイ実行
./deploy/deploy-aws.sh production
```

**GCP Secret Manager**
```bash
# シークレットの作成
gcloud secrets create db-password --data-file=- <<< "your-secure-password"
```

**Azure Key Vault**
```bash
# シークレットの作成
az keyvault secret set \
  --vault-name MyVault \
  --name db-password \
  --value "your-secure-password"
```

#### 方法2: ローカル環境変数ファイル（開発環境向け）

```bash
# 1. .env.exampleをコピー
cp .env.example .env

# 2. .envファイルを編集してパスワードを設定
# DB_PASSWORD=your-secure-password-here

# 3. ローカルデプロイ実行
./deploy/deploy-local.sh
```

### 2. SSL証明書の準備（HTTPS使用時）

#### Let's Encryptを使用する場合
```bash
# Certbotのインストール
sudo apt update
sudo apt install certbot

# 証明書の取得
sudo certbot certonly --standalone -d yourdomain.com
```

#### 自己署名証明書（開発環境用）
```bash
mkdir -p ssl
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout ssl/privkey.pem \
  -out ssl/fullchain.pem
```

### 3. デプロイの実行

```bash
# 本番環境用のDocker Composeで起動
docker-compose -f docker-compose.prod.yml up -d

# 初回のみ: データベースマイグレーション
docker exec -i fortunespinner-mysql mysql -uroot -p${DB_PASSWORD} ${DB_NAME} < migrations/001_initial_schema.sql
```

### 4. 動作確認

```bash
# コンテナの状態確認
docker-compose -f docker-compose.prod.yml ps

# ログの確認
docker-compose -f docker-compose.prod.yml logs -f

# ヘルスチェック
curl http://localhost/api/health
```

## 環境別の設定

### 開発環境
```bash
# 通常のdocker-composeを使用
docker-compose up -d
```

### 本番環境
```bash
# 環境変数を読み込んで起動
docker-compose -f docker-compose.prod.yml up -d
```

## トラブルシューティング

### MySQLが起動しない
```bash
# ボリュームをクリアして再起動
docker-compose down -v
docker-compose up -d
```

### SSL証明書エラー
- nginx.prod.confの証明書パスを確認
- 証明書ファイルの権限を確認: `chmod 644 /path/to/cert`

### ポート競合
- 80/443ポートが使用中の場合は、docker-compose.prod.ymlのポート設定を変更

## セキュリティチェックリスト

- [ ] 強力なDB_PASSWORDを設定（最低12文字、英数字記号混在）
- [ ] .envファイルがGitに含まれていないことを確認
- [ ] 秘匿情報はParameter Store/Secret Manager等で管理
- [ ] SSL証明書が正しく設定されている
- [ ] ファイアウォールで必要なポートのみ開放
- [ ] 定期的なバックアップの設定
- [ ] 環境変数へのアクセス権限を最小限に設定

## バックアップとリストア

### バックアップ
```bash
# データベースのバックアップ
docker exec fortunespinner-mysql mysqldump -uroot -p${DB_PASSWORD} ${DB_NAME} > backup_$(date +%Y%m%d).sql
```

### リストア
```bash
# バックアップからリストア
docker exec -i fortunespinner-mysql mysql -uroot -p${DB_PASSWORD} ${DB_NAME} < backup_20240101.sql
```

## 監視とメンテナンス

### ログローテーション
```bash
# Dockerログの確認
docker logs fortunespinner-backend --tail 100 --follow

# ログサイズの確認
docker inspect fortunespinner-backend | grep -i logpath
```

### アップデート手順
```bash
# 最新コードを取得
git pull origin main

# イメージの再ビルドと再起動
docker-compose -f docker-compose.prod.yml up -d --build
```