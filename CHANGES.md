# Changes from upstream

This repository is a frozen fork of etcd used for Sourcegraph QA pipelines. The following directive was added to the `go.mod` file in the following commits. The rest of the repository content is the same as upstream (but frozen).

```
replace (
    go.uber.org/zap => github.com/sourcegraph-testing/zap v1.12.0
)
```

`aef232fbec9089d4468ff06705a3a7f84ee50ea6` -> `fb38de395ba67f49978b218e099de1c45122fb50`
`33623cc32f8d9f999fd69189d29124d4368c20ab` -> `415ffd5a3ba7a92a07cd96c7d9f4b734f61248f7`
`0ad2e75d529bda74472a1dbb5e488ec095b07fe7` -> `f8307e394c512b4263fc0cd67ccf9fd46f1ad9a5`
