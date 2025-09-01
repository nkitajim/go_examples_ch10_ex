```
% go mod init github.com/nkitajim/go_examples_ch10_ex
% mkdir -p cmd/cli math
% vim cmd/cli/main.go
% vim math/math.go
% go build -o cli ./cmd/cli
% ./cli
1 + 2 = 3
```

# version up
```
% go doc github.com/nkitajim/go_examples_ch10_ex/math
package math // import "github.com/nkitajim/go_examples_ch10_ex/math"

func Add(a int, b int) int

% git add math/math.go 
% git commit -m'update'
[main 7c033dd] update
 2 files changed, 3 insertions(+), 1 deletion(-)
% git push origin main
% git tag v1.0.1
% git push origin v1.0.1
```

# pkgsite
```
$ go install golang.org/x/pkgsite/cmd/pkgsite@latest
```
