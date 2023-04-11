# leetlink

## Usage

```bash
$ leetlink 207
https://leetcode.com/problems/course-schedule/
```

Pipe it to `open` to open it on your browser:

```bash
$ leetlink 207 | xargs open
[... opens on default browser ...]
```

## Installation

```
go install github.com/marianogappa/leetlink@latest
```

## It doesn't work for e.g. 175!

This tool uses Leetcode's stats API endpoint, and it seems 13% of the problems are missing. Looks like many of those are SQL & Bash problems, although there are others missing.

## Disclaimer

It looks like I did this, but actually it was mostly ChatGPT 4 via Bing Chat:

<img width="1257" alt="Screenshot 2023-04-10 at 19 00 53" src="https://user-images.githubusercontent.com/1078546/230880284-7cbdc336-50ca-4c46-be08-f7c906456ea6.png">
