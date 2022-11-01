#!/bin/sh

go install go.k6.io/xk6/cmd/xk6@latest
xk6 build --with github.com/szkiba/xk6-dashboard@latest

if [ -f "tests/load-test-script-$1.js" ]; then 
    script=tests/load-test-script-$1.js
else
    script=tests/load-test-script.js
fi

./k6 run $script --out dashboard             