# httpclout

This repo is going to be an alternative [frontend](https://github.com/bitclout/frontend/) that you can run.

It will try and use as little javascript as possible and put most logic in golang.


# history

a little webapp that just imports [cloutcli](https://github.com/andrewarrow/cloutcli) and calls

```
cloutcli.FollowingFeedPosts(username)
```

Notice you just change "?username=x" in the url to get a different following feed.

![image](https://images.bitclout.com/9a7eec182c96477ea41ee14d6e803d67ab04d4bc9feb76f8f2aff450edf2798d.webp)
