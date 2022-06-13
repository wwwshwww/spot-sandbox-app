# Vite + React + TypeScript + ESLint + Prettier テンプレート

## 開発サーバ立ち上げ

```sh
$ npm run dev
```

## ビルド

```sh
$ npm run build
```

同時に型検査も行われます。

## format/lint

```sh
$ npm run lint:fix
```

📝: husky, lint-staged の設定も入れてあるので、git commit 時にコードフォーマットは掛かります (lint auto-fix は今は入れていません)

## テンプレートの使い方

1. README が置かれているディレクトリをまるごとコピーします
  * この時、もし node_modules が存在する場合、コピーにめちゃくちゃ時間を取られるため、コピー対象に含まれないように注意してください(先に消しておいた方が良いです)
2. `npm ci`
3. You're ready to Go.

