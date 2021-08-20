package main

import (
	"bitbucket_api/internal/bitbucket_client"
	"fmt"
	"log"
	"os"
	"time"
)

// TODO Group by year
// TODO Group by repo

// Variables to hold count
var authorCommitCount = make(map[string]int)

func main()  {
	userName, ok := os.LookupEnv("BITBUCKET_USERNAME")
	if !ok {
		log.Fatalln("Please specify a bitbucket username")
	}

	passWord, ok := os.LookupEnv("BITBUCKET_PASSWORD")
	if !ok {
		log.Fatalln("Please specify a bitbucket password")
	}

	// Your bitbucket role you are assigned to. e.g. admin, contributor, member, owner
	role, ok := os.LookupEnv("BITBUCKET_ROLE")
	if !ok {
		log.Fatalln("Please specify your bitbucket role")
	}

	client := bitbucket_client.NewBitBucketClient(
		fmt.Sprintf("https://api.bitbucket.org/2.0/repositories?role=%s", role),
		userName,
		passWord,
		15*time.Second,
	)

	repos, err := client.GetRepos()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		if len(repos.Next) > 0 {
			for _, itm := range repos.Values {
				commits, err := client.GetCommitDetail(itm.Links.Commits.HRef)
				if err != nil {
					log.Fatalln(err)
				}

				for {
					if len(commits.Next) > 0 {
						for _, c := range commits.Values {
							_, exist := authorCommitCount[c.Author.Raw]
							if exist {
								authorCommitCount[c.Author.Raw]++
							} else {
								authorCommitCount[c.Author.Raw] = 1
							}
						}
					} else {
						break
					}

					commits, err = client.GetCommitDetail(commits.Next)
					if err != nil {
						log.Fatalln(err)
					}
				}
			}

			client.Url = repos.Next
			repos, err = client.GetRepos()
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			fmt.Println("DONE")
			break
		}
	}

	// Printing results
	for author, count := range authorCommitCount {
		fmt.Printf("%60s:\t%5d\n", author, count)
	}
}
