# alertbench

`alertbench` is a simple tool to provision and delete alert rules in Grafana
for the purpose of benchmarking and load testing.

## Installation

You can install `alertbench` with `go install github.com/grobinson-grafana/alertbench`,
or build it from source.

## Usage

`alertbench` uses `example.json` as a template for the provisioned rules. You should edit
this file before running `alertbench` and change both the `datasourceUid` and `folderUID`.
You can also use the `-file` flag to use alert rule definitions other than `example.json`.

```
Usage of alertbench:
  -bearer-token string
      The bearer token ($BEARER_TOKEN)
  -delete
      Delete all provisioned rules
  -file string
      The file with the example rule (default "example.json")
  -offset int
      The offset that rules should be provisioned from
  -password string
      The password ($PASSWORD) (default "password")
  -rules int
      The number of rules that should be provisioned (default 100)
  -rules-per-group int
      The maximum number of rules per rule group (default 10)
  -url string
      URL of the Grafana server (default "http://127.0.0.1:3000")
  -username string
      The username ($USERNAME) (default "admin")
```

### How to provision alert rules

To provision alert rules run `alertbench`. You can use the `-rules`
and `-rules-per-group` flags to set the number of rules that are
provisioned and how these rules are divided into rule groups.

```
alertbench -rules=50
2023/03/22 15:24:02 Provisioning 50 rules, at most 10 rules per group
2023/03/22 15:24:02 Provisioned a154ba37-166e-418b-bb7a-327bfc8e79c0
2023/03/22 15:24:02 Provisioned c698f4bf-40c1-4bc4-8309-00ef72164e08
2023/03/22 15:24:02 Provisioned d81dee76-6be9-4791-887b-cfede5f72efa
2023/03/22 15:24:02 Provisioned d35b940c-072a-4298-b9e7-3c7c8339a269
...
2023/03/22 15:24:06 Provisioned e4665b2f-511b-464b-a554-b633deafc9a4
2023/03/22 15:24:06 Provisioned c8437395-b6e3-4bdf-93f4-7df9f66837d4
2023/03/22 15:24:06 Provisioned d45a5017-a560-46bd-92b6-816fa989af5f
2023/03/22 15:24:06 Provisioned bf2f279d-d349-445c-b853-16be6815db79
2023/03/22 15:24:06 Done 
```

You can also provision further rules using the `-offset` flag to offset
the rules and avoid naming conflicts.

```
alertbench -rules=50 -offset=50
```

### How to delete alert rules

You use the `-delete` flag to delete provisioned rules. This will delete
all provisioned alert rules. The `-rules`, `rules-per-group` and `-offset`
flags have no effect.

```
alertbench -delete
2023/03/22 15:24:26 Deleting provisioned rules
2023/03/22 15:24:26 Deleted a154ba37-166e-418b-bb7a-327bfc8e79c0
2023/03/22 15:24:26 Deleted c698f4bf-40c1-4bc4-8309-00ef72164e08
2023/03/22 15:24:26 Deleted d81dee76-6be9-4791-887b-cfede5f72efa
2023/03/22 15:24:26 Deleted d35b940c-072a-4298-b9e7-3c7c8339a269
...
2023/03/22 15:24:30 Deleted e4665b2f-511b-464b-a554-b633deafc9a4
2023/03/22 15:24:30 Deleted c8437395-b6e3-4bdf-93f4-7df9f66837d4
2023/03/22 15:24:30 Deleted d45a5017-a560-46bd-92b6-816fa989af5f
2023/03/22 15:24:30 Deleted bf2f279d-d349-445c-b853-16be6815db79
2023/03/22 15:24:30 Done
```