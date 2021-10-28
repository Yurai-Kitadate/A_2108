#!/bin/bash

# もし，commit前のテストにしたい時は次を実行してください。
# ln -s `PROJECT_ROOT`/validate-test.sh `PROJECT_ROOT`/.git/hooks/pre-commit

export PUBLIC_KEY="-----BEGIN PUBLIC KEY-----\nMIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQAU7rDEeCHc4TsFe89TvgILlT81rhn\nyrQTPDvCRZb2u8GRYKQIZmnfteVW3use3UV/4tgFWfwKoLzuEjhDFFRrx88Bu46I\n50zegTDQg0FKkWgqlSrYjNku1/upXOfI2QbzGaNcw2tNU1xDFhZ2lGOcqFAKAGvu\nSQvk4OtQo3SNiT8zJnY=\n-----END PUBLIC KEY-----"
export PRIVATE_KEY="-----BEGIN EC PRIVATE KEY-----\nMIHcAgEBBEIAXXCzMM/ExMnMzrN4G75PMBRVZ+NMErY+eBMxR90u5kQTH67u/4qz\nQYQXQ7LwPvwkxJq+jFalVS6ErLQQdRZqepygBwYFK4EEACOhgYkDgYYABABTusMR\n4IdzhOwV7z1O+AguVPzWuGfKtBM8O8JFlva7wZFgpAhmad+15Vbe6x7dRX/i2AVZ\n/AqgvO4SOEMUVGvHzwG7jojnTN6BMNCDQUqRaCqVKtiM2S7X+6lc58jZBvMZo1zD\na01TXEMWFnaUY5yoUAoAa+5JC+Tg61CjdI2JPzMmdg==\n-----END EC PRIVATE KEY-----"
export ISSUER="https://dawn.shinbunbun.info/"

which golangci-lint
if [ $? != 0 ]; then
    echo Please install golangci-lint
    echo https://golangci-lint.run/usage/install/
    exit 1
fi

go test ./...

if [ $? != 0 ]; then
    echo failed to pass the go test
    exit 1
fi

golangci-lint run

if [ $? != 0 ]; then
    echo failed to pass the linter
    exit 1
fi

exit 0
