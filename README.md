# study-release-drafter
This repository is intended for personal learning purposes.

## 目的
release-drafterを使って、GitHubを操作しながら、リリース作業の自動化を体験します。
https://github.com/release-drafter/release-drafter

# 学習項目

## Release Drafterとは

Release Drafterは、GitHub Actionsを使用してリリースノートの自動生成を行うツールです。
PRのラベルやタイトルから自動的にリリースノートを作成し、セマンティックバージョニングもサポートしています。

## プロジェクト構成

```
.
├── .github.sample/                # 参考用のサンプル設定
│   ├── workflows/
│   │   └── release-drafter.yml    # GitHub Actionsワークフロー例
│   └── release-drafter.yml        # Release Drafter設定例
├── main.go                        # シンプルなGo CLIアプリケーション
├── go.mod                         # Goモジュール定義
└── README.md                      # このファイル
```

> **注意**: `.github.sample`ディレクトリには参考用の設定ファイルが含まれています。
> 学習のため、実際の`.github`ディレクトリは自分で作成してください。

## 自分で設定する項目

### 1. GitHub Actions ワークフロー
`.github/workflows/release-drafter.yml`を作成し、以下を設定：
- mainブランチへのpush時にリリースドラフトを更新
- PR作成/更新時にもドラフトを更新

> サンプルは `.github.sample/workflows/release-drafter.yml` を参照

### 2. Release Drafter設定
`.github/release-drafter.yml`を作成し、以下を定義：

**カテゴリ分類:**
- 🚀 Features (feature, enhancement)
- 🐛 Bug Fixes (fix, bugfix, bug)
- 🧰 Maintenance (chore, maintenance)
- 📚 Documentation (documentation, docs)
- ⬆️ Dependencies (dependencies, deps)

**バージョン管理:**
- Major: `major`ラベル
- Minor: `feature`, `enhancement`ラベル
- Patch: その他のラベル

> サンプルは `.github.sample/release-drafter.yml` を参照

### 3. Go CLIアプリケーション（作成済み）

シンプルなCLIツール（`simple-cli`）:
```bash
# ビルド
go build -o simple-cli main.go

# 実行例
./simple-cli                     # Hello, World!
./simple-cli -name Alice        # Hello, Alice!
./simple-cli -version           # simple-cli version 0.1.0
./simple-cli -help              # ヘルプを表示
```

## 学習手順

### ステップ1: Release Drafterの設定
1. `.github/workflows/`ディレクトリを作成
2. `release-drafter.yml`ワークフローファイルを作成
3. `.github/release-drafter.yml`設定ファイルを作成
4. このリポジトリをGitHubにpush
5. GitHub Actionsが有効になっていることを確認

### ステップ2: PRの作成と動作確認

#### 機能追加のPR例
1. 新しいブランチを作成
   ```bash
   git checkout -b feature/add-greeting-emoji
   ```
2. コードを変更（例：絵文字を追加）
3. PRを作成し、`feature`ラベルを付与
4. マージ後、リリースドラフトが更新されることを確認

#### バグ修正のPR例
1. 新しいブランチを作成
   ```bash
   git checkout -b fix/error-handling
   ```
2. エラーハンドリングを追加
3. PRを作成し、`bug`ラベルを付与
4. マージ後、リリースドラフトが更新されることを確認

### ステップ3: リリースの作成
1. GitHubの「Releases」ページへ移動
2. 自動生成されたドラフトを確認
3. 必要に応じて編集し、リリースを公開

## 注意点

- PRにラベルを付けることで、リリースノートのカテゴリ分けが自動化されます
- セマンティックバージョニングに従ったバージョン番号が自動計算されます
- リリースドラフトはmainブランチへのマージごとに更新されます

## 参考リンク

- [Release Drafter公式](https://github.com/release-drafter/release-drafter)
- [GitHub Actions ドキュメント](https://docs.github.com/ja/actions)
- [セマンティックバージョニング](https://semver.org/lang/ja/)
