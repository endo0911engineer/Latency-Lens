# log-analyzer

高速で並列処理が可能なGo製ログ解析ライブラリ＆CLIツール

---

## 概要

`log-analyzer` は、大量のログファイルを高速かつ並列で解析し、ステータスコードやモデル別リクエスト数、時間帯ごとの統計などを集計するツール兼ライブラリです。  
CLIとしても利用可能で、将来的にはリアルタイムログ監視機能も備える予定です。

---

## 特徴

- Goの goroutine を活用した高速並列解析  
- 柔軟なログフォーマット対応（JSON, Apache/Nginx 形式など）  
- JSON / CSV形式で集計結果を出力可能  
- ライブラリとしても利用でき、他サービスへの組み込みが容易  
- （今後）fsnotifyによるリアルタイムログ監視に対応予定

---

## インストール

```bash
go get github.com/endo0911engineer/log-analyzer
```
---

## 使い方
### CLIツール
ターミナルで以下のように実行してください：
```bash
log-analyzer -file access.log
```

複数ファイルも解析可能です：
```bash
log-analyzer -file access1.log -file access2.log
```
### ライブラリ使用例
```bash
package main

import (
    "log"

    "github.com/endo0911engineer/log-analyzer/pkg/parser"
    "github.com/endo0911engineer/log-analyzer/pkg/aggregator"
    "github.com/endo0911engineer/log-analyzer/pkg/output"
)

func main() {
    files := []string{"access1.log", "access2.log"}

    var allEntries [][]parser.LogEntry
    for _, f := range files {
        entries, err := parser.ParseLogFile(f)
        if err != nil {
            log.Fatal(err)
        }
        allEntries = append(allEntries, entries)
    }

    stats := aggregator.AggregateLogEntries(allEntries)

    err := output.PrintJSON(stats)
    if err != nil {
        log.Fatal(err)
    }
}
```

### APIドキュメント

### 今後の予定
- リアルタイムログ監視（fsnotify + goroutine + チャネル）

- WebSocket経由でのリアルタイムダッシュボード表示

- より多様なログフォーマット対応

- CLIオプションの強化と設定ファイル対応



