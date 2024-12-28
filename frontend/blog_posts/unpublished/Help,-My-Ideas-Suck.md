---
title: "Help, My Ideas Suck"
publishedAt: "2024-11-15"
summary: ""
order: 4
---

# Help, My Ideas Suck

> Well, the last stage of grief is acceptance, so at least you're coming to

If you've ever searched **_Coding Project Ideas_** you're sure to come have come across the Coding Course Industrial Complex, and it sucks. Truly.

The thing I've found, is that as you build stuff, you realise **_wow, this dependency sucks_** or is at least annoying. And I think that is one of the ways you can find inspiration.

To list a few personal examples:

- Deployment (AWS, Railway)
- Building (Makefile, Docker)
- Routing (Fiber, Chi)

Is that everything that I use to deploy websites, yes. Not that I dislike any of these things, but there's always something I wish could be changed.

## Auto Docs

Ever browsed a project on github and thought **_wow, these docs are terrible_**. Fear not then, for the next solution to your problems could become your problem to maintain as well!

Generate docs from a project, and display them in either Markdown, HTML, or both. Maybe the user could add amendments to the docs, which are saved when the docs are regenerated.

If you want to see a good example of such a thing, look at the cargo's (Rust Cargo) generated docs.

## Service Deployment

Now, it's not going to be the next AWS. But, if you've spent any time deploying to a cloud provider, or really anything for that matter, you'll know it's bloody annoying.

Basic Outline:

- Servers for hosting (user-provided, cloud-provided, or both)
- Managing site (localhost, provided, or both)
- Configuration (TOML or YAML file)

You want to be able to type a command, or click a button, which triggers deployment.

I think you could make a lightweight Railway.app or Fly.io
