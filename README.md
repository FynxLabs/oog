# OOG

Bot Engine that enables code agnostic plugins

## Why and What

I like to tinker and I don't like being limited to a single code base and a bunch of DSL type rules to create plugins for chatops. So OOG, it's a rebuilding of a previous project I worked on called [SLAPI](https://github.com/ImperialLabs/slapi). It has a lot of similar concepts but was written in ruby. I wanted to build something that requires very little in the way of requirements to extend a chatbot. The first step is to not need a bunch of crap to make it work. So Go was my first choice (yay compiled binaries and no bs install), but I only knew ruby (holy bundler list and version nightmares). So consider the first project a proof of concept while I learned go (but has a lot of info for what to expect here).

So, what kind of requirements am I looking at for extending this chatbot?

A Plugin will require these 3ish things to be functional. More on this after I get the project out of "I'm throwing code into files and clicking run a lot"

- Info about plugin (for help list and such), you will explain what it does. We're not hethens.
    - Either via a endpoint, part of config, or docker labels
- Can Run in docker (there's an or here too, see next line)
    - This is pretty broad, but for the most part. If you can `docker run $image $cmd` it can be a plugin
    - I do have AWS Fargate on my roadmap for potential plugin options, that's like v0.2.0 though... v0.1.0ish is "Hey it turns on and says Pong"
- OR! It needs to be an API and accept a JSON payload full of fun info
    - This could get stupid complicated, so initially I'll probably limit this to a single endpoint like `localhost:40030/command` or `myawesomedomain.com/bot/command` or whatever.
    - Lambda and API Endpoints are great options here!!

## Potentials Questions and Snarky Remarks

- Why use internal libraries? Are you a facist?
    - Ok, I've seen lots of debate about internal pkgs... and I don't care. I like clean and organized and it's a bot. You're not importing this and it's all open source. Take and tinker, don't care.
- Why does your code look like a 3rd grader wrote it?
    - Becuase I code at a 3rd grade level
- OOG?
    - I dunno, I had some specific reason and cool language translation that meant something neat, then forgot and just like saying oog.

## Contributing
Not yet, sit and wait while I finish brain dumping... and breaking it... repeatedly.

Once it's functionalish and I have the bulk of things sorted out and in it's general layout with a roadmap of sorts... docs. Then I'll share it with the unforgiving, opinionated masses. 