```
% go mod init github.com/nkitajim/go_examples_ch10_ex/client
go: creating new go.mod: module github.com/nkitajim/go_examples_ch10_ex/client
go: to add module requirements and sums:
	go mod tidy
% go mod tidy
go: finding module for package github.com/nkitajim/go_examples_ch10_ex/math
go: downloading github.com/nkitajim/go_examples_ch10_ex v0.0.0-20250901055240-5d0c5211725a
go: found github.com/nkitajim/go_examples_ch10_ex/math in github.com/nkitajim/go_examples_ch10_ex v0.0.0-20250901055240-5d0c5211725a
```

```
% git tag v1.0.0
% git push origin v1.0.0
% go get github.com/nkitajim/go_examples_ch10_ex
% go get github.com/nkitajim/go_examples_ch10_ex@v1.0.0
go: downloading github.com/nkitajim/go_examples_ch10_ex v1.0.0
go: upgraded github.com/nkitajim/go_examples_ch10_ex v0.0.0-20250901055240-5d0c5211725a => v1.0.0
```
