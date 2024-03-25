# winget-cli
A CLI wrapper for the Windows Package Manager (winget)

It only fetches a list of packages from a personal Gist I have, displays a checklist of packages and installs the selected ones.

It's only intended for personal use and tinkering.  
But if you want you can set the environment variable `GIST_URL` to a different URL and it should work as long as the returned text has the format `install <packages>\n`.