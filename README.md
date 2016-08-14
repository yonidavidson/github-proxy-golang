# Backend challange (Go)

## Task description
Your goal is to present a GitHub user public github repositories drawn from GH's API. 
You will present that as a REST API. 

## Task details

* Implement a single API endpoint at `/api/gh/:username` (where `username` is a variable, e.g. `yonidavidson`)
* This API endpoint will present the username's public github info according to the description below.
* Given a GH username (e.g. `rantav`) - list all repositories for this user
* The repositories should be sorted by score, where `repo.score = repo.forks + 2 * repo.stargazers_count + repo.watchers`. This score should also be presented in the API response. 
* For each repo, return a few attributes (per your choice), such as the number of stars, followers, repo name of course, possibly forks and other things you think are interesting.
* All API requests should use basic HTTP authentication (username/password)
* Response as json
* You may use goji lib https://github.com/zenazn/goji
* Extra points (not required): Add params such as `&limit` (on the total number of repos) and `&page` (for pagination)
* Extra points: add more info about each repo, such as who are the other committers, and some information about them.
* Extra points: GH have a limit on the number of requests. Authenticate and use a token in order to get more requests  
* Extra points: Add caching, using memcached or redis. Make sure you add them to the setup instructions in case you do. 

If you have any questions regarding the exercise please feel free to ask. (email ran@yodas.com)

## Things to focus on:

* Clear and concise commits with explanatory commit messages
* Instruction how to run the program from source. Preferably a Makefile with two targets: `setup` and `run`, which will do the expected...
* Test coverage is extra nice. (`make test` as another target)
* Aesthetics doesn't hurt. (code aesthetics and API aesthetics) Clean and correct code. Modularity, encapsulation, reusability, dry etc, simple coding best practices.

When you're ready, after committing and pushing to this repo, ping me. ran@yodas.com

Reference: 
* GitHub API: https://developer.github.com/v3/
* goji: https://github.com/zenazn/goji

Thank you and hope to see you on our team!
