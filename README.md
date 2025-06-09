# log-analyzer

**LatencyLens** は、自作の Go アプリケーションの HTTP/gRPC リクエストのレイテンシを計測・可視化する開発支援ツールです。

リアルタイムに P50 / P95 / P99 のパフォーマンス指標を取得し、Web ダッシュボードで直感的に確認できます。

---

## 特徴

- Goアプリに簡単に埋め込めるレイテンシ計測ミドルウェア  
- P50 / P95 / P99 レイテンシの自動集計と可視化 
- JSON / CSV形式で集計結果を出力可能  
- Web ダッシュボード付き  
- 自作の API / サービスの負荷テスト・性能監視に活用可能

---

## 構成
latency-lens/
├── main.go # メトリクス提供サーバー
├── middleware/ # HTTPリクエストのレイテンシ収集ミドルウェア
├── collector/ # レイテンシ集計ロジック
├── stats/ # P50/P95/P99 計算処理
└── ui/
└── index.html # Webダッシュボード（HTML/JS）

---

## 使い方
### 1. アプリケーションへの組込み
```go
wrapped := middleware.HTTPMiddleware(http.DefaultServeMux)
http.ListenAndServe(":3000", wrapped)
```

### 2. レイテンシ収集サーバー起動
```bash
go run main.go
```
→ :8080 で Web UI（/metrics）および HTML が提供されます

### 3. Web UI 表示
index.htmlをブラウザで開く


## ダッシュボード画面
- /metrics に対する定期リクエストで、以下の指標がリアルタイム表示されます：
  | Label    | P50   | P95   | P99   | Count |
| -------- | ----- | ----- | ----- | ----- |
| `/hello` | 50ms  | 150ms | 200ms | 12    |
| `/slow`  | 350ms | 500ms | 600ms | 6     |


### 今後の予定
- gRPC対応
- 計測結果の保存（CSV, JSON）
-  任意リクエストの UI 送信機能の強化
- 外部サービス統合(例: Prometheus) 

### ライセンス
MIT

