# GitHub Style Plaenkler

## Pin post

```
---
pin: true
---
```

## Add new post

Hugo will create a post with `draft: true`, change it to false in order for it to show in the website.

```
hugo new post/title_of_the_post.md
```

## Limit display content

### Approch 1: use summary

```
---
title: "title"
date: 2019-10-22T18:46:47+08:00
draft: false
summary: "The summary content"
---
```

### Approch 2: use `<!--more-->`

Use `<!--more-->` to seperate content that will display in the posts page as abstraction and the rest of the content. This is different from summary, as summary will not appear in the post.
```
---
title: "title"
date: 2019-10-22T18:46:47+08:00
draft: false
---
abstraction show in the post page
<!--more-->
other content
```

## Add last modified data

add to `config.toml`

```toml
lastmod = true

[frontmatter]
  lastmod = ["lastmod", ":fileModTime", ":default"]
```

## Use [gitalk](https://GitHub.com/gitalk/gitalk) to support comments

add to `config.toml`

```toml
enableGitalk = true

  [params.gitalk]
    clientID = "Your client ID" 
    clientSecret = "Your client secret" 
    repo = "repo"
    owner = "Your GitHub username"
    admin = "Your GitHub username"
    id = "location.pathname"
    labels = "gitalk"
    perPage = 30
    pagerDirection = "last"
    createIssueManually = true
    distractionFreeMode = false
```

## Support LaTex

In you post add `math: true` to [front matter](https://gohugo.io/content-management/front-matter/)

```
---
katex: math
---
```

Then the [katex script](https://katex.org/docs/autorender.html) will auto render the string enclosed be delimiters.

```
# replace ... with latex formula
display inline \\( ... \\)
display block $$ ... $$
```

![latex example](images/latex_example.png)

## Support MathJax
you can add MathJax:true to frontmatter

```
mathJax: true
```
## config.toml example

```toml
baseURL = "https://plaenkler.com/"
languageCode = "de-de"
title = "Plaenkler's blog"
theme = "GitGub Style Plaenkler"
googleAnalytics = "UA-123456-789"
pygmentsCodeFences = true
pygmentsUseClasses = true

[params]
  author = "Plaenkler"
  description = "No cookies"
  GitHub = "Plaenkler"
  facebook = "Plaenkler"
  twitter = "Plaenkler"
  linkedin = "Plaenkler"
  instagram = "Plaenkler"
  tumblr = "Plaenkler"
  email = "plaenkler@gmail.com"
  url = "https://Plaenkler.com"
  keywords = "blog, google analytics"
  rss = true
  lastmod = true
  userStatusEmoji = "😀"
  favicon = "/images/GitHub.png"
  location = "Germany"
  enableGitalk = true

  [params.gitalk]
    clientID = "Your client ID" 
    clientSecret = "Your client secret" 
    repo = "repo"
    owner = "Plaenkler"
    admin = "Plaenkler"
    id = "location.pathname"
    labels = "gitalk"
    perPage = 15
    pagerDirection = "last"
    createIssueManually = true
    distractionFreeMode = false

  [[params.links]]
    title = "Link"
    href = "https://GitHub.com/Plaenkler"
  [[params.links]]
    title = "Link2"
    href = "https://plaenkler.com"
    icon = "https://plaenkler.com/images/avatar.png"

[frontmatter]
  lastmod = ["lastmod", ":fileModTime", ":default"]

```

## deploy.sh example

There are various way to deploy to GitHub, here is a link to official [document](https://gohugo.io/hosting-and-deployment/hosting-on-GitHub/).

Here is an sample. Note line 22 have `env HUGO_ENV="production"`, makes sure googleAnalysis is loaded during production, but is not loaded when we are testing it in localhost.

```bash
#!/bin/sh

if [ "`git status -s`" ]
then
    echo "The working directory is dirty. Please commit any pending changes."
    exit 1;
fi

echo "Deleting old publication"
rm -rf public
mkdir public
git worktree prune
rm -rf .git/worktrees/public/

echo "Checking out gh-pages branch into public"
git worktree add -B gh-pages public origin/gh-pages

echo "Removing existing files"
rm -rf public/*

echo "Generating site"
env HUGO_ENV="production" hugo -t GitHub-Style-Plaenkler

echo "Updating gh-pages branch"
cd public && git add --all && git commit -m "Publishing to gh-pages (publish.sh)"

#echo "Pushing to GitHub"
#git push --all
```

Then you can verify the site is working and use `git push --all` to push the change to GitHub. If you don't want to check again every time, you can uncomment the `#git push --all` in the script.