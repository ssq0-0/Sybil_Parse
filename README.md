# Sybil_Parse
LayerZero-Labs Github parser for the presence of your wallet in reports

# Installation
Make sure you have GO installed:

`go version`

#### If you do not have GO installed:

- MacOS:
  ```
  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
  
  brew install go
  
  go version
  ```
- Linux
  ```
  sudo apt update

  sudo apt install golang-go

  go version
  ```

- Windows
  ```
  curl -LO https://go.dev/dl/go1.22.3.windows-amd64.msi

  msiexec /i go1.22.3.windows-amd64.msi /quiet /norestart


  $oldPath = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::Machine)
  $newPath = $oldPath + ";C:\Go\bin"
  [System.Environment]::SetEnvironmentVariable("Path", $newPath, [System.EnvironmentVariableTarget]::Machine)
  exit
  ```

### Start script
- mkdir parse
- cd parse
- git clone https://github.com/ssq0-0/Sybil_Parse
- cd Sybil_Parse

- Change these lines to your addresses
      (for open file in terminal use **nano "filename"**)
  ```go
  	searchStrings := []string{
		  "<b>you_wallet_1</b>",
      "<b>you_wallet_2</b>",
      "<b>you_wallet_3</b>",
      "<b>you_wallet_4</b>",
	}
  ```
  
- Change **githubToken** to your real token
- Start script:
  ```
  go run main.go
  
  ```


## Sybil detected

If your wallet didn't reporting you see:
```

Processing complete.

```

If your wallet was reported you see:
```
Found in Issue #IssueNumber
Matching string: 'Wallet'
Title: 'Title report'
Body:
**etc**
```

Any questions? Write in the channel chat: https://t.me/ssqcrypto
