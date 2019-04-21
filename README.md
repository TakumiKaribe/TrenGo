# trengo

[![CircleCI](https://circleci.com/gh/TakumiKaribe/trengo.svg?style=svg)](https://circleci.com/gh/TakumiKaribe/trengo)

trengo gets repositories that are currently in trending in the world.

It can be acquired according to language and period.

---

### Usage

| Option | Description | Default Value |
----|----|----
| -l [language_name] | search for [language_name] | `all language` |
| -w | search weekly ※1 | `false` |
| -m | search monthly ※1 | `false` |
| -n | num of result | `10` |

※1 Search period arguments is exclusive. (none is daily)
