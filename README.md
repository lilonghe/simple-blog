# Target
a simple blog system.

base on [@firekylin](https://github.com/firekylin/firekylin) design.

# Start
- mv `config.default.yaml` -> `config.yaml` and edit config

# Tools
`export` can export all posts(exclude deleted) as markdown file. the markdown include `hugo` metadata(date, draft, tags, categories...).

linux usage:
```
chmod +x export
./export 
```

will generate all article into posts dir.  
so just run `cp ./posts/* hugo-site/content/posts/`, but static assets need you config nginx path.

adapter hugo theme -> [hugo-theme-simply](https://github.com/lilonghe/hugo-theme-simply)

---
## Record some bug

1. xorm `*string` column can not update except use AllCols().