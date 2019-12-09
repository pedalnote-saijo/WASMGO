# wasm/go

## うまくいっていない点

`go build -o test.wasm`

test.wasmのMIMEタイプが `application/octet-stream; charset=binary`
（`file --mime test.wasm`）

GCSに保存した際は`application/wasm`になるので、それで期待した動作を確認した。CROSは起こったが。。。

macでのMIMEタイプを`application/wasm`にできればlocalhostでも検証できると思う

## 参考

- `https://buildersbox.corp-sansan.com/entry/2019/02/14/113000`
- `https://qiita.com/_x8/items/eacd113ee25bc46b3bd0`
- `https://sgswtky.github.io/post/golang-wasm/`
