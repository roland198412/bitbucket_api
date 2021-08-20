### Bitbucket V2 API - Total commits per user. All repositories

Simple GoLang script to obtain a total commit count per user/ author on all available Bitbucket repositories.

#### USAGE:

Set the following three environmental variables:

```
export BITBUCKET_USERNAME=<bitbucket_username>
export BITBUCKET_PASSWORD=<bitbucket_password>

#Your bitbucket role you are assigned to. e.g. admin, contributor, member, owner
export BITBUCKET_ROLE=admin 
```

Run the script

```
go run main.go
```

