[![Go](https://github.com/mchirico/switch-context/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/mchirico/switch-context/actions/workflows/go.yml)
# Switch-context

This project was originally created by [Dara]( https://github.com/DaraDadachanji)

## Installation

Install Go from the [official website](https://go.dev/)

clone this repository and build the executable. Then move it to your bin folder

```bash
git clone https://github.com/mchirico/scontext.git
cd scontext
go mod tidy
go build
mv ./switch-context /usr/local/bin/switch-context
```

add the following snippet to your bash profile

```bash
function sc() {
    switch-context $1 > /tmp/switchcontext
    source /tmp/switchcontext
}
```

This allows the environment variable changes to persist in the shell session.

## Configuration

create a folder in your home directory named `.switchcontext`
and a file inside named `profiles.yaml`

for example:

```yaml
log:
  filename: switchcontext.log
  maxSize: 10
  maxBackups: 5
  maxAge: 30
profiles:
  usprod:
    env:
      AWS_PROFILE: default
      AWS_REGION: us-east-1
    kube: us-prod
    bash:
      PS1: '\h:\W (usp) \u\$'
  ukprod:
    env:
      AWS_PROFILE: ukprod
      AWS_REGION: eu-west-2
    kube: uk-prod
    bash:
      PS1: '\h:\W (ukp) \u\$'
```

## Usage

Call `sc` and then the name of your profile

`sc usprod`
